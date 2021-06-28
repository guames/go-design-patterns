package factorymethod

import "fmt"

//ExecuteFactoryMethod method to demonstrate the factory method.
func ExecuteFactoryMethod() {
	fmt.Println("Executing Factory Method Pattern.")

	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
