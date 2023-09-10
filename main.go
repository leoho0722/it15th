package main

import (
	"errors"
	"fmt"
	"it15th/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func main() {
	database.ConnectDB()

	album := &database.Album{
		ID:          uuid.New().String(),
		Title:       "The Modern Sound of Betty Carter",
		Artist:      "Betty Carter",
		ReleaseDate: "1960",
	}

	err := album.Create()
	if err != nil {
		panic(err)
	}

	albums, err := album.ReadAll()

	// 判定錯誤是否為 gorm.ErrRecordNotFound 查無資料
	// 是的話，輸出 No record found，否則輸出報錯
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("No record found")
	} else if err != nil {
		panic(err)
	}
	fmt.Println(albums)

	album, err = album.Read(album.ID)

	// 判定錯誤是否為 gorm.ErrRecordNotFound 查無資料
	// 是的話，輸出 No record found，否則輸出報錯
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("No record found")
	} else if err != nil {
		panic(err)
	}
	fmt.Println(*album)

	album.Update(database.Album{
		Title:       "LILAC",
		Artist:      "IU",
		ReleaseDate: "2021/03/25",
	})
	fmt.Println(*album)

	d := &database.Album{}
	err = d.DeleteAll()
	if err != nil {
		panic(err)
	}

	albums, err = album.ReadAll()

	// 判定錯誤是否為 gorm.ErrRecordNotFound 查無資料
	// 是的話，輸出 No record found，否則輸出報錯
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("No record found")
	} else if err != nil {
		panic(err)
	}
	fmt.Println(albums)
}
