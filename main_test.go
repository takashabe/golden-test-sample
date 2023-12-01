package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
)

func TestGolden(t *testing.T) {
	got, err := genXlsx()
	assert.NoError(t, err)
	gotContent, err := io.ReadAll(got)
	assert.NoError(t, err)

	equalFn := func(actual, expected []byte) bool {
		act, err := excelize.OpenReader(bytes.NewBuffer(actual))
		if err != nil {
			t.Error(err)
			return false
		}
		actCols, err := act.GetCols("Sheet1")
		if err != nil {
			t.Error(err)
			return false
		}
		exp, err := excelize.OpenReader(bytes.NewBuffer(expected))
		if err != nil {
			t.Error(err)
			return false
		}
		expCols, err := exp.GetCols("Sheet1")
		if err != nil {
			t.Error(err)
			return false
		}
		if len(actCols) != len(expCols) {
			t.Error("len(actCols) != len(expCols)")
			return false
		}
		for i, col := range expCols {
			for j, cell := range col {
				diff := cmp.Diff(cell, actCols[i][j])
				if diff != "" {
					t.Error(diff)
					return false
				}
			}
		}
		return true
	}

	g := goldie.New(t,
		goldie.WithDiffFn(func(a, b string) string {
			return "Diff is not shown. See golden files"
		}),
		goldie.WithEqualFn(equalFn),
	)
	g.Assert(t, t.Name(), gotContent)
}
