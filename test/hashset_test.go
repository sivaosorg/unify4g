package example_test

import (
	"testing"

	"github.com/sivaosorg/unify4go"
)

func TestNewHashSet(t *testing.T) {
	hashSet := unify4go.NewHashSet[int]()
	if hashSet == nil {
		t.Errorf("Hashset is nil")
		return
	}
}

func TestHashSet_Add(t *testing.T) {
	hashSet := unify4go.NewHashSet[string]()
	hashSet.Add("test1")
	hashSet.Add("test2")
	hashSet.Add("test3")
	hashSet.Add("test4")

	if hashSet.Size() != 4 {
		t.Errorf("Expected size of set to be %d but got %d", 4, hashSet.Size())
	}
}

func TestHashSet_Remove(t *testing.T) {
	hashSet := unify4go.NewHashSet[string]()
	hashSet.Add("test1")
	hashSet.Add("test2")
	hashSet.Remove("test2")
	hashSet.Add("test3")

	if hashSet.Size() != 2 {
		t.Errorf("Expected size of set to be %d but got %d", 2, hashSet.Size())
	}

	if hashSet.Contains("test") {
		t.Errorf("Expected set to not contain test")
	}
}

func TestHashSet_AddAll(t *testing.T) {
	hashSet := unify4go.NewHashSet[string]()
	hashSet.AddAll("test1", "test2", "test3", "test4")

	if hashSet.Size() != 4 {
		t.Errorf("Expected size of set to be %d but got %d", 4, hashSet.Size())
	}
}

func TestHashSet_RemoveAll(t *testing.T) {
	hashSet := unify4go.NewHashSet[string]()
	hashSet.AddAll("test1", "test2", "test3", "test4")
	hashSet.RemoveAll("test1", "test3")
	hashSet.RemoveAll("test1", "test2", "test4")

	if !hashSet.IsEmpty() {
		t.Errorf("Expected set to be empty")
	}
}

func TestHashSet_Clear(t *testing.T) {
	hashSet := unify4go.NewHashSet[string]()
	hashSet.AddAll("test1", "test2", "test3", "test4")
	hashSet.Clear()

	if !hashSet.IsEmpty() {
		t.Errorf("Expected set to be empty")
	}
}

func TestHashSet_IntersectionFirstSetBiggerSize(t *testing.T) {
	hashSetA := unify4go.NewHashSet[int](1, 2, 3, 4)
	hashSetB := unify4go.NewHashSet[int](2, 5)

	NewHashSet := hashSetA.Intersection(hashSetB)

	if NewHashSet.Size() != 1 {
		t.Errorf("Expected intersection size to be %d but got %d", 1, NewHashSet.Size())
	}

	if !NewHashSet.Contains(2) {
		t.Errorf("Expected intersection to contain value 2")
	}

	if hashSetA.Size() != 4 {
		t.Errorf("Original set should remain unchanged after operation")
	}

	if hashSetB.Size() != 2 {
		t.Errorf("Original set should remain unchanged after operation")
	}
}

func TestHashSet_IntersectionFirstSetSmallerSize(t *testing.T) {
	hashSetA := unify4go.NewHashSet[int](2, 5)
	hashSetB := unify4go.NewHashSet[int](1, 2, 3, 4)

	NewHashSet := hashSetA.Intersection(hashSetB)

	if NewHashSet.Size() != 1 {
		t.Errorf("Expected intersection size to be %d but got %d", 1, NewHashSet.Size())
	}

	if !NewHashSet.Contains(2) {
		t.Errorf("Expected intersection to contain value 2")
	}

	if hashSetA.Size() != 2 {
		t.Errorf("Original set should remain unchanged after operation")
	}

	if hashSetB.Size() != 4 {
		t.Errorf("Original set should remain unchanged after operation")
	}
}

func TestHashSet_Union(t *testing.T) {
	hashSetA := unify4go.NewHashSet[int](1, 2, 4)
	hashSetB := unify4go.NewHashSet[int](2, 3)

	NewHashSet := hashSetA.Union(hashSetB)

	if NewHashSet.Size() != 4 {
		t.Errorf("Expected union size to be %d but got %d", 4, NewHashSet.Size())
	}

	if !NewHashSet.Contains(1) {
		t.Errorf("Expected union to contain value 1")
	}

	if !NewHashSet.Contains(2) {
		t.Errorf("Expected union to contain value 2")
	}

	if !NewHashSet.Contains(3) {
		t.Errorf("Expected union to contain value 3")
	}

	if !NewHashSet.Contains(4) {
		t.Errorf("Expected union to contain value 4")
	}

	if hashSetA.Size() != 3 {
		t.Errorf("Original set should remain unchanged after operation")
	}

	if hashSetB.Size() != 2 {
		t.Errorf("Original set should remain unchanged after operation")
	}

}

