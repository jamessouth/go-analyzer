package slices

import (
	"fmt"
	"strings"

	"github.com/jamessouth/go-analyzer/astrav"
	"github.com/jamessouth/go-analyzer/suggester/sugg"
)

// Register registers all suggestion functions for this exercise.
var Register = sugg.Register{
	Funcs: []sugg.SuggestionFunc{

		// examSwitch,
		// examReturns,
		// examIfs,
		examIfTwoConditions,
	},
	Severity: severity,
}

// checks if `switch` was used
// func examSwitch(pkg *astrav.Package, suggs sugg.Suggester) {
// 	nodes := pkg.FindByNodeType(astrav.NodeTypeSwitchStmt)
// 	if len(nodes) != 0 {
// 		suggs.AppendUnique(Switch)
// 	}
// }

// checks for two returns
// func examReturns(pkg *astrav.Package, suggs sugg.Suggester) {
// 	nodes := pkg.FindByNodeType(astrav.NodeTypeReturnStmt)
// 	if len(nodes) != 2 {
// 		suggs.AppendUnique(NotTwoReturns)
// 	}
// }

// checks for only one 'if'
// func examIfs(pkg *astrav.Package, suggs sugg.Suggester) {
// 	nodes := pkg.FindByNodeType(astrav.NodeTypeIfStmt)
// 	if len(nodes) != 1 {
// 		suggs.AppendUnique(NotOneIf)
// 	}
// }

// checks that 'if' has two conditions
func examIfTwoConditions(pkg *astrav.Package, suggs sugg.Suggester) {
	// orNodes := pkg.FindByToken(token.LOR)
	// andNodes := pkg.FindByToken(token.LAND)
	nodes := pkg.FindByNodeType(astrav.NodeTypeIfStmt)
	conditions := nodes[0].Children()[0].GetSourceString()
	fmt.Println(conditions)
	if strings.Count(conditions, "||")+strings.Count(conditions, "&&") != 1 {
		suggs.AppendUnique(NotTwoIfConditions)
	}
}

// func examSliceRuneConv(pkg *astrav.Package, suggs sugg.Suggester) {
// 	nodes := pkg.FindByName("[]rune")
// 	conversions := map[string]int{}
// 	for _, node := range nodes {
// 		if !node.IsNodeType(astrav.NodeTypeCallExpr) {
// 			continue
// 		}

// 		name := node.ChildByNodeType(astrav.NodeTypeIdent).(astrav.Named).NodeName()
// 		conversions[name]++
// 	}
// 	for varName, count := range conversions {
// 		if 1 < count {
// 			suggs.AppendUniquePH(MinSliceRuneConv, map[string]string{
// 				"name": varName,
// 			})
// 		}
// 	}
// }

// // checks code for an extra `if string(s) are empty` condition
// func examExtraIfEmpty(pkg *astrav.Package, suggs sugg.Suggester) {
// 	main := pkg.FuncDeclByName("Distance")
// 	if main == nil {
// 		return
// 	}
// 	ifs := pkg.FindByNodeType(astrav.NodeTypeIfStmt)
// 	if len(ifs) < 3 {
// 		return
// 	}

// 	returns := main.FindByNodeType(astrav.NodeTypeReturnStmt)
// 	for _, node := range returns {
// 		children := node.Children()
// 		if len(children) != 2 {
// 			continue
// 		}
// 		if !children[0].IsNodeType(astrav.NodeTypeBasicLit) ||
// 			children[0].(*astrav.BasicLit).Value != "0" {
// 			continue
// 		}
// 		if !children[1].IsNodeType(astrav.NodeTypeIdent) ||
// 			children[1].(*astrav.Ident).Name != "nil" {
// 			continue
// 		}
// 		suggs.AppendUnique(ExtraIfStringsEmpty)
// 	}
// }

// checks if `strings.TrimSpace` was used
// func examTrimSpace(pkg *astrav.Package, suggs sugg.Suggester) {
// 	nodes := pkg.FindByName("strings.TrimSpace")
// 	if len(nodes) != 0 {
// 		suggs.AppendUnique(sugg.TrimSpaceUsed)
// 	}
// }

// // checks if the hamming distance is made case insensitive -- which is not tested for but should not be done
// func examCaseInsensitive(pkg *astrav.Package, suggs sugg.Suggester) {
// 	nodes := pkg.FindByName("strings.ToLower")
// 	nodes = append(nodes, pkg.FindByName("strings.ToUpper")...)
// 	nodes = append(nodes, pkg.FindByName("unicode.ToLower")...)
// 	nodes = append(nodes, pkg.FindByName("unicode.ToUpper")...)
// 	if len(nodes) != 0 {
// 		suggs.AppendUnique(CaseInsensitive)
// 	}
// }

