package behaviortree

// AlwaysFail executes the wrapped node and
//always returns FAILURE
func AlwaysFail(t *DecoratorNode, n int) int {
	return t.Update(FAILURE, n)
}

// AlwaysSucceed executes the wrapped node and
// always returns SUCCESS
func AlwaysSucceed(t *DecoratorNode, n int) int {
	return t.Update(SUCCESS, n)
}

// Invert executes the wrapped task and
// returns will SUCCESS if the wrapped task fails
// and returns FAILURE if the wrapped task succeeds
func Invert(t *DecoratorNode, n int) int {
	status := t.Node.Tick(n)
	switch status {
	case SUCCESS:
		return t.Update(FAILURE, n)
	case FAILURE:
		return t.Update(SUCCESS, n)
	default:
		return t.Update(status, n)
	}
}

// UntilFailure executes the wrapped Node until it fails
// and then returns SUCCESS
func UntilFailure(t *DecoratorNode, n int) int {
	for t.Node.Tick(n) != FAILURE {
	}
	return t.Update(SUCCESS, n)
}

// UntilSuccess executes the wrapped Node until if succeeds
// and then returns SUCCESS
func UntilSuccess(t *DecoratorNode, n int) int {
	for t.Node.Tick(n) != SUCCESS {
	}
	return t.Update(SUCCESS, n)
}

// Repeat executes the wrapped task N times
// and then return SUCCESS
func Repeat(t *DecoratorNode, n int) int {
	for i := 0; i < t.Limit; i++ {
		t.Node.Tick(n)
	}
	return t.Update(SUCCESS, n)
}

// LimitN executes the wrapped task once per Tick
// it maintains a count of executions and returns
// FAILURE if the count exceeds the specified limit
func LimitN(t *DecoratorNode, n int) int {
	t.Count++
	if t.Count <= t.Limit {
		status := t.Node.Tick(n)
		return t.Update(status, n)
	}
	return t.Update(FAILURE, n)
}
