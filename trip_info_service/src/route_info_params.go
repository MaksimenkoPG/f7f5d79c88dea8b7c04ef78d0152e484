package main

import (
  "errors"
  "net/http"
)

type Point struct {
  Latitude  string `json:"latitude"`
  Longitude string `json:"longitude"`
}

type RouteInfoParams struct {
  Origin      Point `json:"origin"`
  Destination Point `json:"destination"`
}

func BuildRouteInfoParams(r *http.Request) (route_info_params RouteInfoParams, err error) {
  switch {
  case len(r.FormValue("o_latitude")) == 0:
    err = errors.New("o_latitude: origin latitude can't be blank")
    return
  case len(r.FormValue("o_longitude")) == 0:
    err = errors.New("o_longitude: origin longitude can't be blank")
    return
  case len(r.FormValue("d_latitude")) == 0:
    err = errors.New("d_latitude: destination latitude can't be blank")
    return
  case len(r.FormValue("d_longitude")) == 0:
    err = errors.New("d_longitude: destination longitude can't be blank")
    return
  default:
    route_info_params = RouteInfoParams{
      Origin:      Point{r.FormValue("o_latitude"), r.FormValue("o_longitude")},
      Destination: Point{r.FormValue("d_latitude"), r.FormValue("d_longitude")},
    }
  }

  return route_info_params, err
}
