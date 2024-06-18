package router

import (
	"github.com/gin-gonic/gin"
	"github.com/travellingsalesman/src/domain/controller"
)

func NewRouter(salesmanRouter controller.SalesmanController) (*gin.Engine, error) {
	router := gin.Default()
	salesmanRouter.GetSalesmanRoutes(router)

	return router, nil
}
