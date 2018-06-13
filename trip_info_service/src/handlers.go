package main

import (
	"encoding/json"
	"net/http"
)

func TripInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	route_info_params, err := BuildRouteInfoParams(r)
	if err != nil {
		render422(w, err)
		return
	}

	route_info_response, err := GetRouteInfo(route_info_params)
	trip_info_response, err := GetTripInfo(route_info_response)
	if err != nil {
		render422(w, err)
		return
	}

	if err := json.NewEncoder(w).Encode(trip_info_response); err != nil {
		panic(err)
	}
}

func render422(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusUnprocessableEntity, Text: err.Error()}); err != nil {
		panic(err)
	}
}
