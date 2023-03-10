package datagen

import (
	"math/rand"
	"sync"
	"time"
)

var rnd *random

type random struct {
	r     *rand.Rand
	mutex sync.Mutex
}

func Rand() *random {
	if rnd == nil {
		rnd = &random{
			r:     rand.New(rand.NewSource(time.Now().UnixNano())),
			mutex: sync.Mutex{},
		}
	}
	return rnd
}

func (r *random) Int(max int) int {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.r.Intn(max)
}

func (r *random) IntBetween(min, max int) int {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.r.Intn(max-min) + min
}

func (r *random) Float64(max float64) float64 {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.r.Float64() * max
}
