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
	GetAllPlants() ([]models.Plants, error)
}

func NewService() LocalMysqlIntf {
	svc := service{}
	svc.GoMysql = mysql.NewGoMithMysql(common.DB_NAME, common.DB_SERVER, common.DB_USER, common.DB_PASSWORD, common.DB_DRIVER)
	return &svc
}

func (s *service) GetAllPlants() ([]models.Plants, error) {
	s.openDBConnection()
	defer s.DB.Close()
	list := []models.Plants{}
	query := "SELECT * FROM PLANTS"
	results, err := s.DB.Query(query)
	if err != nil {
		return list, err
	}

	for results.Next() {
		plant := models.Plants{}
		err := results.Scan(
			&plant.PlantID,
			&plant.PlantType,
			&plant.Temperature,
			&plant.PhLevels,
			&plant.Humidity,
			&plant.CreatedAt,
		)
		if err != nil {
			return list, err
		}
		list = append(list, plant)
	}

	return list, nil
}

func (s *service) openDBConnection() {
	s.DB = s.GoMysql.GetDBConnection()
}
