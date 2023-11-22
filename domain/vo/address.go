package vo

import (
	"fmt"
	"strings"
)

type Address struct {
	StreetNumber string
	StreetName   string
	SuburbName   string
}

func TryNewAddress(streetNumber, streetName, suburbName string) (Address, error) {
	errors := []string{}

	if streetNumber == "" {
		errors = append(errors, "street number")
	}

	if streetName == "" {
		errors = append(errors, "street name")
	}

	if suburbName == "" {
		errors = append(errors, "suburb name")
	}

	if len(errors) > 0 {
		return Address{}, fmt.Errorf("address should have: %s", strings.Join(errors, ", "))
	}

	return Address{
		StreetNumber: streetNumber,
		StreetName:   streetName,
		SuburbName:   suburbName,
	}, nil
}

func MustNewAddress(streetNumber, streetName, suburbName string) Address {
	address, err := TryNewAddress(streetNumber, streetName, suburbName)
	if err != nil {
		panic(err)
	}

	return address
}
