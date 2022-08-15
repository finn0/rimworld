package common

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"log"
	"sort"
)

// About is the structure for XML Marshal and Unmarshal based on
// https://rimworldwiki.com/wiki/About.xml.
type About struct {
	XMLName                  xml.Name                  `xml:"ModMetaData"`
	Name                     string                    `xml:"name,omitempty"`
	Author                   string                    `xml:"author,omitempty"`
	Authors                  *[]string                 `xml:"authors,omitempty"`
	PackageID                string                    `xml:"packageId,omitempty"`
	URL                      string                    `xml:"url,omitempty"`
	SupportedVersions        *[]string                 `xml:"supportedVersions>li,omitempty"`
	Description              string                    `xml:"description,omitempty"`
	DescriptionsByVersion    *StringsByVersion         `xml:"descriptionsByVersion,omitempty"`
	ModDependencies          *[]ModDependency          `xml:"modDependencies>li,omitempty"`
	ModDependenciesByVersion *ModDependenciesByVersion `xml:"modDependenciesByVersion,omitempty"`
	LoadBefore               *[]string                 `xml:"loadBefore>li,omitempty"`
	LoadBeforeByVersion      *StringsByVersion         `xml:"loadBeforeByVersion,omitempty"`

	ForceLoadBefore           *[]string         `xml:"forceLoadBefore>li,omitempty"`
	LoadAfter                 *[]string         `xml:"loadAfter>li,omitempty"`
	LoadAfterByVersion        *StringsByVersion `xml:"loadAfterByVersion,omitempty"`
	ForceLoadAfter            *[]string         `xml:"forceLoadAfter>li,omitempty"`
	IncompatibleWith          *[]string         `xml:"incompatibleWith>li,omitempty"`
	IncompatibleWithByVersion *StringsByVersion `xml:"incompatibleWithByVersion,omitempty"`
}

type StringsByVersion struct {
	Value map[string][]string
}

func (p *StringsByVersion) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	p.Value = make(map[string][]string)

	var (
		key    string
		values = make([]string, 0)
	)

	for {
		t, err := d.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch tt := t.(type) {
		case xml.StartElement:
			if tt.Name.Local != "li" {
				key = tt.Name.Local
			}

		case xml.EndElement:
			if tt.Name.Local == key {
				p.Value[key] = values
				values = make([]string, 0)
			}

		case xml.CharData:
			tt = bytes.ReplaceAll(tt, []byte("\n"), nil)
			tt = bytes.ReplaceAll(tt, []byte("\r"), nil)
			if n := bytes.TrimSpace(tt); len(n) > 0 {
				values = append(values, string(tt))
			}
		}
		continue
	}

	return nil
}

func (p *StringsByVersion) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(p.Value) == 0 {
		return nil
	}
	sortedVersions := make([]string, 0, len(p.Value))
	for ver := range p.Value {
		sortedVersions = append(sortedVersions, ver)
	}
	sort.Strings(sortedVersions)

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for _, ver := range sortedVersions {
		// ver start
		verToken := xml.StartElement{Name: xml.Name{Local: ver}}
		if err := e.EncodeToken(verToken); err != nil {
			return err
		}

		// dep, sort by version
		if err := e.EncodeElement(p.Value[ver], xml.StartElement{Name: xml.Name{Local: "li"}}); err != nil {
			return err
		}

		// ver ends
		if err := e.EncodeToken(xml.EndElement{Name: verToken.Name}); err != nil {
			return err
		}
	}

	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}

	return nil
}

type ModDependenciesByVersion struct {
	Value map[string][]ModDependency
}

func (m ModDependenciesByVersion) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m.Value) == 0 {
		return nil
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	sortedVersions := make([]string, 0, len(m.Value))
	for ver := range m.Value {
		sortedVersions = append(sortedVersions, ver)
	}
	sort.Strings(sortedVersions)

	for _, ver := range sortedVersions {
		// ver start
		verToken := xml.StartElement{Name: xml.Name{Local: ver}}
		if err := e.EncodeToken(verToken); err != nil {
			return err
		}

		// dep, sort by version
		if err := e.EncodeElement(m.Value[ver], xml.StartElement{Name: xml.Name{Local: "li"}}); err != nil {
			return err
		}

		// ver ends
		if err := e.EncodeToken(xml.EndElement{Name: verToken.Name}); err != nil {
			return err
		}
	}

	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}

	return nil
}

func (m ModDependenciesByVersion) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var (
		key    string
		values = make([]ModDependency, 0)
	)

	for {
		t, err := d.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch tt := t.(type) {
		case xml.StartElement:
			log.Printf("start elem: %v", tt.Name.Local)
			if tt.Name.Local != "li" {
				key = tt.Name.Local
			}

		case xml.EndElement:
			log.Printf("end elem: %v", tt.Name)
			if tt.Name.Local == key {
				m.Value[key] = values
				values = make([]ModDependency, 0)
			}

		case xml.CharData:
			log.Printf("chardata: %s - %d", tt, len(tt))
			tt = bytes.ReplaceAll(tt, []byte("\n"), nil)
			tt = bytes.ReplaceAll(tt, []byte("\r"), nil)
			if n := bytes.TrimSpace(tt); len(n) > 0 {
				var mods []ModDependency
				if err := xml.Unmarshal(tt, &mods); err != nil {
					return err
				}
				if len(mods) > 0 {
					values = append(values, mods...)
				}
			}
		}
		continue
	}

	return nil
}

type ModDependency struct {
	PackageID        string `xml:"packageId"`
	DisplayName      string `xml:"displayName,omitempty"`
	DownloadURL      string `xml:"downloadUrl,omitempty"`
	SteamWorkshopURL string `xml:"steamWorkshopUrl,omitempty"`
}
