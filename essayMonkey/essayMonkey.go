package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Word bank files, delimeter, and output file
const (
	NounsFile               = "EssayMonkeyNouns.txt"
	VerbsFile               = "EssayMonkeyVerbs.txt"
	AdjectivesFile          = "EssayMonkeyAdjectives.txt"
	Delimeter               = ","
	Space                   = " "
	OutputFile              = "essay.txt"
	MinimumWordsPerSentence = 3
	//This number can be modified; however, it could result is very long sentences.
	MaximumWordsPerSentence = 25
)

// Words contains the word banks for each of the sentence components.
type Words struct {
	Nouns      []string
	Verbs      []string
	Adjectives []string
}

// We check that the number of sentences is under the set limit so that we can ensure
// there is never a sentence with the same length.
func main() {
	numberOfParagraphs := flag.Int("paragraphs", 0, "Number of paragraphs to generate.")
	numberOfSentences := flag.Int("sentences", 0, "Number of sentences per paragraph.")
	flag.Parse()
	if *numberOfParagraphs == 0 || *numberOfSentences == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *numberOfSentences > MaximumWordsPerSentence-MinimumWordsPerSentence {
		fmt.Printf("Please enter a number of sentences less than %d to ensure sentence uniqueness.\n", MaximumWordsPerSentence-MinimumWordsPerSentence)
		os.Exit(1)
	}

	output, err := os.Create(OutputFile)
	if err != nil {
		log.Fatalln(err)
	}

	wordBank := Words{}
	err = wordBank.processFiles()
	if err != nil {
		log.Fatalln(err)
	}

	for *numberOfParagraphs > 0 {
		paragraph := wordBank.writeParagraph(*numberOfSentences)
		_, err := output.WriteString(paragraph + "\n")
		if err != nil {
			log.Fatalln(err)
		}
		*numberOfParagraphs = *numberOfParagraphs - 1
	}
}

// processFiles reads each of the essay monkey word files and parses out the words
// into their respective slices and adds them to the Words object.
func (w *Words) processFiles() error {
	nouns, err := getWordBank(NounsFile)
	if err != nil {
		return err
	}
	verbs, err := getWordBank(VerbsFile)
	if err != nil {
		return err
	}
	adjectives, err := getWordBank(AdjectivesFile)
	if err != nil {
		return err
	}
	w.Nouns = nouns
	w.Verbs = verbs
	w.Adjectives = adjectives
	return nil
}

// writeParagraph takes in the number of sentences per paragraph as an argument.
// Each paragraph starts with a tab. It increases the sentence length each sentence
// to ensure that each sentence is not the same length as per the directions.
func (w *Words) writeParagraph(numSentences int) string {
	paragraph := "\t"
	sentenceLength := 0
	for numSentences > 0 {
		if sentenceLength == 0 {
			paragraph = paragraph + w.buildSentence(MinimumWordsPerSentence+sentenceLength)
		} else {
			paragraph = paragraph + Space + w.buildSentence(MinimumWordsPerSentence+sentenceLength)
		}
		sentenceLength++
		numSentences--
	}
	return paragraph + "\n"
}

// buildSentence creates a sentence with the basic sentence structure of:
// noun+verb+adjective. The word at the start of each sentence will start with
// a capital letter.
func (w *Words) buildSentence(sentenceLength int) string {
	sentence := ""
	startOfSentence := true
	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)
	for sentenceLength > 0 {
		if startOfSentence {
			sentence = strings.Title(w.Nouns[getIndex(w.Nouns, random)])
			startOfSentence = false
		} else {
			sentence = sentence + w.Nouns[getIndex(w.Nouns, random)]
		}
		sentenceLength--
		if sentenceLength == 0 {
			break
		}
		sentence = sentence + Space + w.Verbs[getIndex(w.Verbs, random)]
		sentenceLength--
		if sentenceLength == 0 {
			break
		}
		sentence = sentence + Space + w.Adjectives[getIndex(w.Adjectives, random)]
		sentenceLength--
		if sentenceLength > 0 {
			sentence = sentence + Space
		}
	}
	return strings.TrimSuffix(sentence, Space) + "."
}

// getIndex makes sure that the program won't panic when generating a random number
// if for some reason one of the essay monkey files only had one word.
func getIndex(values []string, random *rand.Rand) int {
	if len(values) == 1 {
		return 0
	}
	return random.Intn(len(values) - 1)
}

// getWordBank takes the name of the essay monkey word file and returns a slice containing
// each of the words.
func getWordBank(fileName string) ([]string, error) {
	wordBankString, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return scrubInput(string(wordBankString)), nil
}

// scrubInput removes any extra spaces in the file, around the words, and any unecessary trailing commas.
func scrubInput(wordBankString string) []string {
	wordBankString = strings.TrimSpace(wordBankString)
	wordBankString = strings.TrimSuffix(wordBankString, Delimeter)
	listOfWords := strings.Split(wordBankString, Delimeter)
	for k, v := range listOfWords {
		listOfWords[k] = strings.TrimSpace(v)
	}
	return listOfWords
}
