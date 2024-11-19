package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	node := maelstrom.NewNode()
	values := []any{}
	node.Handle("broadcast", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		values = append(values, body["message"])
		response := map[string]any{"type": "broadcast_ok"}
		return node.Reply(msg, response)
	})
	node.Handle("read", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		response := map[string]any{"type": "read_ok", "messages": values}
		return node.Reply(msg, response)
	})
	node.Handle("topology", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		response := map[string]any{"type": "topology_ok"}
		return node.Reply(msg, response)
	})
	if err := node.Run(); err != nil {
		log.Fatal(err)
	}
}
