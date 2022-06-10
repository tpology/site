package html

import (
	"bytes"
)

// Element is an HTML element.
type Element interface {
	String() string
	DeepCopy() Element
}

// Document is a sequence of Elements.
type Document []Element

// Attribute is a HTML attribute.
type Attribute struct {
	Key, Value string
}

// Tag is an Element representing an HTML tag.
type HtmlTag struct {
	Name      string
	Attribute []Attribute
	Children  Document
}

// TextContent is an Element representing text.
type TextContent struct {
	Value string
}

// htmlVoidTags is a set of tags that have no closing tag.
var htmlVoidTags = map[string]bool{
	"area":   true,
	"base":   true,
	"br":     true,
	"col":    true,
	"embed":  true,
	"hr":     true,
	"img":    true,
	"input":  true,
	"link":   true,
	"meta":   true,
	"param":  true,
	"source": true,
	"track":  true,
	"wbr":    true,
}

// String returns the string representation of the Element.
func (t HtmlTag) String() string {
	var buf bytes.Buffer
	buf.WriteString("<")
	buf.WriteString(t.Name)
	for _, a := range t.Attribute {
		buf.WriteString(" ")
		buf.WriteString(a.Key)
		buf.WriteString("=\"")
		buf.WriteString(a.Value)
		buf.WriteString("\"")
	}
	if htmlVoidTags[t.Name] {
		buf.WriteString(">")
	} else {
		buf.WriteString(">")
		for _, c := range t.Children {
			buf.WriteString(c.String())
		}
		buf.WriteString("</")
		buf.WriteString(t.Name)
		buf.WriteString(">")
	}
	return buf.String()
}

// String returns the string representation of the Element.
func (t TextContent) String() string {
	return t.Value
}

// String returns the string representation of the Element.
func (doc Document) String() string {
	var buf bytes.Buffer
	for _, e := range doc {
		buf.WriteString(e.String())
	}
	return buf.String()
}

// DeepCopy returns a deep copy of the Document.
func (doc Document) DeepCopy() Element {
	var copy Document
	for _, e := range doc {
		copy = append(copy, e.DeepCopy())
	}
	return copy
}

// DeepCopy returns a deep copy of the Element.
func (e *HtmlTag) DeepCopy() Element {
	cpy := &HtmlTag{
		Name:      e.Name,
		Attribute: make([]Attribute, len(e.Attribute)),
		Children:  e.Children.DeepCopy().(Document),
	}
	copy(cpy.Attribute, e.Attribute)
	return cpy
}

// DeepCopy returns a deep copy of the Element.
func (e *TextContent) DeepCopy() Element {
	cpy := &TextContent{
		Value: e.Value,
	}
	return cpy
}

// Tag returns a HtmlTag with the given name and children
func Tag(name string, children ...Element) *HtmlTag {
	return &HtmlTag{
		Name:     name,
		Children: Document(children),
	}
}

// Attr adds an Attribute to a tag
func (t *HtmlTag) Attr(key, value string) *HtmlTag {
	t.Attribute = append(t.Attribute, Attribute{
		Key:   key,
		Value: value,
	})
	return t
}
