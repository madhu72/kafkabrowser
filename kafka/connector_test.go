package kafka

import (
	"fmt"
	"testing"
	"github.com/alecthomas/log4go"
)

func TestConnector(t *testing.T) {
	fmt.Printf("Running Test for Kafka Connector\n")
	var log log4go.Logger
	connector := &KafkaConnector{}
	connector.Init("localhost:9092","test",0,log)
	connector.InitWriter()
	connector.InitReader()

	defer connector.Close()
	fmt.Printf("Writing: KAFKABROWSER::TEST MESSAGE\n")
	connector.WriteMessage(connector.GetMessage([]byte("KAFKABROWSER::TEST MESSAGE")))
	msg, err := connector.ReadMessage()
	if err!=nil {
		t.Error(err)
	} else {
		if string(msg.Value) != "KAFKABROWSER::TEST MESSAGE" {
			t.Errorf("Reader failed to sync with writer")
		} else {
			fmt.Printf("Text Matched with Write:%s\n",string(msg.Value))
		}
	}

}



