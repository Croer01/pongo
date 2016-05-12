package core

//constants
const playerSpeed float32 = 500
const PlayerUp PlayerDirection = 1
const PlayerDown PlayerDirection = -1

//player direction
type PlayerDirection int

type Player struct {
	IdPlayer string `json:"idPlayer"`
	Score int `json:"score"`
	Name string `json:"name"`
	Position float32 `json:"position"`
}

func (this *Player) Move(direction PlayerDirection, elapsedTime float64){
	this.Position += float32(playerSpeed * float32(direction) * float32(elapsedTime));
}