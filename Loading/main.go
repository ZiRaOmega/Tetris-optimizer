package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	start := time.Now()
	RunCmd("clear")
	for time.Since(start) < time.Duration(5)*time.Second {
		fmt.Println("LOADING...\nLaunching\n_____\n" + `⬡⬡⬡⬡⬡` + "\n")
		WaitNClear()
		fmt.Println("LOADING...\ngetting data\n_____\n" + `⬢⬡⬡⬡⬡` + "\n")
		WaitNClear()
		fmt.Println("LOADING...\nsolving\n_____\n" + `⬢⬢⬡⬡⬡` + "\n")
		WaitNClear()
		fmt.Println("LOADING...\npause clope\n_____\n" + `⬢⬢⬢⬡⬡` + "\n")
		WaitNClear()
		fmt.Println("LOADING...\npetit café\n_____\n" + `⬢⬢⬢⬢⬡` + "\n")
		WaitNClear()
		fmt.Println("LOADING...\net z'est partii\n_____\n" + `⬢⬢⬢⬢⬢` + "\n")
		WaitNClear()
	}
}

func WaitNClear() {
	Wait(1.5)
	RunCmd("clear")
}

func RunCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Wait(waitingTime float32) {
	start := time.Now()
	for time.Since(start) < time.Second*time.Duration(waitingTime) {
	}
}
