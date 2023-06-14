package GPT13

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	c := Counter{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	if got := c.Get(); got != 100 {
		t.Errorf("want 100, got %d", got)
	}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Dec()
		}()
	}

	wg.Wait()

	if got := c.Get(); got != 50 {
		t.Errorf("want 50, got %d", got)
	}

	c.Reset()

	if got := c.Get(); got != 0 {
		t.Errorf("Counter.Get() = %d; want 0", got)
	}
}
