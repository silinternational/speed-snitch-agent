package logqueue

import (
	"testing"
	"github.com/silinternational/speed-snitch-agent"
	"time"
	"encoding/json"
	"fmt"
)

var reportedLogs []string

func areStringSlicesEqual(slc1, slc2 []string) bool {
	if len(slc1) != len(slc2) {
		return false
	}

	for index, nextStr := range slc1 {
		if slc2[index] != nextStr {
			return false
		}
	}

	return true
}


type FakeLogger struct {
}

func (f FakeLogger) Process(a, b string, c ...interface{}) error {
	reportedLogs = append(reportedLogs, b)
	return nil
}


func TestManager(t *testing.T) {
	reportedLogs = []string{}
	testLogger := FakeLogger{}

	testLogs := []string {
		"Log11",
		"Log12",
		"Log13",
		FlushLogQueue,
		"Log21",
		FlushLogQueue,
		"Log31",
		"Log32",
		FlushLogQueue,
		FlushLogQueue,
	}

	newLogs := make(chan string, 10000)

	go Manager(newLogs, "fakeLogKey", &agent.LoggerInstance{testLogger})

	for _, nextLog := range testLogs {
		newLogs <- nextLog
	}

	// Give the Manager time to do its work
	time.Sleep(time.Millisecond * 10) // allow time for connection to logentries

	close(newLogs)

	expected := []string {
		"Log11", "Log12", "Log13",
		"Log21",
		"Log31", "Log32",
	}

	var dat map[string]interface{}

	results := []string{}

	for _, nextRaw := range reportedLogs {
		err := json.Unmarshal([]byte(nextRaw), &dat)
		if err != nil {
			t.Errorf("Could not decode the log: %s", nextRaw)
			return
		}

		results = append(results, fmt.Sprintf("%s", dat["log"]))
	}

	if ! areStringSlicesEqual(expected, results) {
		t.Fatalf("Did not get back expected logs.\n  Expected: %s\n.    But Got: %s.", expected, results)
	}
}