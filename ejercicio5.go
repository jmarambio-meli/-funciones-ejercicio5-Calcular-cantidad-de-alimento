package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	animalDog, msg := animal(dog)
	if msg != nil {
		fmt.Println(msg.Error())
	}

	animalCat, msg := animal(cat)
	if msg != nil {
		fmt.Println(msg.Error())
	}

	animalHamster, msg := animal(hamster)
	if msg != nil {
		fmt.Println(msg.Error())
	}

	animalTarantula, msg := animal(tarantula)
	if msg != nil {
		fmt.Println(msg.Error())
	}

	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)
	amount += animalHamster(5)
	amount += animalTarantula(8)

	fmt.Println("amount of food to buy:", amount, "Kg")

}

func animal(animal string) (func(cantidad float64) float64, error) {
	switch animal {
	case dog:
		return animalDog, nil
	case cat:
		return animalCat, nil
	case hamster:
		return animalHamster, nil
	case tarantula:
		return animalTarantula, nil
	default:
		return nil, errors.New("animal no existe")
	}
}

func animalDog(cantidad float64) float64 {
	return cantidad * 10
}

func animalCat(cantidad float64) float64 {
	return cantidad * 5
}

func animalHamster(cantidad float64) float64 {
	return (cantidad * 250) / 1000 //conversion a KG
}

func animalTarantula(cantidad float64) float64 {
	return (cantidad * 150) / 1000 // conversion a KG
}
