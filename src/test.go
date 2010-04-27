package main


import "time"
import "fmt"

var ch  = make(chan string, 10 )
var x string = "init"

func f() {
	fmt.Printf( "before send f sees: \"%s\" in x\n", x )
	time.Sleep( 1000000 )
	x = "hello from f"
	ch <- "f"
	fmt.Printf( "after send f sees \"%s\" in x\n", x )
}

func g() {
	fmt.Printf( "before send g sees: \"%s\" in x\n", x )
	x = "hello from g"
	ch <- "g"
	fmt.Printf( "after send g sees: \"%s\" in x\n", x )
}

func main() {
	var response string
	go f()	
	go g()
	response = <- ch
	fmt.Printf( "After first recieve main hears from goroutine %s and sees \"%s\" in x\n", response, x )
	response = <- ch
	fmt.Printf( "After second recieve main hears from goroutine %s sees \"%s\" in x\n", response, x )
}
