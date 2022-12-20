package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"gitlab.42paris.fr/notion_service/ent"
	"gitlab.42paris.fr/notion_service/pkg/restapi/operations/clusters"
)

type ClustersCheckhost struct {
	Database *ent.Client
}

func (h *ClustersCheckhost) Handle(clusters.PostCheckhostParams, interface{}) middleware.Responder {
	log.Print("C'EST BIEN RECU ----------------------")
	// blablabla
	// if no error send response
	return &clusters.PostCheckhostNoContent{}
}
