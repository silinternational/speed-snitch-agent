package tasks

import (
	"github.com/silinternational/speed-snitch-agent"
	"gopkg.in/robfig/cron.v2"
	"github.com/silinternational/speed-snitch-agent/lib/speedtestnet"
	"fmt"
)

func clearCron(mainCron *cron.Cron) {
	entries := mainCron.Entries()

	for _, nextEntry := range entries {
		mainCron.Remove(nextEntry.ID)
	}

}

func UpdateTasks(
	tasks []agent.Task,
	taskCron *cron.Cron,
	newLogs chan string,
) {
	clearCron(taskCron)

	for index, _ := range tasks {
		task := tasks[index] // Have to do it this way, in order to use it in the closures
		switch task.Type {
		case agent.TypePing:
			taskCron.AddFunc(
				task.Schedule,
				func() {

					spdTestRunner := speedtestnet.SpeedTestRunner{}
					spTestResults, err := spdTestRunner.Run(task.Data)
					if err != nil {
						newLogs <- "Error running latency test: " + err.Error()
					} else {
						newLogs <- fmt.Sprintf("Latency Results: %f milliseconds", spTestResults.Latency.Seconds()*1000)
					}
				},
			)
		case agent.TypeSpeedTest:
			taskCron.AddFunc(
				task.Schedule,
				func() {

					spdTestRunner := speedtestnet.SpeedTestRunner{}
					spTestResults, err := spdTestRunner.Run(task.Data)
					if err != nil {
						newLogs <- "Error running speed test: " + err.Error()
					} else {
						newLogs <- fmt.Sprintf("Latency Results: %f milliseconds", spTestResults.Latency.Seconds()*1000)
						newLogs <- fmt.Sprintf("Download Results: %f Mb/sec", spTestResults.Download)
						newLogs <- fmt.Sprintf("Upload Results: %f Mb/sec", spTestResults.Upload)
					}
				},
			)
		}
	}

}