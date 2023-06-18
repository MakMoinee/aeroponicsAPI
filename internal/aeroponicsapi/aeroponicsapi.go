package aeroponicsapi

import (
	"aeroponicsAPI/internal/aeroponicsapi/models"
	"log"

	"github.com/MakMoinee/go-mith/pkg/gofirebase"
)

type service struct {
	FirebaseApp gofirebase.FirebaseIntf
}

type AeroponicsIntf interface {
	GetPlants() []models.Plants
}

func NewService() AeroponicsIntf {
	svc := service{}

	return &svc
}

func (s *service) GetPlants() []models.Plants {
	log.Println("aeroponicsAPI:GetPlants() invoked ...")
	list := []models.Plants{}

	return list

}
