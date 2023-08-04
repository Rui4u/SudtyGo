package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		a, _ := strconv.Atoi(r.FormValue("a"))
		b, _ := strconv.Atoi(r.FormValue("b"))
		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		_, _ = w.Write(jsonData)
	})
	_ = http.ListenAndServe(":8000", nil)
}
