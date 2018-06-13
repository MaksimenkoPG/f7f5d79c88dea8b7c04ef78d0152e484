package main

var tariff Tariff

type Tariff struct {
  CallPrice         float64 `yaml:"call_price"`
  PricePerMinute    float64 `yaml:"price_per_minute"`
  PricePerKilometer float64 `yaml:"price_per_kilometer"`
  MinimalTotalPrice float64 `yaml:"minimal_total_price"`
}

func init() {
  loadTariff()
}

func loadTariff() {
  tariff = Tariff{
    CallPrice:         250,
    PricePerMinute:    20,
    PricePerKilometer: 20,
    MinimalTotalPrice: 500,
  }
}
