//
// octohug
//
// copies octopress posts to hugo posts
//   converts the header
//   converts categories and tags to hugo format in header
//   if run in the octopress directory, replaces include_file with the contents
//
// http://codebrane.com/blog/2015/09/10/migrating-from-octopress-to-hugo/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var octopressPostsDirectory string
var hugoPostDirectory string

// HeaderSyntax : either yaml or toml
type HeaderSyntax int

const (
	yaml = iota
	toml
)

func (headerSyntax HeaderSyntax) String() string {
	return [...]string{"yaml", "toml"}[headerSyntax]
}

// CHANGE HERE: HEADER SYNTAX:
var useHeaderSyntax HeaderSyntax = yaml

func setHeaderSyntaxBoundary(syntax HeaderSyntax) string {
	switch syntax {
	case yaml:
		return "---"
	case toml:
		return "+++"
	}
	return "+++" // <-- fallback: original script behavior
}

func setHeaderSyntaxKeyValueSymbol(syntax HeaderSyntax) string {
	switch syntax {
	case yaml:
		return ": "
	case toml:
		return " = "
	}
	return " = " // <-- fallback: original script behavior
}

func formatKeyWithSeparator(key string, syntax HeaderSyntax) string {
	return fmt.Sprintf("%s%s", key, setHeaderSyntaxKeyValueSymbol(syntax))
}

func readFile(path string) (string, error) {
	file, fileError := os.Open(path)
	if fileError != nil {
	}
	defer file.Close()
	var buffer []byte
	fileReader := bufio.NewReaderSize(file, 10*1024)
	line, isPrefix, lineError := fileReader.ReadLine()
	for lineError == nil && !isPrefix {
		buffer = append(buffer, line...)
		buffer = append(buffer, byte('\n'))
		line, isPrefix, lineError = fileReader.ReadLine()
	}
	if isPrefix {
		fmt.Fprintln(os.Stderr, "buffer size too small")
		return "", nil
	}

	return string(buffer), nil
}

