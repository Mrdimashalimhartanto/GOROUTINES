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
	// todo: tambahin go untuk lakuin go routine test
	go RunHelloWord()
	fmt.Println("Hello")

	time.Sleep(2 * time.Second)

}

func TestNasabah(t *testing.T) {
	go HalloNasabah()
	fmt.Println("NASABAH ADA")

	time.Sleep(1 * time.Second)
}

func NomorPolis(number int) {
	fmt.Println("Nomor Polis", number)
}

func TestNomorPolis(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		go NomorPolis(i)
	}
	time.Sleep(5 * time.Second)
}
