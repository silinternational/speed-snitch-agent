package agent

import (
	"testing"
	"strings"
	"strconv"
	"time"
)

func TestGetAppConfig(t *testing.T) {

	baseURL := "http://test.base.com"
	apiKey := "some_key"

	testData := [][3]string{
		// test contents,  expected BaseURL,  expected APIKey
		{baseURL + "\r\n" + apiKey, baseURL, apiKey},
		{baseURL + " " + apiKey, baseURL, apiKey},
		{baseURL + " ", baseURL, ""},
		{"", "", ""},
	}

	for index, data := range testData {
		testContents := data[0]
		allResults := GetAppConfig(strings.NewReader(testContents))

		expected := data[1]
		results := allResults.BaseURL

		if expected != results {
			t.Errorf("Bad BaseURL at index: %d. Expected: %s, but got: %s", index, expected, results)
			return
		}

		expected = data[2]
		results = allResults.APIKey

		if expected != results {
			t.Errorf("Bad APIKey at index: %d. Expected: %s, but got: %s", index, expected, results)
			return
		}
	}
}

func TestGetRandomSecondAsString(t *testing.T) {
	for i := 0; i < 100; i++ {
		randomIntStr := GetRandomSecondAsString()
		asInt, _ := strconv.ParseInt(randomIntStr, 10, 64)
		if asInt < 0 || asInt > MaxSecondsOffset {
			t.Errorf("Got back a random second outside valid range, got: %v", asInt)
			t.Fail()
		}
	}
}


func TestIsValidMACAddress(t *testing.T) {
	
	type TestMacData struct {
		macAddr string
		isValid bool
	}

	testData := []TestMacData{
		{macAddr: "08-00-27-10-B8-D0", isValid: true},
		{macAddr: "00:00:00:00:00:0f", isValid: true},
		{macAddr: "00:00:00:00:00:00:0f", isValid: false},
		{macAddr: "11:22:33:44:55", isValid: false},
	}

	for _, nextData := range testData {
		results := IsValidMACAddress(nextData.macAddr)
		expected := nextData.isValid
		if results != expected {
			t.Errorf("Bad is-Valid Mac Address: %s. Expected: %t. But got: %t", nextData.macAddr, expected, results)
			return
		}
	}

}

func TestSpeedTestResults_CleanData(t *testing.T) {
	spResults := SpeedTestResults{
		Latency: time.Duration(1),
		PacketLossPercent: -9,
	}

	spResults.CleanData()
	if spResults.PacketLossPercent != 0.0 {
		t.Errorf("Bad PacketLossPercent. Expected: 0. But got: %.2f.\n%+v", spResults.PacketLossPercent, spResults)
		return
	}

	spResults = SpeedTestResults{
		Latency: time.Duration(2),
		PacketLossPercent: 3,
	}

	spResults.CleanData()
	if spResults.PacketLossPercent != 3.0 {
		t.Errorf("Bad PacketLossPercent. Expected: 3. But got: %.2f.\n%+v", spResults.PacketLossPercent, spResults)
		return
	}
}