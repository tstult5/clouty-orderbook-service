package controllers

import (
//	"encoding/json"
//	"errors"
//	"fmt"
//	"io/ioutil"
	"net/http"
//	"strconv"
//	"github.com/gorilla/mux"
	"github.com/tstult5/clouty-orderbook-service/api/models"
	"github.com/tstult5/clouty-orderbook-service/api/responses"
//	"github.com/tstult5/clouty-orderbook-service/api/utils/formaterror"
)


func (server *Server) GetOrderBooks(w http.ResponseWriter, r *http.Request) {

	orderbook := models.OrderBook{}

	orderbooks, err := orderbook.FindAllOrderBooks(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, orderbooks)
}
