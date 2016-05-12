package server

import (
	"golang.org/x/net/websocket"
	"fmt"
	"pongo/game"
	"pongo/communication"
)

type clientConnection struct {
	Id string
	socket      *websocket.Conn
	sendChannel chan *communication.ServerResponse
	hub         *hub
	match       *game.Match
	name        string
}

func (this *clientConnection) reader() {
	for {
		clientRequest := communication.ClientRequest{}
		err:= websocket.JSON.Receive(this.socket, &clientRequest)

		if err != nil {
			fmt.Errorf("error on parse request %s : %s", this, err)
			break
		}

		clientRequest.IdPlayer = this.Id
		serverResponse := this.match.DoAction(clientRequest)
		this.hub.broadcastChannel <- &serverResponse
	}
	this.socket.Close()
}

func (this *clientConnection) writer() {
	for serverResponse := range this.sendChannel {
		if err := websocket.JSON.Send(this.socket, serverResponse); err != nil {
			fmt.Errorf("error on sent response: %s\n", err)
			break
		}
	}
	this.socket.Close()
}