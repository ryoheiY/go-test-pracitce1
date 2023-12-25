package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//print
	intro()
	//channel
	doneChann := make(chan bool)
	//block
	go readUserInput(doneChann)
	<-doneChann
	//close channel
	close(doneChann)
	fmt.Println("end")
}

func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("------------")
	fmt.Println("Enter a number:")
	prompt()
}

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		result, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(result)
	}
}

func checkNumbers(scan *bufio.Scanner) (string, bool) {
	scan.Scan()

	if strings.EqualFold(scan.Text(), "q") {
		return "", true
	}

	numToCheck, err := strconv.Atoi(scan.Text())
	if err != nil {
		return "Please enter a whole number", false
	}

	_, msg := isPrime(numToCheck)

	return msg, false
}

func prompt() {
	fmt.Println("input:")
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	// use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}
