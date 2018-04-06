package main

import (
	"fmt"
	"github.com/silinternational/speed-snitch-agent"
	"github.com/silinternational/speed-snitch-agent/lib/adminapi"
	"github.com/silinternational/speed-snitch-agent/lib/logentries"
	"github.com/silinternational/speed-snitch-agent/lib/logqueue"
	"github.com/silinternational/speed-snitch-agent/lib/speedtestnet"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

// This does a real latency test unless you use the -short flag
func TestRunLatencyTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	respBody := `{
  "BaseURL": "https://www.sil.org",
  "Version": {
    "Number": "1.0.0",
    "URL": "https://www.sil.org"
  },
  "Tasks": [
    {
      "Type": "speedTest",
      "Schedule": "5 */6 * * *",
      "Data": {
        "StringValues": {
          "testType": "allTests",
          "Host": "nyc.speedtest.sbcglobal.net:8080"
        },
        "IntValues": {
          "serverID": 5029,
          "timeOut": 5
        },
        "FloatValues": {
          "maxSeconds": 6.0
        },
        "IntSlices": {
          "downloadSizes": [245388, 505544],
          "uploadSizes": [32768, 65536]
        }
      }
    }  
  ]
}`

	mux.HandleFunc("/config/"+agent.GetMacAddr(), func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, respBody)
	})

	config, _ := adminapi.GetConfig(server.URL)
	taskData := config.Tasks[0].Data

	speedster := config.Tasks[0].SpeedTestRunner
	spTestResults, err := speedster.Run(taskData)

	if err != nil {
		t.Fatalf("Unexpected Error: \n%s", err.Error())
	}

	results := spTestResults.Latency.Seconds()

	if results <= 0 {
		t.Errorf("Error: Expected a positive Latency result, but got: %f", results)
	} else {
		fmt.Printf(
			"\nLatency test results for server %d ... %f\n",
			taskData.IntValues[speedtestnet.CFG_SERVER_ID],
			results,
		)
	}

	results = spTestResults.Download
	if results != 0 {
		t.Errorf("Error: Expected a zero Download result, but got: %f", results)
	}

	results = spTestResults.Upload
	if results != 0 {
		t.Errorf("Error: Expected a zero Upload result, but got: %f", results)
	}
}

// This does a real call to logentries unless you use the -short flag
func TestLogEntries(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	logEntriesKey := os.Getenv("LOGENTRIES_KEY")

	if logEntriesKey == "" {
		t.Fatal("No LOGENTRIES_KEY env variable available.")
	}

	logger := agent.LoggerInstance{logentries.Logger{}}

	testLogs := []string{
		"Speed Snitch Agent: TestLogEntries ...  log1",
		"Speed Snitch Agent: TestLogEntries ...  log2",
		logqueue.FlushLogQueue,
	}

	newLogs := make(chan string)

	go logqueue.Manager(newLogs, logEntriesKey, &logger)

	for _, nextLog := range testLogs {
		newLogs <- nextLog
	}

	time.Sleep(time.Millisecond * 1000) // allow time for connection to logentries
	close(newLogs)

	println(`TO SEE THE RESULTS OF THIS TEST
Go to the logentries set that matches your LOGENTRIES_KEY env var and look for ...
Speed Snitch Agent: TestLogEntries ...  log1
Speed Snitch Agent: TestLogEntries ...  log2`)
}
