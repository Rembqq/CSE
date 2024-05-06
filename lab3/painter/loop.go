package painter

import (
	"fmt"
	"image"
	"sync"

	"golang.org/x/exp/shiny/screen"
)

// Receiver отримує текстуру, яка була підготовлена в результаті виконання команд у циклі подій.
type Receiver interface {
	Update(t screen.Texture)
}

// Loop реалізує цикл подій для формування текстури отриманої через виконання операцій отриманих з внутрішньої черги.
type Loop struct {
	Receiver Receiver

	next screen.Texture // текстура, яка зараз формується
	prev screen.Texture // текстура, яка була відправлення останнього разу у Receiver

	mq messageQueue

	stop    chan struct{} // створення каналу для закриття потоку обробки операцій
	stopReq bool          // змінна що вказує на завершення виконнання операцій
}

var size = image.Pt(800, 800)

// запускає цикл подій. Цей метод потрібно запустити до того, як викликати на ньому будь-які інші методи.
func (l *Loop) Start(s screen.Screen) {
	l.next, _ = s.NewTexture(size)
	l.prev, _ = s.NewTexture(size)

	l.stop = make(chan struct{}) // ініціалізує канал

	go func() {
		fmt.Println("\nl.stopReq: ", l.stopReq)
		fmt.Println("l.mq.empty()", l.mq.empty())
		for !l.stopReq || !l.mq.empty() { // перевіряє чи не припиняти обробку повідомлень та чи масив операцій не пустий
			fmt.Println("\ninner l.stopReq: ", l.stopReq)
			fmt.Println("inner l.mq.empty(): ", l.mq.empty())
			op, c := l.mq.pull() // дістає операцію з масиву

			fmt.Println("len(l.mq.data): ", len(l.mq.data))
			fmt.Println("len(l.mq.dataCord): ", len(l.mq.dataCord))
			fmt.Println("Coordinates: ", c.X1, c.Y1, c.X2, c.Y2)

			update := op.Do(l.next, c) // виконує операцію
			fmt.Println("Command = update: ", update)
			if update { // перевіряє чи відбулися зміни
				fmt.Println("Updated #1")
				l.Receiver.Update(l.next) // якщо відбулися оновлює текстуру
				l.next, l.prev = l.prev, l.next
				fmt.Println("Updated #2")
			}
		}
		close(l.stop) // закриває канал, та припиняє потік
		l.stop = nil  // кажемо що в каналі немає сигналів
	}()
}

// додає нову операцію у внутрішню чергу.
func (l *Loop) Post(op Operation) {
	l.mq.push(op) //записує операцю в чергу
}

// додає нову координату у внутрішню чергу.
func (l *Loop) PostCord(c Cord) {
	l.mq.pushCord(c)
}

// сигналізує про необхідність завершити цикл та блокується до моменту його повної зупинки.
func (l *Loop) StopAndWait() {
	l.Post(OperationFunc(func(t screen.Texture, c Cord) { //викликає функцію яка виконується
		l.stopReq = true // вказує що потрібно завершити виконання операцій
	}))
	<-l.stop // передає сигнал
}

// черга повідомлень
type messageQueue struct {
	pushSignal chan struct{} //створення каналу для сигналу, потрібен щоб програма при доставанні
	// операцій з черги не зверталася до пустого масиву операцій

	pushCordSignal chan struct{}

	mu sync.Mutex // для роботи з потоками (закриття, відкритя)

	data []Operation // масив операцій де створюється черга

	dataCord []Cord // масив координат де створюється черга
}

// записує повідомлення у чергу
func (mq *messageQueue) push(op Operation) {
	mq.mu.Lock()         // якщо цю функцію викликають два потоки, то той потік який швидше виконав закриття,
	defer mq.mu.Unlock() //не дозволяє іншому взаємодіяти з функцією поки не закінчіть (відкриє)
	// потрібно для коректної взаємодії між потоками
	// інакше потоки будуть пошкоджувати один одному данні

	mq.data = append(mq.data, op) // додає у масив операцій операцію

	if mq.pushSignal != nil { // перевіряє чи був сигнал
		close(mq.pushSignal) // якщо сигнал був, то закриває канал (припиняє передачу сигналу)
		mq.pushSignal = nil  // вказуємо що канал немає сигналів
	}
}

