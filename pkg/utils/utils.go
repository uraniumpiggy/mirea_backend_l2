package utils

import (
	"errors"
	"fmt"
	"strings"

	// "fmt"
	"strconv"
)

type Shape struct {
	Shape uint8
	Color string
	H, W  int
}

func DefineShape(num string) (*Shape, error) {
	number, err := strconv.ParseInt(num, 0, 0)
	if err != nil {
		return nil, errors.New("Parse error")
	}
	s := &Shape{}

	s.Shape = uint8(number & 1)

	var rawHexCode string = fmt.Sprintf("%x", number&16777215)
	if len(rawHexCode) == 6 {
		s.Color = rawHexCode
	} else {
		var temp = []byte{}
		for i := len(rawHexCode); i < 6; i++ {
			temp = append(temp, '0')
		}
		temp = append(temp, []byte(rawHexCode)...)
		s.Color = string(temp)
	}

	s.H = int(number&400) + 10
	number >>= 4
	s.W = int(number&400) + 10
	return s, nil
}

func StringToArray(str string) ([]int, error) {
	var res = []int{}
	for _, val := range strings.Split(str, ",") {
		num, err := strconv.Atoi(val)
		if err != nil {
			return []int{}, err
		}
		res = append(res, num)
	}
	return res, nil
}

func QSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := arr[(l+r)/2]
	i, j := l, r
	for i <= j {
		for arr[j] > mid {
			j--
		}
		for arr[i] < mid {
			i++
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	if j > l {
		QSort(arr, l, j)
	}
	if i < r {
		QSort(arr, i, r)
	}
}
