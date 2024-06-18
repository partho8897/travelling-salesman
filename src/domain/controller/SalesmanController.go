package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/travellingsalesman/src/domain/model"
	"github.com/travellingsalesman/src/domain/service"
)

type SalesmanController interface {
	GetSalesmanRoutes(r *gin.Engine) *gin.Engine
}

type SalesmanControllerImpl struct {
	service service.SalesmanService
}

func NewSalesmanController(service service.SalesmanService) SalesmanController {
	return SalesmanControllerImpl{service: service}
}

func (controller SalesmanControllerImpl) GetSalesmanRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/sales/person/register", controller.RegisterSalesPerson)
	r.POST("/sales/person/:personId/login", controller.LoginSalesPerson)
	r.POST("/sales/person/:personId/logout", controller.LogoutSalesPerson)
	r.GET("/sales/person/:personId", controller.GetSalesPerson)

	r.GET("/sales/analytics/get_unique_destinations", controller.GetUniqueDestinations)
	r.GET("/sales/analytics/get_active_sales_person", controller.GetActiveSalesPerson)

	return r
}

func (controller SalesmanControllerImpl) GetSalesPerson(ctx *gin.Context) {
	id := ctx.Param("personId")
	salesPerson, tsError := controller.service.GetSalesPerson(id)
	if tsError != nil {
		ctx.JSON(500, gin.H{"error": tsError.Error()})
		return
	}

	ctx.JSON(200, salesPerson)
}

func (controller SalesmanControllerImpl) RegisterSalesPerson(ctx *gin.Context) {
	var salesPerson model.SalesPerson
	ctx.BindJSON(&salesPerson)

	tsError := controller.service.RegisterSalesPerson(salesPerson)
	if tsError != nil {
		ctx.JSON(500, gin.H{"error": tsError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "SalesPerson registered successfully"})
}

func (controller SalesmanControllerImpl) LoginSalesPerson(ctx *gin.Context) {
	var area model.Area
	id := ctx.Param("personId")
	ctx.BindJSON(&area)
	tsError := controller.service.LoginSalesPerson(id, area)
	if tsError != nil {
		ctx.JSON(500, gin.H{"error": tsError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "SalesPerson logged in successfully"})
}

func (controller SalesmanControllerImpl) LogoutSalesPerson(ctx *gin.Context) {
	id := ctx.Param("personId")
	tsError := controller.service.LogoutSalesPerson(id)
	if tsError != nil {
		ctx.JSON(500, gin.H{"error": tsError.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "SalesPerson logged out successfully"})
}

func (controller SalesmanControllerImpl) GetUniqueDestinations(ctx *gin.Context) {
	uniqueDestinations, tsError := controller.service.GetUniqueDestinations()
	if tsError != nil {
		ctx.JSON(500, gin.H{"error": tsError.Error()})
		return
	}

	ctx.JSON(200, uniqueDestinations)
}

func (controller SalesmanControllerImpl) GetActiveSalesPerson(ctx *gin.Context) {
	activeSalesPerson, tsError := controller.service.GetActiveSalesPersons()
	if tsError != nil {
		ctx.JSON(500, gin.H{"error": tsError.Error()})
		return
	}

	ctx.JSON(200, activeSalesPerson)
}
