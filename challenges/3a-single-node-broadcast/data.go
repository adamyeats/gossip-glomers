package main

import "sync"

// BroadcastData is a data structure that stores broadcast messages.
type BroadcastData struct {
	ids []int
	mu  sync.RWMutex
}

// GetIDs returns a copy of the list of IDs.
func (d *BroadcastData) GetIDs() []int {
	ids := make([]int, len(d.ids))
	copy(ids, d.ids)
	return ids
}

// AppendID appends an ID to the list of IDs.
func (d *BroadcastData) AppendID(id int) {
	d.ids = append(d.ids, id)
}

// Lock locks the data structure.
func (d *BroadcastData) Lock() {
	d.mu.Lock()
}

// Unlock unlocks the data structure.
func (d *BroadcastData) Unlock() {
	d.mu.Unlock()
}

// RLock locks the data structure for reading.
func (d *BroadcastData) RLock() {
	d.mu.RLock()
}

// RUnlock unlocks the data structure for reading.
func (d *BroadcastData) RUnlock() {
	d.mu.RUnlock()
}

// TopologyData is a data structure that stores the topology.
type TopologyData struct {
	data map[string][]string
	mu   sync.RWMutex
}

// Set sets the topology.
func (d *TopologyData) Set(top map[string][]string) {
	d.data = top
}

// Lock locks the data structure.
func (d *TopologyData) Lock() {
	d.mu.Lock()
}

// Unlock unlocks the data structure.
func (d *TopologyData) Unlock() {
	d.mu.Unlock()
}
