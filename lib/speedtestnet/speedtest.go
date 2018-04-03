package speedtestnet

import (
	"fmt"
	agent "github.com/silinternational/speed-snitch-agent"
	"github.com/silinternational/speed-snitch-agent/lib/httputils"
	"io"
	"net"
	"sync"
	"time"
)

const CFG_SERVER_ID = "serverID"
const CFG_TIME_OUT = "timeOut"
const CFG_DOWNLOAD_SIZES = "downloadSizes"
const CFG_UPLOAD_SIZES = "uploadSizes"
const CFG_MAX_SECONDS = "maxSeconds"
const CFG_TEST_TYPE = "testType"
const CFG_TYPE_LATENCY = "latencyTest"
const CFG_TYPE_DOWNLOAD = "downloadTest"
const CFG_TYPE_UPLOAD = "uploadTest"
const CFG_TYPE_ALL = "allTests"
const MillisecondInt64 = int64(time.Millisecond)

var Task agent.Task

func NewClient() (*agent.Task, error) {
	task := agent.Task{
		Type:     agent.TypeSpeedTest,
		Schedule: "",
		Data: agent.TaskData{
			IntValues: map[string]int{
				CFG_SERVER_ID: 5029,
			},
		},
	}

	return &task, nil
}

func tempLog(text string) {
	println(text)
}

type configuration struct {
	ServerID      int     `json:"ServerID"` // This must be present as an ID value in the serverList.go servers
	Timeout       int     `json:"Timeout"`
	DownloadSizes []int   `json:"DownloadSizes"`
	UploadSizes   []int   `json:"UploadSizes"`
	MaxSeconds    float64 `json:"MaxSeconds"` // After this number of seconds, no new downloads/uploads will be attempted
}

type server struct {
	CC            string
	Country       string
	ID            int
	Latitude      float64
	Longitude     float64
	Name          string
	Sponsor       string
	URL           string
	URL2          string
	Host          string
	Distance      float64
	Speedtest     *speedtest
	Results       agent.SpeedTestResults
	Configuration *configuration
	TCPAddr       *net.TCPAddr
}

// Set the TCPAddr based on the results of calling net.ResolveTCPAddr with the Host
func (s *server) SetTCPAddr() {
	addr, err := net.ResolveTCPAddr("tcp", s.Host)
	if err != nil {
		tempLog(fmt.Sprintf("%s\n", err.Error()))
	}
	s.TCPAddr = addr
}

// Goroutine for downloading data
// the length param refers to the limit to how many seconds the download is allowed to take
func (s *server) Downloader(ci chan int, co chan []int, wg *sync.WaitGroup, start time.Time, length float64) {
	defer wg.Done()
	s.SetTCPAddr()

	conn, err := httputils.DialTimeout("tcp", s.Speedtest.Source, s.TCPAddr, s.Configuration.Timeout)
	if err != nil {
		tempLog(fmt.Sprintf("\nCannot connect to %s\n%s", s.TCPAddr.String(), err.Error()))
	}

	defer conn.Close()

	conn.Write([]byte("HI\n"))
	hello := make([]byte, 1024)
	conn.Read(hello)
	var ask int
	tmp := make([]byte, 1024)

	var out []int

	for size := range ci {
		print(".") // TODO improve on this
		remaining := size

		for remaining > 0 && time.Since(start).Seconds() < length {

			if remaining > 1000000 {
				ask = 1000000
			} else {
				ask = remaining
			}
			down := 0

			conn.Write([]byte(fmt.Sprintf("DOWNLOAD %d\n", ask)))

			for down < ask {
				n, err := conn.Read(tmp)
				if err != nil {
					if err != io.EOF {
						fmt.Printf("ERR: %v\n", err)
					}
					break
				}
				down += n
			}
			out = append(out, down)
			remaining -= down

		}
		print(".") // TODO improve on this
	}

	go func(co chan []int, out []int) {
		co <- out
	}(co, out)
}

