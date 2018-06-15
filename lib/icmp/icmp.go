package icmp

import (
	"fmt"
	"github.com/silinternational/speed-snitch-agent"
	"github.com/sparrc/go-ping"
	"time"
)

const DefaultCount = 10
const DefaultIntervalSeconds = 1
const DefaultTimeoutSeconds = 30

func Ping(host string, count, interval, timeout int) (agent.SpeedTestResults, error) {
	if host == "" {
		return agent.SpeedTestResults{}, fmt.Errorf("host is required for ping")
	}

	if count == 0 {
		count = DefaultCount
	}
	if interval == 0 {
		interval = DefaultIntervalSeconds
	}
	if timeout == 0 {
		timeout = DefaultTimeoutSeconds
	}

	pinger, err := ping.NewPinger(host)
	if err != nil {
		return agent.SpeedTestResults{}, err
	}

	pinger.Count = count
	pinger.Interval = time.Duration(interval) * time.Second
	pinger.Timeout = time.Duration(timeout) * time.Second

	pinger.Run()
	stats := pinger.Statistics()

	return agent.SpeedTestResults{
		Timestamp: time.Now(),
		Latency:   stats.AvgRtt,
	}, nil
}
