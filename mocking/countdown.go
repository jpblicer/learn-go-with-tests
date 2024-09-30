package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type DefaultSleeper struct {}

func(d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

func main () {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

func Countdown (out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
	}

	for i := countdownStart; i > 0; i-- {
	fmt.Fprintln(out, i)
	}


	fmt.Fprint(out, finalWord)
}
