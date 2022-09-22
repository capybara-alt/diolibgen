package mxlibrary

import (
	"encoding/xml"

	"github.com/capybara-alt/libdrawio/libdrawio"
)

const template string = `
<mxGraphModel>
	<root>
		<mxCell id="0" />
		<mxCell id="1" parent="0" />
		<mxCell id="2" value="" style="shape=image;verticalLabelPosition=bottom;verticalAlign=top;imageAspect=0;aspect=fixed;editableCssRules=.*;image=data:image/svg+xml," vertex="1" parent="1">
			<mxGeometry width="" height="" as="geometry" />
		</mxCell>
	</root>
</mxGraphModel>`

func MxGraphModelTemplate() *libdrawio.MxGraphModel {
	mxGraphModelTemplate := &libdrawio.MxGraphModel{}
	xml.Unmarshal([]byte(template), mxGraphModelTemplate)

	return mxGraphModelTemplate
}
