package handler

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/unrolled/render"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// RenderJSON ... JSONをレンダリングする
func RenderJSON(w http.ResponseWriter, status int, v interface{}) {
	r := render.New(render.Options{IndentJSON: true})
	r.JSON(w, status, v)
}

// RenderCSV ... CSVをレンダリングする
func RenderCSV(w http.ResponseWriter, name string, data [][]string) {
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s.csv", name))
	writer := csv.NewWriter(transform.NewWriter(w, japanese.ShiftJIS.NewEncoder()))
	for _, datum := range data {
		writer.Write(datum)
	}
	writer.Flush()
}
