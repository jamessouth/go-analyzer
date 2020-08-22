package go-analyzer/astpatt

import "github.com/jamessouth/astrav"

// Special node types
const (
	// NodeTypeOmit will be omitted including its children.
	NodeTypeOmit astrav.NodeType = "omit"
	// NodeTypeSkip defines a node that is not taken into accound. It's children will be inlined.
	NodeTypeSkip astrav.NodeType = "skip"
)

// Omit is a placeholder node that is to be omitted.
type Omit struct {
	parentNode
}

// Populate populates the pattern node from a given ast node.
func (s *Omit) Populate(ast astrav.Node) {}

func (s *Omit) isType(nodeType astrav.NodeType) bool {
	return NodeTypeOmit == nodeType
}

// Skip is a placeholder node that is to be omitted.
type Skip struct {
	parentNode
}

// Populate populates the pattern node from a given ast node.
func (s *Skip) Populate(ast astrav.Node) {}

func (s *Skip) isType(nodeType astrav.NodeType) bool {
	return NodeTypeSkip == nodeType
}
