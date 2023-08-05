package main

import "time"

type Medium8 struct {
	Id        int64     // 1
	Username  string    // 2
	Password  string    // 3
	Name      string    // 4
	Email     string    // 5
	Phone     string    // 6
	Address   string    // 7
	CreatedAt time.Time // 8
}

func CreateMediumPtr() *Medium8 {
	return &Medium8{}
}

func PassMediumPtr(v *Medium8) *Medium8 {
	return v
}

func CreateMediumVal() Medium8 {
	return Medium8{}
}

func PassMediumVal(v Medium8) Medium8 {
	return v
}
