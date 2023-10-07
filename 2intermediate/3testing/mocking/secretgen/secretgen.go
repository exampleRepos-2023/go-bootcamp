package secretgen

import (
	"math/rand"
	"time"
)

type SG interface {
	Get() int
}

type SecretGen struct {
}

func (sg SecretGen) Get() int {
	time.Sleep(3 * time.Second)
	rand.Seed(time.Now().Unix())
	return rand.Int()
}
