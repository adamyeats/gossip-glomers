package main

type IncomingMessage struct {
	Type string `json:"type"`
}

type OutgoingMessage struct {
	Type     string `json:"type"`
	Messages []int  `json:"messages,omitempty"`
}

type BroadcastMessage struct {
	Type    string `json:"type"`
	Message int    `json:"message"`
}

type TopologyMessage struct {
	Type     string              `json:"type"`
	Topology map[string][]string `json:"topology"`
}
