package jaccard

import (
	"errors"
	"strings"

	"github.com/aaaton/golem"
	"github.com/aaaton/golem/dicts/en"
)

const (
	PrimaryStringKey   = "primaryStringKey"
	SecondaryStringKey = "secondaryStringKey"
)

// NOTE: github.com/aaaton/golem/dicts/en needs to be expanded. Once you install
// the package, go into the directory of the dialect you want:
// Step 1: cd github.com/aaaton/golem/dicts/en
//
// Step 2: Install git lfs by running: brew install git-lfs
// Step 3: git checkout .
// Step 4: git lfs pull

// JaccardSim computes the Jaccard Similarity of two texts
type JaccardSim struct {
	PrimaryText         string          // Store primary text
	SecondaryText       string          // Store text that is being compared
	primarySet          map[string]bool // Set of lemmas from primaryText
	secondarySet        map[string]bool // Set of lemmas from secondaryText
	primaryLemmatizer   *golem.Lemmatizer
	secondaryLemmatizer *golem.Lemmatizer
}

// NewJaccardSim returns a new instance of JaccardSim
func NewJaccardSim(text1 string, text2 string) *JaccardSim {
	lemmatizer1, err := golem.New(en.New())
	if err != nil {
		panic(err)
	}
	lemmatizer2, err := golem.New(en.New())
	if err != nil {
		panic(err)
	}
	return &JaccardSim{
		PrimaryText:         text1,
		SecondaryText:       text2,
		primaryLemmatizer:   lemmatizer1,
		secondaryLemmatizer: lemmatizer2,
	}
}

// GetLemma returns the lemma of a given word
func GetLemma(inp string) string {
	lemmatizer, err := golem.New(en.New())
	if err != nil {
		panic(err)
	}
	word := lemmatizer.Lemma(inp)
	return word
}

func (j *JaccardSim) BuildSets() error {
	go j.buildSet(PrimaryStringKey)
	go j.buildSet(SecondaryStringKey)
	return nil
}

// buildSet builds the sets required to calculate Jaccard Similarity coefficient
func (j *JaccardSim) buildSet(textPos string) (map[string]bool, error) {
	set := make(map[string]bool)

	if textPos != PrimaryStringKey && textPos != SecondaryStringKey {
		return nil, errors.New("textPosition must be either primaryStringKey or secondaryStringKey")
	}

	// Set string and lemmatizer
	s := j.PrimaryText
	lemmatizer := j.primaryLemmatizer
	if textPos == SecondaryStringKey {
		s = j.SecondaryText
		lemmatizer = j.secondaryLemmatizer
	}

	// Check if lemmatizer is nil
	if lemmatizer == nil {
		return nil, errors.New("Lemmatizer is nil")
	}

	// If string is empty, return
	if s == "" {
		return set, nil
	}

	var lemma string
	words := strings.Fields(s)

	// Build map of lemmas
	for _, w := range words {
		lemma = lemmatizer.Lemma(w)
		set[lemma] = true
	}

	// Assign map to corresponding set
	if textPos == SecondaryStringKey {
		j.secondarySet = set
		return set, nil
	}
	j.primarySet = set
	return set, nil
}
