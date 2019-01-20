package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/friendsofgo/bacon-ipsum/generator/bacon"
)

var appName = "bacon-ipsum"

func main() {
	textTypePtr := flag.String("type", "", "Type of the text to generate (Required) [Valid options: all-meat, meat-and-filler]")
	parasPtr := flag.Int("paras", 5, "number of paragraphs")
	sentencesPtr := flag.Int("sentences", 0, "number of sentences (this overrides paragraphs)")
	withLorem := flag.Bool("withLorem", false, "if it is true the first paragraph start with 'Bacon dolor sit amet'")

	flag.Usage = usage

	flag.Parse()

	// Parsing textTypePtr to type TextType
	textType := bacon.TextType(*textTypePtr)

	// Check if is a valid textType or empty
	if !textType.Validate() {
		flag.Usage()
		os.Exit(2)
	}

	// Initialize the new struct bacon Ipsum with the input data
	g := bacon.NewBaconIpsum(
		textType,
		*parasPtr,
		*sentencesPtr,
		*withLorem)

	// Generate the text calling to the selected strategy
	fmt.Println(g.GenerateText())
}

func usage() {
	msg := fmt.Sprintf(`usage: %s [OPTIONS]
	%s is a simple tool to generate random text based on a bacon ipsum API
	`, appName, appName)
	fmt.Println(msg)
	flag.PrintDefaults()
}
