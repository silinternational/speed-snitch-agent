package speedtestnet

import (
	"github.com/silinternational/speed-snitch-agent"
	"testing"
	"net"
	"net/http"
	"net/http/httptest"
	"fmt"
	"strings"
	"time"
	"io/ioutil"
)


func fixture(path string) string {
	b, err := ioutil.ReadFile("../../testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func TestNewClient(t *testing.T) {
	client, _ := NewClient()
	if client.Type != agent.TypeSpeedtest {
		t.Error("Speedtest client type not what was epxected, got ", client.Type)
	}
}


func TestLatencyTest(t *testing.T) {
	mux := http.NewServeMux()
	httpTestServer := httptest.NewServer(mux)
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, `{"Test":"results"}`)
	})

	serverID := 5029

	config := configuration{
		Name:          "Test Agent",
		ID:            "TestA",
		ServerID:      serverID,
		Timeout:       5,
		MaxSeconds:    4.0,
	}

	server := GetServerByID(serverID)
	server.Configuration = &config

	// remove "http://" from the httpTestServer URL
	urlParts := strings.Split(httpTestServer.URL, "//")
	server.Host = strings.Join(urlParts[1:2], "")

	localAddr := net.TCPAddr{}
	source, _ := net.ResolveTCPAddr("tcp", localAddr.String())

	spdTest := speedtest{
		Source: source,
	}

	server.Speedtest = &spdTest
	spdTest.Server = &server

	LatencyTest(&server)
	allResults := server.Results

	results := allResults.Latency

	if results <= time.Duration(0) {
		t.Errorf("Error: Expected a positive time.Duration, but got: %s", results)
	}

}

// This does a real latency test if you comment out the t.SkipNow() call
func TestLatencyTestReal(t *testing.T) {
	t.SkipNow()
	serverID := 5029

	config := configuration{
		Name:          "Test Agent",
		ID:            "TestA",
		ServerID:      serverID,
		Timeout:       5,
		MaxSeconds:    4.0,
	}

	server := GetServerByID(serverID)
	server.Configuration = &config

	localAddr := net.TCPAddr{}
	source, _ := net.ResolveTCPAddr("tcp", localAddr.String())

	spdTest := speedtest{
		Source: source,
	}

	server.Speedtest = &spdTest
	spdTest.Server = &server

	LatencyTest(&server)
	allResults := server.Results

	fmt.Printf("Latency test results for server %d ... %s\n", serverID, allResults.Latency)
}


// This does a real download test if you comment out the t.SkipNow() call
func TestDownloadTestReal(t *testing.T) {
	t.SkipNow()
	serverID := 5029

	config := configuration{
		Name:          "Test Agent",
		ID:            "TestA",
		ServerID:      serverID,
		Timeout:       5,
		DownloadSizes: []int{1024, 2048},
		MaxSeconds:    2.0,
	}

	server := GetServerByID(serverID)
	server.Configuration = &config

	localAddr := net.TCPAddr{}
	source, _ := net.ResolveTCPAddr("tcp", localAddr.String())

	spdTest := speedtest{
		Source: source,
	}

	server.Speedtest = &spdTest
	spdTest.Server = &server

	DownloadTest(&server, config.MaxSeconds, config.DownloadSizes)
	allResults := server.Results

	results := allResults.Download

	if results <= 0 {
		t.Errorf("Error: Expected a positive Download result, but got: %f", results)
	}

}

// This does a real upload test if you comment out the t.SkipNow() call
func TestUploadTestReal(t *testing.T) {
	t.SkipNow()
	serverID := 5029

	config := configuration{
		Name:          "Test Agent",
		ID:            "TestA",
		ServerID:      serverID,
		Timeout:       5,
		UploadSizes:   []int{256, 512},
		MaxSeconds:    2.0,
	}

	server := GetServerByID(serverID)
	server.Configuration = &config

	localAddr := net.TCPAddr{}
	source, _ := net.ResolveTCPAddr("tcp", localAddr.String())

	spdTest := speedtest{
		Source: source,
	}

	server.Speedtest = &spdTest
	spdTest.Server = &server

	UploadTest(&server, config.MaxSeconds, config.UploadSizes)
	allResults := server.Results

	results := allResults.Upload

	if results <= 0 {
		t.Errorf("Error: Expected a positive Upload result, but got: %f", results)
	}

}