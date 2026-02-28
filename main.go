package main

import "fmt"

func main() {
	cad, err := fetchCAD()
	if err != nil {
		fmt.Println("Error fetching CAD:", err)
		return
	}
	fmt.Printf("Fetched %d close approaches\n", cad.Count)

	err = writeCSV("close_approaches.csv", cad.Fields, cad.Data)
	if err != nil {
		fmt.Println("Error writing CAD CSV:", err)
		return
	}
	fmt.Println("Wrote close_approaches.csv")

	sbdb, err := fetchSBDB()
	if err != nil {
		fmt.Println("Error fetching SBDB:", err)
		return
	}
	fmt.Printf("Fetched %d small bodies\n", len(sbdb.Data))

	err = writeCSV("small_bodies.csv", sbdb.Fields, sbdb.Data)
	if err != nil {
		fmt.Println("Error writing SBDB CSV:", err)
		return
	}
	fmt.Println("Wrote small_bodies.csv")
}
