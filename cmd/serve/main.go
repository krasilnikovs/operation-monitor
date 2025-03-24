package main

import (
	"net/http"

	"krasilnikovs.lv/operation-monitor/internal/kernel"
)

func main() {
	r := kernel.LoadWeb()

	http.ListenAndServe(":3000", r)
}
