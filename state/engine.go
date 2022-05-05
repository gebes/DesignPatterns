package main

type EngineState string

const (
	EngineStateOff      = "off"
	EngineStateStarting = "starting"
	EngineStateRunning  = "running"
)

func (receiver *EngineState) Starting() {
	*receiver = EngineStateStarting
}
func (receiver *EngineState) Running() {
	*receiver = EngineStateRunning
}

func (receiver *EngineState) Stop() {
	*receiver = EngineStateOff
}
