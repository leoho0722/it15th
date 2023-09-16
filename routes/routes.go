package routes

import (
	"it15th/controllers"
	"it15th/network"

	"github.com/gin-gonic/gin"
)

type routerRawValue string

const (
	root           routerRawValue = "/"
	createAlbum    routerRawValue = "/album/create"
	getAlbum       routerRawValue = "/album"
	getAlbums      routerRawValue = "/albums"
	updateAlbum    routerRawValue = "/album/update"
	deleteAlbum    routerRawValue = "/album/delete"
	deleteAllAlbum routerRawValue = "/albums/delete"
)

func SetupRoute() {
	app := gin.Default()

	app.GET(string(root), controllers.Root)

	apiRG := app.Group("/api")
	{
		apiRG.POST(string(createAlbum), controllers.Create)
		apiRG.GET(string(getAlbum), controllers.Get)
		apiRG.GET(string(getAlbums), controllers.GetAll)
		apiRG.PUT(string(updateAlbum), controllers.Update)
		apiRG.DELETE(string(deleteAlbum), controllers.Delete)
		apiRG.DELETE(string(deleteAllAlbum), controllers.DeleteAll)
	}

	ip := network.GetLocalNetworkIPAddress()
	err := app.Run(ip + ":8080")
	if err != nil {
		panic(err)
	}
}
