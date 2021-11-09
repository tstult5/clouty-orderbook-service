package controllers

import "github.com/tstult5/clouty-orderbook-service/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
 	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	//Orderbook routes
	s.Router.HandleFunc("/orderbooks", middlewares.SetMiddlewareJSON(s.GetOrderBooks)).Methods("GET")
//	s.Router.HandleFunc("/orderbooks/{id}", middlewares.SetMiddlewareJSON(s.FindOrderBookByID)).Methods("GET")
//	s.Router.HandleFunc("/orderbooks/create", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateOrderbook))).Methods("POST"
//	s.Router.HandleFunc("/orderbooks/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateOrderbook))).Methods("PUT")
//	s.Router.HandleFunc("/orderbooks/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteOrderbook)).Methods("DELETE")

}
