package routes

import (
	"aeroponicsAPI/internal/aeroponicsapi"
	"log"
	"net/http"

	"github.com/MakMoinee/go-mith/pkg/goserve"
	"github.com/MakMoinee/go-mith/pkg/response"
	"github.com/go-chi/cors"
)

const (
	plantsPath = "/plants"
)

type service struct {
	Aeroponics aeroponicsapi.AeroponicsIntf
}

type RoutesIntf interface {
	GetPlants(w http.ResponseWriter, r *http.Request)
}

func Set(httpService *goserve.Service) {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-TOKEN"},
		ExposedHeaders:   []string{"Link", "Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	httpService.Router.Use(cors.Handler)
	handlers := newRoutes()
	initiateRoutes(httpService, handlers)
}

func initiateRoutes(httpService *goserve.Service, handler RoutesIntf) {
	httpService.Router.Get(plantsPath, handler.GetPlants)
}

func newRoutes() RoutesIntf {
	svc := service{}
	svc.Aeroponics = aeroponicsapi.NewService()
	return &svc
}

func (s *service) GetPlants(w http.ResponseWriter, r *http.Request) {
	log.Println("routes:GetPlants() invoked ...")
	plantsList := s.Aeroponics.GetPlants()

	response.Success(w, plantsList)
}
