package main

import (
	"testing"
	"fmt"
	"github.com/silinternational/speed-snitch-agent/lib/speedtestnet"
	"github.com/silinternational/speed-snitch-agent"
	"github.com/silinternational/speed-snitch-agent/lib/logentries"
	"os"
	"github.com/silinternational/speed-snitch-agent/lib/logqueue"
	"time"
)


// This does a real latency test unless you use the -short flag
func TestRunLatencyTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	config := getConfig()
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

	testTracker := logqueue.TestTracker{
		KeepTrack: false,
	}

	logger := agent.LoggerInstance{logentries.Logger{}}

	testLogs := [][2]string {
		{"Type1", "Speed Snitch Agent: TestLogEntries ...  log1"},
		{"Type1", "Speed Snitch Agent: TestLogEntries ...  log2"},
		{"Type1", logqueue.FlushLogQueue},
	}

	newLogs := make(chan [2]string)
	completedLogs := make(chan []string)
	keepOpen := make(chan int, 1)

	go logqueue.Stasher(newLogs, completedLogs)
	go logqueue.Reporter(completedLogs, keepOpen, logEntriesKey, logger, &testTracker)

	for _, nextSet := range testLogs {
		newLogs <- nextSet
	}

	<-keepOpen

	time.Sleep(time.Duration(time.Millisecond * 500)) // allow time for connection to logentries
	close(newLogs)
	close(completedLogs)
	close(keepOpen)

	println(`TO SEE THE RESULTS OF THIS TEST
Go to the logentries set that matches your LOGENTRIES_KEY env var and look for ...
Speed Snitch Agent: TestLogEntries ...  log1
Speed Snitch Agent: TestLogEntries ...  log2`)
}