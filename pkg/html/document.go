package html

import (
	"bytes"
	gohtml "html"
)

// Element is an HTML element.
type Element interface {
	String() string
}

// Document is a sequence of Elements.
type Document []Element

// Attr is a HTML attribute.
type Attr struct {
	Key, Value string
}

// Tag is an Element representing an HTML tag.
type Tag struct {
	Name     string
	Attr     []Attr
	Children Document
}

// Text is an Element representing text.
type Text struct {
	Value string
}

// String returns the string representation of the Element.
func (t Tag) String() string {
	var buf bytes.Buffer
	buf.WriteString("<")
	buf.WriteString(t.Name)
	for _, a := range t.Attr {
		buf.WriteString(" ")
		buf.WriteString(a.Key)
		buf.WriteString("=\"")
		buf.WriteString(gohtml.EscapeString(a.Value))
		buf.WriteString("\"")
	}
	buf.WriteString(">")
	for _, c := range t.Children {
		buf.WriteString(c.String())
	}
	buf.WriteString("</")
	buf.WriteString(t.Name)
	buf.WriteString(">")
	return buf.String()
}

// String returns the string representation of the Element.
func (t Text) String() string {
	return gohtml.EscapeString(t.Value)
}
