package utils

import (
	"sort"
)

type sortable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string
}

type sliceDesc[T sortable] []T

func (x sliceDesc[T]) Len() int           { return len(x) }
func (x sliceDesc[T]) Less(i, j int) bool { return x[i] > x[j] }
func (x sliceDesc[T]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x sliceDesc[T]) Sort() { sort.Sort(x) }

type sliceAsc[T sortable] []T

func (x sliceAsc[T]) Len() int           { return len(x) }
func (x sliceAsc[T]) Less(i, j int) bool { return x[i] < x[j] }
func (x sliceAsc[T]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x sliceAsc[T]) Sort() { sort.Sort(x) }

func SliceDesc[T sortable](s []T) {
	sliceDesc[T](s).Sort()
}

func SliceAsc[T sortable](s []T) {
	sliceAsc[T](s).Sort()
}
