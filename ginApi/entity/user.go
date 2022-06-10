package entity

type User struct {
	Id    int
	Name  string
	Pass  string
	State int
}

func (u User) TableName() string {
	return "user"
}
