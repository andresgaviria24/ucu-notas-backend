package application

import (
	"university/domain/dto"
	"university/domain/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	controller service.Service
}

func InitController(router *gin.Engine) {
	controller := Controller{
		controller: service.InitServiceImpl(),
	}

	student := router.Group("/students")

	student.GET("/ping",
		controller.Ping)
	student.GET("",
		controller.FindCurrentOrdersDetailByIdUser)
	student.POST("/:id",
		controller.CreateNote)
	student.DELETE("/:id",
		controller.CreateNote)
}

func (cdc *Controller) Ping(c *gin.Context) {

	pong := dto.Response{
		Status:      200,
		Description: "Pong",
		Message:     "",
	}

	c.JSON(http.StatusOK, pong)
}

func (cdc *Controller) FindCurrentOrdersDetailByIdUser(c *gin.Context) {

	response, students := cdc.controller.FindCurrentOrdersDetailByIdUser()

	c.JSON(response.Status, students)
}

func (cdc *Controller) CreateNote(c *gin.Context) {

	idStudent := c.Param("id")
	note := dto.CreateNote{}

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.Response{
			Status:      400,
			Description: "La estructura no es correcta",
			Message:     err.Error(),
		})
		return
	}

	response := cdc.controller.CreateNote(note, idStudent)

	c.JSON(response.Status, response)
}

func (cdc *Controller) DeleteNote(c *gin.Context) {

	idStudent := c.Param("id")
	note := dto.CreateNote{}

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.Response{
			Status:      400,
			Description: "La estructura no es correcta",
			Message:     err.Error(),
		})
		return
	}

	response := cdc.controller.DeleteNote(note, idStudent)

	c.JSON(response.Status, response)
}
