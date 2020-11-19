package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main(){
	fmt.Println("Welcome to the best guessing number game of the world!")
	fmt.Println("What is your name?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	fmt.Printf("Welcome %v let's play!\n", name)
	fmt.Println("Set the boundaries...")
	fmt.Println("Minimum number:")
	minScanner := bufio.NewScanner(os.Stdin)
	minScanner.Scan()
	minNumberStr := minScanner.Text()
	fmt.Println("Maximum number:")
	maxScanner := bufio.NewScanner(os.Stdin)
	maxScanner.Scan()
	maxNumberStr := maxScanner.Text()
	minNumber, _ := strconv.Atoi(minNumberStr)
	maxNumber, _ := strconv.Atoi(maxNumberStr)
	fmt.Println("Generating random number...")
	time.Sleep(time.Millisecond * 500)
	randomNum := rand.Intn(maxNumber - minNumber + 1) + minNumber
	fmt.Println("It's time to guess!")
	tryFlag := true
	tryCounter := 0
	for tryFlag{
		tryFlag = guessNumber(randomNum)
		tryCounter += 1
	}
	switch tryCounter {
	case 1 :
		fmt.Printf("It took you %v attempt, awesome!\n", tryCounter)
	default:
		fmt.Printf("It took you %v attempts!\n", tryCounter)
	}

}

func guessNumber(a int)bool{
	fmt.Println("Write your number:")
	guessedNumberScanner := bufio.NewScanner(os.Stdin)
	guessedNumberScanner.Scan()
	guessNumberStr := guessedNumberScanner.Text()
	g, _ := strconv.Atoi(guessNumberStr)
	if g > a{
		fmt.Println("Too high!")
		return true
	} else if g < a{
		fmt.Println("Too low!")
		return true
	} else{
		fmt.Println("Congratulations, you win!")
		return false
	}
}
