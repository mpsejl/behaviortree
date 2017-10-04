/*
Package behaviortree implements a simple framework
for creating and executing Behavior Trees as defined
in <TBD> specification.

The package provides the following features
	* Leaf Nodes:
		Condition
		Action
	* Selector Node
	* Sequence Node
	* Static Nodes:
		Always Succeed
		Always Fail
		Wait
	* Decorators:
		Always Succeed
		Always Fail
		Until Success
		Until Failure
		Invert
		Repeat
		Limit
	* Root Node - Provides functions execute
	* Export - Exports Tree Structure as JSON
-------------------------------------------------------*/
package behaviortree

import "fmt"

// Status Values
const (
	NONE = iota
	SUCCESS
	FAILURE
	RUNNING
	ERROR
)

// BTNode is an interface which defines
// Behavior Tree Node functions
type BTNode interface {
	Tick(int) int
	Update(int, int) int
	State() StateType
	Info() InfoType
}

// ListOfNodes defines an array used for compound nodes
type ListOfNodes []BTNode

// StateType defines Behavior Tree Node State
// that is updated during execution
type StateType struct {
	Status int
	N      int
}

// Update sets the current state of a Behavior Tree Node
func (t *StateType) Update(status int, n int) int {
	t.Status = status
	t.N = n
	return status
}

// State returns the current State of a Behavior Tree Node
func (t *StateType) State() StateType {
	return StateType{Status: t.Status, N: t.N}
}

// InfoType defines Behavior Tree Node Information
// including a label (name) and the type of node
type InfoType struct {
	Label string
	Type  string
}

// Info returns the set Information for a Behavior Tree Node
func (t *InfoType) Info() InfoType {
	return InfoType{Label: t.Label, Type: t.Type}
}

// Version returns package version information
func Version() {
	fmt.Println("BT Version 0.1 Alpha")
}
