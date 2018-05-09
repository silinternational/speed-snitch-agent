package tasks

import (
	"fmt"
	"github.com/silinternational/speed-snitch-agent"
	"github.com/silinternational/speed-snitch-agent/lib/speedtestnet"
	"gopkg.in/robfig/cron.v2"
	"os"
)

func clearCron(mainCron *cron.Cron) {
	entries := mainCron.Entries()

	for _, nextEntry := range entries {
		mainCron.Remove(nextEntry.ID)
	}

}

func UpdateTasks(
	tasks []agent.Task,
	mainCron *cron.Cron,
	newLogs chan agent.TaskLogEntry,
) {
	clearCron(mainCron)

	for index, _ := range tasks {
		task := tasks[index] // Have to do it this way, in order to use it in the closures
		switch task.Type {
		case agent.TypePing:
			mainCron.AddFunc(
				task.Schedule,
				func() {

					spdTestRunner := speedtestnet.SpeedTestRunner{}
					spTestResults, err := spdTestRunner.Run(task.Data)
					if err != nil {
						logEntry := agent.GetTaskLogEntry(agent.TypeError)
						logEntry.ErrorCode = "1525283932"
						logEntry.ErrorMessage = "Error running latency test: " + err.Error()
						newLogs <- logEntry
						fmt.Fprint(os.Stdout, logEntry)
					} else {
						logEntry := agent.GetTaskLogEntry(agent.TypePing)
						logEntry.Latency = spTestResults.Latency.Seconds() * 1000
						logEntry.ServerID = task.Data.IntValues["serverID"]
						newLogs <- logEntry
						fmt.Fprint(os.Stdout, logEntry)
					}
				},
			)
		case agent.TypeSpeedTest:
			mainCron.AddFunc(
				task.Schedule,
				func() {

					spdTestRunner := speedtestnet.SpeedTestRunner{}
					spTestResults, err := spdTestRunner.Run(task.Data)
					if err != nil {
						logEntry := agent.GetTaskLogEntry(agent.TypeError)
						logEntry.ErrorCode = "1525291938"
						logEntry.ErrorMessage = "Error running speed test: " + err.Error()
						newLogs <- logEntry
					} else {
						logEntry := agent.GetTaskLogEntry(agent.TypeSpeedTest)
						logEntry.Download = spTestResults.Download
						logEntry.Upload = spTestResults.Upload
						logEntry.ServerID = task.Data.IntValues["serverID"]
						newLogs <- logEntry
					}
				},
			)
		}
	}

}
