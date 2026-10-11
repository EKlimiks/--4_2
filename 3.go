package main

import ("fmt"
		"time")

func main() {
    semofor := make(chan struct{}, 5)
		tick := time.Tick(200 * time.Millisecond)

	for i := 1; i <=15; i++ {
		go func () {
		semofor <- struct{}{}
		defer func() {<-semofor}()
		fmt.Println("Запрос выполнен",i)
		}()
		<-tick
	}
		}
	
	


