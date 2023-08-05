package main

import "time"

type XLarge24 struct {
	Id            int64     // 1
	UserId        int64     // 2
	ProductId     int64     // 3
	FulfillId     int64     // 4
	DeliveryId    int64     // 5
	DiscountId    int64     // 6
	OrderStatus   string    // 7
	PaymentStatus string    // 8
	PayUpdatedAt  time.Time // 9
	OriginalPrice int       // 10
	DiscountPrice int       // 11
	FinalPrice    int       // 12
	Quantity      int       // 13
	CreatedAt     time.Time // 14
	CreatedBy     string    // 15
	CreatedById   int64     // 16
	UpdatedAt     time.Time // 17
	UpdatedBy     string    // 18
	UpdatedById   int64     // 19
	UpdateNote    string    // 20
	DeletedAt     time.Time // 21
	DeletedBy     string    // 22
	DeletedById   int64     // 23
	DeleteNote    string    // 24
}

func CreateXLargePtr() *XLarge24 {
	return &XLarge24{}
}

func PassXLargePtr(v *XLarge24) *XLarge24 {
	return v
}

func CreateXLargeVal() XLarge24 {
	return XLarge24{}
}

func PassXLargeVal(v XLarge24) XLarge24 {
	return v
}
