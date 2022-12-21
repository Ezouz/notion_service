package handlers

import (
	"log"
	"os"
	"reflect"

	"github.com/go-openapi/runtime/middleware"
	"gitlab.42paris.fr/notion_service/pkg/methods"
	"gitlab.42paris.fr/notion_service/pkg/models"
	"gitlab.42paris.fr/notion_service/pkg/restapi/operations/clusters"
)

type ClustersCheckhost struct {
	// Database *ent.Client
}

func (h *ClustersCheckhost) Handle(compDef clusters.PostCheckhostParams, token interface{}) middleware.Responder {
	// log.Printf("Interface ---------------------- %v", token)
	log.Printf("Hostname ---------------------- %v", compDef.ClusterHost.Hostname)
	log.Printf("IP ---------------------- %v", compDef.ClusterHost.IP)
	log.Printf("MacAddress ---------------------- %v", compDef.ClusterHost.MacAddress)
	log.Printf("Manufacturer ---------------------- %v", compDef.ClusterHost.Manufacturer)
	log.Printf("Model ---------------------- %v", compDef.ClusterHost.Model)
	log.Printf("Serial ---------------------- %v", compDef.ClusterHost.Serial)
	log.Printf("BiosDate ---------------------- %v", compDef.ClusterHost.BiosDate)
	log.Printf("LastRedump ---------------------- %v", compDef.ClusterHost.LastRedump)
	log.Printf("TEST ---------------------- %v", reflect.ValueOf(compDef.ClusterHost))
	// values := reflect.ValueOf(compDef.ClusterHost)
	// types := values.Type()
	// if values.Kind() == reflect.Ptr {
	// 	values = values.Elem()
	// 	log.Printf("TEST ---------------------- %v", values)
	// 	for i := 0; i < values.NumField(); i++ {
	// 		log.Println(types.Field(i).Index[0], types.Field(i).Name, values.Field(i))
	// 	}

	// }

	// getNotionDB
	log.Printf("Cluster DB VLAUE: %s", os.Getenv("DB_CLUSTER"))
	page, err := methods.FetchNotionDB(os.Getenv("DB_CLUSTER"))
	if err != nil {
		log.Printf("Could not connect to Notion Cluster Database: %s", err)
		return &clusters.PostCheckhostBadRequest{
			Payload: &models.ErrorResponse{
				Code:    500,
				Message: "Bad Notion DB ID",
			},
		}
	}
	log.Printf("Cluster DB: %v", reflect.ValueOf(page))
	log.Printf("Cluster DB: %v", reflect.ValueOf(page))
	log.Printf("Cluster DB: %v", reflect.ValueOf(page))

	// Check if Hostname is here
	// if not create new

	// methods.PostToNotionDB()

	// Check All parameters are here
	// if reflect.ValueOf(compDef.ClusterHost) != 8 {
	// 	// if no error
	// 	return &clusters.PostCheckhostBadRequest{
	// 		Payload: &models.ErrorResponse{
	// 			Code:    401,
	// 			Message: "Bad Request",
	// 		},
	// 	}
	// }
	return &clusters.PostCheckhostNoContent{}
}
