package slices

// exercise comments
const (
	NotTwoReturns           = "go.slices.not_two_returns"
	NotOneIf                = "go.slices.not_one_if"
	NotOneIfCondition       = "go.slices.not_one_if_condition"
	NotTwoIfConditions      = "go.slices.not_two_if_conditions"
	Switch                  = "go.slices.switch"
	NotOneAppend            = "go.slices.not_one_append"
	ColonSlice              = "go.slices.colon_slice"
	NotOneMake              = "go.slices.not_one_make"
	NotOneFor               = "go.slices.not_one_for"
	NotThreeClauseLoop      = "go.slices.not_three_clause_loop"
	Range                   = "go.slices.range"
	Recursive               = "go.slices.recursive"
	MakeDoesntHaveThreeArgs = "go.slices.make_doesnt_have_three_args"
	NotOneEllipsis          = "go.slices.not_one_ellipsis"
)

// Severity defines how severe a comment is. A sum over all comments of 5 means no approval.
// The maximum for a single comment is 5. A comment with that severity will block approval.
// When assigning the severity a good guideline is to ask: How many comments of similar severity
// should block approval?
// We can be very strict on automated comments since the student has a very fast feedback loop.
var severity = map[string]int{
	NotTwoReturns:           5,
	NotOneIf:                5,
	NotOneIfCondition:       5,
	NotTwoIfConditions:      5,
	Switch:                  5,
	NotOneAppend:            5,
	ColonSlice:              5,
	NotOneMake:              5,
	NotOneFor:               5,
	NotThreeClauseLoop:      5,
	Range:                   5,
	Recursive:               5,
	MakeDoesntHaveThreeArgs: 5,
	NotOneEllipsis:          2,
}
