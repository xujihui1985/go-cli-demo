package main

import (
    "github.com/deckarep/gosx-notifier"
	"fmt"
)

func main() {
	note := gosxnotifier.NewNotification("hello")
	note.Title = "xxxxx"
	note.Subtitle = "subtitle"
	note.Group = "io.sean.unique.xxxx"
	err := note.Push()
	if err != nil {
		fmt.Println("error xxxx")
	}
}
