package entity

type Like struct {
	Id      int `json:"id" gorm:"autoIncrement"`
	UserId  int `json:"user_id"`
	VideoId int `json:"video_id"`
}
