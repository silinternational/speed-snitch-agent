package speedtestnet

import agent "github.com/silinternational/speed-snitch-agent"

var Task agent.Task

func NewClient() (*agent.Task, error) {
	task := agent.Task{
		Type:     agent.TypeSpeedtest,
		Schedule: "",
		Data: agent.TaskData{
			"URL": "https://speedtest.net",
		},
	}

	return &task, nil
}
