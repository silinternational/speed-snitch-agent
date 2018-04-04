package main

import (
	"os"
	"github.com/silinternational/speed-snitch-agent"
	"github.com/silinternational/speed-snitch-agent/lib/speedtestnet"
	"fmt"
	"gopkg.in/robfig/cron.v2"
	"github.com/silinternational/speed-snitch-agent/lib/tasks"
	"runtime"
	"github.com/silinternational/speed-snitch-agent/lib/logqueue"
)


type FakeLogger struct {
}

func (f FakeLogger) Process(fakeLogKey, logText string, c ...interface{}) error {
	print("\n\tProcessing log: " + logText)
	return nil
}

func main() {
	println("Function Tests - main.main")

	testLogger := FakeLogger{}

	newLogs := make(chan string, 10000)

	go logqueue.Manager(newLogs, "fakeLogKey", &agent.LoggerInstance{testLogger})

	logEntriesKey := os.Getenv("LOGENTRIES_KEY")

	if logEntriesKey == "" {
		fmt.Errorf("No LOGENTRIES_KEY env variable available.")
		return
	}

	config := getConfig()

	mainCron := cron.New()

	tasks.UpdateTasks(config.Tasks, mainCron, newLogs)
	mainCron.Start()

	runtime.Goexit()
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
				Type:     agent.TypePing,
				Schedule: "*/10 * * * * *", // ping every 10 seconds
				Data: agent.TaskData{
					StringValues: map[string]string{
						speedtestnet.CFG_TEST_TYPE: speedtestnet.CFG_TYPE_LATENCY,
					},
					IntValues: map[string]int{
						speedtestnet.CFG_SERVER_ID: 5029,
						speedtestnet.CFG_TIME_OUT:  5,
					},
					FloatValues: map[string]float64{speedtestnet.CFG_MAX_SECONDS: 6},
				},

				SpeedTestRunner: agent.SpeedTestInstance{speedtestnet.SpeedTestRunner{}},
			},
			{
				Type:     logqueue.FlushLogQueue,
				Schedule: "*/27 * * * * *",
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