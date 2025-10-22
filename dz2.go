package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	k := 31 + rand.Intn(6)

	randcolor := fmt.Sprintf("\033[%dm", k) //здесь случайный цвет выводимого результата
	k = k + 1
	randcolor2 := fmt.Sprintf("\033[%dm", k)
	var num int64
	var b bool
	fmt.Printf("%sВведите число\033[0m\t", randcolor2)
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Ошибка ввода:\t", err)
		return
	}

	if num >= 12307 {
		b = true
	}
	for num < 12307 {
		if num < 0 {
			num *= -1
		} else if num%7 == 0 {
			num *= 39
		} else if num%9 == 0 {
			num *= 13
			num++
			continue
		} else {
			num += 2
			num *= 2
		}
		if num%13 == 0 && num%9 == 0 {
			break
			fmt.Printf("service error") //никогда не будет выполнено
		} else {
			num++
		}
	}
	if b {
		fmt.Println("Ожидалось число меньше 12307")
	} else {
		fmt.Printf("%sКонечное число = %v\033[0m", randcolor, num)
	}

}
