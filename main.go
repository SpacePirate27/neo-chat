package main

import "fmt"

func main() {
	data, err := fetchCAD()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Fetched %d close approaches\n", data.Count)

	err = writeCSV("close_approaches.csv", data.Fields, data.Data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Wrote close_approaches.csv")
}
