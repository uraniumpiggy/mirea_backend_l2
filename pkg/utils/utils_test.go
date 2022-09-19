package utils

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestQSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		arr := []int{}
		comp := []int{}
		for j := 0; j < rand.Intn(50); j++ {
			temp := rand.Intn(1000) - 500
			arr = append(arr, temp)
			comp = append(comp, temp)
		}
		sort.Slice(comp, func(i, j int) bool {
			return comp[i] < comp[j]
		})
		QSort(arr, 0, len(arr)-1)
		fmt.Println(arr, comp)
		if !reflect.DeepEqual(arr, comp) {
			t.Error(arr, comp)
		}
	}
}
