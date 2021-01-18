package main

/* Intentionally commented out to simulate failure. We will fix this as part of the exercises :) */
/*
import (
	"fmt"
	"os"
)
*/

func main() {
	hello_from := os.Getenv("HELLO_FROM")

	if hello_from == "" {
		hello_from = "our GitLab meetup"
	}

	fmt.Println("Hello from " + hello_from + "!")
	fmt.Println("Today we learn about GitLab Best Practices and CI/CD magic :)")

	fmt.Println("");

	fmt.Print(GetTanuki(true))

	fmt.Println("");
	fmt.Println("Join us at https://about.gitlab.com/community/")
}

func CreateTmpFile() {
  f, _ := os.Create("tanuki.tmp")
  defer f.Close()

  os.Chmod("tanuki.tmp", 0777)
}
