package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	svc := Service{
		Node: maelstrom.NewNode(),
		broadcast: &BroadcastData{
			ids: make([]int, 0),
		},
		topology: &TopologyData{
			data: make(map[string][]string),
		},
	}

	svc.Node.Handle("broadcast", svc.BroadcastHandler)
	svc.Node.Handle("read", svc.ReadHandler)
	svc.Node.Handle("topology", svc.TopologyHandler)

	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
