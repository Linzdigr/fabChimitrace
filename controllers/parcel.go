package controllers

import (
	"log"
	"net/http"

	"fab-chimix/application-go/utils"

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

	log.Printf("recyclingFacility %s\n", recyclingFacility)

	parcel, err := r.Context().Value("contract").(*gateway.Contract).SubmitTransaction("WithDrawAsset", uuid, recyclingFacility)
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	}

	resp := utils.Message(true, "success")

	// TODO: Marshal it ?

	resp["parcel"] = string(parcel)

	utils.Respond(w, resp, http.StatusOK)
}
