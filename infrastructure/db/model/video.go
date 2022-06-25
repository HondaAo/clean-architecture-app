package model

type Video struct {
	Id        int    `json:"id" gorm:"autoIncrement"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	Script    string `json:"script"`
	View      int    `json:"view"`
	Category  string `json:"category"`
	Series    string `json:"series"`
	EndTime   int    `json:"end_time"`
	Level     string `json:"level"`
	CreatedAt string `json:"created_at"`
}
