package main

// The config received a live update
type configUpdateMsg struct{ err error }

// A module received an update
type moduleUpdateMsg struct{}

// All modules have finished running
type modulesDoneMsg struct{}

// Start spinner animation on modules that are running
type startSpinnerMsg struct{}

// The tab changed
type tabChangeMsg struct{}
