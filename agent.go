package agent

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

const RepoURL = "https://github.com/silinternational/speed-snitch-agent"
const TypePing = "ping"
const TypeSpeedTest = "speedTest"
const Version = "0.0.1"

type Config struct {
	BaseURL string `json:"base_url"`

	Version struct {
		Number string `json:"number"`
		URL    string `json:"url"`
	} `json:"version"`

	Log struct {
		Format      string
		Destination string
	}

	Tasks []Task `json:"tasks"`
}

type Task struct {
	Type     string   `json:"type"`
	Schedule string   `json:"schedule"`
	Data     TaskData `json:"data"`
	SpeedTestRunner
}

type TaskData struct {
	StringValues map[string]string  `json:"StringValues"`
	IntValues    map[string]int     `json:"IntValues"`
	FloatValues  map[string]float64 `json:"FloatValues"`
	IntSlices    map[string][]int   `json:"IntSlices"`
}

type SpeedTestResults struct {
	Download  float64       `json:"download"` // Mb per second
	Upload    float64       `json:"upload"`   // Mb per second
	Latency   time.Duration `json:"latency"`  // Latency in nanoseconds
	Timestamp time.Time     `json:"timestamp"`
	Error     string        `json:"error"`
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

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

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
	return addr
}
