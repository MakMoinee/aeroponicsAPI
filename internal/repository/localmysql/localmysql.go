package localmysql

import (
	"aeroponicsAPI/internal/aeroponicsapi/common"
	"aeroponicsAPI/internal/aeroponicsapi/models"
	"database/sql"

	"github.com/MakMoinee/go-mith/pkg/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type service struct {
	GoMysql mysql.GoMithMysql
	DB      *sql.DB
}

type LocalMysqlIntf interface {
	GetAllPlants() []models.Plants
}

func NewService() LocalMysqlIntf {
	svc := service{}
	svc.GoMysql = mysql.NewGoMithMysql(common.DB_NAME, common.DB_SERVER, common.DB_USER, common.DB_PASSWORD, common.DB_DRIVER)
	return &svc
}

func (s *service) GetAllPlants() []models.Plants {
	list := []models.Plants{}

	return list
}
