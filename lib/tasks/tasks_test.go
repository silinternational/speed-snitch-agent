package tasks

import (
	"testing"
	"strings"
	"strconv"
	"github.com/silinternational/speed-snitch-agent"
)

func TestGetCronScheduleWithRandomSeconds(t *testing.T) {
	// test with sending 5 positions
	results := getCronScheduleWithRandomSeconds("* * * * *")
	parts := strings.Split(results, " ")
	if len(parts) != 6 {
		t.Error("Did not find 6 parts in schedule after call to getCronScheduleWithRandomSeconds, got: ", results)
		t.Fail()
	}

	asInt, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		t.Error("Unable to parse first part of schedule as integer, err:", err.Error())
		t.Fail()
	}
	if asInt < 0 || asInt > agent.MaxSecondsOffset {
		t.Error("First part of schedule outside acceptable range, got: ", asInt)
	}

	// test with sending 6 positions
	results = getCronScheduleWithRandomSeconds("*/30 * * * * *")
	parts = strings.Split(results, " ")
	if len(parts) != 6 {
		t.Error("Did not find 6 parts in schedule after call to getCronScheduleWithRandomSeconds, got: ", results)
		t.Fail()
	}

	if parts[0] == "*/30" {
		t.Error("Seconds part not replaced after call to getCronScheduleWithRandomSeconds, still */30")
		t.Fail()
	}

	asInt, err = strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		t.Error("Unable to parse first part of schedule as integer, err:", err.Error())
		t.Fail()
	}
	if asInt < 0 || asInt > agent.MaxSecondsOffset {
		t.Error("First part of schedule outside acceptable range, got: ", asInt)
	}
}