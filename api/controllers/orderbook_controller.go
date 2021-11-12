package controllers

import (
	"encoding/json"
	"strconv"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/tstult5/clouty-orderbook-service/api/models"
	"github.com/tstult5/clouty-orderbook-service/api/responses"
	"github.com/tstult5/clouty-orderbook-service/api/utils/formaterror"
)


func (server *Server) GetOrders(w http.ResponseWriter, r *http.Request) {
  var err error
	order := models.Order{}
	orders, err := order.FindAllOrders(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, orders)
}

func (server *Server) FindOrderByID(w http.ResponseWriter, r *http.Request) {
  var err error
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	order := models.Order{}
	fetchedOrder, err := order.FindOrderByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, fetchedOrder)
}

func (server *Server) CreateOrder(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	order := models.Order{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	order.Prepare()
	err = order.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	orderCreated, err := order.SaveOrder(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, orderCreated.ID))
	responses.JSON(w, http.StatusCreated, orderCreated)
}
