package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	f, err := os.Open("getpkg.list")
	checkErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("go get ", line)
		cmd := exec.Command("go", "get", line)
		err := cmd.Run()
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
