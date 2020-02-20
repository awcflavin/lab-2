// Package commands contains all the commands of the app.
package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
    "unicode"
)

// AllQuotes prints all quotes.
func AllQuotes() {
	for _, element := range quotes {
		fmt.Println(element)
	}
}

// RandomQuote returns a random quote from the collections of quotes
func RandomQuote() string {
	return quotes.randomQuote()
}

// Persistent prints a Quotes every certain amount of time.
func Persistent(timer int, measure string) {

	// Running a goroutine makes it possible to print a Quotes and wait for a stop message at the same time.
	switch measure {
	case "minute":
		go func() {
			ticking := time.Tick(time.Duration(timer) * time.Minute)
			for range ticking {
				fmt.Println(quotes.randomQuote())
			}
		}()

	case "second":
		go func() {
			ticking := time.Tick(time.Duration(timer) * time.Second)
			for range ticking {
				fmt.Println(quotes.randomQuote())
			}
		}()
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "stop" {
			return
		}
	}
}

// Chat randomly selects a Quotes ending with a "?" and another one ending with a "." and prints them.
func Chat() {
	// split quotes into questions and statements
	var questions Quotes
	var statements Quotes

	for _, quote := range quotes {
		if strings.HasSuffix(quote, "?") {
			questions = append(questions, quote)
		} else if strings.HasSuffix(quote, ".") {
			statements = append(statements, quote)
		}
	}

	// select one for each at random
	fmt.Println(questions.randomQuote())
	fmt.Println(statements.randomQuote())
}

// TalkBack prints a Quotes every time it gets an input message.
func TalkBack() {

	var userName string
	var userReply string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Type 'stop' to stop talking with the Ancestor")
	fmt.Println("Enter your name:")
	_, err := fmt.Scanln(&userName)
	if err != nil {
		log.Fatal(err)
	}

	if userName == "stop" {
		fmt.Println("Good bye nameless fiend!")
		fmt.Println("Bear in mind my last Quotes")
		fmt.Println("Ancestor says: " + quotes.randomQuote())
		return
	}
	fmt.Println("Hi " + userName)

	for {
		fmt.Println("What do you wanna say?")
		userReply, _ = reader.ReadString('\n')
		if userReply == "stop\n" {
			fmt.Println("Good bye " + userName)
			fmt.Println("Bear in mind my last Quotes")
			fmt.Println("Ancestor says: " + quotes.randomQuote())
			break
		}
		fmt.Println("Ancestor says: " + quotes.randomQuote())
	}
}

// Search filters every word in the database to match an input word and prints every Quotes that the word is part of.
func Search(wordToSearch string) {
	// matched is used to know if the search got any matches.
	matched := false

	quoteMatched := make([]string, 0)

	// Quotes is a QuoteType type.
	// wordXSlice is a slice with one word and white-spaces that appear after a filter is applied.
	// wordXFilter is a string inside a wordXSlice.
	// This "for" is a candidate for been transformed into an auxiliary, recursive function.
	for _, quote := range quotes {
		wordFirstSlice := strings.Split(quote, " ")

		for _, wordFirstFilter := range wordFirstSlice {
			wordSecondSlice := strings.Split(wordFirstFilter, ",")

			for _, wordSecondFilter := range wordSecondSlice {
				wordThirdSlice := strings.Split(wordSecondFilter, "!")

				for _, wordThirdFilter := range wordThirdSlice {
					filteredWord := strings.Split(wordThirdFilter, ".")

					// After all the filters are applied the filtered word is compared with the word that is been searched.
					// If a match is found, the Quotes that the filtered word belongs to, is printed.
					if filteredWord[0] == wordToSearch {
						quoteMatched = append(quoteMatched, quote)
						matched = true
					}
				}
			}
		}
	}

	if matched {
		// Defining a nested function makes the next part clearer to see.
		// Candidate for been taking out of the Search func and be used as an auxiliary, separate func instead.
		inside := func(comparingQuote string, unfilteredSlice []string) bool {
			for _, i := range unfilteredSlice {
				if i == comparingQuote {
					return true
				}
			}
			return false
		}

		cleanSlice := make([]string, 0)

		for _, match := range quoteMatched {
			// To be able to use range on a slice, it must have at least one element.
			if len(cleanSlice) == 0 {
				cleanSlice = append(cleanSlice, match)
			}

			// Comparing the element of cleanSlice and unfilteredSlice
			if inside(match, cleanSlice) {
				continue
			} else {
				cleanSlice = append(cleanSlice, match)
			}
		}

		for _, element := range cleanSlice {
			fmt.Println(element + "\n")
		}

	} else {
		fmt.Println("'" + wordToSearch + "'" + " is not present in any quotes.")
	}
}

// NumberOfQuotes returns a given number of random quotes from the collections of quotes
func NumberOfQuotes(quotesNumber int) {
	if quotesNumber < 1 {
		log.Fatal("invalid input param, number must higher than zero")
	}
	index := 0
	for index < quotesNumber {
		fmt.Println(quotes.randomQuote() + "\n")
		index++
	}
}

// ContainsRune function to check if a rune is part of an array
func ContainsRune(arr []rune, r rune) bool {
   for _, a := range arr {
      if a == r {
         return true
      }
   }
   return false
}

// GuessQuote allows user to play a game where they are shown a blanked out quote.
// The user then guesses letters and if the letter is correct it is revealed in the quote.
func Guess() {
    quote := RandomQuote()
    var guess rune
    outString := quote
    done := false
    var guesses[] rune
    reader := bufio.NewReader(os.Stdin)
    var charArray[] string

    
    quoteLength := len(quote)
    
    // Outer loop runs until the quote is guessed completely
    for done != true {
        
        // charArray is reset at the start of each loop
        charArray = []string{}
        
        // Iterate through each character in the quote
        for i:=0; i<quoteLength; i++ {
            
            // each character is stored as a rune in char
            char := []rune(quote)[i]
            
            // if the character is a letter and has not been guessed a blank space is added to charArray
            if unicode.IsLetter(char) && !ContainsRune(guesses, char) {
                charArray = append(charArray, " ")
                
            // otherwise the character is added to charArray
            } else {
                charArray = append(charArray, string(char))
                
            }
        }
        
        // the output string is made by joining the characters in charArray together into a single string
        outString = strings.Join(charArray, "")
                
        
        // If the quote is completely guessed the loop exits
        if outString == quote {
            fmt.Println(quote + "\nCongratulations! You have correctly guessed the quote!")
            done = true
            goto done
        }
        
        // prints the blanked out quote and reads in the next guess
        fmt.Println("\nThe quote is: " + outString + "\nEnter a letter to guess: ")
        guess, _, _= reader.ReadRune()
        
        // excludes the newline character from entry
        if guess != 10 {
            guesses = append(guesses, guess)
        }
        
        done:
    }
}
