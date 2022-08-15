package common

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMarshalAbout(t *testing.T) {
	denp := []ModDependency{
		{
			PackageID:        "1.2.3",
			DisplayName:      "display name",
			DownloadURL:      "http://some-download-url/",
			SteamWorkshopURL: "steam://some-steam-download-url/",
		},
	}

	meta := About{
		Name:              "Test Mod - Plague Gun",
		Author:            "YourNameHere",
		PackageID:         "YourNameHere.PlagueGun",
		SupportedVersions: &[]string{"1.1", "1.2"},
		Description:       `This mod adds a plague gun, a weapon that has a chance to give your enemies the plague.\n\nFor version 1.1.`,
		ModDependencies:   &denp,
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
	}

	b, err := xml.MarshalIndent(&meta, "", "  ")
	require.NoError(t, err)
	t.Logf("%s", string(b))
}

func TestUnmarshal(t *testing.T) {
	s := `<ModMetaData>
          <name>Test Mod - Plague Gun</name>
          <author>YourNameHere</author>
          <packageId>YourNameHere.PlagueGun</packageId>
          <supportedVersions>
            <li>1.1</li>
            <li>1.2</li>
          </supportedVersions>
          <description>This mod adds a plague gun, a weapon that has a chance to give your enemies the plague.\n\nFor version 1.1.</description>
          <modDependencies>
            <li>
              <packageId>1.2.3</packageId>
              <displayName>dep</displayName>
              <downloadUrl>down</downloadUrl>
              <steamWorkshopUrl>stem</steamWorkshopUrl>
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
			<v2>
              <li>2 - us.zoom.finn</li>
              <li>2 - us.ms.fry</li>
            </v2>
          </loadBeforeByVersion>
        </ModMetaData>`

	var about About
	err := xml.Unmarshal([]byte(s), &about)
	require.NoError(t, err)
	t.Logf("%v", about.ModDependenciesByVersion)
	//t.Logf("%v", about.LoadBefore)
	//t.Logf("%v", about.LoadBeforeByVersion)
}
