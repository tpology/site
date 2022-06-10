package layouts

import "github.com/tpology/site/pkg/html"

// Homepage is the homepage layout.
func Homepage() html.Element {
	return html.Tag("html",
		html.Tag("head",
			html.Tag("title", "Homepage"),
		),
		html.Tag("body",
			html.Tag("h1", "Hello, world!"),
		),
	)

}
