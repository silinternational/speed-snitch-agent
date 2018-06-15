package icmp

import (
	"testing"
)

func TestPing(t *testing.T) {
	speedTestResults, err := Ping("google.com", 2, 1, 10)
	if err != nil {
		t.Error("Ping failed:", err)
		t.Fail()
	}
	if speedTestResults.Latency <= 0 {
		t.Errorf("Error running latency test, returned latency is less than or equal to zero: %v", speedTestResults.Latency)
	}

	// test error for bad  host
	_, err = Ping("badhost-alalalalflflwlelelf.com", 2, 1, 10)
	if err == nil {
		t.Error("Ping did not return error for bad hostname")
		t.Fail()
	}
}
