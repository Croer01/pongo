package core

import (
	"math"
)

type Ball struct {
	X float32
	Y float32
}

var counter float64

func init(){
	counter = 0
}

func (this *Ball) Move(elapsedTime float64) {
	this.X = float32(444 + math.Ceil(math.Cos(counter) * 444))
	counter+=elapsedTime;
}
