package main

import (
	"fmt"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Sub struct {
	userChannels map[int]chan *User
}

func main() {
	// userChannel := make(map[string]chan *User, 2)
	userChan := make(chan *User, 1)

	sub := &Sub{
		userChannels: map[int]chan *User{},
	}

	for i := 0; i < 10; i++ {
		sub.userChannels[i] = userChan
	}

	fmt.Println(&sub.userChannels)
}
