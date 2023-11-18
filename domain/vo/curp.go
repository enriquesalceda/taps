package vo

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const expectedCURPItemsLength = 10

var expectedPositionOfPresentItems = map[int]string{
	0: "ID",
	2: "LastLastName",
	3: "FirstLastName",
	4: "FirstName",
	5: "Gender",
	6: "DOB",
	7: "State",
	8: "Number",
}

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
	curp, err := TryParseCURP(rawCURP)
	if err != nil {
		panic(err)
	}
	return curp
}

func TryParseCURP(rawCURP string) (Curp, error) {
	curpData := strings.Split(rawCURP, "|")
	if len(curpData) != expectedCURPItemsLength {
		return Curp{}, errors.New(
			fmt.Sprintf("curp should have %d items, it has %d", expectedCURPItemsLength, len(curpData)),
		)
	}

	var notPresentStrings []string
	for i, v := range curpData {
		_, found := expectedPositionOfPresentItems[i]
		if v == "" && found {
			notPresentStrings = append(notPresentStrings, expectedPositionOfPresentItems[i])
		}
	}
	if len(notPresentStrings) > 0 {
		return Curp{}, errors.New(
			fmt.Sprintf("curp is not including: %s", strings.Join(notPresentStrings, ", ")),
		)
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
