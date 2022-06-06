package entity

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
	Posts    []Post `json:"posts" gorm:"foreignKey:userId"`
}
