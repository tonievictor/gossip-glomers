package main

// message type
//
//	{
//	  "src": "c1",
//	  "dest": "n1",
//	  "body": {
//	    "type": "echo",
//	    "msg_id": 1,
//	    "echo": "Please echo 35"
//	  }
//	}
import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	node := maelstrom.NewNode()

	node.Handle("generate", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		id := fmt.Sprintf("%s%v", time.Now().String(), body["msg_id"])
		body["type"] = "generate_ok"
		body["id"] = id
		body["in_reply_to"] = body["msg_id"]
		return node.Reply(msg, body)
	})

	if err := node.Run(); err != nil {
		log.Fatal(err)
	}
}
