package logqueue

import (
	"testing"
	"github.com/silinternational/speed-snitch-agent"
)

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

func areStringMapsEqual(map1, map2 map[string][]string) bool {
	if len(map1) != len(map2) {
		return false
	}

	for k1, v1 := range map1 {
		v2, ok := map2[k1]
		if !ok {
			return false
		}
		if ! areStringSlicesEqual(v1, v2) {
			return false
		}
	}

	return true
}

type FakeLogger struct {}

func (f FakeLogger) Process(a, b string, c ...interface{}) error {
	// Don't do anything with it
	return nil
}

func TestAppendMapValueStart(t *testing.T) {
	wholeMap := map[string][]string{}
	appendMapValue("NewType", "value1", wholeMap)

	expected := map[string][]string{"NewType": {"value1"}}

	if !areStringMapsEqual(expected, wholeMap) {
		t.Fatalf("Maps are not equal.\n  Expected: %v\n   But got: %v", expected, wholeMap)
	}
}

func TestAppendMapValueNew(t *testing.T) {
	wholeMap := map[string][]string{"OldType": {"oldValue"}}
	appendMapValue("NewType", "value1", wholeMap)

	expected := map[string][]string{"OldType": {"oldValue"}, "NewType": {"value1"}}

	if !areStringMapsEqual(expected, wholeMap) {
		t.Fatalf("Maps are not equal.\n  Expected: %v\n   But got: %v", expected, wholeMap)
	}
}

func TestAppendMapValueAddition(t *testing.T) {
	wholeMap := map[string][]string{"OldType": {"oldValue"}}
	appendMapValue("OldType", "value2", wholeMap)

	expected := map[string][]string{"OldType": {"oldValue", "value2"}}

	if !areStringMapsEqual(expected, wholeMap) {
		t.Fatalf("Maps are not equal.\n  Expected: %v\n   But got: %v", expected, wholeMap)
	}
}


func TestReporter(t *testing.T) {
	reportedLogs := []string{}

	testTracker := TestTracker{
		KeepTrack: true,
	}

	testLogs := [][2]string {
		{"Type1", "Log11"},
		{"Type2", "Log21"},
		{"Type1", "Log12"},
		{"Type2", "Log22"},
		{"Type2", FlushLogQueue}, // Type2 gets flushed first
		{"Type3", "Log31"},
		{"Type1", FlushLogQueue}, // Type1 gets flushed second but will have another log after
		{"Type1", "Log13"},
		{"Type3", FlushLogQueue}, // Type3 gets flushed third and before the second round for Type1
		{"Type1", FlushLogQueue}, // Type1 gets flushed again with the latest log entry
	}

	newLogs := make(chan [2]string)
	completedLogs := make(chan []string)

	keepOpen := make(chan int, 4)

	go Stasher(newLogs, completedLogs)
	go Reporter(completedLogs, keepOpen, "fakeLogKey", agent.LoggerInstance{FakeLogger{}}, &testTracker)

	for _, nextSet := range testLogs {
		newLogs <- nextSet
	}

	close(newLogs)

	logCount := 0
	for {
		logCount += <-keepOpen
		if logCount >= 4 {
			break
		}
	}

	close(keepOpen)
	close(completedLogs)

	expected := []string {
		"Log21", "Log22",
		"Log11", "Log12",
		"Log31",
		"Log13",
	}

	if ! areStringSlicesEqual(expected, testTracker.ReportedLogs) {
		t.Fatalf("Did not get back expected logs.\n  Expected: %s\n.    But Got: %s.", expected, reportedLogs)
	}
}