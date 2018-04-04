package logqueue

import (
	"testing"
	"github.com/silinternational/speed-snitch-agent"
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
	keepOpen := make(chan int)

	go Manager(newLogs, keepOpen, "fakeLogKey", &agent.LoggerInstance{testLogger})

	for _, nextLog := range testLogs {
		newLogs <- nextLog
	}

	logCount := 0
	for {
		logCount += <-keepOpen
		if logCount >= 10 {
			break
		}
	}

	close(keepOpen)
	close(newLogs)

	expected := []string {
		"Log11", "Log12", "Log13",
		"Log21",
		"Log31", "Log32",
	}

	results := reportedLogs

	if ! areStringSlicesEqual(expected, results) {
		t.Fatalf("Did not get back expected logs.\n  Expected: %s\n.    But Got: %s.", expected, results)
	}
}