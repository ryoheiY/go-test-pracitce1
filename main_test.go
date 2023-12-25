package main

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTest := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"not prime", 14, false, "14 is not a prime number because it is divisible by 2"},
		{"negative", -1, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTest {
		fmt.Println("Test_isPrime-testNo:" + e.name)
		actBool, actMsg := isPrime(e.testNum)
		assert.Equal(t, e.expected, actBool)
		assert.Equal(t, e.msg, actMsg)
	}
}

func Test_prompt(t *testing.T) {
	//save a copy of os.stdout
	stdout := os.Stdout

	//create a read and write pipe
	pipe, f, _ := os.Pipe()

	os.Stdout = f

	prompt()

	_ = f.Close()

	os.Stdout = stdout
	all, _ := io.ReadAll(pipe)
	assert.Equal(t, "input:\n", string(all))
}

func Test_intro(t *testing.T) {
	//save a copy of os.stdout
	stdout := os.Stdout

	//create a read and write pipe
	pipe, f, _ := os.Pipe()

	os.Stdout = f

	intro()

	_ = f.Close()

	os.Stdout = stdout
	all, _ := io.ReadAll(pipe)
	fmt.Println("Is it Prime?")
	fmt.Println("------------")
	fmt.Println("Enter a number:")
	assert.Equal(t, "Is it Prime?\n------------\nEnter a number:\ninput:\n", string(all))
}

func Test_checkNumbers(t *testing.T) {
	test := []struct {
		name   string
		input  string
		msg    string
		result bool
	}{
		{"1", "7", "7 is a prime number!", false},
		{"2", "qe", "Please enter a whole number", false},
		{"3", "q", "", true},
	}

	for _, val := range test {
		fmt.Println("Test_checkNumbers:" + val.name)
		input := strings.NewReader(val.input)
		reader := bufio.NewScanner(input)
		res, resBool := checkNumbers(reader)
		assert.Equal(t, val.msg, res)
		assert.Equal(t, val.result, resBool)
	}

}
