package model

type User struct {
	ID uint  `gorm:"AUTO_INCREMENT,primary_key"`
	CreatedAt int64  ` gorm:"column:create_at;not null" `
	Account  string  ` gorm:"column:account;not null;unique" `
	Password string  ` gorm:"not null"`
	NickName string  `gorm:"column:nick_name;not null"`
	Question string  `gorm:"not null"`
	Answer   string  `gorm:"not null"`
}

