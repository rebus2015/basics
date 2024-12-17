package homework

type semaphore chan struct{}

func NewSemaphore(n int) semaphore {
	return make(semaphore, n)
}

func (s semaphore) Acquire(n int) {
	// ваш код
	for i := 0; i < n; i++ {
		s <- struct{}{}
	}

}

func (s semaphore) Release(n int) {
	// ваш код
	for i := 0; i < n; i++ {
		<-s
	}
}
