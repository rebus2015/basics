package main

import (
	"fmt"
	"math"

	//"mathtest/structs"
	"strings"
	"sync"
	"unicode"
)

func main() {

	//fmt.Println(s)
	//D := distance(55.7539, 37.6208, 59.9398, 30.3146)
	//D := distance(64.28, 100.22, 40.71, 74.01)
	//fmt.Printf("Result: %v", D)
	//fmt.Printf("Factorial: %v", fr(5))
	//v := Abs(3)
	//fmt.Println(v)
	//fmt.Println(reverse([]int{1,2,3,7}))
	//toFrequencyMapAsync("Pipa pip pum pam pum pi pppap pim pup pip pi pi")
	StructChanProducer()
	// mapKeyIntersect(map[int]struct{}{1: {}, 2: {}, 3: {}},
	// 	map[int]struct{}{2: {}, 4: {}, 6: {}})
	// toFrequencyMap([]string{"a", "b", "a"})
	//fmt.Println(stringLengthWithoutSpaces("Это какой-то позор!	Или нет?"))
	// a := 3
	// b := 5
	// swap(&a, &b)
	// fmt.Println(a, b)
	// l := structs.NewHashTable(5)
	// l.Add("rirjfne", 1010)
	// l.Add("Arm", 23)
	// l.Add("Bwe", 456)
	// l.Add("ERT", 9234)
	// l.Add("efkjvbe", 4049)
	// l.Print()
	// fmt.Println()
	// l.Add("Bwe", 8000)
	// l.Print()
	// fmt.Println()
	// v, b := l.Get("efkjvbe")
	// fmt.Printf("key = \"efkjvbe\" , val = %v, %v \n", v, b)
	// fmt.Println()
	// fmt.Println("Removing element with key = \"ERT\"")
	// fmt.Printf("Delete result: %v\n", l.Delete("ERT"))
	// fmt.Println()
	// l.Print()

}

type MyStruct struct {
	A *int
	B int
}

func StructChanProducer() {
	in := make(chan MyStruct)
	out := make(chan MyStruct, 10)

	var wg sync.WaitGroup
	wg.Add(2)
	w1 := func(wg *sync.WaitGroup, channel chan<- MyStruct) {
		defer wg.Done()
		fmt.Println("Producer")
		l1 := 100
		s1 := MyStruct{&l1, 101}
		fmt.Println(s1)
		channel <- s1
	}
	//wg.Add(1)
	w2 := func(wg *sync.WaitGroup, channel <-chan MyStruct, myChan chan<- MyStruct) {
		defer wg.Done()
		fmt.Println("Consumer")
		s1 := <-channel

		*s1.A += 11
		myChan <- s1
	}
	go w1(&wg, in)
	go w2(&wg, in, out)
	wg.Wait()
	s2 := <-out
	fmt.Println(s2)

}

type Node struct {
	data interface{}
	next *Node
}

func swap(a *int, b *int) {
	// ваш код
	temp := *a
	*a = *b
	*b = temp
}

func stringLengthWithoutSpaces(str string) int {
	// ваш код
	s := []rune(str)
	l := len(s)
	for _, r := range s {
		if unicode.IsSpace(r) {
			l--
		}
	}
	return l
}

func frequentRune(str string) rune {
	// ваш код
	s := []rune(str)
	freqMap := make(map[rune]int)
	var res rune
	k := 0
	for _, r := range s {
		if _, ok := freqMap[r]; ok {
			freqMap[r]++
			if k < freqMap[r] {
				res = r
			}
		}
		freqMap[r] = 1
	}
	return res
}

func removeSpaces(s string) string {
	// ваш код
	str := []rune(s)
	var res []rune
	for _, r := range str {
		if !unicode.IsSpace(r) {
			res = append(res, r)
		}
	}
	return string(res)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func toFrequencyMap(s []string) map[string]int {
	// ваш код
	m := map[string]int{}
	for _, val := range s {
		if _, ok := m[val]; ok {
			m[val]++
			continue
		}
		m[val] = 1
	}
	fmt.Println(m)
	return m
}

func mapKeyIntersect(m1 map[int]struct{}, m2 map[int]struct{}) []int {
	// ваш код
	intersect := []int{}
	for k := range m1 {
		_, ok := m2[k]
		if ok {
			intersect = append(intersect, k)
		}
	}
	fmt.Println(intersect)
	return intersect
}

type sharedMap struct {
	words map[string]int
	mu    sync.RWMutex
}

func toFrequencyMapAsync(data string) map[string]int {
	// ваш код
	s := strings.Split(data, " ")
	m := sharedMap{
		words: make(map[string]int),
		mu:    sync.RWMutex{},
	}
	wg := sync.WaitGroup{}
	s1 := s[:len(s)/2]
	s2 := s[len(s)/2:]
	for i := 0; i < len(s); i += 2 {

	}

	wg.Add(1)
	go func(str []string) {
		for _, val := range str {
			m.mu.RLock()
			_, ok := m.words[val]
			m.mu.RUnlock()
			if ok {
				m.mu.Lock()
				m.words[val]++
				m.mu.Unlock()
				continue
			}
			m.mu.Lock()
			m.words[val] = 1
			m.mu.Unlock()
		}
		wg.Done()
	}(s1)

	wg.Add(1)
	go func(str []string) {
		for _, val := range str {
			m.mu.RLock()
			_, ok := m.words[val]
			m.mu.RUnlock()
			if ok {
				m.mu.Lock()
				m.words[val]++
				m.mu.Unlock()
				continue
			}
			m.mu.Lock()
			m.words[val] = 1
			m.mu.Unlock()
		}
		wg.Done()

	}(s2)
	wg.Wait()
	fmt.Println(m.words)
	return m.words
}

func removeDuplicates(slice []int) []int {
	// в аш код
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range slice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func reverseSlice(slice []int) []int {
	// ваш код
	n := len(slice)
	reverse := make([]int, n, n)
	j := 0
	for i := n; i != 0; i-- {
		reverse[j] = slice[i-1]
		j++
	}
	return reverse
}

// Abs возвращает абсолютное значение.
// Например: 3.1 => 3.1, -3.14 => 3.14, -0 => 0.
// Покрыть тестами нужно эту функцию.
func Abs(value float64) float64 {
	return math.Abs(value)
}

func f(x int) int {
	res := 1
	for i := 1; i <= x; i++ {
		res = res * i
	}
	return res
}

func fr(x int) int {
	if x == 1 {
		return 1
	}
	res := x * fr(x-1)
	return res
}

func distance(lat1, lon1 float32, lat2, lon2 float32) float32 {
	//var result float32
	R := 6371.009
	C := 3.14 / 180
	cosD := math.Sin(float64(lat1)*C)*math.Sin(float64(lat2)*C) +
		math.Cos(float64(lat1)*C)*math.Cos(float64(lat2)*C)*math.Cos(float64(lon1-lon2)*C)
	return float32(R * math.Acos(cosD))
}

func frequentWord(str string) string {
	// ваш код
	frequency := make(map[string](int))
	words := strings.Split(str, " ")
	var key string = ""
	var val int
	for _, word := range words {
		_, found := frequency[word]
		if found {
			frequency[word]++
			if frequency[word] > val {
				key = word
				val = frequency[word]
			}
			continue
		}
		frequency[word] = 1
	}

	fmt.Println(frequency)
	fmt.Println(key)
	return key
}
