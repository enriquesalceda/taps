package vo

import (
	"errors"
	"strings"
)

type Right string

type rights struct {
	ISSSTE Right
	IMSS   Right
	SEDENA Right
	SEMAR  Right
	PEMEX  Right
	SSA    Right
	NONE   Right
}

var Rights = rights{
	ISSSTE: "ISSSTE",
	IMSS:   "IMSS",
	SEDENA: "SEDENA",
	SEMAR:  "SEMAR",
	PEMEX:  "PEMEX",
	SSA:    "SSA",
	NONE:   "NONE",
}

func MustParseRights(rights string) Right {
	right, err := TryParseRights(rights)
	if err != nil {
		panic(err)
	}
	return right
}

func TryParseRights(rights string) (Right, error) {
	rights = strings.ToUpper(rights)
	switch rights {
	case "ISSSTE":
		return Rights.ISSSTE, nil
	case "IMSS":
		return Rights.IMSS, nil
	case "SEDENA":
		return Rights.SEDENA, nil
	case "SEMAR":
		return Rights.SEMAR, nil
	case "PEMEX":
		return Rights.PEMEX, nil
	case "SSA":
		return Rights.SSA, nil
	case "NONE":
		return Rights.NONE, nil
	default:
		var right Right
		return right, errors.New("invalid rights")
	}
}
