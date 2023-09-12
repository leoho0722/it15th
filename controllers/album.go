package controllers

import (
	"errors"
	"it15th/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Create(c *gin.Context) {
	type param struct {
		Title       string `json:"title"`
		Artist      string `json:"artist"`
		ReleaseDate string `json:"releaseDate"`
	}
	var data param
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Create Failed" + err.Error(),
		})
		return
	}

	album := &database.Album{
		ID:          uuid.New().String(),
		Title:       data.Title,
		Artist:      data.Artist,
		ReleaseDate: data.ReleaseDate,
	}

	err = album.Create()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Create Failed" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create Success",
	})
}

func Get(c *gin.Context) {
	albumId := c.Query("id")

	album := &database.Album{}
	result, err := album.Read(albumId)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"message": "No record found",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get Failed" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get Success",
		"albums":  result,
	})
}

func GetAll(c *gin.Context) {
	album := &database.Album{}
	result, err := album.ReadAll()

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"message": "No record found",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get Failed" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get Success",
		"albums":  result,
	})
}

func Update(c *gin.Context) {
	type param struct {
		Title       string `json:"title"`
		Artist      string `json:"artist"`
		ReleaseDate string `json:"releaseDate"`
	}
	var data param
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Update Failed" + err.Error(),
		})
		return
	}

	album := &database.Album{}
	album.Update(database.Album{
		Title:       data.Title,
		Artist:      data.Artist,
		ReleaseDate: data.ReleaseDate,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Success",
	})
}

func Delete(c *gin.Context) {
	albumId := c.Query("id")

	album := &database.Album{}
	album, err := album.Read(albumId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete Success，No record found",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete Failed" + err.Error(),
		})
		return
	}

	album.Delete()

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Success",
	})
}

func DeleteAll(c *gin.Context) {
	album := &database.Album{}
	err := album.DeleteAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete All Failed" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete All Success",
	})
}
