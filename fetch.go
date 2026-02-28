package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// APIResponse represents the common JSON structure returned by NASA's
// SSD/CNEOS APIs, where fields lists column names and data contains rows
// as string arrays.
type APIResponse struct {
	Signature struct {
		Version string `json:"version"`
		Source  string `json:"source"`
	} `json:"signature"`
	Count  int      `json:"count"`
	Fields []string `json:"fields"`
	Data   [][]any  `json:"data"`
}

// fetchAPI calls the given NASA API URL and decodes the JSON response
// into an APIResponse.
func fetchAPI(url string) (*APIResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http get failed: %w", err)
	}
	defer resp.Body.Close()

	var data APIResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("json decode failed: %w", err)
	}

	return &data, nil
}

// fetchCAD retrieves close approach data from NASA's CAD API.
// It returns all NEO Earth close approaches between 1900-2100 within 0.2 au.
func fetchCAD() (*APIResponse, error) {
	url := "https://ssd-api.jpl.nasa.gov/cad.api?date-min=1900-01-01&date-max=2100-01-01&dist-max=0.2&diameter=true&fullname=true"
	return fetchAPI(url)
}

// fetchSBDB retrieves the full Small-Body Database from NASA's SBDB Query API,
// including orbital elements and physical properties for all known objects.
func fetchSBDB() (*APIResponse, error) {
	url := "https://ssd-api.jpl.nasa.gov/sbdb_query.api?fields=spkid,full_name,kind,neo,pha,H,diameter,albedo,rot_per,spec_T,spec_B,e,a,q,i,om,w,ma,tp,per,n,ad,epoch"
	return fetchAPI(url)
}

// writeCSV writes fields as the header row and data as subsequent rows
// to the given filename in CSV format. Values are converted to strings,
// with nil values written as empty strings.
func writeCSV(filename string, fields []string, data [][]any) error {
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

	row := make([]string, len(fields))
	for _, record := range data {
		for j, val := range record {
			if val == nil {
				row[j] = ""
			} else if num, ok := val.(float64); ok {
				if num == float64(int64(num)) {
					row[j] = fmt.Sprintf("%d", int64(num))
				} else {
					row[j] = fmt.Sprintf("%g", num)
				}
			} else {
				row[j] = fmt.Sprintf("%v", val)
			}
		}
		err = writer.Write(row)
		if err != nil {
			return fmt.Errorf("write row failed: %w", err)
		}
	}

	return nil
}
