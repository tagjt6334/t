package main

import (
	"exec"
	"time"
)

func server() {
	std := exec.Command("python3", "-m", "http.server", "8000")
	std.Run()
}
func cloudflared() {
	std := exec.Command("python3 cloudflared.py")
	std.Run
}
func main() {
	go server()
	go cloudflared()
	for {
		time.Sleep(5 * time.Hour)
	}
}
