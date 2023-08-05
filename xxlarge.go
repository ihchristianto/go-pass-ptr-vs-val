package main

import "time"

type XXLarge32 struct {
	Id               int64     // 1
	UserId           int64     // 2
	ProductId        int64     // 3
	FulfillId        int64     // 4
	FulfillStatus    string    // 5
	FUlfillNote      string    // 6
	DeliveryId       int64     // 7
	DeliveryStatus   string    // 8
	DiscountId       int64     // 9
	OrderStatus      string    // 10
	PaymentStatus    string    // 11
	PayUpdatedAt     time.Time // 12
	OriginalPrice    int       // 13
	DiscountPrice    int       // 14
	FinalPrice       int       // 15
	NormalQuantity   int       // 16
	DiscountQuantity int       // 17
	TotalQuantity    int       // 18
	CreatedAt        time.Time // 19
	CreatedBy        string    // 20
	CreatedById      int64     // 21
	CreateNote       string    // 22
	UpdatedAt        time.Time // 23
	UpdatedBy        string    // 24
	UpdatedById      int64     // 25
	UpdateNote       string    // 26
	DeletedAt        time.Time // 27
	DeletedBy        string    // 28
	DeletedById      int64     // 29
	DeleteNote       string    // 30
	Data1            string    // 31
	Data2            string    // 32
}

func CreateXXLargePtr() *XXLarge32 {
	return &XXLarge32{}
}

func PassXXLargePtr(v *XXLarge32) *XXLarge32 {
	return v
}

func CreateXXLargeVal() XXLarge32 {
	return XXLarge32{}
}

func PassXXLargeVal(v XXLarge32) XXLarge32 {
	return v
}
