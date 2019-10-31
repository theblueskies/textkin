package jaccard

import (
	"sync"
	"testing"

	"github.com/aaaton/golem"
	"github.com/aaaton/golem/dicts/en"
	mapset "github.com/deckarep/golang-set"
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
	l, _ := golem.New(en.New())
	testData := []struct {
		j       JaccardSim
		primary string
		output  mapset.Set
	}{
		{
			JaccardSim{
				PrimaryText:       "concludes tests was lemmas Jaccard needs",
				primaryLemmatizer: l,
			},
			PrimaryStringKey,
			mapset.NewSetFromSlice([]interface{}{
				"conclude", "need", "test", "be", "lemmas", "Jaccard",
			}),
		},
		{
			JaccardSim{
				SecondaryText:       "concludes tests was lemmas Jaccard needs",
				secondaryLemmatizer: l,
			},
			SecondaryStringKey,
			mapset.NewSetFromSlice([]interface{}{
				"conclude", "need", "test", "be", "lemmas", "Jaccard",
			}),
		},
	}
	var wg sync.WaitGroup
	for _, tt := range testData {
		jc := tt.j
		wg.Add(1)
		set, err := jc.buildSet(&wg, tt.primary)
		assert.Nil(t, err)
		assert.Greater(t, set.Cardinality(), 0)
		assert.Equal(t, tt.output, set)
	}
}

func TestBuildSetEmptyString(t *testing.T) {
	l, _ := golem.New(en.New())
	var wg sync.WaitGroup
	wg.Add(1)
	j := JaccardSim{primaryLemmatizer: l}
	set, err := j.buildSet(&wg, PrimaryStringKey)

	assert.Nil(t, err)
	assert.Equal(t, mapset.NewSet(), set)
}

func TestBuildSetBadPositionError(t *testing.T) {
	j := JaccardSim{}
	var wg sync.WaitGroup
	wg.Add(1)
	set, err := j.buildSet(&wg, "IncorrectPositionKey")

	assert.NotNil(t, err)
	assert.Nil(t, set)
}

func TestNewJaccardSim(t *testing.T) {
	j := NewJaccardSim("some text", "random text")

	assert.Equal(t, "some text", j.PrimaryText)
	assert.Equal(t, "random text", j.SecondaryText)
	assert.NotNil(t, j.primaryLemmatizer)
	assert.NotNil(t, j.secondaryLemmatizer)
	assert.Nil(t, j.primarySet)
	assert.Nil(t, j.secondarySet)

	// Test Jaccard confidencce
	confidence := j.GetConfidence()
	assert.Equal(t, 0.3333333333333333, confidence)
}
