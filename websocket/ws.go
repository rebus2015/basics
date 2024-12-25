package websocket

func selectMany(channels []chan int64) chan int64 {
	// ваш код
	res:=make(chan int64,len(channels))

	return res
}