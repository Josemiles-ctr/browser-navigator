package browser

import Playwright "github.com/mxschmitt/playwright-go"

type Manager struct {
	pw *Playwright.Playwright
	browser Playwright.Browser
	context Playwright.BrowserContext
}