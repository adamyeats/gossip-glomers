package main

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

// Service is the main service.
type Service struct {
	Node *maelstrom.Node

	broadcast *BroadcastData
	topology  *TopologyData
}

// Run starts the service.
func (svc *Service) Run() error {
	return svc.Node.Run()
}

// broadcastHandler handles broadcast messages.
func (svc *Service) BroadcastHandler(msg maelstrom.Message) error {
	var body BroadcastMessage

	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	svc.broadcast.Lock()
	defer svc.broadcast.Unlock()

	svc.broadcast.AppendID(body.Message)

	out := OutgoingMessage{
		Type: "broadcast_ok",
	}

	return svc.Node.Reply(msg, out)
}

// readHandler handles read messages.
func (svc *Service) ReadHandler(msg maelstrom.Message) error {
	svc.broadcast.RLock()
	defer svc.broadcast.RUnlock()

	out := OutgoingMessage{
		Type:     "read_ok",
		Messages: svc.broadcast.GetIDs(),
	}

	return svc.Node.Reply(msg, out)
}

// topologyHandler handles topology messages.
func (svc *Service) TopologyHandler(msg maelstrom.Message) error {
	var t TopologyMessage

	if err := json.Unmarshal(msg.Body, &t); err != nil {
		return err
	}

	svc.topology.Lock()
	defer svc.topology.Unlock()

	svc.topology.Set(t.Topology)

	out := OutgoingMessage{
		Type: "topology_ok",
	}

	return svc.Node.Reply(msg, out)
}
