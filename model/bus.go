package model

import (
  "gopkg.in/mgo.v2/bson"
  "time"
  "errors"
  "gopkg.in/mgo.v2"
)

func (this *Bus) Create() error {
  this.Id = bson.NewObjectId()
  this.OpenSeatCount = this.Seats
  this.Status = BUS_STATUS_NOT_READY
  this.CreatedAt, this.UpdatedAt = time.Now(), time.Now()

  for counter := 1; counter <= this.OpenSeatCount; counter++ {
    this.AvailableSeats = append(this.AvailableSeats, counter)
  }

  this.Insert()

  return nil
}

func UpdateBus(id string, status int) (Bus, error) {
  var bus Bus
  BUS := connect("buses")
  if err := BUS.FindId(bson.ObjectIdHex(id)).One(&bus); err != nil {
    return bus, err
  }
  if status < 1 || status > 3 {
    return bus, errors.New("wrong status provided")
  }
  if bus.Status >= status {
    return bus, errors.New("could not revert bus's status")
  }
  changes := mgo.Change{
    Update: bson.M{"$set": bson.M{"status":status}},
    ReturnNew: true,
  }
  if _, err := BUS.FindId(bson.ObjectIdHex(id)).Apply(changes, &bus); err != nil {
    return bus, nil
  }
  if err := ChangePassengersStatus(id, status); err != nil {
    return bus, err
  }
  return bus, nil
}

func (this *Bus) ShowBus(id string) error {
  BUS := connect("buses")
  if err := BUS.FindId(bson.ObjectIdHex(id)).One(this); err != nil {
    return err
  }
  return nil
}

func BusesList() ([]Bus, error) {
  var buses []Bus
  BUS := connect("buses")
  if err := BUS.Find(nil).All(&buses); err != nil {
    return nil, err
  }
  return buses, nil
}

func (this *Bus) Insert() error {
  BUS := connect("buses")
  if err := BUS.Insert(this); err != nil {
    return err
  }

  BUS.Find(bson.M{"id":this.Id}).One(this)

  return nil
}