package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Running winddcutil!")
	command := exec.Command("CMD", "/C", `C:\CODE\winddcutil\toggle-monitor-input.bat`)
	if err := command.Run(); err != nil {
		log.Fatal(err)
	}
}
