package server

import (
  "github.com/IsaacCooke/literarily/data"
)

func main(){
  data.Connect()
  controllers.RunServer()
}
