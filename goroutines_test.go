package GOROUTINES

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWord() {
	fmt.Println("Ups")
}

func HalloNasabah() {
	fmt.Println("hallo nasabah")
}

func TestGoroutine(t *testing.T) {
	go RunHelloWord()
	fmt.Println("Hello")

	time.Sleep(2 * time.Second)

}

func TestNasabah(t *testing.T) {
	go HalloNasabah()
	fmt.Println("NASABAH ADA")

	time.Sleep(1 * time.Second)
}
