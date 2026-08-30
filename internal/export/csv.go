package export

import (
	"encoding/csv"
	"errors"
	"github.com/Kotoninja/expense-tracker/internal/service"
	"os"
	"strconv"
)

func GenerateCSV(data []service.ExpenseDTO) error {
	file, err := os.Create("storage.csv")
	if err != nil {
		return errors.New(err.Error())
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	output := [][]string{{"ID", "Description", "Amount", "Date"}}

	for _, value := range data {
		output = append(output, []string{
			strconv.Itoa(value.ID),
			value.Description,
			strconv.Itoa(value.Amount),
			value.Date,
		})
	}

	for _, record := range output {
		err := writer.Write(record)
		if err != nil {
			return errors.New(err.Error())
		}
	}
	writer.Flush()

	return nil
}
