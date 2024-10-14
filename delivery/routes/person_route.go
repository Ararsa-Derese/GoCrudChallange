package routes

import (
	"person/config"
	"person/delivery/controllers"
	"person/repository"
	"person/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewPersonRoute(env config.Env, r *gin.RouterGroup, timeout time.Duration) {
	personRepository := repository.NewPersonRepository()
	personUseCase := usecase.NewPersonUsecase(personRepository, timeout)
	personController := controllers.NewPersonController(personUseCase)
	r.GET("/", personController.FindAll)
	r.GET("/:id", personController.FindById)
	r.POST("/", personController.Create)
	r.PUT("/:id", personController.Update)
	r.DELETE("/:id", personController.Delete)
}
