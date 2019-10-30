package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	prose "gopkg.in/jdkato/prose.v2"
)

// ProdigyTrainData holds annotated data from www.prodi.gy for supervised training of model
type ProdigyTrainData struct {
	Text   string
	Spans  []prose.LabeledEntity
	Answer string
}

// TrainData holds tagged data for supervised training of model
type TrainData struct {
	Text   string
	Answer string
}

// NewTrainData opens a file and returns an array of training data
func NewTrainData(csvFilePath string) ([]TrainData, error) {
	f, err := os.OpenFile(csvFilePath, os.O_RDONLY, os.ModePerm)
	var tr []TrainData
	defer f.Close()
	if err != nil {
		return []TrainData{}, err
	}

	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] != "class" && record[1] != "sms" {
			tr = append(tr, TrainData{
				Answer: record[0],
				Text:   record[1],
			})
		}
		fmt.Println(record)
	}

	return tr, nil
}
