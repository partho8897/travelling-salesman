package service

import (
	"fmt"
	"github.com/travellingsalesman/src/domain/dao"
	"github.com/travellingsalesman/src/domain/model"
	"github.com/travellingsalesman/src/tserror"
)

type SalesmanServiceImpl struct {
	dao dao.SalesmanDao
}

func NewSalesmanServiceImpl(dao dao.SalesmanDao) SalesmanService {
	return &SalesmanServiceImpl{
		dao: dao,
	}
}

func (service SalesmanServiceImpl) RegisterSalesPerson(person model.SalesPerson) *tserror.TSError {
	salesPerson, getSalesPersonErr := service.dao.GetSalesPerson(person.Id)
	if getSalesPersonErr != nil {
		return getSalesPersonErr
	}

	if salesPerson != nil {
		return tserror.ERR_SALES_PERSON_ALREADY_REGISTERED.WithDetails(fmt.Sprintf("Salesperson with id: %s already registered", person.Id))
	}

	return service.dao.RegisterSalesPerson(person)
}

func (service SalesmanServiceImpl) LoginSalesPerson(personId string, loginArea model.Area) *tserror.TSError {
	salesPerson, getSalesPersonErr := service.dao.GetSalesPerson(personId)
	if getSalesPersonErr != nil {
		return getSalesPersonErr
	}

	if salesPerson == nil {
		return tserror.ERR_SALES_PERSON_ALREADY_REGISTERED.WithDetails(fmt.Sprintf("Salesperson with id: %s doesn't exist", personId))
	}

	return service.dao.LoginSalesPerson(personId, loginArea)
}

func (service SalesmanServiceImpl) LogoutSalesPerson(personId string) *tserror.TSError {
	salesPerson, getSalesPersonErr := service.dao.GetSalesPerson(personId)
	if getSalesPersonErr != nil {
		return getSalesPersonErr
	}

	if salesPerson == nil {
		return tserror.ERR_SALES_PERSON_ALREADY_REGISTERED.WithDetails(fmt.Sprintf("Salesperson with id: %s doesn't exist", personId))
	}

	return service.dao.LogoutSalesPerson(personId)
}

func (service SalesmanServiceImpl) GetSalesPerson(personId string) (*model.SalesPerson, *tserror.TSError) {
	return service.dao.GetSalesPerson(personId)
}

func (service SalesmanServiceImpl) GetUniqueDestinations() (*model.GetUniqueDestinationResponse, *tserror.TSError) {
	uniqueDestination, getUniqueDestinationErr := service.dao.GetUniqueDestinations()
	if getUniqueDestinationErr != nil {
		return nil, getUniqueDestinationErr
	}

	return &model.GetUniqueDestinationResponse{
		UniqueDestinations: uniqueDestination,
	}, nil
}

func (service SalesmanServiceImpl) GetActiveSalesPersons() ([]*model.SalesPerson, *tserror.TSError) {
	destinations, getDestinationErr := service.dao.GetAllSalesPersonLocations()
	if getDestinationErr != nil {
		return nil, getDestinationErr
	}

	activeSalesPersons := make([]*model.SalesPerson, 0)

	for personId, areas := range destinations {
		if len(areas) > 1 {
			salesPerson, getSalesPersonErr := service.dao.GetSalesPerson(personId)
			if getSalesPersonErr != nil {
				return nil, getSalesPersonErr
			}

			activeSalesPersons = append(activeSalesPersons, salesPerson)
		}
	}

	return activeSalesPersons, nil
}
