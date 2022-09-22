package svg_test

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/capybara-alt/diolibgen/svg"
)

func TestInsertStyle(t *testing.T) {
	tests, err := goquery.NewDocumentFromReader(strings.NewReader("<svg></svg>"))
	if err != nil {
		t.Fail()
	}

	s := svg.NewSvg(*tests)
	s.InsertStyle()
	if s.Find("style").Text() != "path {fill: black; stroke: none;}" {
		t.Fail()
	}
}

func TestInsertStyleAlreadyExists(t *testing.T) {
	tests, err := goquery.NewDocumentFromReader(strings.NewReader("<svg><g /><style>path {fill: black; stroke: none;}</style></svg>"))
	if err != nil {
		t.Fail()
	}

	s := svg.NewSvg(*tests)
	s.InsertStyle()
	if s.Find("style").Text() != "path {fill: black; stroke: none;}" {
		t.Fail()
	}
}

func TestNodeToString(t *testing.T) {
	tests, err := goquery.NewDocumentFromReader(strings.NewReader(`<svg><g /><style>path {fill: black; stroke: none;}</style></svg>`))
	if err != nil {
		t.Fail()
	}

	s := svg.NewSvg(*tests)
	str, err := s.NodeToString()
	if err != nil {
		t.Fail()
	}

	t.Log(str)
	if str != "<svg><g></g><style>path {fill: black; stroke: none;}</style></svg>" {
		t.Fail()
	}
}

func TestGetSizeFromViewbox(t *testing.T) {
	tests, err := goquery.NewDocumentFromReader(strings.NewReader(`<svg viewBox="0 0 100 100"><g /><style>path {fill: black; stroke: none;}</style></svg>`))
	if err != nil {
		t.Fail()
	}

	s := svg.NewSvg(*tests)
	if w, h := s.GetSize(); w != 100 || h != 100 {
		t.Fail()
	}
}

func TestGetSizeFromViewboxFloat(t *testing.T) {
	tests, err := goquery.NewDocumentFromReader(strings.NewReader(`<svg viewBox="0 0 100.00 100.00"><g /><style>path {fill: black; stroke: none;}</style></svg>`))
	if err != nil {
		t.Fail()
	}

	s := svg.NewSvg(*tests)
	if w, h := s.GetSize(); w != 100.00 || h != 100.00 {
		t.Fail()
	}
}

func TestGetSizeFromWidthAndHeight(t *testing.T) {
	tests, err := goquery.NewDocumentFromReader(strings.NewReader(`<svg width="25" height="25"><g /><style>path {fill: black; stroke: none;}</style></svg>`))
	if err != nil {
		t.Fail()
	}

	s := svg.NewSvg(*tests)
	if w, h := s.GetSize(); w != 25 || h != 25 {
		t.Fail()
	}
}

func TestGetSizeFromWidthAndHeightWithPx(t *testing.T) {
	tests, err := goquery.NewDocumentFromReader(strings.NewReader(`<svg width="25px" height="25px"><g /><style>path {fill: black; stroke: none;}</style></svg>`))
	if err != nil {
		t.Fail()
	}

	s := svg.NewSvg(*tests)
	if w, h := s.GetSize(); w != 25 || h != 25 {
		t.Fail()
	}
}

func TestGetSizeNotFound(t *testing.T) {
	tests, err := goquery.NewDocumentFromReader(strings.NewReader(`<svg><g /><style>path {fill: black; stroke: none;}</style></svg>`))
	if err != nil {
		t.Fail()
	}

	s := svg.NewSvg(*tests)
	if w, h := s.GetSize(); w != 0 || h != 0 {
		t.Fail()
	}
}

func TestCompress(t *testing.T) {
	tests, err := goquery.NewDocumentFromReader(strings.NewReader(`<svg width="25" height="25"><g /><style>path {fill: black; stroke: none;}</style></svg>`))
	if err != nil {
		t.Fail()
	}

	s := svg.NewSvg(*tests)
	str, err := s.Compress()
	if err != nil {
		t.Fail()
	}

	if str != base64.StdEncoding.EncodeToString([]byte(`<svg width="25" height="25"><g></g><style>path {fill: black; stroke: none;}</style></svg>`)) {
		t.Fail()
	}
}

func TestSetTitle(t *testing.T) {
	tests, err := goquery.NewDocumentFromReader(strings.NewReader(`<svg width="25" height="25"><g /><style>path {fill: black; stroke: none;}</style></svg>`))
	if err != nil {
		t.Fail()
	}

	s := svg.NewSvg(*tests)
	s.SetTitle("test")
	if str, _ := s.NodeToString(); str != `<svg width="25" height="25" title="test"><g></g><style>path {fill: black; stroke: none;}</style></svg>` {
		t.Log(str)
		t.Fail()
	}
}
