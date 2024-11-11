package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Left  *Node
	Value interface{}
	Right *Node
}

func decoding_into_structs() {
	b := []byte(`{"Value": "Father", "Left": {"Value": "Left child"}, "Right": {"Value": "Right
  child"}}`)
	var f Node
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f, f.Left, f.Right)
}
