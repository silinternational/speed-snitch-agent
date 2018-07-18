package adminapi

import (
	"fmt"
	"github.com/silinternational/speed-snitch-agent"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSayHello(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	respBody := ""

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(204)
		fmt.Fprintf(w, respBody)
	})

	apiConfig := agent.APIConfig{
		BaseURL: server.URL,
		APIKey:  "testing",
	}

	startTime := time.Now()
	err := SayHello(apiConfig, startTime)
	if err != nil {
		t.Errorf("Failed to say hello, err: %s", err)
		t.Fail()
	}

}

func TestGetConfig(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	respBody := `{
  "Version": {
    "Number": "1.0.0",
    "URL": "https://www.sil.org"
  },
  "Tasks": [
    {
      "Type": "speedTest",
      "Schedule": "5 */6 * * *",
      "TaskData": {
        "StringValues": {
          "testType": "allTests",
          "serverID": "5029"
        },
        "IntValues": {
          "timeOut": 5
        },
        "FloatValues": {
          "maxSeconds": 6.0
        },
        "IntSlices": {
          "downloadSizes": [245388, 505544],
          "uploadSizes": [32768, 65536]
        }
      }
    }  
  ]
}`

	mux.HandleFunc("/config/"+agent.GetMacAddr(), func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, respBody)
	})

	apiConfig := agent.APIConfig{
		BaseURL: server.URL,
		APIKey:  "testing",
	}

	config, err := GetConfig(apiConfig)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}

	if config.Version.Number != "1.0.0" {
		t.Errorf("Version in config is not what was expected (1.0.0)")
		t.Fail()
	}
}

func TestGetConfigEmptyBody(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	respBody := ""

	mux.HandleFunc("/config/"+agent.GetMacAddr(), func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(204)
		fmt.Fprintf(w, respBody)
	})

	apiConfig := agent.APIConfig{
		BaseURL: server.URL,
		APIKey:  "testing",
	}

	config, err := GetConfig(apiConfig)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}

	expected := agent.Config{}

	if config.Version.Number != expected.Version.Number {
		t.Errorf("Returned config.Version.Number not what was expected, got %s", config.Version.Number)
		t.Fail()
	}

}
