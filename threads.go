package main

import (
	"fmt"
	"time"
	
)
var c int
type Token struct {
	data      string
	recipient int
}

func initial(ch chan Token, token Token) {
	fmt.Println("initializing thread 0")
	ch <- token
}

func thread(past chan Token, next chan Token, num int) {
	fmt.Println("thread", num) 
	time.Sleep(1 * 1e1)
	token := <-past
	if token.recipient == num {
		fmt.Println(token.data, "received in", num, "-th thread")
		c=1
	} 
	next <- token

}

func main() {
	c=0
	k:=11
	n:=20
	var chanmass = make([]chan Token, n)
	var i int 
	for i = range chanmass{
		chanmass[i] = make(chan Token) 
	}
	token := Token{"data", k}
	
	go initial(chanmass[0], token)
	for i := 0; i < n-1; i++ {
		go thread(chanmass[i],chanmass[i+1], i)
	}
	go thread(chanmass[n-1],chanmass[0], n-1)
	time.Sleep(time.Second)
	if c==0 {
		fmt.Println("Wrong adress")
	}
}