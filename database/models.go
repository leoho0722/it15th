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
func (*Album) TableName() string {
	return "Album"
}

// Create 新增資料
//
//	Returns:
//	   error: 錯誤
//
//	// Create a new object of Album
//	album := &database.Album{
//	    ID:          uuid.New().String(),
//	    Title:       "The Modern Sound of Betty Carter",
//	    Artist:      "Betty Carter",
//	    ReleaseDate: "1960",
//	}
//
//	// Use variable album to call method Create()
//	err := album.Create()
//	if err != nil {
//	    panic(err)
//	}
func (a *Album) Create() error {
	result := DB.Create(&a)
	return result.Error
}

// Read 透過 id 讀取指定單筆資料
//
//	 Parameters:
//	   id: 資料的 id
//
//		// Create a new object of Album
//		album := &database.Album{
//		    ID:          uuid.New().String(),
//		    Title:       "The Modern Sound of Betty Carter",
//		    Artist:      "Betty Carter",
//		    ReleaseDate: "1960",
//		}
//
//		// Use variable album to call method Read()
//		album, err = album.Read(album.ID)
//
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//		    fmt.Println("No record found")
//		} else if err != nil {
//		    panic(err)
//		}
//		fmt.Println(*album)
func (a *Album) Read(id string) (*Album, error) {
	result := DB.Where("id = ?", id).First(&a)
	return a, result.Error
}

// ReadAll 讀取全部資料
//
//	 Returns:
//	   []Album: 全部資料
//	   error: 錯誤
//
//		// Create a new object of Album
//		album := &database.Album{
//		    ID:          uuid.New().String(),
//		    Title:       "The Modern Sound of Betty Carter",
//		    Artist:      "Betty Carter",
//		    ReleaseDate: "1960",
//		}
//
//		// Use variable album to call method ReadAll()
//		albums, err := album.ReadAll()
//
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//		    fmt.Println("No record found")
//		} else if err != nil {
//		    panic(err)
//		}
//		fmt.Println(albums)
func (a *Album) ReadAll() ([]Album, error) {
	var albums []Album
	result := DB.Find(&albums)
	return albums, result.Error
}

// Update 更新資料
//
//	 Parameters:
//	   updateData: 更新的資料
//
//		// Create a new object of Album
//		album := &database.Album{
//		    ID:          uuid.New().String(),
//		    Title:       "The Modern Sound of Betty Carter",
//		    Artist:      "Betty Carter",
//		    ReleaseDate: "1960",
//		}
//
//		// Use variable album to call method Update()
//		album.Update(database.Album{
//		    Title:       "LILAC",
//		    Artist:      "IU",
//		    ReleaseDate: "2021/03/25",
//		})
//		fmt.Println(*album)
func (a *Album) Update(updateData interface{}) {
	DB.Model(&a).Updates(updateData)
}

// Delete 刪除資料
//
//	// Create a new object of Album
//	album := &database.Album{
//	    ID:          uuid.New().String(),
//	    Title:       "The Modern Sound of Betty Carter",
//	    Artist:      "Betty Carter",
//	    ReleaseDate: "1960",
//	}
//
//	// Use variable album to call method Delete()
//	album.Delete()
func (a *Album) Delete() {
	DB.Delete(&a)
}

// DeleteAll 刪除全部資料
//
//	 Returns:
//	   error: 錯誤
//
//		// Create a new object of Album
//		album := &database.Album{
//		    ID:          uuid.New().String(),
//		    Title:       "The Modern Sound of Betty Carter",
//		    Artist:      "Betty Carter",
//		    ReleaseDate: "1960",
//		}
//
//		// Use variable album to call method DeleteAll()
//		err = album.DeleteAll()
//		if err != nil {
//		    panic(err)
//		}
//
//		albums, err = album.ReadAll()
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//		    fmt.Println("No record found")
//		} else if err != nil {
//		    panic(err)
//		}
//		fmt.Println(albums)
func (a *Album) DeleteAll() error {
	all, err := a.ReadAll()
	if err != nil {
		return err
	} else {
		for _, v := range all {
			DB.Delete(&v)
		}
	}
	return err
}
