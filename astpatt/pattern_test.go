package go-analyzer/astpatt

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/jamessouth/astrav"
	"github.com/stretchr/testify/assert"
)

type test struct {
	solution   string
	match      bool
	minPercent float64
	maxPercent float64
}

var matchTests = []struct {
	patterns []string
	tests    []test
}{
	{
		patterns: []string{
			"solutions/twofer/patterns/1",
			"solutions/twofer/patterns/2",
		},
		tests: []test{
			{solution: "solutions/twofer/1", match: true, minPercent: 1},
			{solution: "solutions/twofer/2", match: false, maxPercent: 0.9},
			{solution: "solutions/twofer/3", match: true},
			{solution: "solutions/twofer/4", match: true},
			{solution: "solutions/twofer/5", match: false, maxPercent: 0.9},
			{solution: "solutions/twofer/6", match: false},
			{solution: "solutions/twofer/7", match: false},
			{solution: "solutions/twofer/8", match: true},
			{solution: "solutions/twofer/9", match: true},
			{solution: "solutions/twofer/10", match: true},
			{solution: "solutions/twofer/11", match: true},
			{solution: "solutions/twofer/12", match: true},
		},
	},
	{
		patterns: []string{
			"solutions/hamming/patterns/1",
			"solutions/hamming/patterns/2",
		},
		tests: []test{
			{solution: "solutions/hamming/1", match: false},
			{solution: "solutions/hamming/2", match: false},
			{solution: "solutions/hamming/3", match: false},
			{solution: "solutions/hamming/4", match: false},
			{solution: "solutions/hamming/5", match: false},
			{solution: "solutions/hamming/6", match: false, minPercent: 0.9},
			{solution: "solutions/hamming/7", match: false, minPercent: 0.9},
			{solution: "solutions/hamming/8", match: false},
			{solution: "solutions/hamming/9", match: false},
			// TODO: {solution: "solutions/hamming/10", match: false, minPercent: 0.9},
			// TODO: {solution: "solutions/hamming/11", match: false, minPercent: 0.9},
			{solution: "solutions/hamming/12", match: false, minPercent: 0.98},
			{solution: "solutions/hamming/13", match: false, minPercent: 0.98},
		},
	},
	{
		patterns: []string{
			"solutions/raindrops/patterns/1",
			"solutions/raindrops/patterns/2",
			"solutions/raindrops/patterns/3",
		},
		tests: []test{
			{solution: "solutions/raindrops/1", match: false},
			{solution: "solutions/raindrops/2", match: true},
			{solution: "solutions/raindrops/3", match: false},
			// TODO: {solution: "solutions/raindrops/4", match: false, minPercent: 0.9},
			{solution: "solutions/raindrops/5", match: false, minPercent: 0.9},
			{solution: "solutions/raindrops/6", match: false},
			{solution: "solutions/raindrops/7", match: false, maxPercent: 0.85},
			{solution: "solutions/raindrops/8", match: false},
			{solution: "solutions/raindrops/9", match: false},
			{solution: "solutions/raindrops/10", match: false},
			{solution: "solutions/raindrops/11", match: false},
			{solution: "solutions/raindrops/12", match: false},
			// TODO: {solution: "solutions/raindrops/13", match: false, minPercent: 0.9},
			{solution: "solutions/raindrops/14", match: false},
			{solution: "solutions/raindrops/15", match: false, maxPercent: 0.85},
			{solution: "solutions/raindrops/16", match: false, maxPercent: 0.85},
			{solution: "solutions/raindrops/17", match: true},
			{solution: "solutions/raindrops/18", match: true},
		},
	},
	// {
	// 	patterns: []string{"solutions/raindrops/7"},
	// 	tests: []test{
	// 		{solution: "solutions/raindrops/1", match: false},
	// 		{solution: "solutions/raindrops/2", match: false},
	// 		{solution: "solutions/raindrops/3", match: false},
	// 		{solution: "solutions/raindrops/4", match: false},
	// 		{solution: "solutions/raindrops/5", match: false},
	// 		{solution: "solutions/raindrops/6", match: false},
	// 		{solution: "solutions/raindrops/7", match: true},
	// 		{solution: "solutions/raindrops/8", match: false},
	// 		{solution: "solutions/raindrops/9", match: false},
	// 		{solution: "solutions/raindrops/10", match: false},
	// 		{solution: "solutions/raindrops/11", match: false},
	// 		{solution: "solutions/raindrops/12", match: false},
	// 		{solution: "solutions/raindrops/13", match: false},
	// 		{solution: "solutions/raindrops/14", match: false},
	// 		{solution: "solutions/raindrops/15", match: false},
	// 		{solution: "solutions/raindrops/16", match: false},
	// 	},
	// },
	// {
	// 	patterns: []string{"solutions/raindrops/6"},
	// 	tests: []test{
	// 		{solution: "solutions/raindrops/1", match: false},
	// 		{solution: "solutions/raindrops/2", match: false},
	// 		{solution: "solutions/raindrops/3", match: false},
	// 		{solution: "solutions/raindrops/4", match: false},
	// 		{solution: "solutions/raindrops/5", match: false},
	// 		{solution: "solutions/raindrops/6", match: true},
	// 		{solution: "solutions/raindrops/7", match: false},
	// 		{solution: "solutions/raindrops/8", match: true},
	// 		{solution: "solutions/raindrops/9", match: false},
	// 		{solution: "solutions/raindrops/10", match: true},
	// 		{solution: "solutions/raindrops/11", match: true},
	// 		{solution: "solutions/raindrops/12", match: false},
	// 		{solution: "solutions/raindrops/13", match: false},
	// 		{solution: "solutions/raindrops/14", match: false},
	// 		{solution: "solutions/raindrops/15", match: false},
	// 		{solution: "solutions/raindrops/16", match: false},
	// 	},
	// },
	// {
	// 	patterns: []string{"solutions/isogram/9", "solutions/isogram/1"},
	// 	tests: []test{
	// 		{solution: "solutions/isogram/1", match: true},
	// 		{solution: "solutions/isogram/2", match: false},
	// 		{solution: "solutions/isogram/3", match: true},
	// 		{solution: "solutions/isogram/4", match: false},
	// 		{solution: "solutions/isogram/5", match: true},
	// 		{solution: "solutions/isogram/6", match: true},
	// 		{solution: "solutions/isogram/7", match: false},
	// 		{solution: "solutions/isogram/8", match: false},
	// 		{solution: "solutions/isogram/9", match: true},
	// 	},
	// },
}

