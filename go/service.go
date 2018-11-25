package main

import "fmt"

func getNotification(user string) chan string {
	/*
	 * 此处你可以干一些事情，如获取消息等
	 */
	notification := make(chan string)

	go func() {
		notification <- fmt.Sprintf("Hi %s, welcome to weibo.com!", user)
	}()

	return notification
}

func main() {
	jack := getNotification("jack")
	joe := getNotification("joe")

	fmt.Print(<-jack)
	fmt.Print(<-joe)
}
