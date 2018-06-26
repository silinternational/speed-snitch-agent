package tasks

import (
	"github.com/silinternational/speed-snitch-agent"
	"github.com/silinternational/speed-snitch-agent/lib/icmp"
	"github.com/silinternational/speed-snitch-agent/lib/speedtestnet"
	"gopkg.in/robfig/cron.v2"
	"strings"
)

func clearCron(mainCron *cron.Cron) {
	entries := mainCron.Entries()

	for _, nextEntry := range entries {
		mainCron.Remove(nextEntry.ID)
	}

}

func logError(
	errorCode, errorBeginMsg string,
	err error,
	newLogs chan agent.TaskLogEntry,
) agent.TaskLogEntry {

	logEntry := agent.GetTaskLogEntry(agent.TypeError)
	logEntry.ErrorCode = errorCode
	logEntry.ErrorMessage = errorBeginMsg + err.Error()
	newLogs <- logEntry

	return logEntry
}

func UpdateTasks(
	tasks []agent.Task,
	mainCron *cron.Cron,
	newLogs chan agent.TaskLogEntry,
	networkStatus *string,
) {
	clearCron(mainCron)

	for index, _ := range tasks {
		task := tasks[index] // Have to do it this way, in order to use it in the closures
		switch task.Type {
		case agent.TypePing:
			mainCron.AddFunc(
				getCronScheduleWithRandomSeconds(task.Schedule),
				func() {
					if *networkStatus == agent.NetworkOffline {
						return
					}

					spTestResults, err := icmp.Ping(task.NamedServer.ServerHost, 0, 0, 0)
					if err != nil {
						logEntry := agent.GetTaskLogEntry(agent.TypeError)
						logEntry.ErrorCode = "1525283932"
						logEntry.ErrorMessage = "Error running latency test: " + err.Error()
						newLogs <- logEntry
					} else {
						logEntry := agent.GetTaskLogEntry(agent.TypePing)
						logEntry.Latency = spTestResults.Latency.Seconds() * 1000
						logEntry.PacketLossPercent = spTestResults.PacketLossPercent
						logEntry.ServerCountry = task.NamedServer.Country.Code
						logEntry.ServerID = task.NamedServer.SpeedTestNetServerID
						newLogs <- logEntry
					}
				},
			)
		case agent.TypeSpeedTest:
			mainCron.AddFunc(
				getCronScheduleWithRandomSeconds(task.Schedule),
				func() {
					if *networkStatus == agent.NetworkOffline {
						return
					}

					spdTestRunner := speedtestnet.SpeedTestRunner{}
					spTestResults, err := spdTestRunner.Run(task.Data)
					if err != nil {
						logError("1525291938", "Error running speed test: ", err, newLogs)
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
		case agent.TypeReboot:
			mainCron.AddFunc(
				task.Schedule,
				func() {
					err := agent.Reboot()
					if err != nil {
						logEntry := logError("1529675400", "Error rebooting: ", err,  newLogs)
						fmt.Fprint(os.Stdout, logEntry)
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
