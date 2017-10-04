package behaviortree

import (
	"math/rand"
	"reflect"
)

// SelectorNode defines a compound node that executes
// each of its child nodes in turn until one returns
// SUCCESS or RUNNING, otherwise it returns FAILURE
type SelectorNode struct {
	InfoType
	StateType
	Nodes ListOfNodes
}

// NewSelectorNode creates an instance of a Selector
func NewSelectorNode(label string, node ...BTNode) *SelectorNode {
	x := SelectorNode{}
	x.Label = label
	x.Type = reflect.TypeOf(x).Name()
	if len(node) > 0 {
		x.Nodes = append(x.Nodes, node...)
	}
	return &x
}

// Or adds child nodes to the Selector
func (t *SelectorNode) Or(node ...BTNode) *SelectorNode {
	t.Nodes = append(t.Nodes, node...)
	return t
}

// Tick executes the Selector
func (t *SelectorNode) Tick(n int) int {
	for _, node := range t.Nodes {
		status := node.Tick(n)
		if status == RUNNING {
			return t.Update(RUNNING, n)
		} else if status == SUCCESS {
			return t.Update(SUCCESS, n)
		}
	}
	return t.Update(FAILURE, n)
}

// RandomSelectorNode defines a compound node that executes
// each of its child nodes in a random order until one
// returns SUCCESS or RUNNING, otherwise it returns FAILURE
type RandomSelectorNode SequenceNode

// Tick executes the Random Selector
func (t *RandomSelectorNode) Tick(n int) int {
	r := rand.Perm(len(t.Nodes))
	for _, j := range r {
		node := t.Nodes[j]
		status := node.Tick(n)
		if status == RUNNING {
			return t.Update(RUNNING, n)
		} else if status == SUCCESS {
			return t.Update(SUCCESS, n)
		}
	}
	return t.Update(FAILURE, n)
}
