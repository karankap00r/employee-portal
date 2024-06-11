package common

// Status represents the status of a request
type Status int

// Status enum values
const (
	Pending Status = iota
	Approved
	Rejected
)

// StatusToStringMap maps Status to string
var StatusToStringMap = map[Status]string{
	Pending:  "Pending",
	Approved: "Approved",
	Rejected: "Rejected",
}

// StringToStatusMap maps string to Status
var StringToStatusMap = map[string]Status{
	"Pending":  Pending,
	"Approved": Approved,
	"Rejected": Rejected,
}

// String returns the string representation of the Status
func (s Status) String() string {
	return StatusToStringMap[s]
}

// StatusFromString returns the Status from the given string
func StatusFromString(s string) Status {
	return StringToStatusMap[s]
}

// Value returns the integer value of the Status
func (s Status) Value() int {
	return int(s)
}

// IsValid checks if the Status is valid
func (s Status) IsValid() bool {
	return s >= Pending && s <= Rejected
}