func visit(path string, fileInfo os.FileInfo, err error) error {
	if fileInfo.IsDir() {
		return nil
	}

	// Get the base filename of the post
	octopressFilename := filepath.Base(path)

	// Need to strip off the initial date and final .markdown from the post filename
	regex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}-(.*).m(arkdown|d)`)
	matches := regex.FindStringSubmatch(octopressFilename)

	// Ignore non-matching filenames (i.e. do no dereference nil)
	if matches == nil {
		return nil
	}
	octopressFilenameWithoutExtension := matches[1]
	hugoFilename := hugoPostDirectory + "/" + octopressFilenameWithoutExtension + ".md"
	fmt.Printf("%s\n%s\n", path, hugoFilename)

	// Open the octopress file
	octopressFile, octopressFileError := os.Open(path)
	// Nothing to do if we can open the source file
	if octopressFileError != nil {
		fmt.Fprintf(os.Stderr, "Error opening octopress file %s, ignoring\n", path)
		return nil
	}
	defer octopressFile.Close()

	// Create the hugo file
	hugoFile, hugoFileError := os.Create(hugoFilename)
	if hugoFileError != nil {
		fmt.Fprintf(os.Stderr, "could not create hugo file: %v\n", hugoFileError)
		return nil
	}
	defer hugoFile.Close()
	hugoFileWriter := bufio.NewWriter(hugoFile)

	// octopressDateRegex := regexp.MustCompile(`^date:`)
	octopressCategoryOrTagNameRegex := regexp.MustCompile(`^- (.*)`)

	// Read the octopress file line by line
	headerTagSeen := false
	inCategories := false
	firstCategoryAdded := false
	firstInlineCategoryAdded := false
	inTags := false
	firstTagAdded := false
	octopressFileReader := bufio.NewReaderSize(octopressFile, 10*1024)
	octopressLine, isPrefix, lineError := octopressFileReader.ReadLine()
	for lineError == nil && !isPrefix {
		octopressLineAsString := string(octopressLine)

		if octopressLineAsString == "---" {
			headerTagSeen = !headerTagSeen
			if inCategories || inTags {
				hugoFileWriter.WriteString("]\n")
				inCategories = false
				inTags = false
			}
			octopressLineAsString = setHeaderSyntaxBoundary(useHeaderSyntax)
		}

		if strings.Contains(octopressLineAsString, "categories:") {
			inCategories = true
			hugoFileWriter.WriteString(formatKeyWithSeparator("categories", useHeaderSyntax))
			hugoFileWriter.WriteString("[")

			// handle alternative categories syntax: `categories: [foo, bar, baz]`
			// handle multiline?
			if strings.Contains(octopressLineAsString, "[") {
				openingBracketPos := strings.Index(octopressLineAsString, "[")
				closingBracketPos := strings.Index(octopressLineAsString, "]")
				rawCategories := octopressLineAsString[openingBracketPos+1 : closingBracketPos]
				rawCategoriesAsCollection := strings.Split(rawCategories, ", ")

				for _, category := range rawCategoriesAsCollection {
					if firstInlineCategoryAdded {
						hugoFileWriter.WriteString(", ")
					}

					hugoFileWriter.WriteString("\"" + category + "\"")
					firstInlineCategoryAdded = true
				}
			}
		} else if strings.Contains(octopressLineAsString, "tags:") {
			if inCategories {
				inCategories = false
				hugoFileWriter.WriteString("]\n")
			}
			inTags = true
			hugoFileWriter.WriteString(formatKeyWithSeparator("tags", useHeaderSyntax))
			hugoFileWriter.WriteString("[")
		} else if strings.Contains(octopressLineAsString, "keywords: ") {
			inCategories = false
			if inTags {
				hugoFileWriter.WriteString("]\n")
				inTags = false
			}
			hugoFileWriter.WriteString(formatKeyWithSeparator("keywords", useHeaderSyntax))
			hugoFileWriter.WriteString("[")
			parts := strings.Split(octopressLineAsString, ": ")
			keywords := strings.Split(strings.Replace(parts[1], "\"", "", -1), ",")
			firstKeyword := true
			for _, keyword := range keywords {
				if !firstKeyword {
					hugoFileWriter.WriteString(",")
				}
				hugoFileWriter.WriteString("\"" + keyword + "\"")
				firstKeyword = false
			}
			hugoFileWriter.WriteString("]\n")
		} else if inCategories && !inTags {
			fmt.Printf("inCategories and not in tags...")
			fmt.Printf("%s\n", octopressLineAsString)
			matches = octopressCategoryOrTagNameRegex.FindStringSubmatch(octopressLineAsString)
			if len(matches) > 1 {
				if firstCategoryAdded {
					hugoFileWriter.WriteString(", ")
				}
				hugoFileWriter.WriteString("\"" + matches[1] + "\"")
				firstCategoryAdded = true
			}
		} else if octopressLineAsString == "tags:" {
			inTags = true
			hugoFileWriter.WriteString(formatKeyWithSeparator("tags", useHeaderSyntax))
			hugoFileWriter.WriteString("[")
		} else if inTags {
			matches = octopressCategoryOrTagNameRegex.FindStringSubmatch(octopressLineAsString)
			if len(matches) > 1 {
				if firstTagAdded {
					hugoFileWriter.WriteString(", ")
				}
				hugoFileWriter.WriteString("\"" + matches[1] + "\"")
				firstTagAdded = true
			}
			tag := strings.Replace(matches[1], "'", "", -1)
			tag = strings.Replace(tag, "\"", "", -1)
			hugoFileWriter.WriteString("\"" + tag + "\"")
			firstTagAdded = true
		} else if strings.Contains(octopressLineAsString, "date: ") {
			parts := strings.Split(octopressLineAsString, " ")

			// Date
			hugoFileWriter.WriteString(formatKeyWithSeparator("date", useHeaderSyntax))
			// hugoFileWriter.WriteString("\"" + parts[1] + "\"\n")
			hugoFileWriter.WriteString(" " + parts[1] + "T" + parts[2] + ":00\n")

			// Slug
			// octoSlugDate := strings.Replace(parts[1], "-", "/", -1)
			// octoFriendlySlug := octoSlugDate + "/" + octopressFilenameWithoutExtension
			octoFriendlySlug := octopressFilenameWithoutExtension
			hugoFileWriter.WriteString(formatKeyWithSeparator("slug", useHeaderSyntax))
			hugoFileWriter.WriteString("\"" + octoFriendlySlug + "\"\n")

			// Alias
			aliasDate := strings.Replace(parts[1], "-", "/", -1)
			alias := "/blog/" + aliasDate + "/" + octopressFilenameWithoutExtension
			hugoFileWriter.WriteString(formatKeyWithSeparator("aliases", useHeaderSyntax))
			hugoFileWriter.WriteString("[" + alias + "]\n")

		} else if strings.Contains(octopressLineAsString, "title: ") {
			// to keep the urls the same as octopress, the title
			// needs to be the filename
			// ^ previous sentence only applies if not overridden by `config.toml` and usage of `slug` header.
			//
			// parts := strings.Split(octopressFilenameWithoutExtension, "-")
			// hugoFileWriter.WriteString(formatKeyWithSeparator("title", useHeaderSyntax))
			// hugoFileWriter.WriteString("\"")
			// firstPart := true
			// for _, part := range parts {
			// 	if !firstPart {
			// 		hugoFileWriter.WriteString(" ")
			// 	}
			// 	hugoFileWriter.WriteString(part)
			// 	firstPart = false
			// }
			// hugoFileWriter.WriteString("\"\n")
			hugoFileWriter.WriteString(octopressLineAsString + "\n")
		} else if strings.Contains(octopressLineAsString, "description: ") {
			parts := strings.Split(octopressLineAsString, ": ")
			hugoFileWriter.WriteString(formatKeyWithSeparator("description", useHeaderSyntax))
			hugoFileWriter.WriteString(parts[1] + "\n")
		} else if strings.Contains(octopressLineAsString, "layout: ") {
		} else if strings.Contains(octopressLineAsString, "author: ") {
		} else if strings.Contains(octopressLineAsString, "comments: ") {
		} else if strings.Contains(octopressLineAsString, "slug: ") {
		} else if strings.Contains(octopressLineAsString, "wordpress_id: ") {
		} else if strings.Contains(octopressLineAsString, "published: ") {
			hugoFileWriter.WriteString("published = false\n")
		} else if strings.Contains(octopressLineAsString, "include_code") {
			parts := strings.Split(octopressLineAsString, " ")
			// can be:
			// {% include_code [RedViewController.m] lang:objectivec slidernav/RedViewController.m %}
			// or
			// {% include_code [RedViewController.m] slidernav/RedViewController.m %}
			codeFilePath := "source/downloads/code/" + parts[len(parts)-2]
			codeFileContent, _ := readFile(codeFilePath)
			codeFileContent = strings.Replace(codeFileContent, "<", "&lt;", -1)
			codeFileContent = strings.Replace(codeFileContent, ">", "&gt;", -1)
			hugoFileWriter.WriteString("<pre><code>\n" + codeFileContent + "</code></pre>\n")
		} else if strings.Contains(octopressLineAsString, "{% img") {
			parts := strings.Split(octopressLineAsString, " ")
			imageName := parts[2]
			hugoFileWriter.WriteString("\n![img](" + imageName + ")\n")
		} else {
			hugoFileWriter.WriteString(octopressLineAsString + "\n")
		} // if octopressLineAsString == "categories:"

		hugoFileWriter.Flush()
		octopressLine, isPrefix, lineError = octopressFileReader.ReadLine()
	}
	if isPrefix {
		fmt.Fprintln(os.Stderr, "buffer size too small")
	}
	return nil
}

func init() {
	flag.StringVar(&octopressPostsDirectory, "octo", "source/_posts", "path to octopress posts directory")
	// flag.StringVar(&octopressPostsDirectory, "octo", "example-input", "path to octopress posts directory")
	flag.StringVar(&hugoPostDirectory, "hugo", "content/post", "path to hugo post directory")
	// flag.StringVar(&hugoPostDirectory, "hugo", "example-output", "path to hugo post directory")
}

func main() {
	flag.Parse()

	// Check that we can trust octopressPostsDirectory
	if _, err := os.Stat(octopressPostsDirectory); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory: %v\n", err)
		os.Exit(-1)
	}
	os.MkdirAll(hugoPostDirectory, 0777)
	filepath.Walk(octopressPostsDirectory, visit)
}
