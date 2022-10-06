package containers

import (
	"fmt"
	"strings"
)

// For testing purposes
type ContainerTest struct {
	values []interface{}
}

func (container ContainerTest) Empty() bool {
	return len(container.values) == 0
}

func (container ContainerTest) Size() int {
	return len(container.values)
}

func (container ContainerTest) Clear() {
	container.values = []interface{}{}
}

func (container ContainerTest) Values() []interface{} {
	return container.values
}

func (container ContainerTest) String() string {
	str := "ContainerTest\n"
	var values []string
	for _, value := range container.values {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// func TestGetSortedValuesInts(t *testing.T) {
// 	container := ContainerTest{}
// 	GetSortedValues(container, utils.IntComparator)
// 	container.values = []interface{}{5, 1, 3, 2, 4}
// 	values := GetSortedValues(container, utils.IntComparator)
// 	for i := 1; i < container.Size(); i++ {
// 		if values[i-1].(int) > values[i].(int) {
// 			t.Errorf("Not sorted!")
// 		}
// 	}
// }

// func TestGetSortedValuesStrings(t *testing.T) {
// 	container := ContainerTest{}
// 	GetSortedValues(container, utils.StringComparator)
// 	container.values = []interface{}{"g", "a", "d", "e", "f", "c", "b"}
// 	values := GetSortedValues(container, utils.StringComparator)
// 	for i := 1; i < container.Size(); i++ {
// 		if values[i-1].(string) > values[i].(string) {
// 			t.Errorf("Not sorted!")
// 		}
// 	}
// }
