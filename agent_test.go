package agent

import (
	"testing"
	"strings"
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

