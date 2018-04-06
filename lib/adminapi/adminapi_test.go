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

	config := agent.Config{
		BaseURL: server.URL,
	}

	startTime := time.Now()
	err := SayHello(config, startTime)
	if err != nil {
		t.Errorf("Failed to say hello, err: %s", err)
		t.Fail()
	}

}

func TestGetConfig(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	respBody := `{
  "BaseURL": "https://www.sil.org",
  "Version": {
    "Number": "1.0.0",
    "URL": "https://www.sil.org"
  },
  "Tasks": [
    {
      "Type": "speedTest",
      "Schedule": "5 */6 * * *",
      "Data": {
        "StringValues": {
          "testType": "allTests"
        },
        "IntValues": {
          "serverID": 5029,
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

	config, err := GetConfig(server.URL)
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

	config, err := GetConfig(server.URL)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}

	expected := agent.Config{
		BaseURL: server.URL,
	}

	if config.BaseURL != expected.BaseURL {
		t.Errorf("Returned config.BaseURL not what was expected, got %s", config.BaseURL)
		t.Fail()
	}

	emptyVersion := struct {
		Number string `json:"Number"`
		URL    string `json:"URL"`
	}{
		Number: "",
		URL:    "",
	}

	if config.Version != emptyVersion {
		t.Errorf("Returned config.Version not what was expected, got %v", config.Version)
		t.Fail()
	}
}
