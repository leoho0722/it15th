package controllers

import (
	"errors"
	"it15th/database"
	"it15th/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Album API",
	})
}

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

	album := database.Album{
		ID:          uuid.New().String(),
		Title:       data.Title,
		Artist:      data.Artist,
		ReleaseDate: data.ReleaseDate,
	}

	err = repository.Create(album)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Create Failed，Error：" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create Success",
	})
}

func Get(c *gin.Context) {
	albumId := c.Query("id")

	result, err := repository.Read(albumId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get Success，No record found",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get Failed，Error：" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get Success",
		"albums":  result,
	})
}

func GetAll(c *gin.Context) {
	result, err := repository.ReadAll()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get Success，No record found",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get Failed，Error：" + err.Error(),
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
		ID          string `json:"id"`
		Title       string `json:"title"`
		Artist      string `json:"artist"`
		ReleaseDate string `json:"releaseDate"`
	}
	var data param
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Update Failed，Error：" + err.Error(),
		})
		return
	}

	result, err := repository.Read(data.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Update Success，No record found",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Update Failed，Error：" + err.Error(),
		})
		return
	}

	repository.Update(result, database.Album{
		Title:       data.Title,
		Artist:      data.Artist,
		ReleaseDate: data.ReleaseDate,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Success",
	})
}

func Delete(c *gin.Context) {
	type param struct {
		ID string `json:"id"`
	}
	var data param
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete Failed，Error：" + err.Error(),
		})
		return
	}

	repository.Delete(data.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Success",
	})
}

func DeleteAll(c *gin.Context) {
	err := repository.DeleteAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete All Failed，Error：" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete All Success",
	})
}
