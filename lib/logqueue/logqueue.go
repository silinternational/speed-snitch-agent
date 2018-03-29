package logqueue

const FlushLogQueue = "flushLogQueue"

type TestTracker struct {
	KeepTrack bool
	ReportedLogs []string
}

func appendMapValue(logType, newValue string, wholeMap map[string][]string) {

	_, ok := wholeMap[logType]

	if ok {
		wholeMap[logType] = append(wholeMap[logType], newValue)
	} else {
		wholeMap[logType] = []string{newValue}
	}
}

// Stasher listens to the newLogs channel and stores them as they come in
//  based on LogType: LogValue entries.
// When a LogType comes through with flushLogQueue as its value, then
//  it pushes the log set for that LogType to the reporter and also
//  removes its log set from its store of logs.
func Stasher(newLogs chan [2]string, completedLogs chan []string) {
	logStore := map[string][]string{}

	for newLog := range newLogs {
		logType := newLog[0]
		logValue := newLog[1]

		// Check if it's time to flush the queue
		if logValue == FlushLogQueue {
			logSet, ok := logStore[logType]

			// Send the queued logs to the Reporter and remove them from the queue
			if ok {
				completedLogs <- logSet
				logStore[logType] = []string{}
			}
			continue
		}

		appendMapValue(logType, logValue, logStore)

	}
}

// Reporter takes log sets from the completedLogs channel and actually logs them
func Reporter(completedLogs chan []string, tracker *TestTracker) {
	for nextLogSet := range completedLogs {
		for _, nextLog := range nextLogSet {

			// TODO Use a real logger

			if tracker.KeepTrack {
				tracker.ReportedLogs = append(tracker.ReportedLogs, nextLog)
			}
		}
	}
}