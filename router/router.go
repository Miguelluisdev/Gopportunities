package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//  Initialize router
	router := gin.Default()
	 initializeRoutes(router)
	
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
