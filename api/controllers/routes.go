package controllers

import "github.com/rameshrajagopalanbayer/uom-api-go/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	s.Router.HandleFunc("/uoms", middlewares.SetMiddlewareJSON(s.GetUoms)).Methods("GET")

	s.Router.HandleFunc("/uom/{code}", middlewares.SetMiddlewareJSON(s.GetUom)).Methods("GET")

}
