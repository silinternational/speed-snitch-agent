package speedtestnet

import (
	"github.com/silinternational/speed-snitch-agent"
	"testing"
	"net"
	"net/http"
	"net/http/httptest"
	"fmt"
	"strings"
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
	if client.Type != agent.TypeSpeedTest {
		t.Error("Speedtest client type not what was epxected, got ", client.Type)
	}
}


func TestLatencyTestMock(t *testing.T) {
	mux := http.NewServeMux()
	httpTestServer := httptest.NewServer(mux)
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, `{"Test":"results"}`)
	})

	serverID := "5029"

	config := configuration{
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

	results := allResults.Latency.Seconds()

	if results <= 0 {
		t.Errorf("Error: Expected a positive Latency result, but got: %f", results)
	}

}

// This does a real latency test unless you use the -short flag
func TestLatencyTestReal(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	serverID := "5029"

	config := configuration{
		ServerID:      serverID,
		Timeout:       5,
		//MaxSeconds:    4.0,
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

	results := allResults.Latency.Seconds()

	if results <= 0 {
		t.Fatalf("Error: Expected a positive Latency result, but got: %f", results)
	}
	fmt.Printf("\nLatency test results for server %s ... %f\n", serverID, results)
}


// This does a real download test unless you use the -short flag
func TestDownloadTestReal(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	serverID := "5029"

	config := configuration{
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

	fmt.Printf("\nDownload test results for server %s ... %f\n", serverID, results)
}

// This does a real upload test unless you use the -short flag
func TestUploadTestReal(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	serverID := "5029"

	config := configuration{
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
	fmt.Printf("\nUpload test results for server %s ... %f\n", serverID, results)
}


func TestRunTestBadTestType(t *testing.T) {
	emptyTestResults := agent.SpeedTestResults{}
	testType := "BadTestType"

	taskData := agent.TaskData{
		StringValues: map[string]string{
			CFG_TEST_TYPE: testType,
			CFG_SERVER_HOST: "nyc.speedtest.sbcglobal.net:8080",
			CFG_SERVER_ID: "5029",
		},
	}

	spdTestRunner := SpeedTestRunner{}

	spTestResults, err := spdTestRunner.Run(taskData)
	if spTestResults != emptyTestResults {
		t.Fatalf("Error: expected empty test results but got:\n%v", spTestResults)
	}

	expected := fmt.Sprintf("Invalid value in TaskData for testType: %s", testType)
	if err.Error() != expected {
		t.Fatalf(
			"Error: Got wrong error message.\n  Expected: %s\n    But got: %s\n",
			expected,
			err.Error(),
		)
	}
}


// This does a real latency test unless you use the -short flag
func TestRunTestLatencyReal(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	taskData := agent.TaskData{
		StringValues: map[string]string{
			CFG_TEST_TYPE: CFG_TYPE_LATENCY,
			CFG_SERVER_HOST: "nyc.speedtest.sbcglobal.net:8080",
			CFG_SERVER_ID: "5029",
		},
		IntValues: map[string]int {
			CFG_TIME_OUT: 5,
		},
	}

	spdTestRunner := SpeedTestRunner{}

	spTestResults, err := spdTestRunner.Run(taskData)
	if err != nil {
		t.Fatalf("Unexpected Error: \n%s", err.Error())
	}

	results := spTestResults.Latency.Seconds()

	if results <= 0 {
		t.Errorf("Error: Expected a positive Latency result, but got: %f", results)
	} else {
		fmt.Printf("\nLatency test results for server %d ... %f\n", taskData.IntValues[CFG_SERVER_ID], results)
	}

	results = spTestResults.Download
	if results != 0 {
		t.Errorf("Error: Expected a zero Download result, but got: %f", results)
	}

	results = spTestResults.Upload
	if results != 0 {
		t.Errorf("Error: Expected a zero Upload result, but got: %f", results)
	}
}

// This does real speed tests unless you use the -short flag
func TestRunTestAllReal(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	taskData := agent.TaskData{
		StringValues: map[string]string{
			CFG_TEST_TYPE: CFG_TYPE_ALL,
			CFG_SERVER_HOST: "nyc.speedtest.sbcglobal.net:8080",
			CFG_SERVER_ID: "5029",
		},
		IntValues: map[string]int {
			CFG_SERVER_ID: 5029,
			CFG_TIME_OUT: 5,
		},
		FloatValues: map[string]float64 {CFG_MAX_SECONDS: 6},
		IntSlices: map[string][]int{
			CFG_DOWNLOAD_SIZES: {245388, 505544},
			CFG_UPLOAD_SIZES:   {32768, 65536},
		},
	}

	spdTestRunner := SpeedTestRunner{}

	spTestResults, err := spdTestRunner.Run(taskData)
	if err != nil {
		t.Fatalf("Unexpected Error: \n%s", err.Error())
	}

	results := spTestResults.Latency.Seconds()

	if results <= 0 {
		t.Errorf("Error: Expected a positive Latency result, but got: %f", results)
	} else {
		fmt.Printf("\nLatency test results for server %d ... %f\n", taskData.IntValues[CFG_SERVER_ID], results)
	}

	results = spTestResults.Download
	if results <= 0 {
		t.Errorf("Error: Expected a positive Download result, but got: %f", results)
	} else {
		fmt.Printf("\nDownload test results for server %d ... %f\n", taskData.IntValues[CFG_SERVER_ID], results)
	}

	results = spTestResults.Upload
	if results <= 0 {
		t.Errorf("Error: Expected a positive Upload result, but got: %f", results)
	} else {
		fmt.Printf("\nUpload test results for server %d ... %f\n", taskData.IntValues[CFG_SERVER_ID], results)
	}
}