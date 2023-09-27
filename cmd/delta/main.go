package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/mdwhatcott/delta-words/lib"
	"github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/must/osmust"
)

var Version = "dev"

type Config struct {
	letters string
	dict    string
}

func main() {
	var config Config
	flags := flag.NewFlagSet(fmt.Sprintf("%s @ %s", filepath.Base(os.Args[0]), Version), flag.ExitOnError)
	flags.StringVar(&config.letters, "letters", "singer", "The letters to make words out of.")
	flags.StringVar(&config.dict, "dict", "/usr/share/dict/words", "The path of the dictionary file.")
	_ = flags.Parse(os.Args[1:])

	dict := set.Of(strings.Split(strings.ToLower(string(osmust.ReadFile(config.dict))), "\n")...).Slice()
	if config.letters == "" {
		var candidates []string
		for _, word := range dict {
			if len(word) == 6 {
				candidates = append(candidates, word)
			}
		}
		config.letters = candidates[rand.Intn(len(candidates))]
	}

	config.letters = string(funcy.Shuffle([]byte(config.letters)))
	fmt.Println("LETTERS:", config.letters)

	matches := lib.FindMatches(config.letters, dict...)
	lengths := make(map[int][]string)
	for _, match := range matches {
		lengths[len(match)] = append(lengths[len(match)], match)
	}
	for x := 3; x <= len(config.letters); x++ {
		if len(lengths[x]) == 0 {
			continue
		}
		fmt.Println(x)
		for _, match := range lengths[x] {
			fmt.Println(match)
		}
	}
}
