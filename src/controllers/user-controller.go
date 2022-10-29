package controllers

import (
	"example/go-api/src/models"
	repository "example/go-api/src/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReadUsers(c *gin.Context) {
	users, err := repository.ReadUsers()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "cannot list users: " + err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func ReadUser(c *gin.Context) {
	id := c.Param("id")

	new, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	user, err := repository.ReadUser(new)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "cannot find user",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "cannot bind json",
		})
		return
	}

	//! Business rules here
	//!
	//! -

	err = repository.CreateUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func UpdateUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "cannot bind json",
		})
		return
	}

	err = repository.UpdateUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "cannot update user: " + err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	err = repository.DeleteUser(newid)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "cannot delete user",
		})
		return
	}

	c.Status(http.StatusOK)
}