// // checks if an error is being returned right away. It should be.
// func examReturnOnError(pkg *astrav.Package, suggs sugg.Suggester) {
// 	nodes := pkg.FindByName("errors.New")
// 	nodes = append(nodes, pkg.FindByName("fmt.Errorf")...)

// 	for _, node := range nodes {
// 		ifStmt := node.NextParentByType(astrav.NodeTypeIfStmt)
// 		if ifStmt == nil {
// 			continue
// 		}
// 		returns := ifStmt.FindByNodeType(astrav.NodeTypeReturnStmt)
// 		if len(returns) == 0 {
// 			suggs.AppendUnique(ReturnOnError)
// 		}
// 	}
// }

// // looking for error definition in the form of "err := error(nil)"
// func examDefineError(pkg *astrav.Package, suggs sugg.Suggester) {
// 	nodes := pkg.FindByName("error")
// 	for _, node := range nodes {
// 		if !node.IsNodeType(astrav.NodeTypeCallExpr) {
// 			continue
// 		}
// 		for _, child := range node.Children() {
// 			if named, ok := child.(astrav.Named); ok && named.NodeName() == "nil" {
// 				suggs.AppendUnique(DefineEmptyErr)
// 			}
// 		}
// 	}
// }

// // looking for a mixture of runes and bytes. Also using the range index of string iteration as a byte or rune index.
// func examMixture(pkg *astrav.Package, suggs sugg.Suggester) {
// 	loop := pkg.FindFirstByNodeType(astrav.NodeTypeForStmt)
// 	if loop == nil {
// 		loop = pkg.FindFirstByNodeType(astrav.NodeTypeRangeStmt)
// 	}
// 	if loop == nil {
// 		return
// 	}
// 	loopType := getIndexType(loop)

// 	nodes := loop.FindByNodeType(astrav.NodeTypeBinaryExpr)
// 	for _, node := range nodes {
// 		if node.Parent().IsNodeType(astrav.NodeTypeForStmt) || node.Parent().IsNodeType(astrav.NodeTypeRangeStmt) {
// 			continue
// 		}
// 		expr := node.(*astrav.BinaryExpr)
// 		var xType = getUnderlyingValType(expr.X())
// 		var yType = getUnderlyingValType(expr.Y())
// 		if xType != yType {
// 			suggs.AppendUnique(MixtureRunesBytes)
// 			return
// 		}
// 		if loopType != "" && (xType != "" && loopType != xType || yType != "" && loopType != yType) {
// 			suggs.AppendUnique(RuneByteIndex)
// 			return
// 		}
// 		if expr.X().IsValueType("byte") {
// 			suggs.AppendUnique(ComparingBytes)
// 		}
// 	}
// }

// func getIndexType(node astrav.Node) string {
// 	if node.IsNodeType(astrav.NodeTypeRangeStmt) {
// 		rngType := node.(*astrav.RangeStmt).X().ValueType().String()
// 		switch rngType {
// 		case "string":
// 			return "runebyte"
// 		case "[]rune":
// 			return "rune"
// 		case "[]byte":
// 			return "byte"
// 		case "[]string":
// 			return "string"
// 		}
// 		return ""
// 	}
// 	return ""
// }

// func getUnderlyingValType(node astrav.Node) string {
// 	if node.IsNodeType(astrav.NodeTypeCallExpr) {
// 		for _, n := range node.Children() {
// 			if t := getUnderlyingValType(n); t != "" {
// 				return t
// 			}
// 		}
// 	}
// 	if node.IsNodeType(astrav.NodeTypeIdent) && isOneOf(node.(astrav.Named).NodeName(), "byte", "string", "rune") {
// 		return ""
// 	}

// 	if node.IsValueType("byte") {
// 		return "byte"
// 	}
// 	if node.IsValueType("rune") {
// 		return "rune"
// 	}
// 	if node.IsValueType("string") {
// 		return "string"
// 	}
// 	return ""
// }

// func isOneOf(s string, strs ...string) bool {
// 	for _, str := range strs {
// 		if s == str {
// 			return true
// 		}
// 	}
// 	return false
// }

// // check if strings.Split was used
// func examStringSplit(pkg *astrav.Package, suggs sugg.Suggester) {
// 	nodes := pkg.FindByName("strings.Split")
// 	for _, node := range nodes {
// 		if node.GetSourceString() == "strings.Split" {
// 			suggs.AppendUnique(StringsSplitUsed)
// 		}
// 	}
// }

