// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
)

type request struct{
    action string
    respServer chan int
}

func server(reqChan chan request){
    var i = 0
    for{
        select{
        case req := <- reqChan:
            switch req.action{
            case "incr":
                i++
            case "dec":
                i--
            case "get":
                req.respServer <- i
            }
    
        }
    }
    
}

func incrementing(reqChan chan request, doneChan chan bool) {
    //TODO: increment i 1000000 times
    for j := 0; j < 1000000; j++ {
        reqChan <- request{action: "incr", respServer: nil}
    }
    doneChan <- true
}

func decrementing(reqChan chan request, doneChan chan bool) {
    //TODO: decrement i 1000000 times
    for j := 0; j < 1000000; j++ {
        reqChan <- request{action: "dec", respServer: nil}
    }
    doneChan <- true
}



func main() {
    // What does GOMAXPROCS do? What happens if you set it to 1?
    runtime.GOMAXPROCS(2)    
	
    // TODO: Spawn both functions as goroutines
    reqChan := make(chan request)
    doneChan := make(chan bool, 2)

    go server(reqChan)
    go incrementing(reqChan, doneChan)
    go decrementing(reqChan, doneChan)

    <-doneChan
    <-doneChan

    resChan := make(chan int)
    reqChan <- request{action: "get", respServer: resChan}
    i := <- resChan

	
    // We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
    // We will do it properly with channels soon. For now: Sleep.
    //time.Sleep(500*time.Millisecond)
    Println("The magic number is:", i)
}
