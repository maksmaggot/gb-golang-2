package main

import (
	"fmt"
	"time"
)

type CastToFloatError struct {
	time.Time
}

func NewCastToFloatError(time time.Time) error {
	return &CastToFloatError{time}
}

func (c *CastToFloatError) Error() string {
	return fmt.Sprintf("can't cast to float, error time: %s ", c.String())
}

func main() {

	defer func() {
		if v := recover(); v != nil {
			castErr := NewCastToFloatError(time.Now())
			fmt.Printf("processing error: %s\n", castErr.Error())
		}
	}()

	var i interface{} = "hello"
	f := castToFloat(i)
	fmt.Println(f)
}

func castToFloat(i interface{}) float64 {
	return i.(float64)
}
