package main 

import "fmt"

// CarDealerShip definition
type CarDealerShip struct {
	cars map[int]*Car 
}

// Car definition
type Car struct {
	ID int
	Name string
}

// Constructor
func NewCarDealerShip() CarDealerShip{
	cars := make(map[int]*Car)
	return CarDealerShip{
		cars, 
	}
}

// Add a Car
func (c CarDealerShip) Add(car Car) {
	c.cars[car.ID] = &car
}

// Print Cars
func (c CarDealerShip) Print(){
	for _, v := range c.cars {
		fmt.Printf("[%v]\t%v\n", v.ID, v.Name)
	}
}

// Find Car by ID
func (c CarDealerShip) FindByID (ID int) *Car{
	return c.cars[ID]
}

// Delete Car by ID
func (c CarDealerShip) DeleteByID (ID int) {
	delete(c.cars, ID)
}

// Update Car
func (c CarDealerShip) Update (car Car) {
	c.cars[car.ID] = &car
}