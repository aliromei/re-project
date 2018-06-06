package model

func Insert() bool {
  BUS := connect("bus")
  err := BUS.Insert(&Bus{Model: "asdf", Seats: 31, OpenSeatCount: 32})
  if err != nil {
    panic(err)
  }
  return true
}
