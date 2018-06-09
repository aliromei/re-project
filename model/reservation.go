package model

import (
  "gopkg.in/mgo.v2/bson"
  "errors"
  "github.com/aliromei/re-project/authentication"
  "gopkg.in/mgo.v2"
  "time"
)

func Reserve(id string, seat int) error {
  var bus Bus
  BUS := connect("buses")
  if err := BUS.FindId(bson.ObjectIdHex(id)).One(&bus); err != nil {
    return err
  }
  if !(bus.OpenSeatCount > 0) {
    return errors.New("no open seats")
  }
  if bus.Status != BUS_STATUS_NOT_READY {
    return errors.New("bus has gone")
  }
  passenger := Passenger{
    PassengerId: authentication.Id,
    Seat: seat,
  }

  seatFound := false
  for i := range bus.AvailableSeats {
    if bus.AvailableSeats[i] == seat {
      bus.AvailableSeats = append(bus.AvailableSeats[:i], bus.AvailableSeats[i+1:]...)
      bus.OpenSeatCount -= 1
      seatFound = true
      break
    }
  }
  if !seatFound {
    return errors.New("seat not found")
  }
  changeBus := mgo.Change{
    Update: bson.M{
      "$push": bson.M{
        "passengers": passenger,
      },
      "$set": bson.M{
        "availableSeats": bus.AvailableSeats,
        "openSeats":      bus.OpenSeatCount,
        "updatedAt":      time.Now(),
      },
    },
  }

  var user User
  USER := connect("users")
  if err := USER.FindId(bson.ObjectIdHex(authentication.Id)).One(&user); err != nil {
    return err
  }
  if _, err := BUS.FindId(bson.ObjectIdHex(id)).Apply(changeBus, &bus); err != nil {
    return err
  }

  reservation := Reservation{
    BusId: string(bus.Id.Hex()),
    Seat:  seat,
    Origin: Address{
      ProvinceId: bus.Origin.ProvinceId,
      CityId:     bus.Origin.CityId,
    },
    Destination: Address{
      ProvinceId: bus.Destination.ProvinceId,
      CityId:     bus.Destination.CityId,
    },
    Status:      bus.Status,
  }

  user.Reservations = append(user.Reservations, reservation)

  changeUser := mgo.Change{
    Update: bson.M{
      "$set": bson.M{
        "reservations": user.Reservations,
        "updatedAt":   time.Now(),
      },
    },
  }

  if _, err := USER.FindId(bson.ObjectIdHex(authentication.Id)).Apply(changeUser, &user); err != nil {
    return err
  }

  return nil
}
