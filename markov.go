// Takes a text file as input and trains a Markov chain. Once trained, the Generate function can be called to generate a string based off of the training.
package markov

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

// Takes a text file as input. Returns a map[string][]string and [2]int.
// [2]int[0] is the number of words in the source file.
// [2]int[1] is the number of unique words in the source file.
func Train(filename string) (map[string][]string, [2]int) {
	chain := make(map[string][]string)

	inputString, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	splitString := strings.Split(strings.Trim(strings.Replace(string(inputString), "\n", " ", -1), " \n\t"), " ")

	for index, item := range splitString {
		_, keyExists := chain[item]

		if !keyExists {
			chain[item] = []string{}
		}

		if index+1 < len(splitString) {
			chain[item] = append(chain[item], splitString[index+1])
		}
	}

	return chain, [2]int{len(splitString), len(chain)}
}

//Generates text of a specific length from a previously trained chain (obtained from markov.Train(filename))
func Generate(chain map[string][]string, length int) string {
	keys := make([]string, len(chain))

	i := 0
	for item := range chain {
		keys[i] = item
		i++
	}

	output := []string{keys[rand.Intn(len(keys))]}

	for wordInStory := 1; wordInStory < length; wordInStory++ {
		lastWord := output[len(output)-1]
		var nextWord string

		if len(chain[lastWord]) > 0 {
			nextWord = chain[lastWord][rand.Intn(len(chain[lastWord]))]
		} else {
			nextWord = keys[rand.Intn(len(keys))]
		}

		output = append(output, nextWord)
	}

	output[0] = strings.Title(output[0])
	return strings.Join(output, " ")

}
