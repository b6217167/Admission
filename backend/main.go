package main

import (
	"fmt"

	"github.com/b6217167/sa-64-example/controller"
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
			//r.POST("/users", controller.CreateMedicationRacord)
			//protected.PATCH("/nurses", controller.UpdateNurse)
			//protected.DELETE("/nurses/:id", controller.DeleteNurse)

			protected.GET("/admissions", controller.ListAdmission)
			protected.GET("/admission/:id", controller.GetAdmission)
			protected.POST("/admissions", controller.CreateAdmission)
			//protected.PATCH("/admissions", controller.UpdateAdmission)
			//protected.DELETE("/admission/:id", controller.DeleteAdmission)

			protected.GET("/patients", controller.ListPatients)
			protected.GET("/patient/:id", controller.GetPatient)
			protected.POST("/patients", controller.CreatePatient)
			//protected.PATCH("/patients", controller.UpdatePatient)
			//protected.DELETE("/patient/:id", controller.DeletePatient)

			protected.GET("/rooms", controller.ListRoom)
			protected.GET("/rooms/:id", controller.GetRoom)
			protected.POST("/rooms", controller.CreateRoom)
			//protected.PATCH("/rooms", controller.UpdateRoom)
			//protected.DELETE("/rooms/:id", controller.DeleteRoom)

			protected.GET("/roomtypes", controller.ListRoomtypes)
			protected.GET("/roomtypes/:id", controller.GetRoomtype)
			protected.POST("/roomtypes", controller.CreateRoomtype)
			//protected.PATCH("/roomtypes", controller.UpdateRoomtype)
			//protected.DELETE("/roomtypes/:id", controller.DeleteRoomtype)

			protected.GET("/right_treatments", controller.ListRight_Treatments)
			protected.GET("/right_treatments/:id", controller.GetRight_Treatment)
			protected.POST("/right_treatments", controller.CreateRight_Treatment)
			//protected.PATCH("/right_treatments", controller.UpdateRight_Treatment)
			//protected.DELETE("/right_treatments/:id", controller.DeleteRight_Treatment)
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
