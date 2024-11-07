package example_test

import (
	"sync"
	"testing"

	"github.com/sivaosorg/unify4g"
)

func TestGenerateUUID(t *testing.T) {
	var wg sync.WaitGroup
	numTests := 100
	uniqueIDs := make(map[string]struct{})
	mu := sync.Mutex{}

	for i := 0; i < numTests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id, _ := unify4g.GenerateUUID()
			mu.Lock()
			if _, exists := uniqueIDs[id]; exists {
				t.Errorf("Duplicate ID generated: %s", id)
			}
			uniqueIDs[id] = struct{}{}
			mu.Unlock()
		}()
	}

	wg.Wait()
	if len(uniqueIDs) != numTests {
		t.Errorf("Expected %d unique IDs but got %d", numTests, len(uniqueIDs))
	}
}

func TestGenerateTimestampID(t *testing.T) {
	var wg sync.WaitGroup
	numTests := 100
	uniqueIDs := make(map[string]struct{})
	mu := sync.Mutex{}

	for i := 0; i < numTests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := unify4g.GenerateTimestampID()
			mu.Lock()
			if _, exists := uniqueIDs[id]; exists {
				t.Errorf("Duplicate ID generated: %s", id)
			}
			uniqueIDs[id] = struct{}{}
			mu.Unlock()
		}()
	}

	wg.Wait()
	if len(uniqueIDs) != numTests {
		t.Errorf("Expected %d unique IDs but got %d", numTests, len(uniqueIDs))
	}
}

// Recommended
func TestGenerateCryptoID(t *testing.T) {
	var wg sync.WaitGroup
	numTests := 100
	uniqueIDs := make(map[string]struct{})
	mu := sync.Mutex{}

	for i := 0; i < numTests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := unify4g.GenerateCryptoID()
			mu.Lock()
			if _, exists := uniqueIDs[id]; exists {
				t.Errorf("Duplicate ID generated: %s", id)
			}
			uniqueIDs[id] = struct{}{}
			mu.Unlock()
		}()
	}

	wg.Wait()
	if len(uniqueIDs) != numTests {
		t.Errorf("Expected %d unique IDs but got %d", numTests, len(uniqueIDs))
	}
}
