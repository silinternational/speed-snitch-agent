package main

import (
	"fmt"
	"github.com/silinternational/speed-snitch-agent"
	"github.com/silinternational/speed-snitch-agent/lib/adminapi"
	"github.com/silinternational/speed-snitch-agent/lib/icmp"
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
var networkStatus string
var networkOfflineStartTime time.Time

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

	// Log that the node just restarted
	logEntry := agent.GetTaskLogEntry(agent.TypeRestarted)
	newLogs <- logEntry

	taskCron := cron.New()
	tasks.UpdateTasks(config.Tasks, taskCron, newLogs, &networkStatus)
	taskCron.Start()

	sysCron := cron.New()
	helloSchedule := agent.GetRandomSecondAsString() + " * * * * *"
	fmt.Println("Hello schedule:", helloSchedule)
	sysCron.AddFunc( // Say Hello every minute
		helloSchedule,
		func() {
			if networkStatus == agent.NetworkOffline {
				return
			}

			adminapi.SayHello(apiConfig, agentStartTime)
			now := time.Now()
			fmt.Println(now.Format(time.RFC3339), "Just ran Say Hello with version "+agent.Version)
		},
	)

	getConfigSchedule := agent.GetRandomSecondAsString() + " */2 * * * *"
	fmt.Println("Get Config schedule:", getConfigSchedule)
	sysCron.AddFunc( // Get Config every 2 minutes
		getConfigSchedule,
		func() {
			if networkStatus == agent.NetworkOffline {
				return
			}

			now := time.Now()
			config, err := adminapi.GetConfig(apiConfig)
			if err != nil {
				fmt.Printf("\n%s Error getting config from %s\n\t%s\n", now.Format(time.RFC3339), apiConfig.BaseURL, err.Error())
				return
			}
			fmt.Println(now.Format(time.RFC3339), "Just ran GetConfig with version "+agent.Version)

			tasks.UpdateTasks(config.Tasks, taskCron, newLogs, &networkStatus)

			wasNeeded, err := selfupdate.UpdateIfNeeded(
				agent.Version,
				config.Version.Number,
				config.Version.URL,
				true,
			)

			if err != nil {
				fmt.Println(now.Format(time.RFC3339), "Got error trying to self update ...\n\t"+err.Error())
			} else if wasNeeded {
				wd, _ := os.Getwd()
				fmt.Println(now.Format(time.RFC3339), "Self update was needed. Current working directory: "+wd)
			}
		},
	)

	checkNetworkStatusSchedule := agent.GetRandomSecondAsString() + " */5 * * * *"
	fmt.Println("Check network status schedule:", checkNetworkStatusSchedule)
	sysCron.AddFunc(
		checkNetworkStatusSchedule,
		func() {
			fmt.Printf("\nChecking network status...")

			// use the pre-Ping time for networkOfflineStartTime
			tempStartTime := time.Now().UTC()
			
			_, err := icmp.Ping("google.com", 2, 1, 30)
			if err != nil {
				// appears to be offline, change status and start tracking if needed
				if networkStatus != agent.NetworkOffline {
					networkStatus = agent.NetworkOffline
					networkOfflineStartTime = tempStartTime
				}
				fmt.Printf("offline. Started at %s. Down for %v seconds",
					networkOfflineStartTime.Format(time.RFC3339),
					int64(time.Since(networkOfflineStartTime)/time.Second))
			} else if networkStatus == agent.NetworkOffline {
				// reset status
				networkStatus = agent.NetworkOnline

				// appears to be online but status is offline, calculate duration, log it, and reset status
				downtime := time.Since(networkOfflineStartTime)
				log := agent.GetTaskLogEntry("downtime")
				log.DowntimeStart = networkOfflineStartTime.Format(time.RFC3339)
				log.DowntimeSeconds = int64(downtime / time.Second)
				newLogs <- log

				fmt.Printf("online. Restored at %s, Down for %v seconds",
					networkOfflineStartTime.Format(time.RFC3339),
					int64(time.Since(networkOfflineStartTime)/time.Second))
			} else {
				// ensure status is online
				networkStatus = agent.NetworkOnline
				fmt.Printf("online.\n")
			}

		},
	)

	sysCron.Start()

	runtime.Goexit()
}
