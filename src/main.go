package main

import (
	"github.com/PopSquad/BalloonField/src/http"
	"github.com/PopSquad/BalloonField/src/plaza"
	"github.com/PopSquad/BalloonField/src/room"
)

func main() {
	go plaza.NewPlazaServer("0.0.0.0:8000").Start()
	go room.NewRoomServer("0.0.0.0:9000").Start()
	http.NewHTTPServer("0.0.0.0:80").Start()
}
