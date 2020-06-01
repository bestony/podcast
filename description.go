package podcast

import "encoding/xml"

// Description represents text inputs.
type Description struct {
	XMLName xml.Name `xml:"description"`
	Text    string   `xml:",cdata"`
}
