package main

import "homeworks-go-level2/internal/greeting"

func main() {
	greeting := greeting.Greeting{"John"}
	greeting.SayHello()
}
