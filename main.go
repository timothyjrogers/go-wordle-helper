package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"regexp"
	"math"
)

func gen_freqs(words []string) map[string][5]int {
	freqs := make(map[string][5]int)
	letters := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	for i := 0; i < len(letters); i++ {
		freqs[letters[i]] = [5]int{0,0,0,0,0}
	}

	for i := 0; i < len(words); i++ {
		word := words[i]
		for j := 0; j < 5; j++ {
			c := string(word[j])
			var arr = freqs[c]
			arr[j] = arr[j] + 1
			freqs[c] = arr
		}
	}
	return freqs
}

func get_guess() string {
	var guess string
	fmt.Printf("Guess: ")
	fmt.Scanln(&guess)
	guess = strings.ToLower(guess)
	if !check_guess_string(guess) {
		os.Exit(1)
	}
	return guess
}

func check_guess_string(guess string) bool {
	if len(guess) != 5 {
		log.Fatal("Guess must be 5 letters")
		return false
	}
	r, _ := regexp.Compile("[a-z]{5}")
	if !r.MatchString(guess) {
		log.Fatal("Guess must only contain letters")
		return false
	} 
	return true
}

func get_results() string {
	var results string
	fmt.Printf("Enter results: ")
	fmt.Scanln(&results)
	results = strings.ToLower(results)
	if ! check_result_string(results) {
		os.Exit(1)
	}
	return results
}

func check_result_string(results string) bool {
	if len(results) != 5 {
		log.Fatal("Results must be 5 letters")
		return false
	}
	r, _ := regexp.Compile("[xyg]{5}")
	if !r.MatchString(results) {
		log.Fatal("Results must only contain x, y, or g")
		return false
	} 
	return true
}

func check_word(word string, guess string, results string) bool {
	for i := 0; i < 5; i++ {
		if results[i] == 'g' && word[i] != guess[i] {
			return false
		}
		if results[i] == 'x' && strings.Contains(word, string(guess[i])) {
			return false
		}
		if results[i] == 'y' && word[i] == guess[i] {
			return false
		}
		if results[i] == 'y' && !strings.Contains(word, string(guess[i])) {
			return false
		}
	}
	return true
}

func suggest_word(words []string, freqs map[string][5]int) []string {
	scores := make(map[string]int)
	for i := 0; i < len(words); i++ {
		word := words[i]
		for j := 0; j < 5; j++ {
			c := string(word[j])
			scores[word] = scores[word] + freqs[c][j]
		}
	}

	var suggestions []string
	iters := int(math.Min(10.0, float64(len(scores))))
	for i := 0; i < iters; i++ {
		suggestion := ""
		max := 0
		for key, val := range scores {
			if val > max {
				suggestion = key
				max = val
			}
		}
		suggestions = append(suggestions, suggestion)
		delete(scores, suggestion)
	}
	return suggestions
}

func main() {
    content, err := ioutil.ReadFile("./scrabble_5.json")
	if err != nil {
		log.Fatal("Could not open scrabble_5.json")
	}

	var words []string
	json.Unmarshal([]byte(content), &words)
	letter_freq := gen_freqs(words)
	
	for x := 0; x < 6; x++ {
		guess := get_guess()
		results := get_results()
		
		if results == "ggggg" {
			fmt.Printf("Congratulations!")
			break
		}

		var updated_words []string
		for i := 0; i < len(words); i++ {
			if check_word(words[i], guess, results) {
				updated_words = append(updated_words, words[i])
			}
		}
		words = updated_words
		suggestions := suggest_word(words, letter_freq)
		fmt.Printf("Suggested guesses: %s\n", suggestions)
	}
}