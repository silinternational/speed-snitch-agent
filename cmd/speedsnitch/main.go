package main

import (
	"fmt"
	"github.com/silinternational/speed-snitch-agent"
	"github.com/silinternational/speed-snitch-agent/lib/adminapi"
	"os"
	"time"
)

var config agent.Config
var agentStartTime time.Time

func main() {
	agentStartTime = time.Now()
	if len(os.Args) < 2 {
		fmt.Println("You must provide the Admin API BaseURL as the first arguement")
		os.Exit(1)
	}

	config, err := adminapi.GetConfig(os.Args[1])
	if err != nil {
		fmt.Println("Unable to fetch config from admin API:", err)
		os.Exit(1)
	}
}
