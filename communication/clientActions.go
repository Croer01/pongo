package communication

type Action int;

type ClientAction Action;

type ServerAction Action;

const (
	MoveUp ClientAction = iota
	MoveDown
)

const (
	GameStart ServerAction = 1000 + iota
	GameEnd
	MoveBall
)
