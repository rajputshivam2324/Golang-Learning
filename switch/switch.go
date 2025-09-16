package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	fmt.Print("Write ", i, " ", "as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its weekend enjoy")
	default:
		fmt.Println("its weekend guys")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("its first half of the day")
	default:
		fmt.Println("its second half guys")
	}

	whatAmi := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("mai int")
		case bool:
			fmt.Println("mai  bool hu")
		default:
			fmt.Println("mujhe nh jante yeh", t, " apun in sab se alag hai")
		}
	}
	whatAmi(2)
	whatAmi("string")
	whatAmi(false)

}
