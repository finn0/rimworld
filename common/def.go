package common

type RootDef struct {
	Defs []*Def `xml:"Defs"`
}

type Def struct {
	Name string `xml:"defName"`
}

// ┌About (Your mod MUST have an About folder containing an About.xml file. Case sensitive)
// ├─About.xml (Contains info about the mod)
// ├─Preview.png (Image that appears above the mod info in-game)
// │
// ├Assemblies (If your mod uses any DLL files, put them here)
// ├─MyMod.dll
// │
// ├Defs (Contains XML definitions of the mod)
// ├┬ThingDefs
// │├─Things.xml
// │└─Buildings.xml
// │
// ├┬ResearchProjectDefs
// │└─MyProjects.xml
// │
// │
// ├┬Sounds (Put any sound files here. Universally supported formats are .ogg and .wav. .mp3 files are not guaranteed to work)
// ├─MySound.wav
// |
// ├Source
// ├─MyMod.cs (Optionally, put the source code of your mod here)
// │
// ├Patches
// ├─MyPatch.xml
// │
// ├Languages
// ├┬English (Replace with the language name)
// │├┬Keyed
// ││└─Keys.xml
// │├┬Strings
// ││└┬Names
// ││ └─PawnNames.xml
// │├┬DefInjected (NOTE: the folder (and subfolder) names must be specific here and follow the XML structure of the mod)
// ││└┬ThingDef
// ││ └─Thing.xml
// │├─LanguageInfo.xml
// │└─LangIcon.png
// │
// ├Textures (Put any image textures here, preferably in .png format.)
// ├┬Things
// │├─MyMod_ImageA.png
// │└─MyMod_ImageB.png
