package model

type Code struct {
	Id    int64  `gorm:"column:id;type:bigint;primary_key"`
	Phone string `gorm:"column:phone;type:varchar"`
	Code  int    `gorm:"column:code;type:int"`
}

func NewCodeModel() *Code {
	return &Code{}
}
