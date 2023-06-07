package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceAsc_Sort(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, func() []int {
		a := []int{1, 3, 2}
		sliceAsc[int](a).Sort()
		return a
	}())
	assert.Equal(t, []int8{1, 2, 3}, func() []int8 {
		a := []int8{1, 3, 2}
		sliceAsc[int8](a).Sort()
		return a
	}())
	assert.Equal(t, []int16{1, 2, 3}, func() []int16 {
		a := []int16{1, 3, 2}
		sliceAsc[int16](a).Sort()
		return a
	}())
	assert.Equal(t, []int32{1, 2, 3}, func() []int32 {
		a := []int32{1, 3, 2}
		sliceAsc[int32](a).Sort()
		return a
	}())
	assert.Equal(t, []int64{1, 2, 3}, func() []int64 {
		a := []int64{1, 3, 2}
		sliceAsc[int64](a).Sort()
		return a
	}())
	assert.Equal(t, []uint{1, 2, 3}, func() []uint {
		a := []uint{1, 3, 2}
		sliceAsc[uint](a).Sort()
		return a
	}())
	assert.Equal(t, []uint8{1, 2, 3}, func() []uint8 {
		a := []uint8{1, 3, 2}
		sliceAsc[uint8](a).Sort()
		return a
	}())
	assert.Equal(t, []uint16{1, 2, 3}, func() []uint16 {
		a := []uint16{1, 3, 2}
		sliceAsc[uint16](a).Sort()
		return a
	}())
	assert.Equal(t, []uint32{1, 2, 3}, func() []uint32 {
		a := []uint32{1, 3, 2}
		sliceAsc[uint32](a).Sort()
		return a
	}())
	assert.Equal(t, []uint64{1, 2, 3}, func() []uint64 {
		a := []uint64{1, 3, 2}
		sliceAsc[uint64](a).Sort()
		return a
	}())
	assert.Equal(t, []float32{1, 2, 3}, func() []float32 {
		a := []float32{1, 3, 2}
		sliceAsc[float32](a).Sort()
		return a
	}())
	assert.Equal(t, []float64{1, 2.0, 2.1}, func() []float64 {
		a := []float64{1, 2.1, 2.0}
		sliceAsc[float64](a).Sort()
		return a
	}())
}

func TestSliceDesc_Sort(t *testing.T) {
	assert.Equal(t, []int{3, 2, 1}, func() []int {
		a := []int{1, 3, 2}
		sliceDesc[int](a).Sort()
		return a
	}())
	assert.Equal(t, []int8{3, 2, 1}, func() []int8 {
		a := []int8{1, 3, 2}
		sliceDesc[int8](a).Sort()
		return a
	}())
	assert.Equal(t, []int16{3, 2, 1}, func() []int16 {
		a := []int16{1, 3, 2}
		sliceDesc[int16](a).Sort()
		return a
	}())
	assert.Equal(t, []int32{3, 2, 1}, func() []int32 {
		a := []int32{1, 3, 2}
		sliceDesc[int32](a).Sort()
		return a
	}())
	assert.Equal(t, []int64{3, 2, 1}, func() []int64 {
		a := []int64{1, 3, 2}
		sliceDesc[int64](a).Sort()
		return a
	}())
	assert.Equal(t, []uint{3, 2, 1}, func() []uint {
		a := []uint{1, 3, 2}
		sliceDesc[uint](a).Sort()
		return a
	}())
	assert.Equal(t, []uint8{3, 2, 1}, func() []uint8 {
		a := []uint8{1, 3, 2}
		sliceDesc[uint8](a).Sort()
		return a
	}())
	assert.Equal(t, []uint16{3, 2, 1}, func() []uint16 {
		a := []uint16{1, 3, 2}
		sliceDesc[uint16](a).Sort()
		return a
	}())
	assert.Equal(t, []uint32{3, 2, 1}, func() []uint32 {
		a := []uint32{1, 3, 2}
		sliceDesc[uint32](a).Sort()
		return a
	}())
	assert.Equal(t, []uint64{3, 2, 1}, func() []uint64 {
		a := []uint64{1, 3, 2}
		sliceDesc[uint64](a).Sort()
		return a
	}())
	assert.Equal(t, []float32{3, 2, 1}, func() []float32 {
		a := []float32{1, 3, 2}
		sliceDesc[float32](a).Sort()
		return a
	}())
	assert.Equal(t, []float64{2.1, 2.0, 1}, func() []float64 {
		a := []float64{1, 2.1, 2.0}
		sliceDesc[float64](a).Sort()
		return a
	}())
}

func TestAscSort(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3"}, func() []string {
		a := []string{"1", "3", "2"}
		SliceAsc(a)
		return a
	}())
}

func TestDescSort(t *testing.T) {
	assert.Equal(t, []string{"3", "2", "1"}, func() []string {
		a := []string{"1", "3", "2"}
		SliceDesc(a)
		return a
	}())
}
