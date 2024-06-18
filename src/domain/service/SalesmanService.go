package service

import (
	"github.com/travellingsalesman/src/domain/model"
	"github.com/travellingsalesman/src/tserror"
)

type SalesmanService interface {
	RegisterSalesPerson(person model.SalesPerson) *tserror.TSError
	LoginSalesPerson(personId string, loginArea model.Area) *tserror.TSError
	LogoutSalesPerson(personId string) *tserror.TSError
	GetSalesPerson(personId string) (*model.SalesPerson, *tserror.TSError)

	GetUniqueDestinations() (*model.GetUniqueDestinationResponse, *tserror.TSError)
	GetActiveSalesPersons() ([]*model.SalesPerson, *tserror.TSError)
}
