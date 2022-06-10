package html

import "strconv"

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
