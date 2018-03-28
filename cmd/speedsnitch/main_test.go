package main

import (
	"testing"
	"fmt"
	"github.com/silinternational/speed-snitch-agent/lib/speedtestnet"
)


// This does a real latency test if you comment out the t.SkipNow() call
func TestRunLatencyTest(t *testing.T) {
	t.SkipNow()
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
