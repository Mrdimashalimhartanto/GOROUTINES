package GOROUTINES

import (
	"fmt"
	"testing"
	"time"
)

func IdentitasNasabah(number int) {
	fmt.Println("TOTAL NASABAH", number)
}

func TestDataNasabah(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		go IdentitasNasabah(i)
	}
	time.Sleep(2 * time.Second)
}
