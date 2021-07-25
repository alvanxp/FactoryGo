package main

import (
	"fmt"

	".main.go/factory"
)

const carsAmount = 100

func main() {
	factory := factory.New()

	//Hint: change appropriately for making factory give each vehicle once assembled, even though the others have not been assembled yet,
	//each vehicle delivered to main should display testinglogs and assemblelogs with the respective vehicle id
	out := factory.StartAssemblingProcess(carsAmount)
	for vehicleAssemblied := range out {
		fmt.Println("Vehicle Id: ", vehicleAssemblied.Id)
		fmt.Printf("Testinglogs : \n%s\n", vehicleAssemblied.TestingLog)
		fmt.Printf("AssembleLog : \n%s\n", vehicleAssemblied.AssembleLog)
		fmt.Println("---------------")
	}

}
