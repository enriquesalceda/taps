package vo

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const expectedCURPItemsLength = 10

type Curp struct {
	ID            string
	LastLastName  string
	FirstLastName string
	FirstName     string
	Gender        string
	DOB           string
	State         string
	Number        int
}

func MustParseCURP(rawCURP string) Curp {
	curp, err := ParseCURP(rawCURP)
	if err != nil {
		panic(err)
	}
	return curp
}

func ParseCURP(rawCURP string) (Curp, error) {
	curpData := strings.Split(rawCURP, "|")
	if len(curpData) != expectedCURPItemsLength {
		return Curp{}, errors.New(fmt.Sprintf("curp should have 10 items, it has %d", len(curpData)))
	}

	number, err := strconv.Atoi(curpData[8])
	if err != nil {
		return Curp{}, err
	}

	return Curp{
		ID:            curpData[0],
		LastLastName:  curpData[2],
		FirstLastName: curpData[3],
		FirstName:     curpData[4],
		Gender:        curpData[5],
		DOB:           curpData[6],
		State:         curpData[7],
		Number:        number,
	}, nil
}
