package game

import (
	"pongo/communication"
	"pongo/core"
	"time"
	"errors"
	"fmt"
)

func NewMatch() *Match {
	//ball
	ball := core.Ball{X:0, Y:-250}

	return &Match{startTime:time.Now(), Ball:&ball, Players:make([]*core.Player, 2), Dimension:[2]int{32, 18}}
}

type Match struct {
	Players     []*core.Player `json:"players"`
	Ball        *core.Ball `json:"ball"`
	Dimension   [2]int `json:"size"`
	startTime   time.Time
	elapsedTime float64 `json:"-"`
}

//private
func (this *Match) getPlayer(idPlayer string) (*core.Player, error) {
	for _, player := range this.Players {
		if player.IdPlayer == idPlayer {
			return player, nil
		}
	}

	return nil, fmt.Errorf("Player %s not found", idPlayer)
}

func (this *Match) movePlayer(player *core.Player, direction core.PlayerDirection) {
	player.Move(direction, this.elapsedTime)
}

func (this *Match) moveBall() {
	this.Ball.Move(this.elapsedTime);
}

//public
func (this *Match) DoAction(request communication.ClientRequest) communication.ServerResponse {

	player, err := this.getPlayer(request.IdPlayer)

	if err != nil {
		panic(err)
	}

	response := communication.ServerResponse{IdPlayer:request.IdPlayer, Action:communication.Action(request.Action)}

	switch request.Action {
	case communication.MoveUp:
		this.movePlayer(player, core.PlayerUp)
		response.Result = player.Position
	case communication.MoveDown:
		this.movePlayer(player, core.PlayerDown)
		response.Result = player.Position
	default:
		response = communication.UnkownServerResponse
	}

	return response
}

func (this *Match) DoServerAction(action communication.ServerAction) communication.ServerResponse {

	response := communication.ServerResponse{Action: communication.Action(action)}

	switch action {
	case communication.GameStart:
		response.Result = this
	case communication.MoveBall:
		this.moveBall()
		response.Result = this.Ball
	default:
		response = communication.UnkownServerResponse
	}

	return response
}

func (this *Match) IsFullFill() bool {
	players := this.Players
	return players[0] != nil && players[1] != nil
}

func (this *Match) Start(sender func(*communication.ServerResponse)) {
	go func() {
		//send game is start
		gameStartNotification := this.DoServerAction(communication.GameStart)
		sender(&gameStartNotification)

		timeToSleep := time.Second / 60
		cycleStartTime := time.Now()
		for {
			this.elapsedTime = time.Since(cycleStartTime).Seconds()
			serverResponse := this.DoServerAction(communication.MoveBall)

			sender(&serverResponse)
			cycleStartTime = time.Now()
			time.Sleep(timeToSleep)
		}

	}()
}

func (this *Match) AddPlayer(player *core.Player) {
	player.Score = 0
	player.Position = 0

	players := this.Players

	if players[0] == nil {
		players[0] = player
	} else if players[1] == nil {
		players[1] = player
	} else {
		panic(errors.New("Match is full"))
	}
}