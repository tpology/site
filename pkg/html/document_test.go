package html

import "testing"

// Test_Text tests the Text type.
func Test_Text(t *testing.T) {
	txt := &Text{"Hello, world!"}
	if txt.String() != "Hello, world!" {
		t.Errorf("Text.String() = %q, want %q", txt.String(), "Hello, world!")
	}
}

// Test_HtmlTag tests the HtmlTag type.
func Test_HtmlTag(t *testing.T) {
	tag := &HtmlTag{
		Name: "div",
		Attr: []Attr{
			{Key: "class", Value: "foo"},
			{Key: "id", Value: "bar"},
		},
		Children: []Element{
			&Text{"Hello, world!"},
		},
	}
	if tag.String() != "<div class=\"foo\" id=\"bar\">Hello, world!</div>" {
		t.Errorf("HtmlTag.String() = %q, want %q", tag.String(), "<div class=\"foo\" id=\"bar\">Hello, world!</div>")
	}
}

// Test_HtmlTag_Void tests the HtmlTag type for void tags.
func Test_HtmlTag_Void(t *testing.T) {
	tag := &HtmlTag{
		Name: "img",
		Attr: []Attr{
			{Key: "src", Value: "foo.png"},
		},
	}
	if tag.String() != "<img src=\"foo.png\">" {
		t.Errorf("HtmlTag.String() = %q, want %q", tag.String(), "<img src=\"foo.png\">")
	}
}

// Test_HtmlTag_Children tests the HtmlTag type for children.
func Test_HtmlTag_Children(t *testing.T) {
	tag := &HtmlTag{
		Name: "div",
		Attr: []Attr{
			{Key: "class", Value: "foo"},
			{Key: "id", Value: "bar"},
		},
		Children: []Element{
			&HtmlTag{
				Name: "span",
				Attr: []Attr{
					{Key: "class", Value: "baz"},
				},
				Children: []Element{
					&Text{"Hello, world!"},
				},
			},
		},
	}
	if tag.String() != "<div class=\"foo\" id=\"bar\"><span class=\"baz\">Hello, world!</span></div>" {
		t.Errorf("HtmlTag.String() = %q, want %q", tag.String(), "<div class=\"foo\" id=\"bar\"><span class=\"baz\">Hello, world!</span></div>")
	}
}

// Test_Document tests the Document type.
func Test_Document(t *testing.T) {
	doc := Document{
		&HtmlTag{
			Name: "html",
			Attr: []Attr{
				{Key: "lang", Value: "en"},
			},
			Children: []Element{
				&HtmlTag{
					Name: "head",
					Children: []Element{
						&HtmlTag{
							Name: "title",
							Children: []Element{
								&Text{"Hello, world!"},
							},
						},
					},
				},
				&HtmlTag{
					Name: "body",
					Children: []Element{
						&HtmlTag{
							Name: "h1",
							Children: []Element{
								&Text{"Hello, world!"},
							},
						},
					},
				},
			},
		},
	}
	if doc.String() != "<html lang=\"en\"><head><title>Hello, world!</title></head><body><h1>Hello, world!</h1></body></html>" {
		t.Errorf("Document.String() = %q, want %q", doc.String(), "<html lang=\"en\"><head><title>Hello, world!</title></head><body><h1>Hello, world!</h1></body></html>")
	}
}
