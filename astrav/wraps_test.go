package astrav

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectorExpr_PackageName(t *testing.T) {
	n := getTree(t, 1)

	selExpr := n.FindFirstByNodeType(NodeTypeSelectorExpr).(*SelectorExpr)
	pkgIdent := selExpr.PackageName()

	assert.Equal(t, "strings", pkgIdent.Name)
	assert.Equal(t, selExpr, pkgIdent.Parent())
}

func TestFuncType_Params(t *testing.T) {
	n := getTree(t, 1)

	f := n.FindFirstByName("Score").ChildByNodeType(NodeTypeFuncType)
	params := f.(*FuncType).Params()
	assert.NotNil(t, params)
	assert.Equal(t, 1, len(params.List))
}

func TestFuncType_Results(t *testing.T) {
	n := getTree(t, 1)

	f := n.FindFirstByName("Score").ChildByNodeType(NodeTypeFuncType)
	params := f.(*FuncType).Results()
	assert.NotNil(t, params)
	assert.Equal(t, 1, len(params.List))
}

func TestFuncDecl_Params(t *testing.T) {
	n := getTree(t, 1)

	f := n.FindFirstByName("Score")
	params := f.(*FuncDecl).Params()
	assert.NotNil(t, params)
	assert.Equal(t, 1, len(params.List))
}

func TestFuncDecl_Results(t *testing.T) {
	n := getTree(t, 1)

	f := n.FindFirstByName("Score")
	params := f.(*FuncDecl).Results()
	assert.NotNil(t, params)
	assert.Equal(t, 1, len(params.List))
}

func TestForStmt_Init(t *testing.T) {
	n := getPackage(t, 2)

	loop := n.FindFirstByNodeType(NodeTypeForStmt)
	init := loop.(*ForStmt).Init()
	assert.NotNil(t, init)
	assert.NotNil(t, init.FindByName("i"))
	assert.Equal(t, NodeTypeAssignStmt, init.NodeType())
}

func TestForStmt_Cond(t *testing.T) {
	n := getPackage(t, 2)

	loop := n.FindFirstByNodeType(NodeTypeForStmt)
	cond := loop.(*ForStmt).Cond()
	assert.NotNil(t, cond)
	assert.NotNil(t, cond.FindByName("i"))
	assert.NotNil(t, cond.FindByName("num"))
	assert.Equal(t, NodeTypeBinaryExpr, cond.NodeType())
}

func TestForStmt_Post(t *testing.T) {
	n := getPackage(t, 2)

	loop := n.FindFirstByNodeType(NodeTypeForStmt)
	post := loop.(*ForStmt).Post()
	assert.NotNil(t, post)
	assert.NotNil(t, post.FindByName("i"))
	assert.Equal(t, NodeTypeIncDecStmt, post.NodeType())
}

func TestRangeStmt_Key(t *testing.T) {
	n := getPackage(t, 1)

	loop := n.FindFirstByNodeType(NodeTypeRangeStmt)
	k := loop.(*RangeStmt).Key()
	assert.NotNil(t, k)
	assert.Equal(t, "_", k.(*Ident).Name)
	assert.Equal(t, NodeTypeIdent, k.NodeType())
}

func TestRangeStmt_Value(t *testing.T) {
	n := getPackage(t, 1)

	loop := n.FindFirstByNodeType(NodeTypeRangeStmt)
	v := loop.(*RangeStmt).Value()
	assert.NotNil(t, v)
	assert.Equal(t, "c", v.(*Ident).Name)
	assert.Equal(t, NodeTypeIdent, v.NodeType())
}

func TestRangeStmt_X(t *testing.T) {
	n := getPackage(t, 1)

	loop := n.FindFirstByNodeType(NodeTypeRangeStmt)
	x := loop.(*RangeStmt).X()
	assert.NotNil(t, x)
	assert.NotNil(t, x.FindByName("word"))
	assert.Equal(t, NodeTypeCallExpr, x.NodeType())
}
