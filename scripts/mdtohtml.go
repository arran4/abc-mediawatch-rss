package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func main() {
	md, err := ioutil.ReadFile("readme.md")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	htmlOutput := markdown.Render(doc, renderer)

	fullHtml := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
<title>ABC Mediawatch RSS</title>
<meta charset="utf-8">
</head>
<body>
%s
</body>
</html>`, htmlOutput)

	err = ioutil.WriteFile("public/index.html", []byte(fullHtml), 0644)
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}
	fmt.Println("Created public/index.html")
}
