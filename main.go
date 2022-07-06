package main

import (
	"fmt"
	"github.com/00mohamad00/url-shortener/internal/service"
	"github.com/go-redis/redis/v8"
)

func main() {
	redisClient := *redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	urlShrotenerService := service.NewUrlShortener(redisClient)
	token, err := urlShrotenerService.SetUrl("google.com")
	if err != nil {
		fmt.Print("We have error!!!")
		return
	}
	fmt.Printf("Token for google.com is %s \n", token)

	url, err := urlShrotenerService.GetUrl(token)
	if err != nil {
		fmt.Print("We have error!!!")
		return
	}
	fmt.Printf("Url of token %s is %s \n", token, url)
}
