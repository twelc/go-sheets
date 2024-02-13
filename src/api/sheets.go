package api

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
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

// Retrieve a token, saves the token, then returns the generated client.

func GetData() *sheets.ValueRange {
	creds := ParseCredentials()
	// Create a JWT configurations object for the Google service account
	conf := &jwt.Config{
		Email:        creds.Email,
		PrivateKey:   []byte(creds.PrivateKey),
		PrivateKeyID: creds.PrivateKeyID,
		TokenURL:     google.JWTTokenURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/spreadsheets.readonly",
		},
	}

	cont := conf.Client(oauth2.NoContext)

	// Create a service object for Google sheets
	srv, err := sheets.New(cont)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	table_conf := ParseConfig()
	spreadsheetId := table_conf.SheetId
	readRange := fmt.Sprintf("%v!%v", table_conf.TableName, table_conf.Columns)
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}
	return resp
}

func GetAll() [][]string {
	resp := GetData()
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

func GetFiltered(querry string, min string, max string) [][]string {
	resp := GetData()
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
