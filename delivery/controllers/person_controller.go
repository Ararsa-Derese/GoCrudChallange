package controllers

import (
	"net/http"
	"person/usecase"
	"person/domain"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PersonController struct {
	PersonUseCase usecase.PersonUsecase
}

func NewPersonController(p usecase.PersonUsecase) PersonController {
	return PersonController{
		PersonUseCase: p,
	}
}

func (p *PersonController) FindAll(c *gin.Context) {
	persons, err := p.PersonUseCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": domain.ErrInternalError.Error()})
		return
	}
	c.JSON(http.StatusOK, persons)
}

func (p *PersonController) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	person, err := p.PersonUseCase.GetById(uint32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": domain.ErrNotFound.Error()})
		return
	}
	c.JSON(http.StatusOK, person)
}

func (p *PersonController) Create(c *gin.Context) {
	var person domain.PersonConfig
	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrInvalidInput.Error()})
		return
	}
	persn := domain.NewPerson(person)
	newPerson, err := p.PersonUseCase.Create(persn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": domain.ErrInternalError.Error()})
		return
	}
	c.JSON(http.StatusCreated, newPerson)
}

func (p *PersonController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var person domain.PersonConfig
	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrInvalidInput.Error()})
		return
	}
	persn := domain.NewPerson(person)
	updatedPerson, err := p.PersonUseCase.Update(uint32(id),persn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedPerson)
}

func (p *PersonController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = p.PersonUseCase.Remove(uint32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Person deleted successfully"})
}
