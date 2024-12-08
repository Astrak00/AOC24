package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	result1 := part1("input.txt")
	fmt.Println("Part 1:", result1)
	launchPythonCleanupProcess()
	result2 := part1("sal.sal")
	fmt.Println("Part 2:", result2)
}

func launchPythonCleanupProcess() {
	cmd := exec.Command("python3", "ale.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal("Could not run the python cleaner file", err)
	}
}

func part1(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	acc := 0

	for scanner.Scan() {
		line := scanner.Text()
		pattern, _ := regexp.Compile(`mul\x28[0-9]{1,3},[0-9]{1,3}\x29`)
		matches := pattern.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			// Remove the first 4 characters and the last one
			localMatch := strings.TrimPrefix(match[0], "mul(")
			localMatch = strings.Replace(localMatch, ")", "", 1)
			nums := strings.Split(localMatch, ",")
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			acc += num1 * num2
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return acc
}
