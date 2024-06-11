package common

type HolidayState int

const (
	Active HolidayState = iota
	Inactive
)

var HolidayStateToStringMap = map[HolidayState]string{
	Active:   "Active",
	Inactive: "Inactive",
}

var StringToHolidayStateMap = map[string]HolidayState{
	"Active":   Active,
	"Inactive": Inactive,
}

func (s HolidayState) String() string {
	return HolidayStateToStringMap[s]
}

func HolidayStateFromString(s string) HolidayState {
	return StringToHolidayStateMap[s]
}

func (s HolidayState) Value() int {
	return int(s)
}

func (s HolidayState) IsValid() bool {
	return s >= Active && s <= Inactive
}
