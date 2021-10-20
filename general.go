package GoPlus

import (
	"fmt"
	"log"
	"time"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
func HandleLightErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func RemoveFromSliceOrdered(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func RemoveFromSliceUnordered(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

type interval string

// StartTimeTicker ticks when exactly a minute hits or other interval. interval must be in minutes
func StartTimeTicker(outputBroadcast *Broker, interval int) {
	outputBroadcast = NewBroadcast()
	go outputBroadcast.Start()
	for {
		if time.Now().Unix()%60 == 0 {
			for {
				outputBroadcast.Publish(struct{}{})
				time.Sleep(time.Minute * time.Duration(interval))
			}
		}
	}
}
