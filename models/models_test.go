package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadData(t *testing.T) {
	filePath := "testdata/test.csv"
	tr, err := NewTrainData(filePath)
	assert.Nil(t, err)
	fmt.Println(tr)
}
