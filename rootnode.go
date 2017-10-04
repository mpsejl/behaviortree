package behaviortree

import (
	"fmt"
	"reflect"
)

// RootNode defines a Root Node which is the start or
// head of the Behavior Tree
type RootNode struct {
	InfoType
	StateType
	Node BTNode
}

// NewRootNode creates a new Behavior Tree
// and initializes it for execution
func NewRootNode(nm string, node BTNode, step bool) *RootNode {
	x := RootNode{Node: node}
	x.Label = nm
	x.Type = reflect.TypeOf(x).Name()
	x.Status = RUNNING
	x.N = 0
	return &x
}

// Tick executes the root node for the next tick
// and updates the state of the Behavior Tree
func (t *RootNode) Tick(n int) int {
	fmt.Println("Tick:", n)
	t.Status = t.Node.Tick(n)
	return t.Update(t.Status, n)
}

// Step executes a single itteration (tick)
// of the Behavior Tree
func (t *RootNode) Step() int {
	t.N++
	return t.Tick(t.N)
}

// Run ticks the Behavior Tree until
// completion (FAILURE|SUCCESS)
func (t *RootNode) Run() int {
	for ; t.Status == RUNNING; t.Step() {
	}
	return t.Status
}
