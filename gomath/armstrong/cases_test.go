package armstrong

var testCases = []struct {
	description string
	input       int
	expected    bool
}{
	{
		description: "Zero is an Armstrong number",
		input:       0,
		expected:    true,
	},
	{
		description: "Single digit numbers are Armstrong numbers",
		input:       5,
		expected:    true,
	},
	{
		description: "There are no 2 digit Armstrong numbers",
		input:       10,
		expected:    false,
	},
	{
		description: "Three digit number that is an Armstrong number",
		input:       153,
		expected:    true,
	},
	{
		description: "Three digit number that is not an Armstrong number",
		input:       100,
		expected:    false,
	},
	{
		description: "Four digit number that is an Armstrong number",
		input:       9474,
		expected:    true,
	},
	{
		description: "Four digit number that is not an Armstrong number",
		input:       9475,
		expected:    false,
	},
	{
		description: "Seven digit number that is an Armstrong number",
		input:       9926315,
		expected:    true,
	},
	{
		description: "Seven digit number that is not an Armstrong number",
		input:       9926314,
		expected:    false,
	},
}
