package jaccard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLemma(t *testing.T) {
	testData := []struct {
		input  string
		output string
	}{
		{"abducting", "abduct"},
		{"is", "be"},
		{"ridden", "ride"},
		{"sat", "sit"},
	}
	for _, tt := range testData {
		assert.Equal(t, tt.output, GetLemma(tt.input))
	}
}

func TestBuildSets(t *testing.T) {
	testData := []struct {
		j       JaccardSim
		primary string
		output  map[string]bool
	}{
		{
			JaccardSim{PrimaryText: "concludes tests was lemmas Jaccard needs"},
			PrimaryStringKey,
			map[string]bool{
				"conclude": true,
				"need":     true,
				"test":     true,
				"be":       true,
				"lemmas":   true,
				"Jaccard":  true,
			},
		},
		{
			JaccardSim{SecondaryText: "concludes tests was lemmas Jaccard needs"},
			SecondaryStringKey,
			map[string]bool{
				"conclude": true,
				"need":     true,
				"test":     true,
				"be":       true,
				"lemmas":   true,
				"Jaccard":  true,
			},
		},
	}
	for _, tt := range testData {
		jc := tt.j
		set, err := jc.buildSet(tt.primary)
		assert.Nil(t, err)
		assert.NotNil(t, set)
		assert.Equal(t, tt.output, set)
	}
}

func TestBuildSetEmptyString(t *testing.T) {
	j := JaccardSim{}
	set, err := j.buildSet(PrimaryStringKey)
	expectedMap := make(map[string]bool)

	assert.Nil(t, err)
	assert.Equal(t, expectedMap, set)
}

func TestBuildSetBadPositionError(t *testing.T) {
	j := JaccardSim{}
	set, err := j.buildSet("IncorrectPositionKey")

	assert.NotNil(t, err)
	assert.Nil(t, set)
}
