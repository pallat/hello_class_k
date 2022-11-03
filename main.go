package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pallat/hello_kbao/generator"
)

func addAll[T int | int32 | float64](a ...T) T {
	var r T
	for _, v := range a {
		r += v
	}
	return r
}

func mainOfGeneric() {
	n := []int32{1, 2, 3, 4, 5, 6}

	fmt.Println(addAll(n...))
}

func mainOfEvenOdd() {
	evenCh := make(chan int)
	oddCh := make(chan int)
	even := newLabelWithNumber("even")
	odd := newLabelWithNumber("odd")

	go even(evenCh)
	go odd(oddCh)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
}

func newLabelWithNumber(label string) func(ch chan int) {
	return func(ch chan int) {
		for {
			fmt.Printf("%s: %d\n", label, <-ch)
		}
	}
}

func mainOfChanFibonacci() {
	quitSig := make(chan struct{})
	ch := make(chan int)
	go fibonacciCH(ch, quitSig)

	for i := 0; i < 12; i++ {
		fmt.Println(<-ch)
	}

	quitSig <- struct{}{}
}

func fibonacciCH(ch chan<- int, quitSig chan struct{}) {
	a, b := 0, 1
	for {
		select {
		case ch <- a:
			a, b = b, a+b
		case <-quitSig:
			fmt.Println("gracefully")
			return
		}
	}
}

func mainOfChannel() {
	var ch chan int

	ch = make(chan int, 2)

	ch <- 1
	ch <- 2

	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

var wg sync.WaitGroup

func mainOfGo() {
	wg.Add(3)

	start := time.Now()
	go lazyPrint("1")
	go lazyPrint("2")
	go lazyPrint("3")

	wg.Wait()
	fmt.Println(time.Since(start))
}

func lazyPrint(s string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(s)
	wg.Done()
}

func x() func() int {
	a := 0
	return func() int {
		defer func() { a++ }()
		return a
	}
}

func mainOfClosure() {
	f := fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

type intFunc func() int

func fibonacci() intFunc {
	a, b := 0, 1

	return func() int {
		defer func() {
			a, b = b, a+b
		}()

		return a
	}
}

func add(a, b int) int {
	return a + b
}

func one() int {
	return 1
}
func two() int {
	return 2
}

func hof(a, b func() int) func() int {
	return func() int {
		return a() + b()
	}
}

func mainOgHOF() {
	var plus = func(a, b int) int {
		return a + b
	}
	fmt.Println(plus(1, 2))

	r := hof(one, two)
	fmt.Println(r())
}

type Bool bool

func (*Bool) String() string {
	return "I'm bool"
}

func Println(a any) error {
	switch v := a.(type) {
	case fmt.Stringer:
		fmt.Println(v.String())
	}
	return nil
}

func mainOfInterface() {
	var b bool
	fmt.Println(b)

	var i *Bool

	// i = Bool(true)

	fmt.Println(i)
}

func mainOfDefer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	log.Fatal("fatal")

	risk()
	fmt.Println(double(4))

}

func risk() {

	var s []Int
	s[0] = 10
}

func double(n int) (nn int) {
	defer func() {
		nn = n * 2
	}()

	return
}

type user struct {
	UserID int64  `json:"userId"`
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func mainOfStructUser() {
	var u user
	if err := json.Unmarshal([]byte(jsonString), &u); err != nil {
		log.Panic(err)
	}

	fmt.Printf("%#v\n", u)
}

type Int int

func (i Int) String(int) string {
	return strconv.Itoa(int(i))
}

func (i *Int) parseString(s string) {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	*i = Int(n)
}

func mainOfTypeInt() {
	var i Int
	i.parseString("16")

	fmt.Printf("%d\n", i)
}

type String string

func (s *String) toLowerCase() {
	*s = String(strings.ToLower(string(*s)))
}
func mainOfLowercase() {
	var s String = "PALLAT"
	s.toLowerCase()
	fmt.Println(s)

}

var jsonString = `{
	"userId": 1,
	"id": 1,
	"title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
	"body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
  }`

func mainOfJsonParse() {

	var m map[string]any

	if err := json.Unmarshal([]byte(jsonString), &m); err != nil {
		log.Panic(err)
	}

	fmt.Printf("%#v\n", m)
	fmt.Println(m["userId"])
}

func mainOfCouple() {
	s := "abcdefghi"
	fmt.Println(couple(s))
}

func couple(s string) (c []string) {
	for s += "*"; len(s) > 1; s = s[2:] {
		c = append(c, s[:2])
	}
	return
}

func mainOfCap() {
	primes := [...]int{2, 3, 5, 7, 11}

	slice := primes[0:3:3]

	fmt.Printf("%T %T %v %d\n", primes, slice, slice, cap(slice))

	slice = append(slice, 4)

	fmt.Printf("%T %T %v %d\n", primes, slice, slice, cap(slice))
	fmt.Println(primes[3])
}

func mainOfSlice() {

	// primes := [...]int{2, 3, 5, 7, 11}
	var primes []int

	primes = make([]int, 2, 5) // array 2 element on 5 cap

	primes[0] = 1
	fmt.Println(primes)
	fmt.Printf("%p len=%d cap=%d", primes, len(primes), cap(primes))

	primes = append(primes, 3, 6, 9, 10)
	fmt.Println(primes)
	fmt.Printf("%p len=%d cap=%d", primes, len(primes), cap(primes))

	// for i, prime := range primes {
	// 	fmt.Println(prime)
	// 	primes[i] = 0
	// }

	fmt.Printf("%T %v\n", primes, primes)
}

func mainOfPackage() {
	fmt.Println(generator.Title)
}

func mainOfPointer() {
	var p *int
	i := 10
	p = &i

	fmt.Println(p, *p)

	var s string
	flag.StringVar(&s, "name", "", "input your name")
}

func power(b, x int) int {
	r := 1
	for i := 0; i < x; i++ {
		r *= b
	}
	return r
}

func prime(max int) {
	for i := 2; i <= max; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			fmt.Print(i, " ")
		}
	}
}

func IsCorrect() bool {
	return true
}

func hello(name string) {
	fmt.Println("Hello", name)
}

func squareArea(a float64) float64 {
	return a * a
}
