package tasks

import (
	"fmt"
	"github.com/silinternational/speed-snitch-agent"
	"github.com/silinternational/speed-snitch-agent/lib/speedtestnet"
	"gopkg.in/robfig/cron.v2"
	"os"
	"strings"
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
				getCronScheduleWithRandomSeconds(task.Schedule),
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
						logEntry.ServerCountry = task.NamedServer.Country.Code
						logEntry.ServerID = task.NamedServer.SpeedTestNetServerID
						newLogs <- logEntry
						fmt.Fprint(os.Stdout, logEntry)
					}
				},
			)
		case agent.TypeSpeedTest:
			mainCron.AddFunc(
				getCronScheduleWithRandomSeconds(task.Schedule),
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
						logEntry.ServerCountry = task.NamedServer.Country.Code
						logEntry.ServerID = task.NamedServer.SpeedTestNetServerID
						newLogs <- logEntry
					}
				},
			)
		}
	}

}

// getCronScheduleWithRandomSeconds takes cron schedule string and adds or replaces the first element as needed
// with a random number between 0 and 60
func getCronScheduleWithRandomSeconds(schedule string) string {
	parts := strings.Split(schedule, " ")
	withSec := ""
	if len(parts) == 6 {
		parts[0] = agent.GetRandomSecondAsString()
		withSec = strings.Join(parts, " ")
	} else {
		withSec = agent.GetRandomSecondAsString() + " " + schedule
	}

	return withSec
}
