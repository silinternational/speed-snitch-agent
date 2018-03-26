package main

import agent "github.com/silinternational/speed-snitch-agent"

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
				Type:     agent.TypeSpeedtest,
				Schedule: "",
				Data: agent.TaskData{
					"URL": "something",
				},
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
