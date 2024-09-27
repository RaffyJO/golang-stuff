package car

import (
	"learn-protobuf/protogen/car"
	"log"

	"github.com/google/uuid"
)

func ValidateCar() {
	car := car.Car{
		CarId:           uuid.New().String(),
		Brand:           "Toyota",
		Model:           "123 Toyota",
		Price:           10000,
		ManufactureYear: 2021,
	}

	err := car.Validate()

	if err != nil {
		log.Fatalln("Valiidation failed", err)
	}

	log.Println(&car)
}
