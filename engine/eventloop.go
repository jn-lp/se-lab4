package engine

import "sync"

type EventLoop struct {
	sync.Mutex

	messageQueue []Command
	receive      chan struct{}
	ready        bool

	pause  bool
	finish chan struct{}
}

func (loop *EventLoop) popCommand() Command {
	loop.Lock()
	defer loop.Unlock()

	if len(loop.messageQueue) == 0 {
		loop.ready = true
		loop.Unlock()
		<-loop.receive
		loop.Lock()
	}

	cmd := loop.messageQueue[0]
	loop.messageQueue[0] = nil
	loop.messageQueue = loop.messageQueue[1:]

	return cmd
}

func (loop *EventLoop) run() {
	for !loop.pause || len(loop.messageQueue) != 0 {
		loop.popCommand().Execute(loop)
	}
	loop.finish <- struct{}{}
}

func (loop *EventLoop) Start() {
	loop.receive = make(chan struct{})
	loop.finish = make(chan struct{}, 1)

	go loop.run()
}

func (loop *EventLoop) Post(command Command) {
	loop.Lock()
	defer loop.Unlock()
	loop.messageQueue = append(loop.messageQueue, command)

	if loop.ready {
		loop.ready = false
		loop.receive <- struct{}{}
	}
}

func (loop *EventLoop) AwaitFinish() {
	loop.Post(pauseCommand(func(h Handler) {}))
	<-loop.finish
}
