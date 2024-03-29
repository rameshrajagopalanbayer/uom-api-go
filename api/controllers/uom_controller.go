package controllers

import (
	"github.com/rameshrajagopalanbayer/uom-api-go/api/models"
	"github.com/rameshrajagopalanbayer/uom-api-go/api/responses"
	"net/http"
)

func (server *Server) GetUoms(w http.ResponseWriter, r *http.Request) {

	uom := models.Uom{}

	uoms, err := uom.FindAllUoms(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, uoms)
}
