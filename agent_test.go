package agent

import (
	"testing"
	"strings"
	"strconv"
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
		if asInt < 0 || asInt > 60 {
			t.Errorf("Got back a random second outside valid range, got: %v", asInt)
			t.Fail()
		}
	}
}
