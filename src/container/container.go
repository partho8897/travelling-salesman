package container

import (
	"github.com/travellingsalesman/src/domain/controller"
	"github.com/travellingsalesman/src/domain/dao"
	"github.com/travellingsalesman/src/domain/service"
	"github.com/travellingsalesman/src/router"
	"go.uber.org/dig"
)

func BuildContainer() (*dig.Container, error) {
	container := dig.New()

	// Register dao
	err := getError(container.Provide(dao.NewSalesmanDaoImpl), nil)

	// Register service
	err = getError(container.Provide(service.NewSalesmanServiceImpl), err)

	// Register controller
	err = getError(container.Provide(controller.NewSalesmanController), err)

	// Register router
	err = getError(container.Provide(router.NewRouter), err)

	return container, nil
}

func getError(previousError error, newErr error) error {
	if previousError != nil {
		return previousError
	} else {
		return newErr
	}
}
