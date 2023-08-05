package main

type Small4 struct {
	Id    int64  // 1
	Name  string // 2
	Email string // 3
	Phone string // 4
}

func CreateSmallPtr() *Small4 {
	return &Small4{}
}

func PassSmallPtr(v *Small4) *Small4 {
	return v
}

func CreateSmallVal() Small4 {
	return Small4{}
}

func PassSmallVal(v Small4) Small4 {
	return v
}
