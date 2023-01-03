package main

import (
	"fmt"
	"github.com/silinternational/speed-snitch-agent"
	"github.com/silinternational/speed-snitch-agent/lib/adminapi"
	"github.com/silinternational/speed-snitch-agent/lib/logqueue"
	"github.com/silinternational/speed-snitch-agent/lib/speedtestnet"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunLatencyTest(t *testing.T) {

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
      "TaskData": {
        "StringValues": {
          "testType": "latencyTest",
          "serverID": "16976",
          "Host": "speedtest.nyc.rr.com:8080"
        },
        "IntValues": {
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

	apiConfig := agent.APIConfig{
		BaseURL: server.URL,
		APIKey:  "testing",
	}

	config, _ := adminapi.GetConfig(apiConfig)
	taskData := config.Tasks[0].TaskData

	speedster := agent.SpeedTestInstance{speedtestnet.SpeedTestRunner{}}

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

func TestLogEntries(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc("/log/"+agent.GetMacAddr()+"/testLog", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "")
	})

	testLogs := []agent.TaskLogEntry{
		{
			Timestamp: 11111111,
			EntryType: "testLog",
		},
		{
			Timestamp: 22222222,
			EntryType: "testLog",
		},
	}

	newLogs := make(chan agent.TaskLogEntry)

	apiConfig := agent.APIConfig{
		BaseURL: server.URL,
		APIKey:  "testing",
	}

	go logqueue.Manager(apiConfig, newLogs)

	for _, nextLog := range testLogs {
		newLogs <- nextLog
	}

	time.Sleep(time.Millisecond * 1000) // allow time for connection
	close(newLogs)

}
