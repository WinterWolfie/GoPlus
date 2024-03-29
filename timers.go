package GoPlus

import "time"

func GlobalClock(minuteTicker *Broker) {
	minuteTicker = NewBroadcast()
	go minuteTicker.Start()
	go func() {
		for {
			if time.Now().Unix()%60 == 0 {
				for {
					//fmt.Println("GlobalClock(): tick")
					minuteTicker.Publish(struct{}{})
					time.Sleep(time.Minute)
				}
			}
		}
	}()
}

func (broker *Broker) GlobalClock() {
	go broker.Start()
	go func() {
		for {
			if time.Now().Unix()%60 == 0 {
				for {
					//fmt.Println("GlobalClock(): tick")
					broker.Publish(struct{}{})
					time.Sleep(time.Minute)
				}
			}
		}
	}()
}
