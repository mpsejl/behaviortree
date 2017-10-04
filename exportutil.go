package behaviortree

import (
	"fmt"
	"strings"
)

// Export1 walks the Behavior Tree and returns a
// JSON Graph Structure containing Nodes and Edges
func Export1(node BTNode) string {

	var gnodes []string
	var gedges []string

	walkTree1(node, 0, &gnodes, &gedges)

	return fmt.Sprintf(`{"Nodes":[%s],"Edges":[%s]}`, strings.Join(gnodes, ","), strings.Join(gedges, ","))

}

// walkTree1 is a recursive function that walks the
// Behavior Tree and creates nodes and edges
func walkTree1(node BTNode, level int, gnodes *[]string, gedges *[]string) int {

	const FMTGNODE = `{"id":%d,"type":"%s","label":"%s"}`
	const FMTGEDGE = `{"source":%d,"target":%d}`

	l := level

	info := node.Info()
	isource := len(*gnodes) + 1
	*gnodes = append(*gnodes, fmt.Sprintf(FMTGNODE, isource, info.Type, info.Label))

	switch node.(type) {
	case *SequenceNode:
		nodes := node.(*SequenceNode).Nodes
		for _, x := range nodes {
			itarget := walkTree1(x, l+1, gnodes, gedges)
			*gedges = append(*gedges, fmt.Sprintf(FMTGEDGE, isource, itarget))

		}
	case *SelectorNode:
		nodes := node.(*SelectorNode).Nodes
		for _, x := range nodes {
			itarget := walkTree1(x, l+1, gnodes, gedges)
			*gedges = append(*gedges, fmt.Sprintf(FMTGEDGE, isource, itarget))
		}
	case *DecoratorNode:
		itarget := walkTree1(node.(*DecoratorNode).Node, l+1, gnodes, gedges)
		*gedges = append(*gedges, fmt.Sprintf(FMTGEDGE, isource, itarget))
	case *RootNode:
		itarget := walkTree1(node.(*RootNode).Node, l+1, gnodes, gedges)
		*gedges = append(*gedges, fmt.Sprintf(FMTGEDGE, isource, itarget))
	default:
	}

	return isource

}

// Export2 walks the Behavior Tree and returns a
// Parent Child JSON Structure
func Export2(node BTNode) string {
	return walkTree2(node, 0)

}

// walkTree2 is a recursive function that walks the
// Behavior Tree and creates parent/child data
func walkTree2(node BTNode, level int) string {

	const FMTGPARENT = `{"type":"%s","name":"%s","status":%d, "n":%d, "children":[%s]}`
	const FMTGLEAF = `{"type":"%s","name":"%s", "status":%d, "n":%d}`

	info := node.Info()
	state := node.State()

	children := ""
	var a []string

	switch node.(type) {
	case *SequenceNode:
		nodes := node.(*SequenceNode).Nodes
		a = make([]string, len(nodes))
		for i, x := range nodes {
			a[i] = walkTree2(x, level+1)
		}
	case *SelectorNode:
		nodes := node.(*SelectorNode).Nodes
		a = make([]string, len(nodes))
		for i, x := range nodes {
			a[i] = walkTree2(x, level+1)
		}
	case *DecoratorNode:
		a = make([]string, 1)
		a[0] = walkTree2(node.(*DecoratorNode).Node, level+1)
	case *RootNode:
		a = make([]string, 1)
		a[0] = walkTree2(node.(*RootNode).Node, level+1)
	default:
	}

	if len(a) > 0 {
		children = fmt.Sprintf(FMTGPARENT, info.Type, info.Label, state.Status, state.N, strings.Join(a, ","))
	} else {
		children = fmt.Sprintf(FMTGLEAF, info.Type, info.Label, state.Status, state.N)
	}

	return children

}
