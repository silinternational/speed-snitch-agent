package logqueue

import (
	"github.com/silinternational/speed-snitch-agent"
)

const FlushLogQueue = "flushLogQueue"

type TestTracker struct {
	KeepTrack bool
	ReportedLogs []string
}


// Manager listens to the newLogs channel and stores them as they come in.
// When a new log comes through with flushLogQueue as its value, then
//  it processes all the logs in its store and removes them.
func Manager(
	newLogs chan string,
	logServiceKey string,
	logger *agent.LoggerInstance,
) {
	logQueue := []string{}

	for newLog := range newLogs {

		// If it's not time to flush the queue, append the new log to the queue
		if newLog != FlushLogQueue {
			logQueue = append(logQueue, newLog)
			continue
		}

		// Otherwise, flush the queue
		for _, nextLog := range logQueue {
			_ = logger.Process(logServiceKey, nextLog)
		}
		logQueue = []string{}
	}
}
