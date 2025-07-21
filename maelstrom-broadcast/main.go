package main

import (
	"encoding/json"
	"log"

	"github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	node := maelstrom.NewNode()

	var messages []float64
	node.Handle("broadcast", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		m := body["message"].(float64)
		messages = append(messages, m)

		resbody := make(map[string]string)

		resbody["type"] = "broadcast_ok"
		return node.Reply(msg, resbody)
	})

	node.Handle("read", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		body["type"] = "read_ok"
		body["messages"] = messages
		return node.Reply(msg, body)
	})

	node.Handle("topology", func(msg maelstrom.Message) error {
		body := make(map[string]any)
		body["type"] = "topology_ok"
		return node.Reply(msg, body)
	})

	if err := node.Run(); err != nil {
		log.Fatal(err)
	}
}
