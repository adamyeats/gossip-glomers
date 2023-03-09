package main

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type IncomingMessage struct {
	Type string `json:"type"`
}

type OutgoingMessage struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// Service is the main service.
type Service struct {
	Node *maelstrom.Node
}

// Run starts the service.
func (svc *Service) Run() error {
	return svc.Node.Run()
}

// handler is the callback for the "generate" command.
func (svc *Service) handler(msg maelstrom.Message) error {
	var body IncomingMessage

	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	out := OutgoingMessage{
		Type: "generate_ok",
		ID:   uuid.New().String(),
	}

	return svc.Node.Reply(msg, out)
}

func main() {
	svc := Service{
		Node: maelstrom.NewNode(),
	}

	svc.Node.Handle("generate", svc.handler)

	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
