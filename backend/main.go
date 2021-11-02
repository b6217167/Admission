package main

import (
	"fmt"

	controller "github.com/b6217167/sa-64-example/controller/Admission"
	"github.com/b6217167/sa-64-example/middlewares"

	"github.com/b6217167/sa-64-example/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	/*gin.SetMode(gin.ReleaseMode)*/
	r := gin.Default()
	r.Use(CORSMiddleware())

	// User Routes
	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// User Routes
			protected.GET("/nurses", controller.ListNurses)
			protected.GET("/nurse/:id", controller.GetNurse)
			protected.POST("/nurses", controller.CreateNurse)

			protected.GET("/admissions", controller.ListAdmission)
			protected.GET("/admission/:id", controller.GetAdmission)
			protected.POST("/admissions", controller.CreateAdmission)

			protected.GET("/patients", controller.ListPatients)
			protected.GET("/patient/:id", controller.GetPatient)
			protected.POST("/patients", controller.CreatePatient)

			protected.GET("/rooms", controller.ListRoom)
			protected.GET("/rooms/:id", controller.GetRoom)
			protected.POST("/rooms", controller.CreateRoom)

			protected.GET("/roomtypes", controller.ListRoomtypes)
			protected.GET("/roomtypes/:id", controller.GetRoomtype)
			protected.POST("/roomtypes", controller.CreateRoomtype)

			protected.GET("/right_treatments", controller.ListRight_Treatments)
			protected.GET("/right_treatments/:id", controller.GetRight_Treatment)
			protected.POST("/right_treatments", controller.CreateRight_Treatment)

		}
	}

	//User Routes
	r.POST("/nurses/create", controller.CreateNurse)

	//Authentication Routes
	r.POST("/login", controller.Login)

	//Run the server
	r.Run()
	fmt.Print("Test main.go")

}
func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}
