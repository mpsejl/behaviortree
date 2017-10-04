package behaviortree

import (
	"math/rand"
	"reflect"
)

// SequenceNode defines a a compound node that executes
// each of its child nodes in turn until one returns
// FAILURE or RUNNING, otherwise it returns SUCCESS
type SequenceNode struct {
	InfoType
	StateType
	Nodes ListOfNodes
}

// NewSequenceNode creates an instance of a Sequence
func NewSequenceNode(nm string, node ...BTNode) *SequenceNode {
	x := SequenceNode{}
	x.Label = nm
	x.Type = reflect.TypeOf(x).Name()
	if len(node) > 0 {
		x.Nodes = append(x.Nodes, node...)
	}
	return &x
}

// And adds child nodes to the Sequence
func (t *SequenceNode) And(node ...BTNode) *SequenceNode {
	t.Nodes = append(t.Nodes, node...)
	return t
}

// Tick executes the Sequence
func (t *SequenceNode) Tick(n int) int {
	for _, node := range t.Nodes {
		status := node.Tick(n)
		if status == RUNNING {
			return t.Update(RUNNING, n)
		} else if status == FAILURE {
			return t.Update(FAILURE, n)
		}
	}
	return t.Update(SUCCESS, n)
}

// RandomSequenceNode defines a behavior that executes
// each of its child nodes in a random order  until 
// one returns FAILURE or RUNNING, otherwise it 
// returns SUCCESS
type RandomSequenceNode SequenceNode

// Tick executes the Random Sequence
func (t *RandomSequenceNode) Tick(n int) int {
	r := rand.Perm(len(t.Nodes))
	for _, j := range r {
		node := t.Nodes[j]
		status := node.Tick(n)
		if status == RUNNING {
			return t.Update(RUNNING, n)
		} else if status == FAILURE {
			return t.Update(FAILURE, n)
		}
	}
	return t.Update(SUCCESS, n)
}