// Goroutine for uploading data
// the length param refers to the limit to how many seconds the upload is allowed to take
func (s *server) Uploader(ci chan int, co chan []int, wg *sync.WaitGroup, start time.Time, length float64) {
	defer wg.Done()
	s.SetTCPAddr()

	conn, err := httputils.DialTimeout("tcp", s.Speedtest.Source, s.TCPAddr, s.Configuration.Timeout)
	if err != nil {
		tempLog(fmt.Sprintf("\nCannot connect to %s\n", s.TCPAddr.String()))
	}

	defer conn.Close()

	conn.Write([]byte("HI\n"))
	hello := make([]byte, 1024)
	conn.Read(hello)

	var give int
	var out []int
	for size := range ci {
		print(".") // TODO improve on this
		remaining := size

		for remaining > 0 && time.Since(start).Seconds() < length {
			if remaining > 100000 {
				give = 100000
			} else {
				give = remaining
			}
			header := []byte(fmt.Sprintf("UPLOAD %d 0\n", give))
			data := make([]byte, give-len(header))

			conn.Write(header)
			conn.Write(data)
			up := make([]byte, 24)
			conn.Read(up)

			out = append(out, give)
			remaining -= give
		}
		print(".") // TODO improve on this
	}

	go func(co chan []int, out []int) {
		co <- out
	}(co, out)
}

type speedtest struct {
	Server  *server
	Source  *net.TCPAddr
	Timeout time.Duration
}

// Tests the latency of the server in question and registers the value
// in server.Speedtest.Results.Latency
func LatencyTest(server *server) {
	server.SetTCPAddr()

	conn, err := httputils.DialTimeout(
		"tcp",
		server.Speedtest.Source,
		server.TCPAddr,
		server.Configuration.Timeout,
	)
	if err != nil {
		server.Results.Error = fmt.Sprintf(
			"Error running latency test for server %s.\n%s\n",
			server.Name,
			err.Error(),
		)
		return
	}

	defer conn.Close()

	conn.Write([]byte("HI\n"))
	hello := make([]byte, 1024)
	conn.Read(hello)

	sum := time.Duration(0)
	for j := 0; j < 3; j++ {
		resp := make([]byte, 1024)
		start := time.Now()
		conn.Write([]byte(fmt.Sprintf("PING %d\n", start.UnixNano()/MillisecondInt64)))
		conn.Read(resp)
		total := time.Since(start)
		sum += total
	}
	server.Results.Latency = sum / 3

}

// DownloadTest - controls Downloader goroutine
// @param server - a pointer to a server
// @param length - the maximum number of seconds this upload is allowed to take
// @param sizes - if this is left empty, then ten default sizes will be
//                used each bigger than the previous one
func DownloadTest(server *server, length float64, sizes []int) {
	if len(sizes) < 1 {
		sizes = []int{245388, 505544, 1118012, 1986284, 4468241, 7907740, 12407926, 17816816, 24262167, 31625365}
	}

	ci := make(chan int)
	co := make(chan []int)
	wg := new(sync.WaitGroup)
	start := time.Now()

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go server.Downloader(ci, co, wg, start, length)
	}

	for _, size := range sizes {
		for i := 0; i < 4; i++ {
			ci <- size
		}
	}

	close(ci)
	wg.Wait()

	downDuration := time.Since(start)

	var totalSize int
	for i := 0; i < 8; i++ {
		chunks := <-co
		for _, chunk := range chunks {
			totalSize += chunk
		}
	}

	downBits := float64(totalSize) * 8
	server.Results.Download = downBits / downDuration.Seconds() / 1000 / 1000 // Mb/sec
}

// Function that controls Uploader goroutine
// @param server - a pointer to a speedtest.Server
// @param length - the maximum number of seconds this upload is allowed to take
// @param sizes - if this is left empty, then seven default sizes will be
//                used each bigger than the previous one
func UploadTest(server *server, length float64, sizes []int) {
	if len(sizes) < 1 {
		sizes = []int{32768, 65536, 131072, 262144, 524288, 1048576, 7340032}
	}

	ci := make(chan int)
	co := make(chan []int)
	wg := new(sync.WaitGroup)
	start := time.Now()

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go server.Uploader(ci, co, wg, start, length)
	}

	var tmp int
	for _, size := range sizes {
		for i := 0; i < 4; i++ {
			tmp += size
			ci <- size
		}
	}
	close(ci)
	wg.Wait()

	upDuration := time.Since(start)

	var totalSize int
	for i := 0; i < 8; i++ {
		chunks := <-co
		for _, chunk := range chunks {
			totalSize += chunk
		}
	}

	upBits := float64(totalSize) * 8
	server.Results.Upload = upBits / upDuration.Seconds() / 1000 / 1000
}

func getTestType(taskData agent.TaskData) (string, error) {

	var testType string
	var ok bool
	testType, ok = taskData.StringValues[CFG_TEST_TYPE]
	if !ok {
		return "", fmt.Errorf("taskData.StringValues is missing an entry for %s", CFG_TEST_TYPE)
	}

	validTypes := []string{
		CFG_TYPE_ALL,
		CFG_TYPE_DOWNLOAD,
		CFG_TYPE_LATENCY,
		CFG_TYPE_UPLOAD,
	}

	for _, nextValid := range validTypes {
		if testType == nextValid {
			return testType, nil
		}
	}

	return "", fmt.Errorf("Invalid value in TaskData for %s: %s", CFG_TEST_TYPE, testType)
}

