package wfs

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

func (s *Server) GetCapabilities() (*GetCapabilities, error) {
	res, err := s.Client.Get(s.Url + `?service=wfs&version=1.1.0&request=GetCapabilities`)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var buf bytes.Buffer
	io.Copy(&buf, res.Body)
	return ParseGetCapabilities(buf.Bytes())
}

func ParseGetCapabilities(bs []byte) (*GetCapabilities, error) {
	var res GetCapabilities
	err := xml.Unmarshal(bs, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

type GetCapabilities struct {
	Title      string         `xml:"ServiceIdentification>Title"`
	Operations []OwsOperation `xml:"OperationsMetadata>Operation"`
	Features   []FeatureType  `xml:"FeatureTypeList>FeatureType"`
}

func (g *GetCapabilities) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "GetCapabilities %q\n", g.Title)
	for _, op := range g.Operations {
		fmt.Fprintf(&buf, "  %s:\n", op.Name)
		for _, p := range op.OwsParameters {
			fmt.Fprintf(&buf, "    %s: %#v\n", p.Name, p.Values)
		}
	}
	fmt.Fprintf(&buf, "Features:\n")
	for _, f := range g.Features {
		fmt.Fprintf(&buf, "%v\n", f)
	}
	return buf.String()
}

type OwsOperation struct {
	Name          string         `xml:"name,attr"`
	OwsParameters []OwsParameter `xml:"Parameter"`
}
type OwsParameter struct {
	Name   string   `xml:"name,attr"`
	Values []string `xml:"Value"`
}

type FeatureType struct {
	Name       string   `xml:"Name"`
	Title      string   `xml:"Title`
	Abstract   string   `xml:"Abstract"`
	Keywords   []string `xml:"Keywords>Keyword"`
	DefaultSRS string   `xml:"DefaultSRS"`
}

func (f FeatureType) String() string {
	return fmt.Sprintf("%q %q %v",
		f.Name, f.Title, f.Keywords)
}
