package guns

import "fmt"

func client() {
	ak47, _ := GetGun("ak47")
	musket, _ := GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun %s, \n Power: %d", g.getName(), g.getPower())
}
