package logqueue

import (
	"fmt"
	"github.com/silinternational/speed-snitch-agent"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc("/log/"+agent.GetMacAddr()+"/ping", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "")
	})

	reportedLogs = []string{}

	testLogs := []agent.TaskLogEntry{
		{
			EntryType:     agent.TypePing,
			Latency:       12.123,
			Timestamp:     1525877951,
			NamedServerID: 1234,
		},
	}

	apiConfig := agent.APIConfig{
		BaseURL: server.URL,
		APIKey:  "",
	}
	newLogs := make(chan agent.TaskLogEntry, 10000)

	go Manager(apiConfig, newLogs)

	for _, nextLog := range testLogs {
		newLogs <- nextLog
	}

	// Give the Manager time to do its work
	time.Sleep(time.Millisecond * 10) // allow time for connection to logentries

	close(newLogs)
}