const runOnly = ""

func TestPattern_Match(t *testing.T) {
	for _, group := range matchTests {
		var valid []*Pattern
		for _, folder := range group.patterns {
			pkg, err := getPackage(folder)
			if err != nil {
				t.Error(err)
				continue
			}
			perms := ExtractPatternPermutations(pkg)
			for _, perm := range perms {
				perm.Name = folder
				valid = append(valid, perm)
			}
		}

		for _, tt := range group.tests {
			if runOnly != "" && runOnly != tt.solution {
				continue
			}

			t.Run(tt.solution, func(t *testing.T) {
				pkg, err := getPackage(tt.solution)
				if err != nil {
					t.Error(err)
				}

				diff, percent, ok := DiffPatterns(valid, pkg)

				fmt.Printf("\t\t%s: %.1f Percent\n", tt.solution, percent*100)

				// check if it matches 100% and if it should
				assert.Equal(t, tt.match, ok, fmt.Sprintf("solution failed: %s\n%s", tt.solution, diff))

				// check if greated minPercent
				assert.Condition(t, func() bool {
					return tt.minPercent <= percent
				}, "solution failed: %s\n\tpercentage should be greater than %.2f. Is %.2f.",
					tt.solution, tt.minPercent, percent,
				)

				// check if less than maxPercent. Only if maxPercent != 0
				if tt.maxPercent != 0 {
					assert.Condition(t, func() bool {
						return percent < tt.maxPercent
					}, "solution failed: %s\n\tpercentage should be less than %.2f. Is %.2f.",
						tt.solution, tt.maxPercent, percent,
					)
				}
			})
		}
	}
}

func getPackage(path string) (*astrav.Package, error) {
	folder := astrav.NewFolder(http.Dir(path), "")
	packages, err := folder.ParseFolder()
	if err != nil {
		return nil, err
	}
	for _, pkg := range packages {
		return pkg, nil
	}

	return nil, errors.New("no go package found")
}

