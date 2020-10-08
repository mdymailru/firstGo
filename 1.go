package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	rnd := rand.Intn(100) + 1
	fmt.Println(rnd)

	//fmt.Println(reflect.TypeOf(now))
	//fmt.Println(time.Now().Date())
	var inputStr string
	var err error
	var input int

	for i := 1; i <= 10; i++ {
		fmt.Print(i, "-я Попытка ")
		fmt.Print("Введите число 1-100: ")
		inputStr, err = bufio.NewReader(os.Stdin).ReadString('\n')
		input, err = strconv.Atoi(strings.TrimSpace(inputStr))
		if err != nil {
			log.Fatal(err)
		}

		if input > rnd {
			fmt.Println("Bolshe")
		} else if input < rnd {
			fmt.Println("Menshe")
		} else {
			fmt.Println("Ugadal")
			break
		}
	}
}
