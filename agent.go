package agent

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

const TypePing = "ping"
const TypeSpeedTest = "speedTest"
const Version = "0.0.2"

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

// getMacAddr gets the MAC hardware
// address of the host machine
func GetMacAddr() string {
	addr := ""
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return strings.ToLower(addr)
}
