package mxlibrary_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/capybara-alt/diolibgen/mxlibrary"
	"github.com/capybara-alt/diolibgen/svg"
	"github.com/capybara-alt/libdrawio/libdrawio"
)

func TestGenerateMxLibrary(t *testing.T) {
	tests := []string{
		`<svg title="test1" viewBox="0 0 25 25"><style>path {fill: black; stroke: none;}</style></svg>`,
		`<svg title="test2" viewBox="0 0 25 25"><style>path {fill: black; stroke: none;}</style></svg>`,
		`<svg title="test3" viewBox="0 0 25 25"><style>path {fill: black; stroke: none;}</style></svg>`,
		`<svg title="test4" viewBox="0 0 25 25"><style>path {fill: black; stroke: none;}</style></svg>`,
	}

	testSvgs := make([]svg.Svg, len(tests))
	for index, str := range tests {
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(str))
		testSvgs[index] = *svg.NewSvg(*doc)
	}

	lib := mxlibrary.NewMxlibrary()
	mxlib, _ := lib.GenerateMxlibrary(testSvgs)
	if mxlib == nil {
		t.Fail()
	}

	libobjs := make([]libdrawio.MxLibObject, len(tests))
	json.Unmarshal([]byte(mxlib.Value), &libobjs)
	if libobjs[0].Title != "test1" || libobjs[0].H != 25 || libobjs[0].W != 25 {
		t.Fail()
	}
	if libobjs[1].Title != "test2" || libobjs[1].H != 25 || libobjs[1].W != 25 {
		t.Fail()
	}
	if libobjs[2].Title != "test3" || libobjs[2].H != 25 || libobjs[2].W != 25 {
		t.Fail()
	}
	if libobjs[3].Title != "test4" || libobjs[3].H != 25 || libobjs[3].W != 25 {
		t.Fail()
	}
}

func TestGenerateMxLibrarySvgsNil(t *testing.T) {
	tests := []string{
		`<svg title="test1" viewBox="0 0 25 25"><style>path {fill: black; stroke: none;}</style></svg>`,
		`<svg title="test2" viewBox="0 0 25 25"><style>path {fill: black; stroke: none;}</style></svg>`,
		`<svg title="test3" viewBox="0 0 25 25"><style>path {fill: black; stroke: none;}</style></svg>`,
		`<svg title="test4" viewBox="0 0 25 25"><style>path {fill: black; stroke: none;}</style></svg>`,
	}

	testSvgs := make([]svg.Svg, len(tests))
	for index, str := range tests {
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(str))
		testSvgs[index] = *svg.NewSvg(*doc)
	}

	lib := mxlibrary.NewMxlibrary()
	testSvgs = nil
	mxlib, _ := lib.GenerateMxlibrary(testSvgs)
	if mxlib == nil {
		t.Fail()
	}
}

func TestGenerateMxLibrarySvgsEmpty(t *testing.T) {
	tests := []string{}

	testSvgs := make([]svg.Svg, len(tests))
	for index, str := range tests {
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(str))
		testSvgs[index] = *svg.NewSvg(*doc)
	}

	lib := mxlibrary.NewMxlibrary()
	mxlib, _ := lib.GenerateMxlibrary(testSvgs)
	if mxlib.Value != "[]" {
		t.Fail()
	}
}
