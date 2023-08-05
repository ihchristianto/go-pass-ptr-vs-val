package main

import "time"

type Large16 struct {
	Id          int64     // 1
	UserId      int64     // 2
	ProductId   int64     // 3
	DeliveryId  int64     // 4
	DiscountId  int64     // 5
	FinalPrice  int       // 6
	Quantity    int       // 7
	CreatedAt   time.Time // 8
	CreatedBy   string    // 9
	CreatedById int64     // 10
	UpdatedAt   time.Time // 11
	UpdatedBy   string    // 12
	UpdatedById int64     // 13
	DeletedAt   time.Time // 14
	DeletedBy   string    // 15
	DeletedById int64     // 16
}

func CreateLargePtr() *Large16 {
	return &Large16{}
}

func PassLargePtr(v *Large16) *Large16 {
	return v
}

func CreateLargeVal() Large16 {
	return Large16{}
}

func PassLargeVal(v Large16) Large16 {
	return v
}
