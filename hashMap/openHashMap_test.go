package hashMap
import (
	"testing"
	"reflect"
	"fmt"
)

func Test_OpenHashMap(t *testing.T) {
	cmap := NewOpenHashMap()
	BasicTestHashMap(t,cmap)
}

func Test_OpenHashMapResize(t *testing.T) {
	cmap := NewOpenHashMap()
	TestHashMapResize(t, cmap)
	if !reflect.DeepEqual(cmap.Cap, uint32(16)) {
		t.Log(fmt.Sprintf("expect: %0d ", uint32(16)) + fmt.Sprintf("but get: %0d ", cmap.Cap))
		t.Fail()
	}
}

func Test_OpenHashMapDelete(t *testing.T) {
	cmap := NewOpenHashMap()
	for i:=0;i<4; {
		TestHashMapDelete(t, cmap)
		if !reflect.DeepEqual(cmap.Count, uint32(0)) {
			t.Log(fmt.Sprintf("expect: %0d ", 0) + fmt.Sprintf("but get:%0d ", cmap.Count))
			t.Fail()
		}
		if !reflect.DeepEqual(cmap.Cap, uint32(0)) {
			t.Log(fmt.Sprintf("expect: %0d ", 0) + fmt.Sprintf("but get:%0d ", cmap.Cap))
			t.Fail()
		}
		i++
	}
}

func BenchmarkOpenHashMap_HashInsert(b *testing.B) {
	BenchmarkHashMapInsert(b, NewOpenHashMap())
}

func BenchmarkOpenHashMap_HashInsertDelete(b *testing.B) {
	BenchmarkHashMapInsertDelete(b, NewOpenHashMap())
}

func BenchmarkOpenHashMap_HashGet(b *testing.B) {
	BenchmarkHashMapGet(b, NewOpenHashMap())
}