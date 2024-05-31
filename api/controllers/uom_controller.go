package controllers

import (
	"github.com/gorilla/mux"
	"github.com/rameshrajagopalanbayer/uom-api-go/api/models"
	"github.com/rameshrajagopalanbayer/uom-api-go/api/responses"
	"github.com/rameshrajagopalanbayer/uom-api-go/api/uomcache"
	"log"
	"net/http"
)

// var globalVar uomCacheInstance  *uomCache = uomcache.NewUomCache()
var uomCache = uomcache.NewUomCache()

func (server *Server) GetUoms(w http.ResponseWriter, r *http.Request) {

	cachedUoms := uomCache.GetAll()

	log.Println("cachedUoms", len(cachedUoms))

	if len(cachedUoms) != 0 {
		responses.JSON(w, http.StatusOK, cachedUoms)
		return
	}

	uom := models.Uom{}

	uoms, err := uom.FindAllUoms(server.DB)

	uomCache.UpdateAll(*uoms)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, uoms)
}

//func getCachedUoms(server *Server) *[]models.Uom {
//	uom := models.Uom{}
//	uoms, _ := uom.FindAllUoms(server.DB)
//	elementMap := make(map[string]string)
//	for i := 0; i < len(uoms); i += 2 {
//		elementMap[uoms[i].Code] = uoms[i]
//	}
//	return uoms
//}

func (server *Server) GetUom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["code"]
	uom := models.Uom{}

	uoms, err := uom.FindAllUoms(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	for _, value := range *uoms {
		if value.Code == code {
			responses.JSON(w, http.StatusOK, value)
			return
		}
	}

	//responses.JSON(w, http.StatusOK, []string{})
}
