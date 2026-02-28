package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type CADResponse struct {
	Signature struct {
		Version string `json:"version"`
		Source  string `json:"source"`
	} `json:"signature"`
	Count  int        `json:"count"`
	Fields []string   `json:"fields"`
	Data   [][]string `json:"data"`
}

func fetchCAD() (*CADResponse, error) {
	url := "https://ssd-api.jpl.nasa.gov/cad.api?date-min=1900-01-01&date-max=2100-01-01&dist-max=0.2&diameter=true&fullname=true"


	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http get failed: %w", err)
	}
	defer resp.Body.Close()

	var data CADResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("json decode failed: %w", err)
	}

	return &data, nil
}

func writeCSV(filename string, fields []string, data [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create file failed: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(fields)
	if err != nil {
		return fmt.Errorf("write header failed: %w", err)
	}

	for _, row := range data {
		err = writer.Write(row)
		if err != nil {
			return fmt.Errorf("write row failed: %w", err)
		}
	}

	return nil
}
