package agent

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
	"bufio"
	"golang.org/x/crypto/openpgp"
	"fmt"
)

const TypePing = "ping"
const TypeSpeedTest = "speedTest"
const TypeError = "error"
const Version = "0.0.2.1"
const ExeFileName = "speedsnitch"

const ConfigPath = "/boot/AppConfig"
const ConfigFileName = "speedsnitch.txt"

const GPGKeyFileName = "gpg.pubkey"

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
	Type     string   `json:"Type"`
	Schedule string   `json:"Schedule"`
	Data     TaskData `json:"Data"`
	SpeedTestRunner
}

type TaskData struct {
	StringValues map[string]string  `json:"StringValues"`
	IntValues    map[string]int     `json:"IntValues"`
	FloatValues  map[string]float64 `json:"FloatValues"`
	IntSlices    map[string][]int   `json:"IntSlices"`
}

type TaskLogEntry struct {
	Timestamp    int64   `json:"Timestamp"`
	EntryType    string  `json:"EntryType"`
	ServerID     int     `json:"ServerID,omitempty"`
	Upload       float64 `json:"Upload,omitempty"`
	Download     float64 `json:"Download,omitempty"`
	Latency      float64 `json:"Latency,omitempty"`
	ErrorCode    string  `json:"ErrorCode,omitempty"`
	ErrorMessage string  `json:"ErrorMessage,omitempty"`
}

type SpeedTestResults struct {
	Download  float64       `json:"Download,omitempty"` // Mb per second
	Upload    float64       `json:"Upload,omitempty"`   // Mb per second
	Latency   time.Duration `json:"Latency,omitempty"`  // Latency in nanoseconds
	Timestamp time.Time     `json:"Timestamp"`
	Error     string        `json:"Error"`
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

// GetAppConfig accepts an io.Reader for testing purposes.
//  If the io.Reader param is nil, then it uses the default
//  config file to provide an custom APIConfig
func GetAppConfig(reader io.Reader) APIConfig {
	apiConfig := APIConfig{}

	// If no (test) reader is provided, get the default config file as the reader
	if reader == nil {
		configFilePath := ConfigPath + "/" + ConfigFileName
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
func VerifyFileSignature(directory, targetFile, signedFile string) error {
	keyFilePath := ConfigPath + "/" + GPGKeyFileName

	// If there is no key, then don't try to verify it
	_, err := os.Stat(keyFilePath)
	if os.IsNotExist(err) {
		return nil
	}

	keyRingReader, err := os.Open(keyFilePath)
	if err != nil {
		return err
	}

	signature, err := os.Open(signedFile)
	if err != nil {
		return err
	}

	verificationTarget, err := os.Open(targetFile)
	if err != nil {
		return err
	}

	keyring, err := openpgp.ReadArmoredKeyRing(keyRingReader)
	if err != nil {
		return fmt.Errorf("Error Reading Armored Key Ring: %s", err.Error())
	}

	_, err = openpgp.CheckArmoredDetachedSignature(keyring, verificationTarget, signature)
	if err != nil {
		return err
	}

	return nil
}