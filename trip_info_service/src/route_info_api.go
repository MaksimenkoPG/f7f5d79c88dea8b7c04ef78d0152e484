package main

import (
  "encoding/json"
  "fmt"
  "github.com/streadway/amqp"
  "log"
  "os"
)

type RouteInfoResponse struct {
  Distance float64 `json:"distance"`
  Duration float64 `json:"duration"`
  Status   string  `json:"status"`
}

func GetRouteInfo(route_info_params RouteInfoParams) (route_info_response RouteInfoResponse, err error) {
  conn, err := amqp.Dial(os.Getenv("TRIP_INFO_SERVICE_AMQP"))
  failOnError(err, "Failed to connect to RabbitMQ")
  defer conn.Close()

  ch, err := conn.Channel()
  failOnError(err, "Failed to open a channel")
  defer ch.Close()

  q, err := ch.QueueDeclare(
    "",    // name
    false, // durable
    false, // delete when unused
    true,  // exclusive
    false, // noWait
    nil,   // arguments
  )
  failOnError(err, "Failed to declare a queue")

  msgs, err := ch.Consume(
    q.Name, // queue
    "",     // consumer
    true,   // auto-ack
    false,  // exclusive
    false,  // no-local
    false,  // no-wait
    nil,    // args
  )
  failOnError(err, "Failed to register a consumer")

  corrId, err := GenerateRandomString(32)
  if err != nil {
    fmt.Println(err)
    return
  }

  route_info_json, err := json.Marshal(route_info_params)
  if err != nil {
    fmt.Println(err)
    return
  }

  err = ch.Publish(
    "",          // exchange
    "rpc_queue", // routing key
    false,       // mandatory
    false,       // immediate
    amqp.Publishing{
      ContentType:   "application/json",
      CorrelationId: corrId,
      ReplyTo:       q.Name,
      Body:          []byte(route_info_json),
    })
  failOnError(err, "Failed to publish a message")

  for d := range msgs {
    if corrId == d.CorrelationId {
      var data map[string]interface{}

      err := json.Unmarshal([]byte(d.Body), &data)
      if err != nil {
        fmt.Println("There was an error:", err)
      } else {
        route_info_response = RouteInfoResponse{
          Distance: float64(data["distance"].(float64)),
          Duration: float64(data["duration"].(float64)),
          Status:   string(data["status"].(string)),
        }
      }

      break
    }
  }

  return route_info_response, err
}

func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}
