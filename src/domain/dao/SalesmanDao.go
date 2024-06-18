package dao

import (
	"github.com/travellingsalesman/src/domain/model"
	"github.com/travellingsalesman/src/tserror"
)

type SalesmanDao interface {
	RegisterSalesPerson(person model.SalesPerson) *tserror.TSError
	LoginSalesPerson(personId string, loginArea model.Area) *tserror.TSError
	LogoutSalesPerson(personId string) *tserror.TSError
	GetSalesPerson(personId string) (*model.SalesPerson, *tserror.TSError)

	GetUniqueDestinations() (int, *tserror.TSError)
	GetAllSalesPersonLocations() (map[string][]*model.Area, *tserror.TSError)
}
