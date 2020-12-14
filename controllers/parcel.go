package controllers

import (
	"log"
	"net/http"

	"fab-chimix/application-go/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

var GetParcel = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	uuid := params["uuid"]

	parcel, err := r.Context().Value("contract").(*gateway.Contract).EvaluateTransaction("ReadAsset", uuid)
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	}

	resp := utils.Message(true, "success")

	// TODO: Marshal it ?

	resp["parcel"] = string(parcel)

	utils.Respond(w, resp, http.StatusOK)
}

var WidthDrawParcel = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	uuid := params["uuid"]
	recyclingFacility := r.FormValue("recyclingFacility")

	_, err := r.Context().Value("contract").(*gateway.Contract).SubmitTransaction("WithDrawAsset", uuid, recyclingFacility)
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	}

	resp := utils.Message(true, "success")

	// TODO: Marshal it ?

	utils.Respond(w, resp, http.StatusOK)
}

var GetAllParcels = func(w http.ResponseWriter, r *http.Request) {

	parcel, err := r.Context().Value("contract").(*gateway.Contract).EvaluateTransaction("GetAllAssets")
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	}

	resp := utils.Message(true, "success")

	resp["parcels"] = string(parcel)

	utils.Respond(w, resp, http.StatusOK)
}

var CreateParcel = func(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	weight := r.FormValue("weight")
	volume := r.FormValue("volume")
	warehouse := r.FormValue("warehouse")
	intransit := r.FormValue("in-transit")
	recyclingFacility := r.FormValue("recyclingFacility")

	uuid := uuid.New().String()

	_, err := r.Context().Value("contract").(*gateway.Contract).SubmitTransaction("CreateAsset", uuid, code, weight, volume, warehouse, intransit, recyclingFacility)
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	}

	resp := utils.Message(true, "success")

	resp["id"] = uuid

	utils.Respond(w, resp, http.StatusOK)
}
