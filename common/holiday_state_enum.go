package common

// HolidayState represents the state of a holiday
type HolidayState int

// HolidayState enum values
const (
	Active HolidayState = iota
	Inactive
)

// HolidayStateToStringMap maps HolidayState to string
var HolidayStateToStringMap = map[HolidayState]string{
	Active:   "Active",
	Inactive: "Inactive",
}

// StringToHolidayStateMap maps string to HolidayState
var StringToHolidayStateMap = map[string]HolidayState{
	"Active":   Active,
	"Inactive": Inactive,
}

// String returns the string representation of the HolidayState
func (s HolidayState) String() string {
	return HolidayStateToStringMap[s]
}

// HolidayStateFromString returns the HolidayState from the given string
func HolidayStateFromString(s string) HolidayState {
	return StringToHolidayStateMap[s]
}

// Value returns the integer value of the HolidayState
func (s HolidayState) Value() int {
	return int(s)
}

// IsValid checks if the HolidayState is valid
func (s HolidayState) IsValid() bool {
	return s >= Active && s <= Inactive
}
