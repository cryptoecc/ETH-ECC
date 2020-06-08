package miner

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestWritingCSV(t *testing.T) {

	csvFile, _ := os.OpenFile("output.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	wr := csv.NewWriter(csvFile)

	wr.Write([]string{"A", "0.25"})
	wr.Write([]string{"B", "55.70"})
	wr.Write([]string{"C", "60.70"})
	wr.Flush()
}
