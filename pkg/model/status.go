package model

// Status to record the status of a item
type Status int

func (s Status) String() string {
	switch s {
	case StatusUnknown:
		return "未設定"
	case StatusInside:
		return "館內"
	case StatusLending:
		return "借出中"
	case StatusMissing:
		return "遺失"
	}
	return ""
}

const (
	// StatusUnknown is the default status
	StatusUnknown Status = iota
	// StatusInside means the book is in the library
	StatusInside
	// StatusLending means the book is lending by someone
	StatusLending
	// StatusMissing means the book is missing
	StatusMissing
)
