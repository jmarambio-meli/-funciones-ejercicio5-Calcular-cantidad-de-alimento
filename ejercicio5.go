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
	var amount float64
	animalDog, msg := animal(tarantula)
	if msg != nil {
		fmt.Println(msg.Error())
	} else {
		amount += animalDog(10)
		fmt.Printf("La cantidad de comida que debe comprar es %.2f KG\n", amount)
	}

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