type SpeedTestRunner struct{}

// Run ensures all the needed agent.TaskData values are present and valid and
//   runs a Latency Test as well as any other requested test(s)
//
//   Here are the required values ...
//     taskData.IntValues:
//       - CFG_SERVER_ID ... must match one in the serverlist.go file
//       - CFG_TIME_OUT
//     taskData.StringValues:
//       - CFG_TEST_TYPE ...  CFG_TYPE_ALL | CFG_TYPE_DOWNLOAD | CFG_TYPE_LATENCY | CFG_TYPE_UPLOAD
//     taskData.FloatValues:
//       - CFG_MAX_SECONDS ... (not required for Latency Tests)
//     taskData.IntSlices:
//       - CFG_DOWNLOAD_SIZES ... (only required for Download Tests)
//       - CFG_UPLOAD_SIZES ... (only required for Upload Tests)
func (s SpeedTestRunner) Run(taskData agent.TaskData) (agent.SpeedTestResults, error) {
	emptyResults := agent.SpeedTestResults{}
	var ok bool

	testConfig := configuration{}

	// Get the ID of the speedtestnet server
	testConfig.ServerID, ok = taskData.IntValues[CFG_SERVER_ID]

	if !ok {
		return emptyResults, fmt.Errorf("taskData.IntValues is missing an entry for %s", CFG_SERVER_ID)
	}

	// Get the configuration for the speedtestnet server
	emptyServer := server{}
	testServer := GetServerByID(testConfig.ServerID)
	if testServer == emptyServer {
		return emptyResults, fmt.Errorf("Could not find speedtestnet server with ID %d", testConfig.ServerID)
	}
	testServer.Configuration = &testConfig

	// Get the requested test type (Latency, Download, Upload, All)
	// Note that a latency test is performed in all cases
	var testType string
	var err error
	testType, err = getTestType(taskData)

	if err != nil {
		return emptyResults, err
	}

	// Get the desired Time-out value for the test
	timeOut, ok := taskData.IntValues[CFG_TIME_OUT]
	if !ok {
		return emptyResults, fmt.Errorf("taskData.IntValues is missing an entry for %s", CFG_TIME_OUT)
	}
	testConfig.Timeout = timeOut

	// Set the Source value for the speedtest
	localAddr := net.TCPAddr{}
	source, _ := net.ResolveTCPAddr("tcp", localAddr.String())

	spdTest := speedtest{
		Source: source,
	}

	// Mutually connect the speedtest and the test server
	testServer.Speedtest = &spdTest
	spdTest.Server = &testServer

	// Run Latency Test in all cases
	LatencyTest(&testServer)
	if testType == CFG_TYPE_LATENCY {
		return testServer.Results, nil
	}

	// Get the desired MaxSeconds value
	maxSeconds, ok := taskData.FloatValues[CFG_MAX_SECONDS]
	if !ok {
		return emptyResults, fmt.Errorf("taskData.IntValues is missing an entry for %s", CFG_MAX_SECONDS)
	}
	testConfig.MaxSeconds = maxSeconds

	// If requested, get the Download Sizes and run the Download test
	if testType == CFG_TYPE_DOWNLOAD || testType == CFG_TYPE_ALL {
		downloadSizes, ok := taskData.IntSlices[CFG_DOWNLOAD_SIZES]
		if !ok {
			return emptyResults, fmt.Errorf("taskData.IntSlices is missing an entry for %s", CFG_DOWNLOAD_SIZES)
		}
		testConfig.DownloadSizes = downloadSizes
		DownloadTest(&testServer, testConfig.MaxSeconds, testConfig.DownloadSizes)
	}

	// If requested, get the Upload Sizes and run the Upload test
	if testType == CFG_TYPE_UPLOAD || testType == CFG_TYPE_ALL {
		uploadSizes, ok := taskData.IntSlices[CFG_UPLOAD_SIZES]
		if !ok {
			return emptyResults, fmt.Errorf("taskData.IntSlices is missing an entry for %s", CFG_UPLOAD_SIZES)
		}
		testConfig.UploadSizes = uploadSizes
		UploadTest(&testServer, testConfig.MaxSeconds, testConfig.UploadSizes)
	}

	return testServer.Results, nil
}
