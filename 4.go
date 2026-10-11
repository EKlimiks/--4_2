package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	semofor := make(chan struct{}, 3)
	urls := []string{"https://ktk-45.ru/", "https://vk.ru/", "https://2gis.ru/", "https://www.twitch.tv/"}
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func() {
			semofor <- struct{}{}
			defer func() { <-semofor; wg.Done() }()
			r, errors := http.Get(url)
			if errors != nil {
				fmt.Println(errors)
				return
			}
			fmt.Printf("%v Статус %v\n", url, r.Status)
		}()
	}
	wg.Wait()
}
