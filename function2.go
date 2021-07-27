package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result, err := sqrt(26)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined")
	}

	return math.Sqrt(x), nil

}
