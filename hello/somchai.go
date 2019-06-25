package hello

import "github.com/endkarn/demo-go-module/message"

func Greeting() string {
	return message.GoodAfternoon() + " Somchai."
}