// // check if bytes are being compared. Comment on how this won't work with utf8.
// func examComparingBytes(pkg *astrav.Package, suggs sugg.Suggester) {
// 	if suggs.HasSuggestion(MixtureRunesBytes) {
// 		return
// 	}
// 	nodes := pkg.FindByNodeType(astrav.NodeTypeBinaryExpr)
// 	for _, node := range nodes {
// 		if node.NextParentByType(astrav.NodeTypeForStmt) == nil && node.NextParentByType(astrav.NodeTypeRangeStmt) == nil ||
// 			node.Parent().IsNodeType(astrav.NodeTypeForStmt) || node.Parent().IsNodeType(astrav.NodeTypeRangeStmt) {
// 			continue
// 		}
// 		expr := node.(*astrav.BinaryExpr)
// 		if expr.X().IsValueType("byte") {
// 			suggs.AppendUnique(ComparingBytes)
// 			continue
// 		}
// 		if isByteIndexToRuneConv(expr.X()) && isByteIndexToRuneConv(expr.Y()) {
// 			suggs.AppendUnique(CompBytesInDisguise)
// 		}
// 	}
// }

// func isByteIndexToRuneConv(node astrav.Node) bool {
// 	return node.IsValueType("rune") &&
// 		node.IsNodeType(astrav.NodeTypeCallExpr) &&
// 		node.ChildByNodeType(astrav.NodeTypeIndexExpr) != nil
// }

// // check if Distance function exists
// func examMainFunc(pkg *astrav.Package, suggs sugg.Suggester) {
// 	main := pkg.FuncDeclByName("Distance")
// 	if main == nil {
// 		suggs.AppendUnique(MissingEntryFunc)
// 		return
// 	}

// 	if len(main.Params().Children()) != 2 {
// 		suggs.AppendUnique(FuncSignatureChanged)
// 	}
// 	if len(main.Results().Children()) != 2 {
// 		suggs.AppendUnique(FuncSignatureChanged)
// 	}
// }

// // check if returns of Distance function have been tempered with
// func examReturns(pkg *astrav.Package, suggs sugg.Suggester) {
// 	main := pkg.FindFirstByName("Distance")
// 	if main == nil {
// 		suggs.AppendUnique(MissingEntryFunc)
// 		return
// 	}

// 	returns := main.FindByNodeType(astrav.NodeTypeReturnStmt)
// 	for _, ret := range returns {
// 		if len(ret.Children()) == 0 {
// 			suggs.AppendUnique(NakedReturns)
// 			continue
// 		}
// 		if len(ret.Children()) < 2 {
// 			continue
// 		}

// 		if ret.Children()[1].ValueType().String() == "nil" {
// 			continue
// 		}
// 		switch node := ret.Children()[0].(type) {
// 		case *astrav.BasicLit:
// 			if node.Value != "0" {
// 				suggs.AppendUnique(ZeroValueOnErr)
// 			}
// 		case *astrav.UnaryExpr:
// 			lit := node.FindFirstByNodeType(astrav.NodeTypeBasicLit)
// 			if lit == nil {
// 				continue
// 			}
// 			if lit.(*astrav.BasicLit).Value != "0" {
// 				suggs.AppendUnique(ZeroValueOnErr)
// 			}
// 		}
// 	}
// }

// // Check if the distance counter was declared too early -- e.g. before the error check
// func examDeclareWhenNeeded(pkg *astrav.Package, suggs sugg.Suggester) {
// 	if suggs.HasSuggestion(InvertIf) {
// 		return
// 	}

// 	main := pkg.FindFirstByName("Distance")
// 	if main == nil {
// 		return
// 	}
// 	returns := main.FindByNodeType(astrav.NodeTypeReturnStmt)
// 	for _, ret := range returns {
// 		for _, child := range ret.Children() {
// 			if !child.IsNodeType(astrav.NodeTypeIdent) {
// 				continue
// 			}
// 			returnVar := child.(*astrav.Ident)
// 			if returnVar.Obj == nil {
// 				continue
// 			}

// 			varDecl := main.FindFirstByName(returnVar.Name).Parent()

// 			// variable not declared in the same block as the return statement
// 			if varDecl.IsNodeType(astrav.NodeTypeAssignStmt) {
// 				if !returnVar.NextParentByType(astrav.NodeTypeBlockStmt).Contains(varDecl) {
// 					suggs.AppendUniquePH(DeclareWhenNeeded, map[string]string{
// 						"returnVar": returnVar.Name,
// 					})
// 				}
// 			}

