package logentries


import (
	"fmt"
	"github.com/bsphere/le_go"
)

type Logger struct {} // to be used with agent.LoggerInstance

// Process logs to the logentries online log service
func (l Logger) Process(logKey, text string, a ...interface{}) error {
	if logKey == "" {
		return fmt.Errorf("No Log Key provided")
	}

	formattedText := fmt.Sprintf(text, a...)

	le, err := le_go.Connect(logKey)

	if err != nil {
		fmt.Printf("Got Error in Logger: %s\n", err.Error())
		return err
	}

	defer le.Close()
	le.Println(formattedText)
	return nil
}