func Test_ExtractPatternPermutations(t *testing.T) {
	pkg, err := getPackage("solutions/permutations/1")
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		ast *astrav.Package
	}
	tests := []struct {
		name string
		args args
		want []Node
	}{
		{
			name: "permutations",
			args: args{ast: pkg},
			want: []Node{
				&Pattern{
					parentNode: parentNode{
						Nodes: []Node{
							&DefaultNode{parentNode{
								Nodes: []Node{
									&DefaultNode{parentNode{
										Nodes: []Node{
											&DefaultNode{parentNode{
												Nodes:    nil,
												NodeType: "*astrav.ValueSpec",
												Code:     "",
											}},
										},
										NodeType: "*astrav.GenDecl",
										Code:     "",
									}},
									&DefaultNode{parentNode{
										Nodes: []Node{
											&DefaultNode{parentNode{
												Nodes: []Node{
													&DefaultNode{parentNode{
														Nodes:    nil,
														NodeType: "*astrav.BasicLit",
														Code:     "",
													}},
												},
												NodeType: "*astrav.ValueSpec",
												Code:     "",
											}},
										},
										NodeType: "*astrav.GenDecl",
										Code:     "",
									}},
									&FuncDecl{
										parentNode: parentNode{
											Nodes: []Node{
												&DefaultNode{parentNode{
													Nodes: []Node{
														&DefaultNode{parentNode{
															Nodes:    nil,
															NodeType: "*astrav.FieldList",
															Code:     "",
														}},
													},
													NodeType: "*astrav.FuncType",
													Code:     "",
												}},
												&DefaultNode{parentNode{
													Nodes:    nil,
													NodeType: "*astrav.BlockStmt",
													Code:     "",
												}},
											},
											NodeType: "*astrav.FuncDecl",
											Code:     "",
										},
										Name: "Test",
									},
								},
								NodeType: "*astrav.File",
								Code:     "",
							}},
						},
						NodeType: "*astrav.Package",
					},
				},
				&Pattern{
					parentNode: parentNode{
						Nodes: []Node{
							&DefaultNode{parentNode{
								Nodes: []Node{
									&DefaultNode{parentNode{
										Nodes: []Node{
											&DefaultNode{parentNode{
												Nodes:    nil,
												NodeType: "*astrav.ValueSpec",
												Code:     "",
											}},
										},
										NodeType: "*astrav.GenDecl",
										Code:     "",
									}},
									&FuncDecl{
										parentNode: parentNode{
											Nodes: []Node{
												&DefaultNode{parentNode{
													Nodes: []Node{
														&DefaultNode{parentNode{
															Nodes:    nil,
															NodeType: "*astrav.FieldList",
															Code:     "",
														}},
													},
													NodeType: "*astrav.FuncType",
													Code:     "",
												}},
												&DefaultNode{parentNode{
													Nodes:    nil,
													NodeType: "*astrav.BlockStmt",
													Code:     "",
												}},
											},
											NodeType: "*astrav.FuncDecl",
											Code:     "",
										},
										Name: "Test",
									},
									&DefaultNode{parentNode{
										Nodes: []Node{
											&DefaultNode{parentNode{
												Nodes: []Node{
													&DefaultNode{parentNode{
														Nodes:    nil,
														NodeType: "*astrav.BasicLit",
														Code:     "",
													}},
												},
												NodeType: "*astrav.ValueSpec",
												Code:     "",
											}},
										},
										NodeType: "*astrav.GenDecl",
										Code:     "",
									}},
								},
								NodeType: "*astrav.File",
								Code:     "",
							}},
						},
						NodeType: "*astrav.Package",
					},
				},
				&Pattern{
					parentNode: parentNode{
						Nodes: []Node{
							&DefaultNode{parentNode{
								Nodes: []Node{
									&DefaultNode{parentNode{
										Nodes: []Node{
											&DefaultNode{parentNode{
												Nodes: []Node{
													&DefaultNode{parentNode{
														Nodes:    nil,
														NodeType: "*astrav.BasicLit",
														Code:     "",
													}},
												},
												NodeType: "*astrav.ValueSpec",
												Code:     "",
											}},
										},
										NodeType: "*astrav.GenDecl",
										Code:     "",
									}},
									&DefaultNode{parentNode{
										Nodes: []Node{
											&DefaultNode{parentNode{
												Nodes:    nil,
												NodeType: "*astrav.ValueSpec",
												Code:     "",
											}},
										},
										NodeType: "*astrav.GenDecl",
										Code:     "",
									}},
									&FuncDecl{
										parentNode: parentNode{
											Nodes: []Node{
												&DefaultNode{parentNode{
													Nodes: []Node{
														&DefaultNode{parentNode{
															Nodes:    nil,
															NodeType: "*astrav.FieldList",
															Code:     "",
														}},
													},
													NodeType: "*astrav.FuncType",
													Code:     "",
												}},
												&DefaultNode{parentNode{
													Nodes:    nil,
													NodeType: "*astrav.BlockStmt",
													Code:     "",
												}},
											},
											NodeType: "*astrav.FuncDecl",
											Code:     "",
										},
										Name: "Test",
									},
								},
								NodeType: "*astrav.File",
								Code:     "",
							}},
						},
						NodeType: "*astrav.Package",
					},
				},
				&Pattern{
					parentNode: parentNode{
						Nodes: []Node{
							&DefaultNode{parentNode{
								Nodes: []Node{
									&DefaultNode{parentNode{
										Nodes: []Node{
											&DefaultNode{parentNode{
												Nodes: []Node{
													&DefaultNode{parentNode{
														Nodes:    nil,
														NodeType: "*astrav.BasicLit",
														Code:     "",
													}},
												},
												NodeType: "*astrav.ValueSpec",
												Code:     "",
											}},
										},
										NodeType: "*astrav.GenDecl",
										Code:     "",
									}},
									&FuncDecl{
										parentNode: parentNode{
											Nodes: []Node{
												&DefaultNode{parentNode{
													Nodes: []Node{
														&DefaultNode{parentNode{
															Nodes:    nil,
															NodeType: "*astrav.FieldList",
															Code:     "",
														}},
													},
													NodeType: "*astrav.FuncType",
													Code:     "",
												}},
												&DefaultNode{parentNode{
													Nodes:    nil,
													NodeType: "*astrav.BlockStmt",
													Code:     "",
												}},
											},
											NodeType: "*astrav.FuncDecl",
											Code:     "",
										},
										Name: "Test",
									},
									&DefaultNode{parentNode{
										Nodes: []Node{
											&DefaultNode{parentNode{
												Nodes:    nil,
												NodeType: "*astrav.ValueSpec",
												Code:     "",
											}},
										},
										NodeType: "*astrav.GenDecl",
										Code:     "",
									}},
								},
								NodeType: "*astrav.File",
								Code:     "",
							}},
						},
						NodeType: "*astrav.Package",
					},
				},
				&Pattern{
					parentNode: parentNode{
						Nodes: []Node{
							&DefaultNode{parentNode{
								Nodes: []Node{
									&FuncDecl{
										parentNode: parentNode{
											Nodes: []Node{
												&DefaultNode{parentNode{
													Nodes: []Node{
														&DefaultNode{parentNode{
															Nodes:    nil,
															NodeType: "*astrav.FieldList",
															Code:     "",
														}},
													},
													NodeType: "*astrav.FuncType",
													Code:     "",
												}},
												&DefaultNode{parentNode{
													Nodes:    nil,
													NodeType: "*astrav.BlockStmt",
													Code:     "",
												}},
											},
											NodeType: "*astrav.FuncDecl",
											Code:     "",
										},
										Name: "Test",
									},
									&DefaultNode{parentNode{
										Nodes: []Node{
											&DefaultNode{parentNode{
												Nodes: []Node{
													&DefaultNode{parentNode{
														Nodes:    nil,
														NodeType: "*astrav.BasicLit",
														Code:     "",
													}},
												},
												NodeType: "*astrav.ValueSpec",
												Code:     "",
											}},
										},
										NodeType: "*astrav.GenDecl",
										Code:     "",
									}},
									&DefaultNode{parentNode{
										Nodes: []Node{
											&DefaultNode{parentNode{
												Nodes:    nil,
												NodeType: "*astrav.ValueSpec",
												Code:     "",
											}},
										},
										NodeType: "*astrav.GenDecl",
										Code:     "",
									}},
								},
								NodeType: "*astrav.File",
								Code:     "",
							}},
						},
						NodeType: "*astrav.Package",
					},
				},
				&Pattern{
					parentNode: parentNode{
						Nodes: []Node{
							&DefaultNode{parentNode{
								Nodes: []Node{
									&FuncDecl{
										parentNode: parentNode{
											Nodes: []Node{
												&DefaultNode{parentNode{
													Nodes: []Node{
														&DefaultNode{parentNode{
															Nodes:    nil,
															NodeType: "*astrav.FieldList",
															Code:     "",
														}},
													},
													NodeType: "*astrav.FuncType",
													Code:     "",
												}},
												&DefaultNode{parentNode{
													Nodes:    nil,
													NodeType: "*astrav.BlockStmt",
													Code:     "",
												}},
											},
											NodeType: "*astrav.FuncDecl",
											Code:     "",
										},
										Name: "Test",
									},
									&DefaultNode{parentNode{
										Nodes: []Node{
											&DefaultNode{parentNode{
												Nodes:    nil,
												NodeType: "*astrav.ValueSpec",
												Code:     "",
											}},
										},
										NodeType: "*astrav.GenDecl",
										Code:     "",
									}},
									&DefaultNode{parentNode{
										Nodes: []Node{
											&DefaultNode{parentNode{
												Nodes: []Node{
													&DefaultNode{parentNode{
														Nodes:    nil,
														NodeType: "*astrav.BasicLit",
														Code:     "",
													}},
												},
												NodeType: "*astrav.ValueSpec",
												Code:     "",
											}},
										},
										NodeType: "*astrav.GenDecl",
										Code:     "",
									}},
								},
								NodeType: "*astrav.File",
								Code:     "",
							}},
						},
						NodeType: "*astrav.Package",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractPatternPermutations(tt.args.ast)
			if len(got) != len(tt.want) {
				t.Fatalf("permutations() = %v, want %v", got, tt.want)
			}
			for i, pattern := range got {
				if !pattern.Match(tt.want[i]) {
					t.Errorf("permutations(%d) = %v, want %v", i, pattern, tt.want[i])
				}
			}
		})
	}
}
