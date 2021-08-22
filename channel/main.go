package main

import (
	"fmt"
)

// type User struct {
// 	ID        int       `json:"id"`
// 	Email     string    `json:"email"`
// 	Name      string    `json:"name"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }

// type Sub struct {
// 	userChannels map[int]chan *User
// }

// func main() {
// 	// userChannel := make(map[string]chan *User, 2)
// 	userChan := make(chan *User, 1)

// 	sub := &Sub{
// 		userChannels: map[int]chan *User{},
// 	}

// 	for i := 0; i < 10; i++ {
// 		sub.userChannels[i] = userChan
// 	}

// 	fmt.Println(&sub.userChannels)
// }

// func main() {
// 	jobs := make(chan int, 5)
// 	done := make(chan bool)

// 	go func() {
// 		for {
// 			j, more := <-jobs
// 			if more {
// 				fmt.Println("received job", j)
// 			} else {
// 				fmt.Println("received all jobs")
// 				done <- true
// 				return
// 			}
// 		}
// 	}()

// 	for j := 1; j <= 5; j++ {
// 		jobs <- j
// 		fmt.Println("sent job", j)
// 	}
// 	close(jobs)
// 	fmt.Println("sent all jobs")

// 	<-done
// }

// チャネルに値を渡す
func sendValue(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}

// チャネルから値を受け取る
func receiveValue(c <-chan int) {
	for v := range c {
		fmt.Println("チャネルから受け取った値：", v)
	}
}

func main() {
	c := make(chan int)
	go sendValue(c)
	receiveValue(c)

	fmt.Println("処理を終了します。")
}
