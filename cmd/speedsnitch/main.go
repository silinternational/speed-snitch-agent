package main

import (
	agent "github.com/silinternational/speed-snitch-agent"
	"github.com/silinternational/speed-snitch-agent/lib/speedtestnet"
)

var config agent.Config

func main() {

	config = getConfig()
}

func getConfig() agent.Config {

	return agent.Config{
		Version: struct {
			Latest string
			URL    string
		}{
			Latest: "1.0.0",
			URL:    "https://github.com/silinternational/speed-snitch-agent/raw/1.0.0/dist",
		},
		Tasks: []agent.Task{
			{
				Type:     agent.TypeSpeedTest,
				Schedule: "",
				Data: agent.TaskData{
					StringValues: map[string]string{
						speedtestnet.CFG_TEST_TYPE: speedtestnet.CFG_TYPE_LATENCY,
					},
					IntValues: map[string]int{
						speedtestnet.CFG_SERVER_ID: 5029,
						speedtestnet.CFG_TIME_OUT:  5,
					},
					FloatValues: map[string]float64{speedtestnet.CFG_MAX_SECONDS: 6},
					IntSlices: map[string][]int{
						speedtestnet.CFG_DOWNLOAD_SIZES: {245388, 505544},
						speedtestnet.CFG_UPLOAD_SIZES:   {32768, 65536},
					},
				},

				SpeedTestRunner: agent.SpeedTestInstance{speedtestnet.SpeedTestRunner{}},
			},
		},
		Log: struct {
			Format      string
			Destination string
		}{
			Format:      "LogTypeFile",
			Destination: "/var/log/speed-snitch",
		},
	}
}
