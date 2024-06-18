package dao

import (
	"github.com/travellingsalesman/src/domain/model"
	"github.com/travellingsalesman/src/tserror"
)

type SalesmanDaoImpl struct {
	salesPerson           map[string]*model.SalesPerson
	salesPersonsLocations map[string][]*model.Area
	uniqueDestinations    map[model.Area]interface{}
}

func NewSalesmanDaoImpl() SalesmanDao {
	return &SalesmanDaoImpl{
		salesPerson:           make(map[string]*model.SalesPerson),
		salesPersonsLocations: make(map[string][]*model.Area),
		uniqueDestinations:    make(map[model.Area]interface{}),
	}
}

func (dao SalesmanDaoImpl) RegisterSalesPerson(person model.SalesPerson) *tserror.TSError {
	dao.salesPerson[person.Id] = &person

	return nil
}

func (dao SalesmanDaoImpl) GetSalesPerson(personId string) (*model.SalesPerson, *tserror.TSError) {
	salesPerson, isPresent := dao.salesPerson[personId]
	if isPresent {
		return salesPerson, nil
	}

	return nil, nil
}

func (dao SalesmanDaoImpl) LoginSalesPerson(personId string, loginArea model.Area) *tserror.TSError {
	areas, isPresent := dao.salesPersonsLocations[personId]
	if !isPresent {
		areas = make([]*model.Area, 0)
	}

	areas = append(areas, &loginArea)
	dao.salesPersonsLocations[personId] = areas
	dao.uniqueDestinations[loginArea] = new(interface{})

	return nil
}

func (dao SalesmanDaoImpl) LogoutSalesPerson(personId string) *tserror.TSError {
	delete(dao.salesPersonsLocations, personId)

	return nil
}

func (dao SalesmanDaoImpl) GetUniqueDestinations() (int, *tserror.TSError) {
	return len(dao.uniqueDestinations), nil
}

func (dao SalesmanDaoImpl) GetAllSalesPersonLocations() (map[string][]*model.Area, *tserror.TSError) {
	return dao.salesPersonsLocations, nil
}
