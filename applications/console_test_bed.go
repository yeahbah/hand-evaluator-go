package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"holdemHand"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	printMenu()
	for {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		fmt.Println(input)
		switch input {
		case "1":

			go func() {
				fiveCardRunIteration()
			}()

		case "Q":
			fmt.Println("bye")
			return

		default:
			fmt.Print("Do something.\n\n")
			printMenu()
		}
	}
}

func printMenu() {
	fmt.Print("+++ Keith Rule Hand Evaluator in Go +++\n\n")
	fmt.Println("What do you want to do?")
	fmt.Println("1 - Run Benchmarks")
	fmt.Println("Q - Quit")

}

func fiveCardRunIteration() {
	fmt.Println("Five card run interation benchmark...")
	start := time.Now()
	handTypes := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	count := 0

	holdemHand.HandsRange2(5, func(mask uint64) {
		handTypes[holdemHand.EvaluateType(mask)]++
		count++
	})
	Assert(func() bool { return handTypes[holdemHand.HighCard] != 1302540 }, "Unexpected HighCard count")
	Assert(func() bool { return handTypes[holdemHand.Pair] != 1098240 }, "Unexpected Pair count")
	Assert(func() bool { return handTypes[holdemHand.TwoPair] != 123552 }, "Unexpected TwoPair count")
	Assert(func() bool { return handTypes[holdemHand.Trips] != 54912 }, "Unexpected Trips count")
	Assert(func() bool { return handTypes[holdemHand.Straight] != 10200 }, "Unexpected Straight count")
	Assert(func() bool { return handTypes[holdemHand.Flush] != 5108 }, "Unexpected Flush count")
	Assert(func() bool { return handTypes[holdemHand.FullHouse] != 3744 }, "Unexpected Fullhouse count")
	Assert(func() bool { return handTypes[holdemHand.FourOfAKind] != 624 }, "Unexpected FourOfAKind count")
	Assert(func() bool { return handTypes[holdemHand.StraightFlush] != 40 }, "Unexpected StraightFlush count")

	endTime := time.Since(start)

	fmt.Printf("Elapsed: %v\n", endTime.Seconds())
	fmt.Printf("Total hands: %d\n", count)

	handsPerSecond := float64(count) / endTime.Seconds()
	fmt.Printf("Hands/s: %f\n", handsPerSecond)
}

type AssertFunc func() bool

func Assert(assertFunc AssertFunc, msg string) {
	if assertFunc() {
		log.Fatal(msg)
	}
}
