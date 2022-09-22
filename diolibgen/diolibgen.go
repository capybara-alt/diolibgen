package diolibgen

import (
	"path/filepath"
	"strings"

	"github.com/capybara-alt/diolibgen/mxlibrary"
	"github.com/capybara-alt/diolibgen/svg"
)

func Run() error {
	files, err := filepath.Glob(filepath.Join("input", "*.svg"))
	if err != nil {
		return err
	}

	svgreader := svg.NewSvgReader()
	svglist := make([]svg.Svg, len(files))
	for index, file := range files {
		doc, err := svgreader.OpenAndRead(file)
		if err != nil {
			return err
		}
		svg := svg.NewSvg(*doc)
		svg.InsertStyle()
		svg.SetTitle(strings.ReplaceAll(filepath.Base(file), ".svg", ""))
		svglist[index] = *svg
	}

	mxlib := mxlibrary.NewMxlibrary()
	mxlibrary, err := mxlib.GenerateMxlibrary(svglist)
	if err != nil {
		return err
	}

	if err = mxlibrary.Write(filepath.Join("output", "mxlibrary.xml")); err != nil {
		return err
	}

	return nil
}
