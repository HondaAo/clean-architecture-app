package entity

type Video struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Script   string `json:"script"`
	View     int    `json:"view"`
	Category string `json:"category"`
	Series   string `json:"series"`
	End      int    `json:"end"`
	Level    string `json:"level"`
}
