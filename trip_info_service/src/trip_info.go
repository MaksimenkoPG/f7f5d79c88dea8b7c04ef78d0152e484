package main

import (
  "errors"
  "fmt"
  "math"
)

type TripInfoResponse struct {
  TotalPrice float64 `json:"total_price"`
}

func GetTripInfo(route_info_response RouteInfoResponse) (trip_info_response TripInfoResponse, err error) {
  call_price := tariff.CallPrice
  distance_price := (route_info_response.Distance / 1000) * tariff.PricePerKilometer
  duration_price := (route_info_response.Duration / 60) * tariff.PricePerMinute
  total_price := math.Round(call_price + distance_price + duration_price)

  if total_price < tariff.MinimalTotalPrice {
    err_message := fmt.Sprintf(
      "total_price %v is too small, minimal total_price is %v",
      total_price,
      tariff.MinimalTotalPrice,
    )
    err = errors.New(err_message)
  } else {
    trip_info_response = TripInfoResponse{TotalPrice: total_price}
  }

  return trip_info_response, err
}
