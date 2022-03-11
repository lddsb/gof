package singleton_test

import (
	"github.com/lddsb/gof/singleton"
	"testing"
)

func TestGetInstance(t *testing.T) {
	ins1 := singleton.GetInstance()
	ins2 := singleton.GetInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetInstance() != singleton.GetInstance() {
				b.Errorf("test failed")
			}
		}
	})
}
