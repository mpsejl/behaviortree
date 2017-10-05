package behaviortree

// DecoratorNode type defines a Behavior Tree Decorator
type DecoratorNode struct {
	InfoType
	StateType
	Node      BTNode
	Limit     int
	Count     int
	decorator func(*DecoratorNode, int) int
}

// Decorators Implemented
const (
	ALWAYSFAIL = iota
	ALWAYSSUCCEED
	INVERT
	UNTILFAIL
	UNTILSUCCESS
	REPEAT
	LIMIT
)

// Tick executes the Node through the Decorator function
func (t *DecoratorNode) Tick(n int) int {
	return t.decorator(t, n)
}

// NewDecorator creates a new Decorator instance for
// the specified function and node
func NewDecorator(decoratorid int, label string, node BTNode, limit int) *DecoratorNode {
	x := DecoratorNode{}
	x.Label = label
	x.Node = node
	x.Limit = limit
	switch decoratorid {
	case ALWAYSFAIL:
		x.decorator = AlwaysFail
	case ALWAYSSUCCEED:
		x.decorator = AlwaysSucceed
	case INVERT:
		x.decorator = Invert
	case UNTILFAIL:
		x.decorator = UntilFailure
	case UNTILSUCCESS:
		x.decorator = UntilSuccess
	case REPEAT:
		x.decorator = Repeat
	case LIMIT:
		x.decorator = LimitN
	}

	x.Type = "D:" + getFuncName(x.decorator)

	return &x
}

// SetNode updates the Decorators wrapped Node
func (t *DecoratorNode) SetNode(node BTNode) {
	t.Node = node
}
