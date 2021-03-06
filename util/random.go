package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

//RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//RandomOwner generates a random owner name
func RandomName() string {
	return RandomString(6)
}

//RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func RandomNumber() int64 {
	return int64(rangeIn(10000000, 99999999))
}

func RandomAddress() string {
	return RandomString(10)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomStatus() string {
	status := []string{"Delivered", "Hold"}
	n := len(status)

	return status[rand.Intn(n)]
}

func RandomPizzaType() string {
	p_type := []string{"Cheese", "Veggie", "Pepperoni", "Meat", "Margherita", "Hawaiian", "Buffalo"}
	n := len(p_type)

	return p_type[rand.Intn(n)]
}

func RandomPaymentStatus() string {
	payment := []string{"Paid", "NotPaid"}
	n := len(payment)

	return payment[rand.Intn(n)]
}
