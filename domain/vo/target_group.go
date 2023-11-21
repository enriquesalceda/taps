package vo

import "errors"

type TargetGroup struct {
	SixToFiftyNineMonthsOld bool
	SixtyMonthsAndMore      bool
}

func TryBuildTargetGroup(sixToFiftyNineMonthsOld bool, sixtyMonthsAndMore bool) (TargetGroup, error) {
	if sixToFiftyNineMonthsOld == sixtyMonthsAndMore {
		return TargetGroup{}, errors.New("target group values cannot be the same")
	}

	return TargetGroup{
		SixToFiftyNineMonthsOld: sixToFiftyNineMonthsOld,
		SixtyMonthsAndMore:      sixtyMonthsAndMore,
	}, nil
}

func MustBuildTargetGroup(sixToFiftyNineMonthsOld bool, sixtyMonthsAndMore bool) TargetGroup {
	targetGroup, err := TryBuildTargetGroup(sixToFiftyNineMonthsOld, sixtyMonthsAndMore)
	if err != nil {
		panic(err)
	}

	return targetGroup
}
