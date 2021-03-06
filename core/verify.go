package core

import (
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// VerifyReflection is check reflected xss pattern
func VerifyReflection(body, payload string) bool {
	if strings.Contains(body, payload) {
		return true
	} else {
		return false
	}
}

// VerifyDOM is check success inject on code
func VerifyDOM(body io.ReadCloser) bool {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(body)
	check := false
	if err != nil {
		fmt.Println(err)
		return false
	}
	// Find the review items
	doc.Find(".dalfox").Each(func(i int, s *goquery.Selection) {
		fmt.Println("zfzf")
		check = true
	})
	if check {
		return true
	} else {
		doc.Find("dalfox").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			check = true
		})
		if check {
			return true
		} else {
			return false
		}
	}
}
