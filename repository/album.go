package repository

import "it15th/database"

// Create 新增資料
//
//	Returns:
//	   error: 錯誤
//
//	Example:
//	// Create a new object of Album
//	album := database.Album{
//	    ID:          uuid.New().String(),
//	    Title:       "The Modern Sound of Betty Carter",
//	    Artist:      "Betty Carter",
//	    ReleaseDate: "1960",
//	}
//
//	// Use variable album to call repository function Create()
//	err := repository.Create(album)
//	if err != nil {
//	    panic(err)
//	}
func Create(a database.Album) error {
	result := database.DB.Create(&a)
	return result.Error
}

// Read 透過 id 讀取指定單筆資料
//
//	 Parameters:
//	   id: 資料的 id
//
//	 Example:
//		// Create a new object of Album
//		album := database.Album{
//		    ID:          uuid.New().String(),
//		    Title:       "The Modern Sound of Betty Carter",
//		    Artist:      "Betty Carter",
//		    ReleaseDate: "1960",
//		}
//
//		// Use variable album to call repository function Read()
//		album, err = repository.Read(album.ID)
//
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//		    fmt.Println("No record found")
//		} else if err != nil {
//		    panic(err)
//		}
func Read(id string) (database.Album, error) {
	var a database.Album
	result := database.DB.Where("id = ?", id).First(&a)
	return a, result.Error
}

// ReadAll 讀取全部資料
//
//	 Returns:
//	   []Album: 全部資料
//	   error: 錯誤
//
//	 Example:
//		// Call repository function ReadAll()
//		albums, err := repository.ReadAll()
//
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//		    fmt.Println("No record found")
//		} else if err != nil {
//		    panic(err)
//		}
func ReadAll() ([]database.Album, error) {
	var albums []database.Album
	result := database.DB.Find(&albums)
	return albums, result.Error
}

// Update 更新資料
//
//	 Parameters:
//	   updateData: 更新的資料
//
//	 Example:
//		// Create a new object of Album
//		album := database.Album{
//		    ID:          uuid.New().String(),
//		    Title:       "The Modern Sound of Betty Carter",
//		    Artist:      "Betty Carter",
//		    ReleaseDate: "1960",
//		}
//
//		// Use variable album to call repository function Update()
//		repository.Update(album, database.Album{
//		    Title:       "LILAC",
//		    Artist:      "IU",
//		    ReleaseDate: "2021/03/25",
//		})
func Update(a database.Album, updateData interface{}) {
	database.DB.Model(&a).Updates(updateData)
}

// Delete 刪除資料
//
//	Example:
//	// Create a new object of Album
//	album := database.Album{
//	    ID:          uuid.New().String(),
//	    Title:       "The Modern Sound of Betty Carter",
//	    Artist:      "Betty Carter",
//	    ReleaseDate: "1960",
//	}
//
//	// Use variable album to call repository function Delete()
//	repository.Delete(album.ID)
func Delete(id string) {
	database.DB.Delete(&database.Album{ID: id})
}

// DeleteAll 刪除全部資料
//
//	 Returns:
//	   error: 錯誤
//
//	 Example:
//		// Call repository function DeleteAll()
//		err := repository.DeleteAll()
//		if err != nil {
//		    panic(err)
//		}
func DeleteAll() error {
	all, err := ReadAll()
	if err != nil {
		return err
	} else {
		for _, v := range all {
			database.DB.Delete(&v)
		}
	}
	return err
}
