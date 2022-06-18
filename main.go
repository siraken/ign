package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ign <language name>")
		os.Exit(1)
	}

	name := os.Args[1]

	downloadFile(name, fmt.Sprintf("https://raw.githubusercontent.com/github/gitignore/main/%s.gitignore", name))
}

func downloadFile(name string, url string) {

	cSuccess := color.New(color.FgWhite).Add(color.BgGreen)
	cError := color.New(color.FgWhite).Add(color.BgRed)

	res, err := http.Get(url)

	if err != nil {
		cError.Println(err)
		os.Exit(1)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		cError.Printf("Error: %s\n", res.Status)
		os.Exit(1)
	}

	out, err := os.Create(".gitignore")

	if err != nil {
		cError.Println(err)
		os.Exit(1)
	}

	defer out.Close()

	_, err = io.Copy(out, res.Body)

	if err != nil {
		cError.Println(err)
		os.Exit(1)
	}

	cSuccess.Printf("Downloaded successfully: %s.gitignore\n", name)
}
