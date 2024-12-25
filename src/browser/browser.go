package browser

import (
	"errors"
	"fmt"

	"github.com/playwright-community/playwright-go"
)

// fetchHTMLPage:
// Accepts a URL and if successful it returns the HTML content and nil error,
// otherwise an empty string as the HTML content and an error.
// By default, the playwright browser used is firefox, this function assumes that
// you already have this installed.
func FetchHTMLPage(url string) (content string, err error) {
	pw, err := playwright.Run()
	if err != nil {
		return "", errors.New("Could not start playwright")
	}
	browser, err := pw.Firefox.Launch()
	if err != nil {
		return "", errors.New("Could not launch browser")
	}
	defer browser.Close()

	context, err := browser.NewContext(playwright.BrowserNewContextOptions{IgnoreHttpsErrors: playwright.Bool(true)})
	if err != nil {
		return "", errors.New("Could not create new context")
	}
	defer context.Close()

	page, err := context.NewPage()
	if err != nil {
		return "", errors.New("Could not create page")
	}
	if _, err := page.Goto(url); err != nil {
		return "", errors.New(fmt.Sprintf("Could not go to: %s", url))
	}
	content, err = page.Content()
	if err != nil {
		return "", errors.New("Could not get content")
	}
	return content, nil
}
