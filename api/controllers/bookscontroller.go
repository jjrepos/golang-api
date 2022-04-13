package controllers

import (
	"net/http"
	"strconv"

	"jjrepos/gonang/api/dtos"
	"jjrepos/gonang/api/services"

	"github.com/gin-gonic/gin"
)

func FindBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, services.FindBooks())
}

func CreateBook(ctx *gin.Context) {
	var input dtos.CreateBookDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created := services.CreateBook(input)
	ctx.JSON(http.StatusCreated, created)
}

func UpdateBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No such book exists"})
		return
	}
	var input dtos.UpdateBookDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book, err := services.UpdateBook(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, book)
}

func FindBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No such book exists"})
		return
	}
	book, err := services.FindBook(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func DeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No such book exists"})
		return
	}
	if err = services.DeleteBook(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}
