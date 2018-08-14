package agent

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/openpgp"
	"io"
	"math/big"
	"net"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"
)

const TypePing = "ping"
const TypeSpeedTest = "speedTest"
const TypeReboot = "reboot"
const TypeRestarted = "restarted" // For sending a log after a restart

const TypeError = "error"
const Version = "0.0.11"
const ExeFileName = "speedsnitch"
const MaxSecondsOffset = 50
const NetworkOnline = "online"
const NetworkOffline = "offline"
const ConfigFileName = "speedsnitch.txt"

type APIConfig struct {
	BaseURL string
	APIKey  string
}

type Config struct {
	Version struct {
		Number string `json:"Number"`
		URL    string `json:"URL"`
	} `json:"Version"`

	Log struct {
		Format      string
		Destination string
	}

	Tasks []Task `json:"Tasks"`
}

type Task struct {
	Type          string   `json:"Type"`
	Schedule      string   `json:"Schedule"`
	NamedServerID uint     `json:"NamedServerID"`
	ServerHost    string   `json:"ServerHost"`
	TaskData      TaskData `json:"TaskData"`
	SpeedTestRunner
}

type TaskData struct {
	StringValues map[string]string  `json:"StringValues"`
	IntValues    map[string]int     `json:"IntValues"`
	FloatValues  map[string]float64 `json:"FloatValues"`
	IntSlices    map[string][]int   `json:"IntSlices"`
}

type TaskLogEntry struct {
	Timestamp         int64   `json:"Timestamp"`
	EntryType         string  `json:"EntryType"`
	ServerCountry     string  `json:"ServerCountry,omitempty"`
	NamedServerID     uint    `json:"NamedServerID,omitempty"`
	Upload            float64 `json:"Upload,omitempty"`
	Download          float64 `json:"Download,omitempty"`
	Latency           float64 `json:"Latency,omitempty"`
	PacketLossPercent float64 `json:"PacketLossPercent,omitempty"`
	ErrorCode         string  `json:"ErrorCode,omitempty"`
	ErrorMessage      string  `json:"ErrorMessage,omitempty"`
	DowntimeStart     string  `json:"DowntimeStart,omitempty"`
	DowntimeSeconds   int64   `json:"DowntimeSeconds,omitempty"`
}

type SpeedTestResults struct {
	Download          float64       `json:"Download,omitempty"`          // Mb per second
	Upload            float64       `json:"Upload,omitempty"`            // Mb per second
	Latency           time.Duration `json:"Latency,omitempty"`           // Latency in nanoseconds
	PacketLossPercent float64       `json:"PacketLossPercent,omitempty"` // Percentage of package loss on ping
	Timestamp         time.Time     `json:"Timestamp"`
	Error             string        `json:"Error"`
}

func (s *SpeedTestResults)CleanData() {
	if s.PacketLossPercent < 0 {
		s.PacketLossPercent = 0
	}
}

type SpeedTestRunner interface {
	Run(TaskData) (SpeedTestResults, error)
}

type SpeedTestInstance struct {
	SpeedTestRunner
}

// Any struct that implements a Process method - for swapping which Logging service we use
type LogReporter interface {
	Process(string, string, ...interface{}) error
}

// Needed to be able to swap in a customized logging struct that implements a Logger
// To use this ...
//   ` // mycustomlogapp.go
//   `type Logger struct {}`
//   `func (l Logger) Process(logKey, text string, a ...interface{}) { ... }`
//
//   `// main.go
//   `logger := agent.LoggerInstance{mycustomlogapp.Logger{}}`
type LoggerInstance struct {
	LogReporter
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string, mode os.FileMode) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	err = os.Chmod(filepath, mode)
	if err != nil {
		return err
	}

	return nil
}

// IsValidMacAddress checks whether the input is ...
//   - 12 hexacedimal digits OR
//   - 6 pairs of hexadecimal digits separated by colons and/or hyphens
func IsValidMACAddress(mAddr string) bool {
	controller := "^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"
	match, _ := regexp.MatchString(controller, mAddr)

	// no separators
	if !match {
		match, _ = regexp.MatchString("^([0-9A-Fa-f]{12})$", mAddr)
	}

	return match
}

// getMacAddr gets the lowest (alphabetically) MAC hardware
// address of the host machine
func GetMacAddr() string {
	addr := ""
	interfaces, err := net.Interfaces()
	lowestAddress := "zz:zz:zz:zz:zz:zz"

	if err == nil {
		for _, i := range interfaces {
			if bytes.Compare(i.HardwareAddr, nil) != 0 {
				addr = i.HardwareAddr.String()

				if !IsValidMACAddress(addr) {
					continue
				}

				if addr < lowestAddress {
					lowestAddress = addr
				}

			}
		}
	}
	return strings.ToLower(lowestAddress)
}

// GetTimeNow returns the current UTC time in the RFC3339 format
func GetTimeNow() string {
	t := time.Now().UTC()
	return t.Format(time.RFC3339)
}

func GetTaskLogEntry(entryType string) TaskLogEntry {
	return TaskLogEntry{
		Timestamp: time.Now().UTC().Unix(),
		EntryType: entryType,
	}
}

func getCustomAppConfigPath() string {
	paths := []string{
		"C:/ProgramData/speedsnitch/AppConfig",
		"/boot/AppConfig",
		"~/Library/speedsnitch/AppConfig",
	}

	for _, path := range paths {
		_, err := os.Stat(path)
		if err == nil {
			return path
		}
	}
	return ""
}

// GetAppConfig accepts an io.Reader for testing purposes.
//  If the io.Reader param is nil, then it uses the default
//  config file to provide an custom APIConfig
func GetAppConfig(reader io.Reader) APIConfig {
	apiConfig := APIConfig{}

	// If no (test) reader is provided, get the default config file as the reader
	if reader == nil {
		configPath := getCustomAppConfigPath()
		if configPath == "" {
			return apiConfig
		}

		configFilePath := configPath + "/" + ConfigFileName
		var err error
		reader, err = os.Open(configFilePath)
		if err != nil {
			return apiConfig
		}
	}

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	apiConfig.BaseURL = scanner.Text()

	scanner.Scan()
	apiConfig.APIKey = scanner.Text()

	return apiConfig
}

// VerifyFileSignature only checks the signature of the target file if there is a gpg key to use
func VerifyFileSignature(directory, targetFile, signedFile string, keys []io.Reader) error {
	signature, err := os.Open(signedFile)
	if err != nil {
		return err
	}

	verificationTarget, err := os.Open(targetFile)
	if err != nil {
		return err
	}

	for _, keyReader := range keys {

		keyring, err := openpgp.ReadArmoredKeyRing(keyReader)
		if err != nil {
			continue
		}

		_, err = openpgp.CheckArmoredDetachedSignature(keyring, verificationTarget, signature)
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("None of the current keys are able to verify the signature.")
}

func GetRandomSecondAsString() string {
	val, err := rand.Int(rand.Reader, big.NewInt(MaxSecondsOffset))
	if err != nil {
		return "15"
	}
	return fmt.Sprintf("%v", val)
}

// Reboot checks the GOOS and GOARCH and if they are valid, reboots the system
func Reboot() error {

	goSys := runtime.GOOS + " " + runtime.GOARCH

	switch goSys {
	case "linux arm":
		return exec.Command("reboot").Run()
	default:
		return fmt.Errorf("Not Implemented: rebooting %s.", goSys)
	}

}
