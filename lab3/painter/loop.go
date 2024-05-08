package painter

import (
	"image"
	"sync"

	"golang.org/x/exp/shiny/screen"
)

type Receiver interface {
	Update(t screen.Texture)
}

type Loop struct {
	Receiver Receiver

	next screen.Texture
	prev screen.Texture

	mq messageQueue

	stop    chan struct{}
	stopReq bool
}

var size = image.Pt(800, 800)

func (l *Loop) Start(s screen.Screen) {
	l.next, _ = s.NewTexture(size)
	l.prev, _ = s.NewTexture(size)

	l.stop = make(chan struct{})

	go func() {
		for !l.stopReq || !l.mq.empty() {
			op := l.mq.pull()

			update := op.Do(l.next)
			if update {
				l.Receiver.Update(l.next)
				l.next, l.prev = l.prev, l.next
			}
		}
		close(l.stop)
		l.stop = nil
	}()
}

func (l *Loop) Post(op Operation) {
	l.mq.push(op)
}

func (l *Loop) StopAndWait() {
	l.Post(OperationFunc(func(t screen.Texture) {
		l.stopReq = true
	}))
	<-l.stop
}

type messageQueue struct {
	pushSignal chan struct{}

	mu sync.Mutex

	data []Operation
}

func (mq *messageQueue) push(op Operation) {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	mq.data = append(mq.data, op)

	if mq.pushSignal != nil {
		close(mq.pushSignal)
		mq.pushSignal = nil
	}
}

func (mq *messageQueue) pull() Operation {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	for len(mq.data) == 0 {
		mq.pushSignal = make(chan struct{})
		mq.mu.Unlock()
		<-mq.pushSignal
		mq.mu.Lock()
	}

	res := mq.data[0]
	mq.data[0] = nil
	mq.data = mq.data[1:]

	return res
}

func (mq *messageQueue) empty() bool {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	return len(mq.data) == 0
}
