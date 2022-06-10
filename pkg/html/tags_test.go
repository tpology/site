package html

import "testing"

// tagTestCase is a test case for the HtmlTag type.
type tagTestCase struct {
	name string
	tag  *HtmlTag
	str  string
}

// tagTestCases is a slice of tagTestCase.
var tagTestCases = []tagTestCase{
	{
		name: "AccessKey",
		tag:  Tag("div").AccessKey("a"),
		str:  `<div accesskey="a"></div>`,
	},
	{
		name: "Class",
		tag:  Tag("div").Class("foo"),
		str:  `<div class="foo"></div>`,
	},
	{
		name: "ContentEditable",
		tag:  Tag("div").ContentEditable("true"),
		str:  `<div contenteditable="true"></div>`,
	},
	{
		name: "Dir",
		tag:  Tag("div").Dir("ltr"),
		str:  `<div dir="ltr"></div>`,
	},
	{
		name: "Draggable",
		tag:  Tag("div").Draggable("true"),
		str:  `<div draggable="true"></div>`,
	},
	{
		name: "Hidden",
		tag:  Tag("div").Hidden("true"),
		str:  `<div hidden="true"></div>`,
	},
	{
		name: "Id",
		tag:  Tag("div").Id("foo"),
		str:  `<div id="foo"></div>`,
	},
	{
		name: "Lang",
		tag:  Tag("div").Lang("en"),
		str:  `<div lang="en"></div>`,
	},
	{
		name: "SpellCheck",
		tag:  Tag("div").SpellCheck("true"),
		str:  `<div spellcheck="true"></div>`,
	},
	{
		name: "Style",
		tag:  Tag("div").Style("color: red"),
		str:  `<div style="color: red"></div>`,
	},
	{
		name: "TabIndex",
		tag:  Tag("div").TabIndex(1),
		str:  `<div tabindex="1"></div>`,
	},
	{
		name: "Title",
		tag:  Tag("div").Title("foo"),
		str:  `<div title="foo"></div>`,
	},
	{
		name: "Translate",
		tag:  Tag("div").Translate("yes"),
		str:  `<div translate="yes"></div>`,
	},
}

// Test_tagTestCases runs the tagTestCases.
func Test_tagTestCases(t *testing.T) {
	for _, tc := range tagTestCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.tag.String() != tc.str {
				t.Errorf("\nwant: %s\ngot:  %s", tc.str, tc.tag.String())
			}
		})
	}
}
