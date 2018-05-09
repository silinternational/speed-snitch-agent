package logqueue

import (
	"fmt"
	"github.com/silinternational/speed-snitch-agent"
	"github.com/silinternational/speed-snitch-agent/lib/adminapi"
	"os"
)

const FlushLogQueue = "flushLogQueue"

type TestTracker struct {
	KeepTrack    bool
	ReportedLogs []string
}

func Manager(apiConfig agent.APIConfig, newLogs chan agent.TaskLogEntry) {
	logQueue := []agent.TaskLogEntry{}

	for newLog := range newLogs {

		logQueue = append(logQueue, newLog)
		retryQueue := []agent.TaskLogEntry{}

		for _, nextLog := range logQueue {
			err := adminapi.Log(apiConfig, nextLog)
			if err != nil {
				retryQueue = append(retryQueue, newLog)
				fmt.Fprint(os.Stderr, err.Error())
			}
		}
		logQueue = retryQueue
	}
}
