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
	pinger.SetPrivileged(true)

	pinger.Run()
	stats := pinger.Statistics()

	if stats.PacketsRecv == 0 {
		return agent.SpeedTestResults{}, fmt.Errorf("zero ping packets received")
	} else if stats.AvgRtt == 0 {
		return agent.SpeedTestResults{}, fmt.Errorf("average RTT for ping reported as zero, something went wrong")
	}

	return agent.SpeedTestResults{
		Timestamp:         time.Now(),
		Latency:           stats.AvgRtt,
		PacketLossPercent: stats.PacketLoss,
	}, nil
}
