package adminapi

import (
	"encoding/json"
	"fmt"
	"github.com/silinternational/speed-snitch-agent"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

type Hello struct {
	ID      string `json:"id"`
	Version string `json:"version"`
	Uptime  int64  `json:"uptime"`
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
func SayHello(apiConfig agent.APIConfig, agentStartTime time.Time) error {
	helloBody := Hello{
		ID:      agent.GetMacAddr(),
		Version: agent.Version,
		Uptime:  time.Since(agentStartTime).Nanoseconds() / 1000000,
		OS:      runtime.GOOS,
		Arch:    runtime.GOARCH,
	}
	helloJson, err := json.Marshal(helloBody)
	if err != nil {
		return fmt.Errorf("unable to marshal json for /hello call")
	}

	resp, err := CallAPI("POST", apiConfig.BaseURL+"/hello", string(helloJson), getCallApiHeaders(apiConfig))
	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("call to /hello did not return 204, got: %v", resp.StatusCode)
	}

	return nil
}

// GetConfig fetches config from admin api
func GetConfig(apiConfig agent.APIConfig) (agent.Config, error) {
	url := fmt.Sprintf("%s/config/%s", apiConfig.BaseURL, agent.GetMacAddr())

	resp, err := CallAPI("GET", url, "", getCallApiHeaders(apiConfig))
	if err != nil {
		return agent.Config{}, err
	}

	// If successful but empty response, agent is not yet configured so only set BaseURL on config and return
	if resp.StatusCode == 204 {
		return agent.Config{}, nil
	}

	if resp.StatusCode != 200 {
		return agent.Config{}, fmt.Errorf("unable to get config, got status code %d", resp.StatusCode)
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

	return config, nil
}

func getCallApiHeaders(apiConfig agent.APIConfig) map[string]string {
	return map[string]string{"x-api-key": apiConfig.APIKey}
}

// Process logs to the logentries online log service
func Log(apiConfig agent.APIConfig, logEntry agent.TaskLogEntry) error {

	// Write to stdout/stderr
	var localOutput *os.File
	if logEntry.EntryType == "error" {
		localOutput = os.Stderr
	} else {
		localOutput = os.Stdout
	}
	fmt.Fprintf(localOutput, "%+v\n", logEntry)

	logEntryJson, err := json.Marshal(logEntry)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		fmt.Fprintf(os.Stderr, "%+v\n", logEntry)
	}

	url := fmt.Sprintf("%s/log/%s/%s", apiConfig.BaseURL, agent.GetMacAddr(), logEntry.EntryType)
	resp, err := CallAPI("POST", url, string(logEntryJson), getCallApiHeaders(apiConfig))

	// If there is an error, then resp won't be usable below
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "log api call response: %s\n", resp.Status)
	return nil
}
