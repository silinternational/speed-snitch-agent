package main

import (
	"github.com/silinternational/speed-snitch-agent"
	"gopkg.in/robfig/cron.v2"
	"github.com/silinternational/speed-snitch-agent/lib/tasks"
	"runtime"
	"github.com/silinternational/speed-snitch-agent/lib/logqueue"
	"github.com/silinternational/speed-snitch-agent/lib/adminapi"
	"fmt"
	"os"
	"time"
	"github.com/silinternational/speed-snitch-agent/lib/logentries"
)


type FakeLogger struct {
}

func (f FakeLogger) Process(fakeLogKey, logText string, c ...interface{}) error {
	print("\n\tProcessing log: " + logText)
	return nil
}

func main() {
	println("Function Tests - main.main")

	startTime := time.Now()
	newLogs := make(chan string, 10000)

	logEntriesKey := os.Getenv("LOGENTRIES_KEY")

	if logEntriesKey == "" {
		println("\n\n *** Just logging to console, since no LOGENTRIES_KEY env variable is available.")
		testLogger := FakeLogger{}
		go logqueue.Manager(newLogs, "fakeLogKey", &agent.LoggerInstance{testLogger})
	} else {
		logger := agent.LoggerInstance{logentries.Logger{}}
		go logqueue.Manager(newLogs, logEntriesKey, &agent.LoggerInstance{logger})
	}


	baseURL := "http://fillup.proxy.beeceptor.com"
	config, err := adminapi.GetConfig(baseURL)

	if err != nil {
		fmt.Printf("\nError getting config from %s\n\t%s", baseURL, err.Error())
		os.Exit(1)
	}

	testCron := cron.New()

	tasks.UpdateTasks(config.Tasks, testCron, newLogs)
	testCron.Start()


	sysCron := cron.New()

	sysCron.AddFunc( // Say Hello every 15 seconds
		"*/15 * * * * *",
		func() {
			adminapi.SayHello(config, startTime)
		},
	)

	sysCron.AddFunc( // Get Config every 3 minutes
		"*/3 * * * *",
		func() {
			config, err := adminapi.GetConfig(baseURL)
			if err != nil {
				fmt.Printf("\nError getting config from %s\n\t%s", baseURL, err.Error())
				return
			}
			tasks.UpdateTasks(config.Tasks, testCron, newLogs)
		},
	)

	sysCron.AddFunc( // Flush log every minute
		"* * * * *",
		func() {
			newLogs <- logqueue.FlushLogQueue
		},
	)

	sysCron.Start()

	runtime.Goexit()


}

func clearCron(thisCron *cron.Cron) {
	entries := thisCron.Entries()

	for _, nextEntry := range entries {
		thisCron.Remove(nextEntry.ID)
	}

}