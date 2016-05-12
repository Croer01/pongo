package server

import (
	"golang.org/x/net/websocket"
	"pongo/game"
	"pongo/communication"
	"fmt"
	"pongo/core"
	"github.com/ventu-io/go-shortid"
)

type hub struct {
	activeConnections map[*clientConnection]bool
	broadcastChannel  chan *communication.ServerResponse
	registerChannel   chan *clientConnection
	unregisterChannel chan *clientConnection
	matches             []*game.Match
}

func NewHub() *hub {
	return &hub{
		broadcastChannel:   make(chan *communication.ServerResponse),
		registerChannel:    make(chan *clientConnection),
		unregisterChannel:  make(chan *clientConnection),
		activeConnections: make(map[*clientConnection]bool),
		matches: make([]*game.Match,0),
	}
}

func (this *hub) Run() {

	for {
		select {
		case connection := <-this.registerChannel:
			this.registerAndAddToMatch(connection)

		case connection := <-this.unregisterChannel:
			if _, ok := this.activeConnections[connection]; ok {
				delete(this.activeConnections, connection)
				close(connection.sendChannel)
			}

		case response := <-this.broadcastChannel:
			for connection := range this.activeConnections {
				select {
				case connection.sendChannel <- response:
				default:
					delete(this.activeConnections, connection)
					close(connection.sendChannel)
				}
			}
		}
	}
}

func (this *hub) RegisterConnection(socket *websocket.Conn, session *Session) {
	id, _ := shortid.Generate()
	//wrap socket connection
	connection := &clientConnection{
		Id: id,
		sendChannel: make(chan *communication.ServerResponse, 256),
		socket: socket,
		hub: this,
		name: session.Nick,
	}

	//register connection on hub
	this.registerChannel <- connection

	//unregister connection on hub
	defer func() {
		this.unregisterChannel <- connection
	}()

	//open connection's IO channels
	go connection.writer()
	connection.reader()
}

//private
func (this *hub) registerAndAddToMatch(connection *clientConnection) {
	connection.match = this.findOrCreateMatch();
	connection.match.AddPlayer(&core.Player{Name:connection.name, IdPlayer:connection.Id})
	this.activeConnections[connection] = true
}

func (this *hub) findOrCreateMatch() *game.Match {
	for _, match := range this.matches {
		if !match.IsFullFill() {
			match.Start(func(response *communication.ServerResponse) {
				this.broadcastChannel <- response
			});
			fmt.Println("Find match to play")
			return match
		}
	}

	fmt.Println("create new match")
	match := game.NewMatch();
	this.matches = append(this.matches, match)
	return match
}
