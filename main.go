package main

import (
	"fmt"
)

func otherGreetings() func() string {
	return func() string {
		return "Hello World!!!"
	}
}

func main(){
	greetings := func() {
		fmt.Println("Hi There!!!")
	}

	greetAgain := otherGreetings()

	greetings()

	fmt.Println(greetAgain())
}