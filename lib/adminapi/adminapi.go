package adminapi

import (
	"encoding/json"
	"fmt"
	"github.com/silinternational/speed-snitch-agent"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"
)

type Hello struct {
	ID      string `json:"id"`
	Version string `json:"version"`
	Uptime  string `json:"uptime"`
	OS      string `json:"os"`
	Arch    string `json:"arch"`
}

// CallAPI creates a http.Request object, attaches headers to it and makes the
// requested api call.
func CallAPI(method, url, postData string, headers map[string]string) (*http.Response, error) {
	var err error
	var req *http.Request

	if postData != "" {
		req, err = http.NewRequest(method, url, strings.NewReader(postData))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode >= 300 {
		return resp, fmt.Errorf(
			"API returned an error. \n\tMethod: %s, \n\tURL: %s, \n\tCode: %v, \n\tStatus: %s \n\tBody: %s",
			method, url, resp.StatusCode, resp.Status, postData)
	}

	return resp, nil
}

// SayHello makes a POST call to /hello with id, version, and update
func SayHello(config agent.Config, agentStartTime time.Time) (bool, error) {
	helloBody := Hello{
		ID:      agent.GetMacAddr(),
		Version: agent.Version,
		Uptime:  string(time.Since(agentStartTime) / 1000000 / 60),
		OS:      runtime.GOOS,
		Arch:    runtime.GOARCH,
	}
	helloJson, err := json.Marshal(helloBody)
	if err != nil {
		return false, fmt.Errorf("Unable to marshal json for /hello call")
	}

	resp, err := CallAPI("POST", config.BaseURL+"/hello", string(helloJson), map[string]string{})
	if err != nil {
		return false, err
	}

	if resp.StatusCode != 204 {
		return false, fmt.Errorf("Call to /hello did not return 204, got: %v", resp.StatusCode)
	}

	return true, nil
}

// GetConfig fetches config from
func GetConfig(baseURL string) (agent.Config, error) {
	url := fmt.Sprintf("%s/config/%s", baseURL, agent.GetMacAddr())

	resp, err := CallAPI("GET", url, "", map[string]string{})
	if err != nil {
		return agent.Config{}, err
	}

	if resp.StatusCode != 200 {
		return agent.Config{}, fmt.Errorf("Unable to get config, got status code %s", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return agent.Config{}, err
	}

	var config agent.Config
	err = json.Unmarshal(body, &config)
	if err != nil {
		return agent.Config{}, err
	}

	config.BaseURL = baseURL

	return config, nil
}

//func getDemoConfig() agent.Config {
//
//	return agent.Config{
//		Version: struct {
//			Number string
//			URL    string
//		}{
//			Number: "1.0.0",
//			URL:    "https://github.com/silinternational/speed-snitch-agent/raw/1.0.0/dist",
//		},
//		Tasks: []agent.Task{
//			{
//				Type:     agent.TypeSpeedTest,
//				Schedule: "",
//				Data: agent.TaskData{
//					StringValues: map[string]string{
//						speedtestnet.CFG_TEST_TYPE: speedtestnet.CFG_TYPE_LATENCY,
//					},
//					IntValues: map[string]int{
//						speedtestnet.CFG_SERVER_ID: 5029,
//						speedtestnet.CFG_TIME_OUT:  5,
//					},
//					FloatValues: map[string]float64{speedtestnet.CFG_MAX_SECONDS: 6},
//					IntSlices: map[string][]int{
//						speedtestnet.CFG_DOWNLOAD_SIZES: {245388, 505544},
//						speedtestnet.CFG_UPLOAD_SIZES:   {32768, 65536},
//					},
//				},
//
//				SpeedTestRunner: agent.SpeedTestInstance{speedtestnet.SpeedTestRunner{}},
//			},
//		},
//		Log: struct {
//			Format      string
//			Destination string
//		}{
//			Format:      "LogTypeFile",
//			Destination: "/var/log/speed-snitch",
//		},
//	}
//}
