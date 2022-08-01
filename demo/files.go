package demo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func FileDemo() {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(2)
	//start go routines

	go printFile(ch, &wg)
	go readFile(ch, &wg)

	wg.Wait()
}

func readFile(ch chan<- string, wg *sync.WaitGroup) {
	file, err := os.Open("main.go")
	if err != nil {
		log.Fatal("error opening a file")
	}
	defer file.Close()

	// var bytes []byte = make([]byte, 5)

	// for {
	// 	data, err := file.Read(bytes)
	// 	if err != nil || err == io.EOF {
	// 		//fmt.Println(err)
	// 		break
	// 	}
	// 	ch <- string(bytes[:data])
	// }

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		ch <- sc.Text()
	}
	close(ch)
	wg.Done()
}

func printFile(ch <-chan string, wg *sync.WaitGroup) {
	for st := range ch {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(st)
	}
	wg.Done()
}
