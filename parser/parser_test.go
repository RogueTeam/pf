package parser_test

import (
	"encoding/json"
	"testing"

	"github.com/RogueTeam/pf/parser"
	"github.com/RogueTeam/pf/parser/testsuite"
	"github.com/stretchr/testify/assert"
)

func Test_Parse(t *testing.T) {
	t.Run("Samples", func(t *testing.T) {
		for _, filename := range testsuite.SampleFiles() {
			t.Run(filename, func(t *testing.T) {
				assertions := assert.New(t)

				file, err := testsuite.Samples.Open(filename)
				assertions.Nil(err, "failed to open filename")
				defer file.Close()

				conf, err := parser.ParseReader(file)
				if !assertions.Nil(err, "failed to parse configuration") {
					return
				}

				c, _ := json.MarshalIndent(conf, "", "\t")
				t.Log(string(c))
			})
		}
	})
}
