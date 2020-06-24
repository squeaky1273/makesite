package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"html/template"
	"os"
	"strings"
	"github.com/bregydoc/gtranslate"
)

type content struct {
	Content string
}

func main() {
	filePtr := flag.String("file", "", "filename")
	dirPtr := flag.String("dir", "", "directory")
	flag.Parse()
	if *dirPtr != "" {
		files, err := ioutil.ReadDir(*dirPtr)
		if err != nil {
			panic(err)
		}
		for _, f := range files {
			name := f.Name()
			if isTxtFile(name) == true {
				renderTemplate("template.tmpl", readFile(name))
				writeTemplateToFile("template.tmpl", name)
			}		
		}
	}

	if *filePtr != "" {
		renderTemplate("template.tmpl", readFile(*filePtr))
		writeTemplateToFile("template.tmpl", *filePtr)
	}

	text := "This is a go mod test."
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: "en",
			To:   "ja",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("en: %s | ja: %s \n", text, translated)

}

func readFile(name string) string {
	fileContents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func renderTemplate(filename string, data string) {
	c := content{Content: data}
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))

	var err error
	err = t.Execute(os.Stdout , c)
	if err != nil {
		panic(err)
	}
}

func writeTemplateToFile(filename string, data string) {
	c := content{Content: data}
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))
	
	fileName := strings.Split(filename, ".")[0] + ".html"
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, c)
	if err != nil {
		panic(err)
	}

}

func isTxtFile(filename string) bool {
	if strings.Contains(filename, ".") {
		return strings.Split(filename, ".")[1] == "txt"
	} else {
		return false
	}
}