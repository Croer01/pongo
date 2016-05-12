package communication

type ServerResponse struct {
	IdPlayer string `json:"idPlayer"`
	Action   Action `json:"action"`
	Result   interface{} `json:"result"`
}
var UnkownServerResponse ServerResponse

func init() {
	UnkownServerResponse = ServerResponse{}
}