package app

var VERSION = "dev"

type (
	GameRequest struct {
		ClientID int     `json:"client_id"`
		Fee      float64 `json:"fee"`
		Pair     Pair    `json:"pair"`
	}

	GameResponse struct {
		GameStatus GameStatus `json:"game_status"`
		Amount     float64    `json:"amount"`
	}

	GameStatus = string
)

const (
	FAIL_GAME     GameStatus = "fail"
	WIN_FREE_GAME GameStatus = "free game"
	WIN_GAME      GameStatus = "win"
)
