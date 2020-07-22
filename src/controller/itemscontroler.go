package controller

import (
	"log"
	"net/http"

	"github.com/arshabbir/docker1/src/domain/store"
	"github.com/arshabbir/docker1/src/services"
	"github.com/arshabbir/docker1/src/utils"
	"github.com/gin-gonic/gin"
)

type itemcontroller struct {
	service services.ItemService
	router  *gin.Engine
}

type ItemController interface {
	Start()
	Create(c *gin.Context)
	Search(c *gin.Context)
	Ping(c *gin.Context)
	Love(c *gin.Context)
}

func NewItemsController() ItemController {

	return &itemcontroller{service: services.NewItemService(), router: gin.Default()}
}

func (ic *itemcontroller) Ping(c *gin.Context) {

	c.String(http.StatusOK, ic.service.PingItemService())

}

func (ic *itemcontroller) Create(c *gin.Context) {

	var item store.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiErr{Status: http.StatusBadRequest, Msg: "Bad Requet"})
		return

	}

	if err := ic.service.Create(&item); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return

	}
	c.JSON(http.StatusOK, utils.ApiErr{Status: http.StatusOK, Msg: "Iteam Creation Successfull"})
	return

}

func (ic *itemcontroller) Search(c *gin.Context) {

	//var item store.Item

	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, utils.ApiErr{Status: http.StatusBadRequest, Msg: "Bad Requet"})
		return

	}

	item, err := ic.service.Search(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return

	}

	//c.JSON()

	c.JSON(http.StatusOK, item)
	return
}

func (ic *itemcontroller) Love(c *gin.Context) {

	c.File("./love.gif")

	return

}

func (ic *itemcontroller) Start() {

	//Define Mappings

	log.Println("Starting the Gin router. . .")

	ic.router.GET("/ping", ic.Ping)
	ic.router.POST("/create", ic.Create)
	ic.router.GET("/search/:id", ic.Search)

	ic.router.GET("/love", ic.Love)

	ic.router.Run(":8080")

}
