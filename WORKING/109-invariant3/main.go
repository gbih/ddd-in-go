package main

import "fmt"

type Session struct {
	ready  bool
	orders []int
}

func (s Session) createSession(i []int) interface{} {
	var session Session
	fmt.Println("i is ", i)
	if len(i) > 0 {
		fmt.Println("OK")
		session = Session{orders: i, ready: true}
		return session
	}
	panic("No proper data")
}

//---------------

func main() {
	fmt.Println("--------------------------")

	var t *Session
	fmt.Printf("%#v\n", t)

	// var a = (*Session)(nil) // (*main.Session)(nil)
	// fmt.Printf("%#v\n", a)

	b := (*Session)(nil) // (*main.Session)(nil)
	fmt.Printf("%#v\n", b)

	if t == nil {
		t.createSession([]int{1, 2, 3})
	}

	if b == nil {
		b = new(Session)
		b.orders = []int{3, 3, 4}
		b.ready = true
	}
	fmt.Println("t is ", t)
	// if (*Session)(nil) {
	// 	fmt.Println("NIL")
	// }

	// var t *Session
	// fmt.Printf("%v", t)

	// if s == nil {
	// 	s = new(Session)
	// 	// do stuff (populate s)
	// }

	// var s Session
	// s.createSession([]int{1, 2, 3})
	// result := s.createSession([]int{1, 2, 3})
	// fmt.Println(result)
}
