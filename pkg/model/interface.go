package model

type Merger interface {
	Equal(interface{}) bool
}
