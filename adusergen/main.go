package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func convertToDotFormat(name string) string {
	splitName := strings.Split(name, " ")
	return strings.Join(splitName, ".")
}

func convertToDotFormatLower(name string) string {
	splitName := strings.Split(name, " ")
	return strings.ToLower(strings.Join(splitName, "."))
}

func convertToNoSpaceFormat(name string) string {
	return strings.ReplaceAll(name, " ", "")
}

func convertToNoSpaceFormatLower(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", ""))
}

func convertToReverseFormat(name string) string {
	splitName := strings.Split(name, " ")
	reverseName := []string{splitName[1], splitName[0]}
	return strings.Join(reverseName, "")
}

func convertToReverseFormatLower(name string) string {
	splitName := strings.Split(name, " ")
	reverseName := []string{splitName[1], splitName[0]}
	return strings.ToLower(strings.Join(reverseName, ""))
}

func convertToDotReverseFormat(name string) string {
	splitName := strings.Split(name, " ")
	reverseName := []string{splitName[1], splitName[0]}
	return strings.Join(reverseName, ".")
}

func convertToDotReverseFormatLower(name string) string {
	splitName := strings.Split(name, " ")
	reverseName := []string{splitName[1], splitName[0]}
	return strings.ToLower(strings.Join(reverseName, "."))
}

func convertToShortNameFormat(name string) string {
	splitName := strings.Split(name, " ")
	firstLetter := string(splitName[0][0])
	return firstLetter + strings.ToLower(splitName[1])
}

func convertToShortNameFormatLower(name string) string {
	splitName := strings.Split(name, " ")
	firstLetter := string(splitName[0][0])
	return strings.ToLower(firstLetter) + strings.ToLower(strings.ToLower(splitName[1]))
}

func convertToDotShortName(name string) string {
	splitName := strings.Split(name, " ")
	firstLetter := string(splitName[0][0])
	return firstLetter + "." + strings.ToLower(splitName[1])
}

func convertToDotShortNameLower(name string) string {
	splitName := strings.Split(name, " ")
	firstLetter := string(splitName[0][0])
	return strings.ToLower(firstLetter + "." + strings.ToLower(splitName[1]))
}

func convertToDashFormat(name string) string {
	splitName := strings.Split(name, " ")
	return strings.Join(splitName, "-")
}

func convertToDashFormatLower(name string) string {
	splitName := strings.Split(name, " ")
	return strings.ToLower(strings.Join(splitName, "-"))
}

func convertToDashShortFormat(name string) string {
	splitName := strings.Split(name, " ")
	firstLetter := string(splitName[0][0])
	return firstLetter + "." + strings.ToLower(splitName[1])
}

func convertToDashShortFormatLower(name string) string {
	splitName := strings.Split(name, " ")
	firstLetter := string(splitName[0][0])
	return strings.ToLower(firstLetter + "." + strings.ToLower(splitName[1]))
}

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	flag.Usage = func() {
		h := []string{
			"Usage: adusergen [names]",
			"",
			"Examples:",
			"\t$ cat first_last_names.txt | adusergen",
			"",
		}

		fmt.Fprint(os.Stderr, strings.Join(h, "\n"))
	}

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		rawName := sc.Text()
		wg.Add(1)
		go func() {
			defer wg.Done()
			formats := []func(string) string{
				convertToDotFormat,
				convertToNoSpaceFormat,
				convertToReverseFormat,
				convertToDotReverseFormat,
				convertToShortNameFormat,
				convertToDotShortName,
				convertToDashFormat,
				convertToDashShortFormat,
				convertToDotFormatLower,
				convertToNoSpaceFormatLower,
				convertToReverseFormatLower,
				convertToDotReverseFormatLower,
				convertToShortNameFormatLower,
				convertToDotShortNameLower,
				convertToDashFormatLower,
				convertToDashShortFormatLower,
			}
			for _, format := range formats {
				fmt.Fprintf(os.Stdout, "%s\n", format(rawName))
			}
		}()
	}
	if err := sc.Err(); err != nil {
		log.Fatalln(err)
	}
	wg.Wait()
}
