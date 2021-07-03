package GoPlus

import (
	"fmt"
	"time"
)

type Broker struct {
	stopCh    chan struct{}
	publishCh chan interface{}
	subCh     chan chan interface{}
	unsubCh   chan chan interface{}
}

func NewBroadcast() *Broker {
	return &Broker{
		stopCh:    make(chan struct{}),
		publishCh: make(chan interface{}, 10),
		subCh:     make(chan chan interface{}, 1),
		unsubCh:   make(chan chan interface{}, 1),
	}
}

func (b *Broker) Start() {
	subs := map[chan interface{}]struct{}{}
	for {
		select {
		case <-b.stopCh:
			return
		case msgCh := <-b.subCh:
			subs[msgCh] = struct{}{}
		case msgCh := <-b.unsubCh:
			delete(subs, msgCh)
		case msg := <-b.publishCh:
			for msgCh := range subs {
				// msgCh is buffered, use non-blocking send to protect the broker:
				select {
				case msgCh <- msg:
				default:
				}
			}
		}
	}
}

func (b *Broker) Stop() {
	close(b.stopCh)
}

func (b *Broker) Subscribe() chan interface{} {

	msgCh := make(chan interface{}, 10)
	b.subCh <- msgCh
	return msgCh
}

func (b *Broker) Unsubscribe(msgCh chan interface{}) {
	b.unsubCh <- msgCh
}

func (b *Broker) Publish(msg interface{}) {
	b.publishCh <- msg
}

func example() {
	// Create and start a broker:
	b := NewBroadcast()
	go b.Start()

	// Create and subscribe 3 clients:
	clientFunc := func(id int) {
		msgCh := b.Subscribe()
		for {
			fmt.Printf("Client %d got message: %v\n", id, <-msgCh)
		}
	}
	for i := 0; i < 3; i++ {
		go clientFunc(i)
	}

	// Start publishing messages:
	go func() {
		for msgId := 0; ; msgId++ {
			b.Publish(fmt.Sprintf("msg#%d", msgId))
			time.Sleep(300 * time.Millisecond)
		}
	}()

	time.Sleep(time.Second)
}
