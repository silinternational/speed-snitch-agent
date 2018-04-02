package agent

import (
	"io"
	"net/http"
	"os"
	"time"
)

const RepoURL = "https://github.com/silinternational/speed-snitch-agent"
const TypePing = "ping"
const TypeSpeedTest = "speedTest"
const Version = "0.0.1"


type Config struct {
	Version struct {
		Latest string
		URL    string
	}

	Log struct {
		Format      string
		Destination string
	}

	Tasks []Task
}

type Status struct {
	Version string
	Uptime  string
}

type Task struct {
	Type     string
	Schedule string
	Data     TaskData
	SpeedTestRunner
}

type TaskData struct {
	StringValues map[string]string
	IntValues map[string]int
	FloatValues map[string]float64
	IntSlices map[string][]int
}

type SpeedTestResults struct {
	Download  float64  // Mb per second
	Upload    float64  // Mb per second
	Latency   time.Duration // seconds
	Timestamp time.Time
	Error     string
}

type SpeedTestRunner interface {
	Run(TaskData) (SpeedTestResults, error)
}

type SpeedTestInstance struct {
	SpeedTestRunner
}

// Any struct that implements a Process method - for swapping which Logging service we use
type LogReporter interface{
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
