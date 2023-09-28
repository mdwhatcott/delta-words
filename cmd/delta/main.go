package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mdwhatcott/delta-words/lib"
	"github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/go-set/v2/set"
	"github.com/mdwhatcott/must/osmust"
	"github.com/mdwhatcott/tui/v2"
)

var Version = "dev"

func main() {
	flags := flag.NewFlagSet(fmt.Sprintf("%s @ %s", filepath.Base(os.Args[0]), Version), flag.ExitOnError)
	_ = flags.Parse(os.Args[1:])

	ui := tui.New()

	path := "/usr/share/dict/words"
	dict := set.Of(strings.Split(strings.ToLower(string(osmust.ReadFile(path))), "\n")...).Slice()
	var candidates []string
	for _, word := range dict {
		if strings.Contains(word, " ") {
			continue
		}
		if len(word) == 6 {
			candidates = append(candidates, word)
		}
	}
	fmt.Println("candidates:", len(candidates))

	const minOptions = 5
	goodLetters := set.Make[string](0)
	for {
		letters := sortString(candidates[rand.Intn(len(candidates))])
		lengths := funcy.SlicedIndexBy(funcy.ByLength[string], lib.FindMatches(letters, dict...))
		if len(lengths[3]) < minOptions || len(lengths[4]) < minOptions || len(lengths[5]) < minOptions || len(lengths[6]) < minOptions {
			continue
		}
		fmt.Println()
		fmt.Println("LETTERS:", letters)
		for _, match := range lengths[6] {
			fmt.Println(match)
		}
		if ui.NoYes(fmt.Sprintf("%d/%d - Keep?", goodLetters.Len(), 10)) {
			goodLetters.Add(letters)
		}
		if goodLetters.Len() == 10 {
			result := goodLetters.Slice()
			sort.Strings(result)
			fmt.Println(result)
			break

		}
	}
}

func sortString(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}
