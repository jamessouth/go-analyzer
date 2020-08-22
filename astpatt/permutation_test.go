package go-analyzer/astpatt

import (
	"reflect"
	"testing"
)

func TestExtractPatternPermutations(t *testing.T) {
	tests := []struct {
		dir string
	}{
		{dir: "solutions/permutations/1"},
		{dir: "solutions/permutations/2"},
		{dir: "solutions/permutations/3"},
	}
	for _, tt := range tests {
		t.Run(tt.dir, func(t *testing.T) {
			pkg, err := getPackage(tt.dir)
			if err != nil {
				t.Error(t)
				return
			}
			patt := ExtractPattern(pkg)
			perms := ExtractPatternPermutations(pkg)
			for _, perm := range perms {
				if perm.Match(patt) {
					return
				}
			}
			t.Error("no perms pattern equals the extracted pattern")
		})
	}
}

func TestPerm(t *testing.T) {
	type args struct {
		a []Node
	}
	defaultNode := &DefaultNode{}
	funcDecl := &FuncDecl{}
	selectorExpr := &SelectorExpr{}
	tests := []struct {
		name string
		args args
		want [][]Node
	}{
		{
			name: "empty",
			args: args{a: []Node{}},
			want: [][]Node{nil},
		},
		{
			name: "permutating",
			args: args{a: []Node{
				defaultNode,
				funcDecl,
				selectorExpr,
			}},
			want: [][]Node{
				{
					defaultNode,
					funcDecl,
					selectorExpr,
				},
				{
					defaultNode,
					selectorExpr,
					funcDecl,
				},
				{
					funcDecl,
					defaultNode,
					selectorExpr,
				},
				{
					funcDecl,
					selectorExpr,
					defaultNode,
				},
				{
					selectorExpr,
					funcDecl,
					defaultNode,
				},
				{
					selectorExpr,
					defaultNode,
					funcDecl,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Perm(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Perm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combinations(t *testing.T) {
	type args struct {
		pNodes [][]Node
	}
	node1 := &DefaultNode{}
	node2 := &DefaultNode{}
	node3 := &SelectorExpr{}
	node4 := &Omit{}
	node5 := &Omit{}
	node6 := &Skip{}
	tests := []struct {
		name             string
		args             args
		wantCombinations [][]Node
	}{
		{
			name: "cominations",
			args: args{
				pNodes: [][]Node{
					{node1, node5, node6},
					{node2},
					{node3, node4},
				},
			},
			wantCombinations: [][]Node{
				{node1, node2, node3},
				{node5, node2, node3},
				{node6, node2, node3},
				{node1, node2, node4},
				{node5, node2, node4},
				{node6, node2, node4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCombinations := combinations(tt.args.pNodes); !reflect.DeepEqual(gotCombinations, tt.wantCombinations) {
				t.Errorf("combinations() = %v, want %v", gotCombinations, tt.wantCombinations)
			}
		})
	}
}

func Test_permutations(t *testing.T) {
	tests := []struct {
		name    string
		dir     string
		wantLen int
	}{
		{
			name:    "sample 2",
			dir:     "solutions/permutations/2",
			wantLen: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkg, err := getPackage(tt.dir)
			if err != nil {
				t.Error(t)
				return
			}
			if got := permutations(pkg); len(got) != tt.wantLen {
				t.Errorf("permutations().len = %v, want %v", len(got), tt.wantLen)
			}
		})
	}
}
