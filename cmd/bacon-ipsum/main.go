package main

import (
	"flag"
	"fmt"
)

func main() {
	textTypePtr := flag.String("type", "", "Type of the text to generate")
	parasPtr := flag.Int("paras", 5, "number of paragraphs")
	sentencesPtr := flag.Int("sentences", 0, "number of senteces (this overrides paragraphs)")
	withLorem := flag.Bool("withLorem", false, "if it is true the first paragraph start with 'Bacon dolor sit amet'")
	flag.Parse()

	fmt.Println(
		"type:", *textTypePtr,
		"paras:", *parasPtr,
		"sentences:", *sentencesPtr,
		"withLorem:", *withLorem,
	)
}
