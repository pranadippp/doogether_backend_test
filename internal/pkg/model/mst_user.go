package model

//model Users struct
type Users struct {
	ID       int    `json:"id" gorm:"autoIncrement;column:id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

//func for tablename
func (m *Users) TableName() string {
	return "user"
}
