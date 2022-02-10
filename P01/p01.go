package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {

	// read csv file
	file, err := os.Open("P01/p01.csv")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// we have to create 3 different files one with the vowels and one with the consonants and one with the numbers

	// create a new file for the vowels
	vowel_file, err := os.Create("P01/vowels.txt")

	if err != nil {
		panic(err)
	}
	defer vowel_file.Close()

	// create a new file for the consonants
	consonant_file, err := os.Create("P01/consonants.txt")

	if err != nil {
		panic(err)
	}
	defer consonant_file.Close()

	// create a new file for the numbers
	number_file, err := os.Create("P01/numbers.txt")

	if err != nil {
		panic(err)
	}
	defer number_file.Close()

	// cycle through the file and write the lines to the appropriate file word by word
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, ",")

		for _, word := range words {
			// cycle character by character
			for _, char := range word {
				// check if the character is a vowel
				if strings.ContainsAny("aeiou", string(char)) {
					// write the character to the vowel file
					vowel_file.WriteString(string(char) + " ")
				} else if strings.ContainsAny("0123456789", string(char)) {
					// write the character to the number file
					number_file.WriteString(string(char) + " ")
				} else {
					// write the character to the consonant file
					consonant_file.WriteString(string(char) + " ")
				}
			}
			// write a new line to the files
			vowel_file.WriteString("\n")
			consonant_file.WriteString("\n")
			number_file.WriteString("\n")

		}
	}
}