// 			// there is another return inbetween
// 			for _, rt := range returns {
// 				if rt == ret {
// 					continue
// 				}
// 				if varDecl.Pos() <= rt.Pos() && rt.Pos() <= returnVar.Pos() {
// 					suggs.AppendUniquePH(DeclareWhenNeeded, map[string]string{
// 						"returnVar": returnVar.Name,
// 					})
// 				}
// 			}
// 		}
// 	}
// }

// // Check for an if that can be inverted so the error case is checked first
// func examInvertIf(pkg *astrav.Package, suggs sugg.Suggester) {
// 	for _, ifNode := range pkg.FindByNodeType(astrav.NodeTypeIfStmt) {
// 		loop := ifNode.FindFirstByNodeType(astrav.NodeTypeRangeStmt)
// 		if loop == nil {
// 			loop = ifNode.FindFirstByNodeType(astrav.NodeTypeForStmt)
// 		}
// 		binExpr := ifNode.ChildByNodeType(astrav.NodeTypeBinaryExpr)
// 		if binExpr == nil {
// 			continue
// 		}
// 		condition := binExpr.(*astrav.BinaryExpr)
// 		if loop != nil && condition != nil && condition.Op.String() == "==" {
// 			suggs.AppendUnique(InvertIf)
// 		}
// 	}
// }

// // Look for rune to byte conversion
// func examRuneToByte(pkg *astrav.Package, suggs sugg.Suggester) {
// 	nodes := pkg.FindByName("byte")
// 	for _, node := range nodes {
// 		parentType := node.Parent().NodeType()
// 		if parentType != astrav.NodeTypeCallExpr {
// 			continue
// 		}
// 		for _, n := range node.Siblings() {
// 			if n.ValueType().String() == "rune" {
// 				suggs.AppendUnique(RuneToByte)
// 			}
// 		}
// 	}
// }

// // Look for conversion to string in order to compare runes and bytes
// func examMultipleStringConversions(pkg *astrav.Package, suggs sugg.Suggester) {
// 	rngNode := pkg.FindFirstByNodeType(astrav.NodeTypeRangeStmt)
// 	if rngNode == nil {
// 		return
// 	}

// 	count := 0
// 	for _, node := range rngNode.FindByName("string") {
// 		if node.Parent().IsNodeType(astrav.NodeTypeCallExpr) {
// 			count++
// 		}
// 	}
// 	if 1 < count {
// 		suggs.AppendUnique(ToStringConversion)
// 	}
// }

// // make sure ++ was used, if not comment
// func examIncrease(pkg *astrav.Package, suggs sugg.Suggester) {
// 	loop := pkg.FindFirstByNodeType(astrav.NodeTypeRangeStmt)
// 	if loop == nil {
// 		loop = pkg.FindFirstByNodeType(astrav.NodeTypeForStmt)
// 	}
// 	if loop == nil {
// 		return
// 	}
// 	for _, node := range loop.FindByNodeType(astrav.NodeTypeBinaryExpr) {
// 		if node.(*astrav.BinaryExpr).Op.String() == "+" {
// 			suggs.AppendUnique(IncreaseOperator)
// 		}
// 	}
// }

// // Check error message format for capitalization and punctuation
// func examErrorMessage(pkg *astrav.Package, suggs sugg.Suggester) {
// 	checkErrMessage(pkg.FindByName("errors.New"), suggs)
// 	checkErrMessage(pkg.FindByName("fmt.Errorf"), suggs)
// }

// func checkErrMessage(nodes []astrav.Node, suggs sugg.Suggester) {
// 	for _, node := range nodes {
// 		if node.NodeType() == astrav.NodeTypeIdent {
// 			continue
// 		}
// 		errMsgNode := node.(*astrav.SelectorExpr).Parent().FindFirstByNodeType(astrav.NodeTypeBasicLit)
// 		if errMsgNode == nil {
// 			continue
// 		}

// 		errText := errMsgNode.(*astrav.BasicLit).Value

// 		// check punctuation
// 		if strings.HasSuffix(errText, ".") {
// 			suggs.AppendUnique(ErrorMsgFormat)
// 			continue
// 		}

// 		// check if first letter is capitalized and second not.
// 		var isUpper bool
// 		for i, rn := range strings.Split(errText, " ")[0] {
// 			// first letter is " or `
// 			if i == 2 {
// 				if isUpper && !unicode.IsUpper(rn) {
// 					suggs.AppendUnique(ErrorMsgFormat)
// 				}
// 				break
// 			}
// 			isUpper = unicode.IsUpper(rn)
// 		}
// 	}
// }
