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
	"github.com/silinternational/speed-snitch-agent/lib/selfupdate"
	"os/exec"
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


	//baseURL := "http://fillup.proxy.beeceptor.com"
	baseURL := "http://demo7457258.mockable.io"
	config, err := adminapi.GetConfig(baseURL)

	if err != nil {
		fmt.Printf("\nError getting config from %s\n\t%s", baseURL, err.Error())
		os.Exit(1)
	}

	testCron := cron.New()

	tasks.UpdateTasks(config.Tasks, testCron, newLogs)
	testCron.Start()


	sysCron := cron.New()

	sysCron.AddFunc( // Say Hello 15 seconds past every minute
		"15 * * * * *",
		func() {
			adminapi.SayHello(config, startTime)

			newLogs <- "Just ran Say Hello with version " + agent.Version
		},
	)

	sysCron.AddFunc( // Get Config every 3 minutes
		"*/2 * * * *",
		func() {
			config, err := adminapi.GetConfig(baseURL)
			if err != nil {
				fmt.Printf("\nError getting config from %s\n\t%s", baseURL, err.Error())
				return
			}
			newLogs <- "Just ran GetConfig with version " + agent.Version

			tasks.UpdateTasks(config.Tasks, testCron, newLogs)

			wasNeeded, err := selfupdate.UpdateIfNeeded(agent.Version, config.Version.Number, config.Version.URL)
			if err != nil {
				newLogs <- "Got error trying to self update ...\n\t" + err.Error()
			} else if wasNeeded {
				wd, _ := os.Getwd()
				newLogs <- "Self update was needed. Current working directory: " + wd
				newLogs <- logqueue.FlushLogQueue
			}
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


// RebootRasPi checks if the GOOS is linux and the GOARCH is arm and
//  if so, executes the reboot command
func RebootRasPi() {
	goOS := runtime.GOOS
	goArch := runtime.GOARCH

	if goOS == "linux" && goArch == "arm"{
		_ = exec.Command("reboot").Run()
	}
}