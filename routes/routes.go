package routes

import (
	"it15th/controllers"

	"github.com/gin-gonic/gin"
)

type routerRawValue string

const (
	createAlbum    routerRawValue = "/album/create"
	getAlbum       routerRawValue = "/album"
	getAlbums      routerRawValue = "/albums"
	updateAlbum    routerRawValue = "/album/update"
	deleteAlbum    routerRawValue = "/album/delete"
	deleteAllAlbum routerRawValue = "/albums/delete"
)

func SetupRoute() {
	app := gin.Default()

	apiRG := app.Group("/api")
	{
		apiRG.POST(string(createAlbum), controllers.Create)
		apiRG.GET(string(getAlbum), controllers.Get)
		apiRG.GET(string(getAlbums), controllers.GetAll)
		apiRG.PUT(string(updateAlbum), controllers.Update)
		apiRG.DELETE(string(deleteAlbum), controllers.Delete)
		apiRG.DELETE(string(deleteAllAlbum), controllers.DeleteAll)
	}

	err := app.Run(":8080")
	if err != nil {
		panic(err)
	}
}
