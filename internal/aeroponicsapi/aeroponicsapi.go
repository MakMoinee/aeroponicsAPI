package aeroponicsapi

import (
	"aeroponicsAPI/internal/aeroponicsapi/models"
	"aeroponicsAPI/internal/repository/localmysql"
	"log"

	"github.com/MakMoinee/go-mith/pkg/gofirebase"
)

type service struct {
	FirebaseApp gofirebase.FirebaseIntf
	LocalMySql  localmysql.LocalMysqlIntf
}

type AeroponicsIntf interface {
	GetPlants() ([]models.Plants, error)
}

func NewService() AeroponicsIntf {
	svc := service{}
	svc.LocalMySql = localmysql.NewService()

	return &svc
}

func (s *service) GetPlants() ([]models.Plants, error) {
	log.Println("aeroponicsAPI:GetPlants() invoked ...")
	list, err := s.LocalMySql.GetAllPlants()
	return list, err

}
