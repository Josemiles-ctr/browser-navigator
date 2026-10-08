package browser

import Playwright "github.com/mxschmitt/playwright-go"

type Manager struct {
	pw *Playwright.Playwright
	browser Playwright.Browser
	context Playwright.BrowserContext
}

func NewManager() (*Manager, error) {
	pw, err := Playwright.Run()
	if err != nil {
		return nil, err
	}

	browser, err := pw.Chromium.LaunchPersistentContext(
		"/Chromium-profile",
		Playwright.BrowserTypeLaunchPersistentContextOptions{
			Headless: Playwright.Bool(false),
		},
	)

	if err != nil {
		return nil, err
	}
	
	return &Manager{
		pw: pw,
		context: browser,
	}, nil
}

func (m *Manager) Close() error {
	if err := m.context.Close(); err != nil {
		return err
	}
	if err := m.pw.Stop(); err != nil {
		return err
	}
	return nil
}

func (m *Manager) Navigate(url string) (Playwright.Page, error) {
	page, err := m.context.NewPage()
	if err != nil {
		return nil, err
	}

	_, err = page.Goto(url)
	if err != nil {
		return nil, err
	}

	return page, nil
}