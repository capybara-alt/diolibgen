package svg_test

import (
	"os"
	"strings"
	"testing"

	"github.com/capybara-alt/diolibgen/svg"
)

func TestRead(t *testing.T) {
	tests := strings.NewReader(`<svg viewBox="0 0 24 24"><g /></svg>`)
	reader := svg.NewSvgReader()
	doc, err := reader.Read(tests)
	if err != nil {
		t.Fail()
	}
	svg := doc.Find("svg")
	if svg == nil {
		t.Fail()
	}
	if _, exists := svg.Attr("viewBox"); !exists {
		t.Fail()
	}
}

func TestReadNotSvg(t *testing.T) {
	tests := strings.NewReader(`<html><head /><body><h1>Hello world</h1></body></html>`)
	reader := svg.NewSvgReader()
	doc, err := reader.Read(tests)
	if err != nil {
		t.Fail()
	}
	svg := doc.Find("svg")
	if _, exists := svg.Attr("viewBox"); exists {
		t.Fail()
	}
}

func TestReadError(t *testing.T) {
	reader := svg.NewSvgReader()
	f, _ := os.Open("")
	_, err := reader.Read(f)
	if err == nil {
		t.Fail()
	}
}
