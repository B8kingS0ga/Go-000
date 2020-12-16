package dao

type Dao struct {
}

func (d *Dao) GetUser() string {
	return "mike"
}

func NewDao() Dao {
	return Dao{}
}
