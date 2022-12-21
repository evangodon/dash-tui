package main

type configUpdateMsg struct{ err error }
type moduleUpdateMsg struct{}
type modulesDone struct{}
type startSpinner struct{}
