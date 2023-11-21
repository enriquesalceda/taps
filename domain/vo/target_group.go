package vo

import "errors"

type TargetGroup struct {
	SixToFiftyNineMonthsOld bool
	SixtyMonthsAndMore      bool
}

func TryParseTargetGroup(sixToFiftyNineMonthsOld bool, sixtyMonthsAndMore bool) (TargetGroup, error) {
	if sixToFiftyNineMonthsOld == sixtyMonthsAndMore {
		return TargetGroup{}, errors.New("target group values cannot be the same")
	}

	return TargetGroup{
		SixToFiftyNineMonthsOld: sixToFiftyNineMonthsOld,
		SixtyMonthsAndMore:      sixtyMonthsAndMore,
	}, nil
}

func MustParseTargetGroup(sixToFiftyNineMonthsOld bool, sixtyMonthsAndMore bool) TargetGroup {
	targetGroup, err := TryParseTargetGroup(sixToFiftyNineMonthsOld, sixtyMonthsAndMore)
	if err != nil {
		panic(err)
	}

	return targetGroup
}
