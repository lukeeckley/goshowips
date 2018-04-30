package main

import (
        "fmt"
        "regexp"
        "bufio"
        _ "io"
        "os"
)

func main() {
        scanner := bufio.NewScanner(os.Stdin)
        var source string = ""
        var results []string

        for scanner.Scan() {
                source += scanner.Text()
        }
        if err := scanner.Err(); err != nil {
                fmt.Fprintln(os.Stderr, "reading standard input:", err)
        }

        re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

        submatchall := re.FindAllString(source, -1)
        for _, element := range submatchall {
                // append to slice only if IP isn't already in slice
                if !contains(results, element) {
                        results = append(results, element)
                }
        }
        // Print one IP per line
        for _, element := range results {
                fmt.Println(element)
        }
}

func contains(s []string, e string) bool {
        for _, a := range s {
                if a == e {
                        return true
                }
        }
        return false
}
