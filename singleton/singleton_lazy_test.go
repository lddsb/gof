package singleton_test

import (
	"github.com/lddsb/gof/singleton"
	"testing"
)

func TestGetLazyInstance(t *testing.T) {
	ins1 := singleton.GetInstance()
	ins2 := singleton.GetInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func BenchmarkGetLazyInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetLazyInstance() != singleton.GetLazyInstance() {
				b.Errorf("test failed")
			}
		}
	})
}
