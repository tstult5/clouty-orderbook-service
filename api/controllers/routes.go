package controllers

import (
  "github.com/tstult5/clouty-orderbook-service/api/middlewares"
  _ "github.com/tstult5/clouty-orderbook-service/docs"
  httpSwagger "github.com/swaggo/http-swagger"
)

func (s *Server) initializeRoutes() {

	// Home Route
 	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	//Orderbook routes
	s.Router.HandleFunc("/orders", middlewares.SetMiddlewareJSON(s.GetOrders)).Methods("GET")
  s.Router.HandleFunc("/order/{id}", middlewares.SetMiddlewareJSON(s.FindOrderByID)).Methods("GET")
	s.Router.HandleFunc("/order/create", middlewares.SetMiddlewareJSON(s.CreateOrder)).Methods("POST")
//	s.Router.HandleFunc("/order/{id}", middlewares.SetMiddlewareJSON(s.UpdateOrder)).Methods("PUT")
//	s.Router.HandleFunc("/order/{id}", middlewares.SetMiddlewareJSON(s.DeleteOrder)).Methods("DELETE")
  s.Router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}
