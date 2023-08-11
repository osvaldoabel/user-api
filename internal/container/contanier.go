package container

import (
	"log"

	"github.com/osvaldoabel/user-api/configs"
	"github.com/osvaldoabel/user-api/internal/repository/user"
	"github.com/osvaldoabel/user-api/pkg/database"
)

type DependencyContainer struct {
	UserRepo user.UserDBRepository
}

func NewDependenciesContainer(conf configs.Conf) DependencyContainer {
	conn, err := database.InitDB(conf)
	if err != nil {
		log.Fatal(err)
	}

	return DependencyContainer{
		UserRepo: user.NewUserDBRepository(conn),
	}
}
