package entity

type User struct {
	Id        int    `json:"id" gorm:"autoIncrement"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Icon      string `json:"icon"`
	UpdatedAt string `json:"updated_at"`
}
