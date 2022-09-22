package svg

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

const DEFAULT_STYLE = "path {fill: black; stroke: none;}"
const SVG_INPUT_DIR = "./input/*.svg"

type Svg struct {
	*goquery.Selection
}

// Create new Svg instance
// Set filepath to specify the path of the SVG file to be read
func NewSvg(doc goquery.Document) *Svg {
	s := new(Svg)
	s.Selection = doc.Find("svg")

	return s
}

// insert style to svg
func (s *Svg) InsertStyle() {
	s.Find("style").Remove()
	s.Selection = s.Selection.AppendNodes(&html.Node{
		Type: html.ElementNode,
		Data: "style",
		FirstChild: &html.Node{
			Type: html.TextNode,
			Data: DEFAULT_STYLE,
		},
	})
}

func (s *Svg) NodeToString() (string, error) {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	if err := goquery.Render(writer, s.Selection); err != nil {
		return "", err
	}
	if err := writer.Flush(); err != nil {
		return "", err
	}

	return b.String(), nil
}

// Compress svg
func (s *Svg) Compress() (string, error) {
	str, err := s.NodeToString()
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString([]byte(str)), nil
}

// Get svg width and height from attributes
func (s *Svg) GetSize() (width, height int) {
	if width, exists := s.Selection.Attr("width"); exists {
		w, _ := strconv.Atoi(width)
		height, _ := s.Selection.Attr("height")
		h, _ := strconv.Atoi(height)
		return w, h
	}

	if viewBox, exists := s.Selection.Attr("viewBox"); exists {
		viewBoxes := regexp.MustCompile(`(\s|,)`).Split(viewBox, 4)
		w, _ := strconv.Atoi(viewBoxes[2])
		h, _ := strconv.Atoi(viewBoxes[3])
		return w, h
	}

	return 0, 0
}

func (s *Svg) SetTitle(title string) {
	s.Selection = s.SetAttr("title", title)
}
