package common

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

var _testRaw = `<ModMetaData>
  <name>Test Mod</name>
  <author>foo</author>
  <authors>
    <li>foo</li>
    <li>bar</li>
  </authors>
  <packageId>test.mod</packageId>
  <url>some-url</url>
  <supportedVersions>
    <li>1.1</li>
    <li>1.2</li>
  </supportedVersions>
  <description>a description\nin multiple lines.</description>
  <descriptionsByVersion>
    <v1.1>desc of v1.1</v1.1>
    <v1.2>desc of v1.1</v1.2>
  </descriptionsByVersion>
  <modDependencies>
    <li>
      <packageId>1.2.3</packageId>
      <displayName>display name</displayName>
      <downloadUrl>http://some-download-url/</downloadUrl>
      <steamWorkshopUrl>steam://some-steam-download-url/</steamWorkshopUrl>
    </li>
  </modDependencies>
  <modDependenciesByVersion>
    <v1.1>
      <li>
        <packageId>p1</packageId>
        <displayName>display-name</displayName>
        <downloadUrl>download-url</downloadUrl>
        <steamWorkshopUrl>steam-url</steamWorkshopUrl>
      </li>
    </v1.1>
    <v1.2>
      <li>
        <packageId>p1</packageId>
        <displayName>display-name</displayName>
        <downloadUrl>download-url</downloadUrl>
        <steamWorkshopUrl>steam-url</steamWorkshopUrl>
      </li>
      <li>
        <packageId>p2</packageId>
        <displayName>display-name</displayName>
        <downloadUrl>download-url</downloadUrl>
        <steamWorkshopUrl>steam-url</steamWorkshopUrl>
      </li>
    </v1.2>
  </modDependenciesByVersion>
  <loadBefore>
    <li>foo</li>
    <li>bar</li>
  </loadBefore>
  <loadBeforeByVersion>
    <v1>
      <li>us.zoom.finn</li>
      <li>us.ms.fry</li>
    </v1>
  </loadBeforeByVersion>
  <forceLoadBefore>
    <li>mod1</li>
    <li>mod2</li>
  </forceLoadBefore>
  <loadAfter>
    <li>mod1</li>
    <li>mod2</li>
  </loadAfter>
  <loadAfterByVersion>
    <v1.1>
      <li>mod1</li>
      <li>mod2</li>
    </v1.1>
    <v1.2>
      <li>mod3</li>
      <li>mod4</li>
    </v1.2>
  </loadAfterByVersion>
  <forceLoadAfter>
    <li>mod1</li>
    <li>mod2</li>
  </forceLoadAfter>
  <incompatibleWith>
    <li>mod5</li>
    <li>mod6</li>
  </incompatibleWith>
  <incompatibleWithByVersion>
    <v1.1>
      <li>mod1</li>
      <li>mod2</li>
    </v1.1>
    <v1.2>
      <li>mod3</li>
      <li>mod4</li>
    </v1.2>
  </incompatibleWithByVersion>
</ModMetaData>`

var _testRawStruct = About{
	XMLName:           xml.Name{Local: "ModMetaData"},
	Name:              "Test Mod",
	Author:            "foo",
	Authors:           &[]string{"foo", "bar"},
	PackageID:         "test.mod",
	URL:               "some-url",
	SupportedVersions: &[]string{"1.1", "1.2"},
	Description:       `a description\nin multiple lines.`,
	DescriptionsByVersion: &StringByVersion{Value: map[string]string{
		"v1.1": "desc of v1.1",
		"v1.2": "desc of v1.1",
	}},
	ModDependencies: &[]ModDependency{
		{
			PackageID:        "1.2.3",
			DisplayName:      "display name",
			DownloadURL:      "http://some-download-url/",
			SteamWorkshopURL: "steam://some-steam-download-url/",
		},
	},
	ModDependenciesByVersion: &ModDependenciesByVersion{
		Value: map[string][]ModDependency{
			"v1.1": {
				{
					PackageID:        "p1",
					DisplayName:      "display-name",
					DownloadURL:      "download-url",
					SteamWorkshopURL: "steam-url",
				},
			},
			"v1.2": {
				{
					PackageID:        "p1",
					DisplayName:      "display-name",
					DownloadURL:      "download-url",
					SteamWorkshopURL: "steam-url",
				},
				{
					PackageID:        "p2",
					DisplayName:      "display-name",
					DownloadURL:      "download-url",
					SteamWorkshopURL: "steam-url",
				},
			},
		},
	},
	LoadBefore: &[]string{"foo", "bar"},
	LoadBeforeByVersion: &StringsByVersion{
		Value: map[string][]string{
			"v1": {"us.zoom.finn", "us.ms.fry"},
		},
	},
	ForceLoadBefore: &[]string{"mod1", "mod2"},
	LoadAfter:       &[]string{"mod1", "mod2"},
	LoadAfterByVersion: &StringsByVersion{
		Value: map[string][]string{
			"v1.1": {"mod1", "mod2"},
			"v1.2": {"mod3", "mod4"},
		},
	},
	ForceLoadAfter:   &[]string{"mod1", "mod2"},
	IncompatibleWith: &[]string{"mod5", "mod6"},
	IncompatibleWithByVersion: &StringsByVersion{map[string][]string{
		"v1.1": {"mod1", "mod2"},
		"v1.2": {"mod3", "mod4"},
	}},
}

func TestMarshalAbout(t *testing.T) {
	b, err := xml.MarshalIndent(&_testRawStruct, "", "  ")
	require.NoError(t, err)
	require.Equal(t, _testRaw, string(b))
}

func TestUnmarshal(t *testing.T) {
	var about About
	err := xml.Unmarshal([]byte(_testRaw), &about)
	require.NoError(t, err)
	require.EqualValues(t, _testRawStruct, about)
}
