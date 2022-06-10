package entity

type Friend struct {
	Uid int
	Fid int
	Sta int
}

func (f Friend) TableName() string {
	return "friend"
}
