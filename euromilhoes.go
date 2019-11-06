package main

import (
	"fmt"
	"math/rand"
	"time"
	"bufio"
	"os"
	"log"
)

var (
    outfile, _ = os.Create("euromilhoes.log")
    l = log.New(outfile, "", 0)
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	n  :=  1
	fmt.Print("Numeros: ")
	l.Print("Numeros: ")
	for n <= 5 {
		if n == 5 {
			randNum := getRand(50)
			fmt.Print(randNum)
			l.Print(randNum)
			n = n + 1
		} else {
			randNum := getRand(50)
			fmt.Print(randNum, ",")
			l.Print(randNum, ",")
			n = n + 1
		}
	}
	fmt.Print("\nEstrelas: ")
	l.Print("\nEstrelas: ")
	e := 1
	for  e  <=  2  {
		if e == 2 {
			randNum := getRand(12)
			fmt.Print(randNum)
			l.Print(randNum)
			e = e + 1
		} else {
			randNum := getRand(12)
			fmt.Print(randNum, ",")
			l.Print(randNum, ",")
			e = e + 1
		}
	}
	
	fmt.Print("\nPressione 'ENTER' para continuar...")
	bufio.NewReader(os.Stdin).ReadBytes('\n') 
}

func getRand(Range int) int {
	return 1 + r.Intn(Range - 1)
}
