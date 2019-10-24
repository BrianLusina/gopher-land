// Given a name, return a string with the name
// if the name is not provided, return a placeholder string "you"

package twofer

func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
