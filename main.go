package main

import "fmt"

func main() {

	// Generate Car Dealership
	Uzcudun := NewCarDealerShip()

	// Add cars
	Uzcudun.Add(Car{0, "Etios"})
	Uzcudun.Add(Car{1, "Hilux"})
	Uzcudun.Add(Car{2, "Corolla"})

	// Print Cars
	Uzcudun.Print()

	// Look for specific car
	c0 := Uzcudun.FindByID (0)
	if c0 != nil {
		fmt.Println("Se encontro el ID = 0")
	}  else {
		fmt.Println("No se encontro el ID = 0")
	}

	// Delete specific car
	Uzcudun.DeleteByID(0)

	// Print Cars
	Uzcudun.Print()

	// Update specific car
	Uzcudun.Update(Car{1, "RAV4"})

	// Print Cars
	Uzcudun.Print()
}