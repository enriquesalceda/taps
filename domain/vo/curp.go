package vo

import (
	"strconv"
	"strings"
)

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

func ParseCURP(rawCURP string) Curp {
	curpData := strings.Split(rawCURP, "|")

	number, _ := strconv.Atoi(curpData[8])

	return Curp{
		ID:            curpData[0],
		LastLastName:  curpData[2],
		FirstLastName: curpData[3],
		FirstName:     curpData[4],
		Gender:        curpData[5],
		DOB:           curpData[6],
		State:         curpData[7],
		Number:        number,
	}
}
