package mxlibrary

import (
	"strconv"

	"github.com/capybara-alt/diolibgen/svg"
	"github.com/capybara-alt/libdrawio/libdrawio"
)

type Mxlibrary struct{}

func NewMxlibrary() *Mxlibrary {
	mxlib := new(Mxlibrary)

	return mxlib
}

// Create MxGraphModel from svg file
// Create MxLibObject from MxGraphModel and svg file
func (l *Mxlibrary) GenerateMxlibrary(svgs []svg.Svg) (*libdrawio.MxLibrary, error) {
	mxlibrary := libdrawio.NewMxLibrary()
	mxlibobjs := make([]libdrawio.MxLibObject, len(svgs))
	for index, svg := range svgs {
		w, h := svg.GetSize()
		title := svg.AttrOr("title", "")
		compressed, err := svg.Compress()
		if err != nil {
			return nil, err
		}
		mxGraphModel, err := l.CreateMxGraphModel(compressed, w, h)
		if err != nil {
			return nil, err
		}

		mxlibobj, err := mxlibrary.MakeMxLibObj(mxGraphModel, title, h, w)
		if err != nil {
			return nil, err
		}
		mxlibobjs[index] = *mxlibobj
	}

	if err := mxlibrary.MakeMxLibrary(mxlibobjs); err != nil {
		return nil, err
	}

	return mxlibrary, nil
}

func (l *Mxlibrary) CreateMxGraphModel(compressedStr string, width, height int) (*libdrawio.MxGraphModel, error) {
	mxGraphModelTemplate := MxGraphModelTemplate()
	mxGraphModelTemplate.Content.MxCells[2].Style += compressedStr
	mxGraphModelTemplate.Content.MxCells[2].Geo.Width = strconv.Itoa(width)
	mxGraphModelTemplate.Content.MxCells[2].Geo.Height = strconv.Itoa(height)

	return mxGraphModelTemplate, nil
}
