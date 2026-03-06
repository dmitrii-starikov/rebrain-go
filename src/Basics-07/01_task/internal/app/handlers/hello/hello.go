package hello

import (
	"encoding/json"
	"module06/internal/pkg/util"
	"net/http"
)

// Response - repsonse data struct
type Response struct {
	Msgs []string
}

// Handler - hello handler
func Handler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Msgs: make([]string, 0, 1000)}

	for i := 0; i < 3000; i++ {
		msg := util.Pad("Hello, ", i, " World!")
		resp.Msgs = append(resp.Msgs, msg)
	}

	jsonRes, _ := json.Marshal(resp)
	w.Write(jsonRes)
}
