package slices

// exercise comments
const (
	NoElse            = "go.slices.no_else"
	TwoReturns        = "go.slices.two_returns"
	OneIf             = "go.slices.one_if"
	OneConditionInIf  = "go.slices.one_condition_in_if"
	TwoConditionsInIf = "go.slices.two_conditions_in_if"
	NoSwitch          = "go.slices.no_switch"
	OneAppend         = "go.slices.one_append"
	NoColonSlice      = "go.slices.no_colon_slice"
	OneMake           = "go.slices.one_make"
	OneFor            = "go.slices.one_for"
	NoRange           = "go.slices.no_range"
	NotRecursive      = "go.slices.not_recursive"
	MakeHasThreeArgs  = "go.slices.make_has_three_args"
)

// Severity defines how severe a comment is. A sum over all comments of 5 means no approval.
// The maximum for a single comment is 5. A comment with that severity will block approval.
// When assigning the severity a good guideline is to ask: How many comments of similar severity
// should block approval?
// We can be very strict on automated comments since the student has a very fast feedback loop.
var severity = map[string]int{
	MissingEntryFunc:     5,
	FuncSignatureChanged: 5,
	MixtureRunesBytes:    3,
	RuneByteIndex:        2,
	DeclareWhenNeeded:    1,
	ErrorMsgFormat:       1,
	IncreaseOperator:     2,
	InvertIf:             3,
	NakedReturns:         3,
	ZeroValueOnErr:       2,
	RuneToByte:           3,
	ToStringConversion:   3,
	StringsSplitUsed:     5,
	ComparingBytes:       0,
	DefineEmptyErr:       1,
	ReturnOnError:        3,
	CaseInsensitive:      3,
	ExtraIfStringsEmpty:  0,
	MinSliceRuneConv:     2,
	CompBytesInDisguise:  5,
}
