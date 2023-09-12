package database

type Album struct {
	ID          string `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	ReleaseDate string `json:"releaseDate"`
}

// TableName 設定資料表名稱
//
//	Returns:
//	  string: 資料表名稱
func (Album) TableName() string {
	return "Album"
}
