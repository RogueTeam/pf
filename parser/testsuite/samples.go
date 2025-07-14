package testsuite

import (
	"embed"
	"path"
)

var SamplesDir = "samples"

//go:embed samples
var Samples embed.FS

func SampleFiles() (filenames []string) {
	entries, _ := Samples.ReadDir(SamplesDir)
	filenames = make([]string, 0, len(entries))
	for _, entry := range entries {
		filenames = append(filenames, path.Join(SamplesDir, entry.Name()))
	}
	return filenames
}
