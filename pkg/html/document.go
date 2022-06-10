package html

import (
	"bytes"
	gohtml "html"
	"strconv"
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
		buf.WriteString(gohtml.EscapeString(a.Value))
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
	return gohtml.EscapeString(t.Value)
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
func Tag(name string, children ...Element) Element {
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

// AccessKey sets the access key attribute of a tag
func (t *HtmlTag) AccessKey(key string) *HtmlTag {
	return t.Attr("accesskey", key)
}

// Class sets the class attribute of a tag
func (t *HtmlTag) Class(class string) *HtmlTag {
	return t.Attr("class", class)
}

// ContentEditable sets the contenteditable attribute of a tag
func (t *HtmlTag) ContentEditable(value string) *HtmlTag {
	return t.Attr("contenteditable", value)
}

// Dir sets the dir attribute of a tag
func (t *HtmlTag) Dir(dir string) *HtmlTag {
	return t.Attr("dir", dir)
}

// Draggable sets the draggable attribute of a tag
func (t *HtmlTag) Draggable(value string) *HtmlTag {
	return t.Attr("draggable", value)
}

// Hidden sets the hidden attribute of a tag
func (t *HtmlTag) Hidden(value string) *HtmlTag {
	return t.Attr("hidden", value)
}

// Id sets the id attribute of a tag
func (t *HtmlTag) Id(id string) *HtmlTag {
	return t.Attr("id", id)
}

// Lang sets the lang attribute of a tag
func (t *HtmlTag) Lang(lang string) *HtmlTag {
	return t.Attr("lang", lang)
}

// SpellCheck sets the spellcheck attribute of a tag
func (t *HtmlTag) SpellCheck(value string) *HtmlTag {
	return t.Attr("spellcheck", value)
}

// Style sets the style attribute of a tag
func (t *HtmlTag) Style(style string) *HtmlTag {
	return t.Attr("style", style)
}

// TabIndex sets the tabindex attribute of a tag
func (t *HtmlTag) TabIndex(index int) *HtmlTag {
	return t.Attr("tabindex", strconv.Itoa(index))
}

// Title sets the title attribute of a tag
func (t *HtmlTag) Title(title string) *HtmlTag {
	return t.Attr("title", title)
}

// Translate sets the translate attribute of a tag
func (t *HtmlTag) Translate(value string) *HtmlTag {
	return t.Attr("translate", value)
}
