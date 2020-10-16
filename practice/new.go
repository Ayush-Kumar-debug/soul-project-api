package main

import (
	"fmt"
	"time"
	"gopkg.in/robfig/cron.v2"
)

func main() {
    a :=1;
	c := cron.New()
    c.AddFunc("@every 0h0m2s", func() {
         fmt.Println("Every 2 second")
         fmt.Println(a);
         a = a+1; 
        })
	c.Start()

	// Added time to see output
	time.Sleep(10 * time.Second)

	c.Stop() // Stop the scheduler (does not stop any jobs already running).
	now = time.Now()
	fmt.Println(now)
}