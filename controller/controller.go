package controller

import (
	"net/http"

	model "gin-crud/entity"
	services "gin-crud/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service *services.Services
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func NewRouter(service *services.Services) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.New() 	//router := gin.Default()
	router.Use(gin.Logger())
	router.SetTrustedProxies(nil)


	c := Controller{Service: service}

	book := router.Group("/book")
	book.GET("/", c.GetAllBooks)
	book.GET("/:id", c.GetBook)
	book.POST("/", c.CreateBook)
	book.PUT("/:id", c.UpdateBook)
	book.DELETE("/:id", c.DeleteBook)

	router.GET("/", c.HelloWorld)

	return router
}

func (c *Controller) HelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello HR!")
}

func (c *Controller) GetAllBooks(ctx *gin.Context) {
	var response Response

	books, err := c.Service.GetAllBooks()
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = books
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) GetBook(ctx *gin.Context) {
	var response Response
	id := ctx.Param("id")

	book, err := c.Service.GetBook(id)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = book
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) CreateBook(ctx *gin.Context) {
	var response Response
	var request model.Book
	ctx.Header("Content-Type", "application/json")

	if err := ctx.BindJSON(&request); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	book, code, err := c.Service.CreateBook(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = code
		ctx.JSON(code, response)
		return
	}

	response.Data = book
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) UpdateBook(ctx *gin.Context) {
	var response Response
	var request model.Book
	ctx.Header("Content-Type", "application/json")
	id := ctx.Param("id")

	if err := ctx.BindJSON(&request); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	request.ID = id

	book, err := c.Service.UpdateBook(request)
	if err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = book
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) DeleteBook(ctx *gin.Context) {
	var response Response
	id := ctx.Param("id")

	if err := c.Service.DeleteBook(id); err != nil {
		response.Data = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = nil
	response.Status = http.StatusOK

	ctx.JSON(http.StatusOK, response)
}
