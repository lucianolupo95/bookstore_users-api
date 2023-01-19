package app

import (
	"github.com/gin-gonic/gin" //El paquete gin es un framework Web para poder usar HTTP
)

var (
	router = gin.Default() //Crea un Router
)

func StartApplication() {
	mapUrls()           //Trae todas las rutas almacenadas en url_mapping.go
	router.Run(":8080") //Corre el Router en el Puerto 8080
}
