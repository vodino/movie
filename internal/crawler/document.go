package crawler

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Element struct {
	Selection *goquery.Selection
}

func NewElement(selection *goquery.Selection) *Element {
	return &Element{Selection: selection}
}

func (element *Element) Text() string{
	return element.Selection.Text();
}

func (element *Element) Content() string {
	content, _ := element.Selection.Html()
	return content
}

func (element *Element) ChildContent(selector string) string {
	html, _ := element.Selection.Find(selector).Html()
	return html
}

func (element *Element) ChildContents(selector string) (values []string) {
	element.Selection.Find(selector).Each(func(_ int, selection *goquery.Selection) {
		content, _ := selection.Html()
		values = append(values, strings.TrimSpace(content))
	})
	return values
}

func (element *Element) ChildText(selector string) string {
	return strings.TrimSpace(element.Selection.Find(selector).Text())
}

func (element *Element) ChildTexts(selector string) (values []string) {
	element.Selection.Find(selector).Each(func(_ int, selection *goquery.Selection) {
		values = append(values, strings.TrimSpace(selection.Text()))
	})
	return values
}

func (element *Element) Attribute(k string) string {
	if value, ok := element.Selection.Attr(k); ok {
		return value
	}
	return ""
}

func (element *Element) ChildAttribute(selector, name string) string {
	if attr, ok := element.Selection.Find(selector).Attr(name); ok {
		return strings.TrimSpace(attr)
	}
	return ""
}

func (element *Element) ChildAttributes(selector, attrName string) (result []string) {
	element.Selection.Find(selector).Each(func(_ int, s *goquery.Selection) {
		if attr, ok := s.Attr(attrName); ok {
			result = append(result, strings.TrimSpace(attr))
		}
	})
	return
}

func (element *Element) ForEach(selector string, callback func(int, *Element)) {
	element.Selection.Find(selector).Each(func(index int, selection *goquery.Selection) {
		for _, node := range selection.Nodes {
			callback(index, NewElement(goquery.NewDocumentFromNode(node).Selection))
		}
	})
}

func (element *Element) ForEachWithBreak(selector string, callback func(int, *Element) bool) {
	element.Selection.Find(selector).EachWithBreak(func(index int, selection *goquery.Selection) bool {
		for _, node := range selection.Nodes {
			if callback(index, NewElement(goquery.NewDocumentFromNode(node).Selection)) {
				return true
			}
		}
		return false
	})
}
