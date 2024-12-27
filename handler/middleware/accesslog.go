package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Log struct {
	Timestamp time.Time
	Latency   int64
	Path      string
	OS        string
}

func AccessLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// ハンドラーを呼び出して処理を実行
		h.ServeHTTP(w, r)
		// time.Sleep(2 * time.Second)
		latency := time.Since(startTime).Milliseconds()
		os, _ := r.Context().Value(OSContextKey{}).(string)
		// if !ok {
		// 	os = "Unknown"
		// }
		logEntry := Log{
			Timestamp: startTime,
			Latency:   latency,
			Path:      r.URL.Path,
			OS:        os,
		}
		logJSON, err := json.Marshal(logEntry)
		if err != nil {
			fmt.Println("Error marshaling log:", err)
			return
		}
		fmt.Println(string(logJSON))
	})
}
