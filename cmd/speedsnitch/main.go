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

		fmt.Printf("Using Custom ApiConfig with BaseURL: %s\n", apiConfig.BaseURL)
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
	helloSchedule := agent.GetRandomSecondAsString() + " * * * * *"
	fmt.Println("Hello schedule:", helloSchedule)
	sysCron.AddFunc( // Say Hello every minute
		helloSchedule,
		func() {
			adminapi.SayHello(apiConfig, agentStartTime)
			now := time.Now()
			fmt.Println(now.Format(time.RFC3339), "Just ran Say Hello with version " + agent.Version)
		},
	)

	getConfigSchedule := agent.GetRandomSecondAsString() + " */2 * * * *"
	fmt.Println("Get Config schedule:", getConfigSchedule)
	sysCron.AddFunc( // Get Config every 2 minutes
		getConfigSchedule,
		func() {
			now := time.Now()
			config, err := adminapi.GetConfig(apiConfig)
			if err != nil {
				fmt.Printf("\n%sError getting config from %s\n\t%s\n", now.Format(time.RFC3339), apiConfig.BaseURL, err.Error())
				return
			}
			fmt.Println(now.Format(time.RFC3339),"Just ran GetConfig with version " + agent.Version)

			tasks.UpdateTasks(config.Tasks, taskCron, newLogs)

			wasNeeded, err := selfupdate.UpdateIfNeeded(
				agent.Version,
				config.Version.Number,
				config.Version.URL,
				true,
			)

			if err != nil {
				fmt.Println(now.Format(time.RFC3339), "Got error trying to self update ...\n\t" + err.Error())
			} else if wasNeeded {
				wd, _ := os.Getwd()
				fmt.Println(now.Format(time.RFC3339), "Self update was needed. Current working directory: " + wd)
			}
		},
	)

	sysCron.Start()

	runtime.Goexit()
}
