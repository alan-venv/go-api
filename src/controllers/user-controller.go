package controllers

import (
	"example/go-api/src/models"
	"example/go-api/src/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Repository repositories.IUserRepository
}

func (self UserController) ReadAll(c *gin.Context) {
	users, err := self.Repository.ReadAll()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "cannot list users: " + err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func (self UserController) Read(c *gin.Context) {
	id := c.Param("id")

	user, err := self.Repository.Read(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "cannot find user!",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (self UserController) Create(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":   "cannot bind json",
			"details": err.Error(),
		})
		return
	}

	//! Business rules here
	//! Duplicated records, etc.
	//! -

	err = self.Repository.Create(user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (self UserController) Update(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Cannot bind json: " + err.Error(),
		})
		return
	}

	err = self.Repository.Update(user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}

func (self UserController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := self.Repository.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "cannot find or delete user",
		})
		return
	}

	c.Status(http.StatusOK)
}
