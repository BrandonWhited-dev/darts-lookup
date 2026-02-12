package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"bullseye.com/internal/read"
)

const inventoryFileName = "Darts.xlsx"

func main() {
	fmt.Println("Reading Files...")

	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	exe, err = filepath.EvalSymlinks(exe)
	if err != nil {
		log.Fatal(err)
	}

	baseDir := filepath.Dir(exe)
	inventoryPath := filepath.Join(baseDir, inventoryFileName)

	inventory, err := read.ReadItems(inventoryPath)
	if err != nil {
		log.Panic("Failure to open file, ", err)
	}
	fmt.Printf("Finished Reading Files...\n\n")
	fmt.Println("Type \"help\" for more information")
	startRepl(inventory)
}

func startRepl(items []read.Item) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Search > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		switch words[0] {
		case "help":
			fmt.Println("Search by Title: Type \"t\" or \"title\" <title>")
			fmt.Println("Search by SKU: Type \"sku\" <sku>")
			fmt.Println("Search by UPC: Type \"upc\" <sku>")
			fmt.Println("* Or just type a title or sku")
		case "t", "title":
			title := strings.Join(words[1:], " ")
			searchTitleAll(title, items)
		case "sku":
			searchSku(words[1], items)
		case "upc":
			searchUpc(words[1], items)
		default:
			title := strings.Join(words, " ")
			searchSku(words[0], items)
			searchTitle(title, items)
		}
	}
}

func searchSku(sku string, items []read.Item) {
	for _, item := range items {
		if strings.Contains(strings.ToLower(item.SKU), sku) {
			fmt.Printf("\n%s\n", item.ToString())
		}
	}
}

func searchUpc(upc string, items []read.Item) {
	for _, item := range items {
		if strings.Contains(strings.ToLower(item.UPC), upc) {
			fmt.Printf("\n%s\n", item.ToString())
		}
	}
}

func searchTitle(title string, items []read.Item) {
	prev := ""
	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Title), title) {
			if prev != item.Title {
				prev = item.Title
				fmt.Printf("%s\n", item.ToString())
			}
		}
	}
}

func searchTitleAll(title string, items []read.Item) {
	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Title), title) {
				fmt.Printf("%s\n", item.ToString())
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
