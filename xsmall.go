package main

type XSmall2 struct {
	Username string // 1
	Password string // 2
}

func CreateXSmallPtr() *XSmall2 {
	return &XSmall2{}
}

func PassXSmallPtr(v *XSmall2) *XSmall2 {
	return v
}

func CreateXSmallVal() XSmall2 {
	return XSmall2{}
}

func PassXSmallVal(v XSmall2) XSmall2 {
	return v
}
