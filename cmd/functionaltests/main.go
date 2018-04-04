package main

import (
	"github.com/silinternational/speed-snitch-agent"
	"github.com/silinternational/speed-snitch-agent/lib/speedtestnet"
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
				Schedule: "10,20,30,50 * * * * *", // ping every 10 seconds
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
				Type:     agent.TypeSpeedTest,
				Schedule: "40 * * * * *", // ping every 10 seconds
				Data: agent.TaskData{
					StringValues: map[string]string{
						speedtestnet.CFG_TEST_TYPE: speedtestnet.CFG_TYPE_ALL,
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
			{
				Type:     logqueue.FlushLogQueue,
				Schedule: "55 * * * * *",
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