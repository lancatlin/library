package main

import (
	"github.com/jinzhu/gorm"
)

const (
	StatusInside = iota
	StatusLending
	StatusMissing
)

type Catalog struct {
	gorm.Model
	Name           string
	Authors        []*Author
	Publisher      *Publisher
	Year           int
	Classification *Classification
	ClassificationNumber string
	Items []*Item
}

type Author struct {
	gorm.Model
	Name string
}

type Publisher struct {
	gorm.Model
	Name string
}

type Classification struct {
	gorm.Model
	Name string
}

type Item struct {
	gorm.Model
	ID string `gorm:"primary_key"`
	Catalog *Catalog
	Status int
	NewBook string
}

