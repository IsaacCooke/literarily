package main

import (
	"github.com/IsaacCooke/literarily/data"
	"github.com/IsaacCooke/literarily/services"
)

func main(){
  data.Connect()
  services.RunServer()
}
