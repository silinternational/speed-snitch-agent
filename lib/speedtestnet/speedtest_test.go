package speedtestnet

import (
	"github.com/silinternational/speed-snitch-agent"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, _ := NewClient()
	if client.Type != agent.TypeSpeedtest {
		t.Error("Speedtest client type not what was epxected, got ", client.Type)
	}
}