// дістає повідомлення з черги
func (mq *messageQueue) pull() (Operation, Cord) {
	mq.mu.Lock()         // якщо цю функцію викликають два потоки, то той потік який швидше виконав закриття,
	defer mq.mu.Unlock() //не дозволяє іншому взаємодіяти з функцією поки не закінчіть (відкриє)

	for len(mq.data) == 0 { // перевіряємо чи масив операцій пустий (чи він nil)
		mq.pushSignal = make(chan struct{}) // ініціалізуємо канал
		mq.mu.Unlock()                      // відкриваємо потік щоб взаємодіяти з каналом
		<-mq.pushSignal                     // створює сигнал, після чого виконується push, який робить масив з однією операцією
		mq.mu.Lock()                        // закриває потік, після чого знову робить перевірку масиву (теоретичний результат, масив буде не пустим)
	}

	res := mq.data[0]     //дістає операцію з  масиву
	mq.data[0] = nil      //звільняє пам'ять у масиві де була взята операція
	mq.data = mq.data[1:] // каже що масив починається  з другого аргументу

	for len(mq.dataCord) == 0 { // перевіряємо чи масив операцій пустий (чи він nil)
		mq.pushCordSignal = make(chan struct{}) // ініціалізуємо канал
		mq.mu.Unlock()                          // відкриваємо потік щоб взаємодіяти з каналом
		<-mq.pushCordSignal                     // створює сигнал, після чого виконується push, який робить масив з однією операцією
		mq.mu.Lock()                            // закриває потік, після чого знову робить перевірку масиву (теоретичний результат, масив буде не пустим)
	}

	var c Cord
	if len(mq.dataCord) != 0 {
		//mq.mu.Unlock()
		//c = mq.pullCord()
		//mq.mu.Lock()
		c = mq.dataCord[0]                                //дістає операцію з  масиву
		mq.dataCord[0] = Cord{X1: 0, Y1: 0, X2: 0, Y2: 0} //звільняє пам'ять у масиві де була взята координата
		mq.dataCord = mq.dataCord[1:]
	}
	return res, c //  повертає операцію яку потрібно виконати
}

// перевіряє чи черга пуста
func (mq *messageQueue) empty() bool {
	mq.mu.Lock()         // якщо цю функцію викликають два потоки, то той потік який швидше виконав закриття,
	defer mq.mu.Unlock() //не дозволяє іншому взаємодіяти з функцією поки не закінчіть (відкриє)

	return len(mq.data)+len(mq.dataCord) == 0 // якщо черга пуста повертає true
}

func (mq *messageQueue) pushCord(c Cord) {
	fmt.Println("Pushed Coords: ", c.X1, c.Y1, c.X2, c.Y2)
	mq.dataCord = append(mq.dataCord, c) // додає у масив координат координати

	if mq.pushCordSignal != nil { // перевіряє чи був сигнал
		close(mq.pushCordSignal) // якщо сигнал був, то закриває канал (припиняє передачу сигналу)
		mq.pushCordSignal = nil  // вказуємо що канал немає сигналів
	}
}

/*func (mq *messageQueue) pullCord() Cord {
	mq.mu.Lock()         // якщо цю функцію викликають два потоки, то той потік який швидше виконав закриття,
	defer mq.mu.Unlock() //не дозволяє іншому взаємодіяти з функцією поки не закінчіть (відкриє)

	res := mq.dataCord[0]                             //дістає операцію з  масиву
	mq.dataCord[0] = Cord{X1: 0, Y1: 0, X2: 0, Y2: 0} //звільняє пам'ять у масиві де була взята координата
	mq.dataCord = mq.dataCord[1:]                     // каже що масив починається  з другого аргументу
	return res                                        //  повертає координати
}*/
