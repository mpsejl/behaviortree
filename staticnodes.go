package behaviortree

import "time"

// SuccessNode defines a node that always returns SUCCESS
type SuccessNode struct {
	InfoType
	StateType
}

// Tick returns SUCCESS
func (t *SuccessNode) Tick(n int) int {
	return t.Update(SUCCESS, n)
}

// FailureNode defines a behavior that always returns FAILURE
type FailureNode struct {
	InfoType
	StateType
}

// Tick returns FAILURE
func (t *FailureNode) Tick(n int) int {
	return t.Update(FAILURE, n)
}

// WaitNode defines a behavior that sleeps for
// a specified number of ms and then returns SUCCESS
type WaitNode struct {
	InfoType
	StateType
	durms time.Duration
}

// Tick sleeps and then returs SUCCESS
func (t *WaitNode) Tick(n int) int {
	time.Sleep(t.durms * time.Millisecond)
	return t.Update(SUCCESS, n)
}
