package handlers

import (
	"fmt"
	"log"
	"os"

	"github.com/dstotijn/go-notion"
	"github.com/go-openapi/runtime/middleware"
	"gitlab.42paris.fr/notion_service/pkg/methods"
	"gitlab.42paris.fr/notion_service/pkg/models"
	"gitlab.42paris.fr/notion_service/pkg/restapi/operations/clusters"
	"gitlab.42paris.fr/notion_service/pkg/utils"
)

type ClustersCheckhost struct {
	// Database *ent.Client
}

func (h *ClustersCheckhost) Handle(compDef clusters.PostCheckhostParams, token interface{}) middleware.Responder {
	// log.Printf("Hostname ---------------------- %v", compDef.ClusterHost.Hostname)
	// log.Printf("IP ---------------------- %v", compDef.ClusterHost.IP)
	// log.Printf("MacAddress ---------------------- %v", compDef.ClusterHost.MacAddress)
	// log.Printf("Manufacturer ---------------------- %v", compDef.ClusterHost.Manufacturer)
	// log.Printf("Model ---------------------- %v", compDef.ClusterHost.Model)
	// log.Printf("Serial ---------------------- %v", compDef.ClusterHost.Serial)
	// log.Printf("BiosDate ---------------------- %v", compDef.ClusterHost.BiosDate)
	// log.Printf("LastRedump ---------------------- %v", compDef.ClusterHost.LastRedump)

	// getNotionDB & Check if Hostname is here
	row, err := utils.FindRowWithTitleInNotionDB(os.Getenv("DB_CLUSTER"), compDef.ClusterHost.Hostname, nil)
	if err != nil {
		return &clusters.PostCheckhostBadRequest{
			Payload: err,
		}
	}

	if row != nil {
		// Update row if exists
		props, ok := row.Properties.(notion.DatabasePageProperties)
		if ok == false {
			log.Print("String Error")
		}
		dbProps, err := utils.UpdateNotionRowProperties(props, compDef.ClusterHost)
		if err != nil {
			return &clusters.PostCheckhostBadRequest{
				Payload: &models.ErrorResponse{
					Code:    404,
					Message: "Bad Request",
				},
			}
		}

		params := notion.UpdatePageParams{
			DatabasePageProperties: dbProps,
		}

		_, err = methods.UpdateNotionPage(row.ID, params)
		if err != nil {
			return &clusters.PostCheckhostBadRequest{
				Payload: &models.ErrorResponse{
					Code:    404,
					Message: fmt.Sprintf("Couldn't update Notion Page: %s", err),
				},
			}
		}

	} else {
		// or create new row // post page
		// get template
		template, err := methods.FetchNotionPage(os.Getenv("DB_CLUSTER_TEMPLATE"))
		if err != nil {
			return &clusters.PostCheckhostBadRequest{
				Payload: &models.ErrorResponse{
					Code:    404,
					Message: "Template Page for new entry Not Found",
				},
			}
		}
		props, ok := template.Properties.(notion.DatabasePageProperties)
		if ok == false {
			log.Print("String Error")
		}
		// fill template properties
		dbProps, err := utils.FillNotionPropertyFromTemplate(props, compDef.ClusterHost)
		if err != nil {
			return &clusters.PostCheckhostBadRequest{
				Payload: &models.ErrorResponse{
					Code:    404,
					Message: "Bad Request",
				},
			}
		}

		params := notion.CreatePageParams{
			ParentType:             "database_id",
			ParentID:               os.Getenv("DB_CLUSTER"),
			DatabasePageProperties: &dbProps,
		}

		methods.PostToNotionDB(os.Getenv("DB_CLUSTER"), params)
	}
	return &clusters.PostCheckhostNoContent{}
}
