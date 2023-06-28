package task

import (
	"encoding/json"
	"game/api/websocket"
	"game/internal/message"
	"game/internal/sdk"
	"time"
)

// InitTask initializes the task
func InitTask() {
	ticker := time.Tick(time.Second * 2)

	t := &sdk.JinseSdk{}

	for range ticker {
		list, err := t.GetNews()
		if err == nil {
			v, e := json.Marshal(message.FormatNewsMessage(list))
			if e == nil {
				websocket.WriteMessage(string(v))
			}
		}
	}
}
