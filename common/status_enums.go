package common

type Status int

const (
	Pending Status = iota
	Approved
	Rejected
)

var StatusToStringMap = map[Status]string{
	Pending:  "Pending",
	Approved: "Approved",
	Rejected: "Rejected",
}

var StringToStatusMap = map[string]Status{
	"Pending":  Pending,
	"Approved": Approved,
	"Rejected": Rejected,
}

func (s Status) String() string {
	return StatusToStringMap[s]
}

func StatusFromString(s string) Status {
	return StringToStatusMap[s]
}

func (s Status) Value() int {
	return int(s)
}

func (s Status) IsValid() bool {
	return s >= Pending && s <= Rejected
}
