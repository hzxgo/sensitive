package main

import (
	"flag"
	"sensitive/trie"

	"github.com/hzxgo/log"
)

func main() {
	word := flag.String("c", "", "sensitive word detection")
	flag.Parse()

	sens := trie.NewTrie()
	if err := sens.AddByLocalFile("./dict/blacklist.txt"); err != nil {
		log.Errorf("load blacklist failed | %v", err)
		return
	}

	if ok, sensitive := sens.IsValidate(*word); ok {
		log.Infof("input: [%s] is ok.")
	} else {
		log.Warnf("input: [%s] is not allow, sensitive_word: %s", *word, sensitive)
	}
}
