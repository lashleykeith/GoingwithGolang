package main

import (
		"github.com/skratchdot/open-golang/open"
		"fmt"
		"time"
	)

func delaySecond(n time.Duration){
	time.Sleep(n * time.Second)
}

func main() {

		total_breaks := 3
		break_count := 0
		fmt.Println("The time is now ", time.Now())
		for break_count < total_breaks {
			delaySecond(3)
			open.RunWith("https://www.youtube.com/watch?v=neV3EPgvZ3g", "firefox")	
			break_count = break_count + 1
		}

}

//https://github.com/toqueteos/webbrowser
// First you must install As simple as: go get -u github.com/toqueteos/webbrowser
//https://github.com/skratchdot/open-golang
//http://golangcookbook.blogspot.kr/2012/11/while-loop.html
// https://www.socketloop.com/references/golang-time-sleep-function-example