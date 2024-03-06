package lib

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func AppendIfMissing(slice []int, i int) []int {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
func GetClient(config *Config) *sheets.Service {
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(config.credentials), option.WithScopes(sheets.SpreadsheetsScope))

	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	return srv
}

// Retrieve a token, saves the token, then returns the generated client.
func GetData(config Config, range_ string) *sheets.ValueRange {

	srv := GetClient(&config)

	spreadsheetId := config.sheetid
	readRange := fmt.Sprintf("%v!%v", config.table_name, range_)
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}
	return resp
}

func GetAll(config Config, range_ string) [][]string {

	resp := GetData(config, range_)
	table := [][]string{}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			rrow := []string{}
			for _, cell := range row {
				rrow = append(rrow, cell.(string))
			}
			table = append(table, rrow)
		}
	}

	return table
}

func GetFiltered(querry string, min string, max string, range_ string, config Config) [][]string {
	resp := GetData(config, range_)
	table := [][]string{}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			rrow := []string{}
			if min != "" {
				min, _ := strconv.Atoi(min)
				val, _ := strconv.Atoi(row[2].(string))
				if min > val {
					continue
				}
			}
			if max != "" {
				max, _ := strconv.Atoi(max)
				val, _ := strconv.Atoi(row[2].(string))
				if max < val {
					continue
				}
			}
			adds := false
			for _, cell := range row {
				if querry == "" || strings.Contains(cell.(string), querry) {
					fmt.Println(cell.(string), querry, strings.Contains(cell.(string), querry))
					adds = true
				}
				rrow = append(rrow, cell.(string))
			}
			if adds {
				table = append(table, rrow)
			}

		}
	}

	return table
}

func AppendData(name string, distrct string, value string, ind int, config Config) {
	header := GetData(config, "A:A")
	row := -1
	for i, val := range header.Values {
		if val[0].(string) == name {
			row = i + 1
			break
		}
	}
	srv := GetClient(&config)
	var vr sheets.ValueRange
	vr.Values = append(vr.Values, []interface{}{value})

	if row == -1 {
		row = len(header.Values) + 1
		var svr sheets.ValueRange
		svr.Values = append(svr.Values, []interface{}{name, distrct})
		srv.Spreadsheets.Values.Update(config.sheetid, fmt.Sprintf("history!A%v:B%v", row, row), &svr).ValueInputOption("RAW").Do()
	}
	fmt.Println(row, ind)

	srv.Spreadsheets.Values.Update(config.sheetid, fmt.Sprintf("history!R%vC%v:R%vC%v", row, ind+3, row, ind+4), &vr).ValueInputOption("RAW").Do()

}

func SetLine(values []interface{}, range_ string, config Config) {
	srv := GetClient(&config)
	var vr sheets.ValueRange
	vr.Values = append(vr.Values, values)
	srv.Spreadsheets.Values.Update(config.sheetid, fmt.Sprintf("%v!%v", config.table_name, range_), &vr).ValueInputOption("RAW").Do()
}
