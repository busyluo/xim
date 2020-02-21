package entity

const (
	GenderUnknown = 0 // 性别位置
	GenderMale    = 1 // 男性
	GenderFemale  = 2 // 女性
	GenderOther   = 3 // 其他
)

type User struct {
	Id       int32
	Phone    string
	Nickname string
	Email    string
	Gender   int8
}