func TestHashSet_Difference(t *testing.T) {
	hashSetA := unify4go.NewHashSet[int](1, 2, 4)
	hashSetB := unify4go.NewHashSet[int](2, 3)

	NewHashSet := hashSetA.Difference(hashSetB)

	if NewHashSet.Size() != 2 {
		t.Errorf("Expected difference size to be %d but got %d", 2, NewHashSet.Size())
	}

	if !NewHashSet.Contains(1) {
		t.Errorf("Expected union to contain value 1")
	}

	if !NewHashSet.Contains(4) {
		t.Errorf("Expected union to contain value 4")
	}

	if hashSetA.Size() != 3 {
		t.Errorf("Original set should remain unchanged after operation")
	}

	if hashSetB.Size() != 2 {
		t.Errorf("Original set should remain unchanged after operation")
	}
}

func TestHashSet_ToString(t *testing.T) {
	hashSet := unify4go.NewHashSet[string]("a", "b")
	expectedString1 := "a,b"
	expectedString2 := "b,a"

	result := hashSet.String()

	if result != expectedString1 && result != expectedString2 {
		t.Errorf("Expected string to be %s or %s but got %s", expectedString1, expectedString2, result)
	}
}

func TestHashSet_ToSlice(t *testing.T) {
	hashSet := unify4go.NewHashSet[int](1, 2, 3)
	slice := hashSet.Slice()

	if len(slice) != hashSet.Size() {
		t.Errorf("Expected slice length to be %d but got %d", hashSet.Size(), len(slice))
	}
}

func BenchmarkHashSetAdd100(b *testing.B) {
	hashSet := unify4go.NewHashSet[int]()
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			hashSet.Add(j)
		}
	}
}

func BenchmarkHashSetAdd10000(b *testing.B) {
	hashSet := unify4go.NewHashSet[int]()
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			hashSet.Add(j)
		}
	}
}

func BenchmarkHashSetAdd1000000(b *testing.B) {
	hashSet := unify4go.NewHashSet[int]()
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000000; j++ {
			hashSet.Add(j)
		}
	}
}

func BenchmarkHashSetRemove100(b *testing.B) {
	hashSet := unify4go.NewHashSet[int]()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for j := 0; j < 100; j++ {
			hashSet.Add(j)
		}
		b.StartTimer()
		for j := 0; j < 100; j++ {
			hashSet.Remove(j)
		}
	}
}

func BenchmarkHashSetRemove10000(b *testing.B) {
	hashSet := unify4go.NewHashSet[int]()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for j := 0; j < 10000; j++ {
			hashSet.Add(j)
		}
		b.StartTimer()
		for j := 0; j < 10000; j++ {
			hashSet.Remove(j)
		}
	}
}

func BenchmarkHashSetRemove1000000(b *testing.B) {
	hashSet := unify4go.NewHashSet[int]()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for j := 0; j < 1000000; j++ {
			hashSet.Add(j)
		}
		b.StartTimer()
		for j := 0; j < 1000000; j++ {
			hashSet.Remove(j)
		}
	}
}

func BenchmarkHashSetContains100(b *testing.B) {
	hashSet := unify4go.NewHashSet[int]()
	for j := 0; j < 100; j++ {
		hashSet.Add(j)
	}
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = hashSet.Contains(50)
	}
}

func BenchmarkHashSetContains10000(b *testing.B) {
	hashSet := unify4go.NewHashSet[int]()
	for j := 0; j < 10000; j++ {
		hashSet.Add(j)
	}
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = hashSet.Contains(5000)
	}
}

func BenchmarkHashSetContains1000000(b *testing.B) {
	hashSet := unify4go.NewHashSet[int]()
	for j := 0; j < 1000000; j++ {
		hashSet.Add(j)
	}
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = hashSet.Contains(500000)
	}
}
