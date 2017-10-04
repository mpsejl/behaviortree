package behaviortree

import "reflect"

// ActionNode defines a leaf node that runs
// a task (calls a function)
type ActionNode struct {
	InfoType
	StateType
	Action func() int
}

// NewActionNode creates an ActionNode instance
// that will execute the specified function
func NewActionNode(f func() int) *ActionNode {
	x := ActionNode{}
	x.Label = getFuncName(f)
	x.Type = reflect.TypeOf(x).Name()
	x.Action = f
	return &x
}

// Tick executes the defined Action for this node
// and returns the status as SUCCESS, FAILURE, or RUNNING
func (t *ActionNode) Tick(n int) int {
	status := t.Action()
	return t.Update(status, n)
}

// ConditionNode defines a leaf node that evaluates
// a boolean condition
type ConditionNode struct {
	InfoType
	StateType
	Condition func() bool
}

// NewConditionNode creates a Condition Node instance
// that will evaluate the specified function
func NewConditionNode(nm string, f func() bool) *ConditionNode {
	x := ConditionNode{}
	x.Label = getFuncName(f)
	x.Type = reflect.TypeOf(x).Name()
	x.Condition = f
	return &x
}

// Tick evaluates the defined Condition for this node
// and returns the status as SUCCESS (true) or FAILURE (false)
func (t *ConditionNode) Tick(n int) int {
	if t.Condition() {
		return t.Update(SUCCESS, n)
	}
	return t.Update(FAILURE, n)
}
