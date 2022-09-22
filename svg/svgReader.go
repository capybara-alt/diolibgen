package svg

import (
	"io"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type SvgReader struct{}

func NewSvgReader() *SvgReader {
	s := new(SvgReader)

	return s
}

// Open file and return goquery.Document
func (s *SvgReader) OpenAndRead(filepath string) (*goquery.Document, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return s.Read(file)
}

// Create new goquery.Document instance from io.Reader
func (s *SvgReader) Read(r io.Reader) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(r)
}
