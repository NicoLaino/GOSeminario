package main

import "testing"

// Test Car Creation
func TestCarDealerShipAdd(t *testing.T) {
	d := NewCarDealerShip()
	c0 := d.FindByID(0)

	// Case car exists
	if c0 != nil {
		t.Error("El auto con ID=0 ya existe!")
	}

	d.Add(Car{0, "Etios"})
	c0 = d.FindByID(0)

	// Case creation failed
	if c0 == nil{
		t.Error("El auto con ID=0 no fue agregado!")
	}

	// Case incorrect name
	if c0.Name != "Etios" {
		t.Error("El auto con ID=0 no tiene el nombre correcto!")
	}
}