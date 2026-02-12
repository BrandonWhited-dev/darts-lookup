package read

import (
	"fmt"
	"github.com/thedatashed/xlsxreader"
)

type Item struct {
	Title       string
	SKU         string
	UPC         string
	Price    string
}

func (i *Item) ToString() string {
	return fmt.Sprintf("Title: %s\nPrice: %s\nSKU: %s\nUPC: %s\n", i.Title, i.Price, i.SKU, i.UPC)
}

func (i *Item) ToStringShort() string {
	return fmt.Sprintf("Title: %s\nSKU: %s\nPrice: %s\n", i.Title, i.SKU, i.Price)
}

func ReadItems(path string) ([]Item, error) {
	xl, err := xlsxreader.OpenFile(path)
	if err != nil {
		return []Item{}, err
	}
	defer xl.Close()

	var items []Item
	first := true
	for row := range xl.ReadRows(xl.Sheets[0]) {
		if first == true {
			first = false
			continue
		}
		if len(row.Cells) == 4 {
			items = append(items, Item{
				Title:       row.Cells[0].Value,
				SKU:         row.Cells[1].Value,
				UPC:         row.Cells[2].Value,
				Price:    row.Cells[3].Value,
			})
		}
	}
	return items, nil
}
