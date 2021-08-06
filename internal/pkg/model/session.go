package model

type Session struct {
	ID          int    `json:"id" gorm:"autoIncrement;column:ID"`
	UserID      int    `json:"user_id" gorm:"column:userID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Start       string `json:"start"`
	Duration    int    `json:"duration"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
}

func (m *Session) TableName() string {
	return "session"
}
