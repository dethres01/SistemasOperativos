package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("p01.csv")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	vowel_file, err := os.Create("vowels.txt")

	if err != nil {
		panic(err)
	}
	defer vowel_file.Close()

	consonant_file, err := os.Create("consonants.txt")

	if err != nil {
		panic(err)
	}
	defer consonant_file.Close()

	number_file, err := os.Create("numbers.txt")

	if err != nil {
		panic(err)
	}
	defer number_file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, ",")

		for _, word := range words {

			for _, char := range word {

				if strings.ContainsAny("aeiou", string(char)) {

					vowel_file.WriteString(string(char) + " ")
				} else if strings.ContainsAny("0123456789", string(char)) {

					number_file.WriteString(string(char) + " ")
				} else {

					consonant_file.WriteString(string(char) + " ")
				}
			}
			vowel_file.WriteString("\n")
			consonant_file.WriteString("\n")
			number_file.WriteString("\n")

		}
	}
}
