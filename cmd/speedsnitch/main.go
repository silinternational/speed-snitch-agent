package main

import (
	"fmt"
	"github.com/silinternational/speed-snitch-agent"
	"github.com/silinternational/speed-snitch-agent/lib/adminapi"
	"github.com/silinternational/speed-snitch-agent/lib/logqueue"
	"github.com/silinternational/speed-snitch-agent/lib/selfupdate"
	"github.com/silinternational/speed-snitch-agent/lib/tasks"
	"gopkg.in/robfig/cron.v2"
	"os"
	"runtime"
	"time"
)

var apiConfig agent.APIConfig
var agentStartTime time.Time

func main() {
	agentStartTime = time.Now()
	if len(os.Args) < 3 {
		fmt.Println("You must provide the Admin API BaseURL as the first argument and API Key as second argument")
		os.Exit(1)
	}

	apiConfig.BaseURL = os.Args[1]
	apiConfig.APIKey = os.Args[2]

	customApiConfig := agent.GetAppConfig(nil)
	if customApiConfig.BaseURL != "" && customApiConfig.APIKey != "" {
		apiConfig = customApiConfig

		fmt.Println("Using Custom ApiConfig with BaseURL: %s", apiConfig.BaseURL)
	}

	config, err := adminapi.GetConfig(apiConfig)
	if err != nil {
		fmt.Println("Unable to fetch config from admin API:", err)
		os.Exit(1)
	}

	newLogs := make(chan agent.TaskLogEntry, 10000)

	go logqueue.Manager(apiConfig, newLogs)

	taskCron := cron.New()
	tasks.UpdateTasks(config.Tasks, taskCron, newLogs)
	taskCron.Start()

	sysCron := cron.New()

	sysCron.AddFunc( // Say Hello 15 seconds past every minute
		"15 * * * * *",
		func() {
			adminapi.SayHello(apiConfig, agentStartTime)

			fmt.Println("Just ran Say Hello with version " + agent.Version)
		},
	)

	sysCron.AddFunc( // Get Config every 2 minutes
		"*/2 * * * *",
		func() {
			config, err := adminapi.GetConfig(apiConfig)
			if err != nil {
				fmt.Printf("\nError getting config from %s\n\t%s", apiConfig.BaseURL, err.Error())
				return
			}
			fmt.Println("Just ran GetConfig with version " + agent.Version)

			tasks.UpdateTasks(config.Tasks, taskCron, newLogs)

			wasNeeded, err := selfupdate.UpdateIfNeeded(
				agent.Version,
				config.Version.Number,
				config.Version.URL,
				true,
			)

			if err != nil {
				fmt.Println("Got error trying to self update ...\n\t" + err.Error())
			} else if wasNeeded {
				wd, _ := os.Getwd()
				fmt.Println("Self update was needed. Current working directory: " + wd)
			}
		},
	)

	sysCron.Start()

	runtime.Goexit()
}
