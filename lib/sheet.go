package tracking

import (
	"context"
	"io/ioutil"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

type SheetClient struct {
	srv           *sheets.Service
	spreadsheetID string
}

func NewSheetClient(ctx context.Context) (*SheetClient, error) {
	secretBlob, err := ioutil.ReadFile(secretPath)
	if err != nil {
		return nil, err
	}

	// read & write permission
	jwt, err := google.JWTConfigFromJSON(secretBlob, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, err
	}
	srv, err := sheets.New(jwt.Client(ctx))
	if err != nil {
		return nil, err
	}

	spreadsheetID, err := GetConfigString("spread_sheet_id")
	if err != nil {
		return nil, err
	}

	return &SheetClient{
		srv:           srv,
		spreadsheetID: spreadsheetID,
	}, nil
}

func (s *SheetClient) Get(range_ string) ([][]interface{}, error) {
	resp, err := s.srv.Spreadsheets.Values.Get(s.spreadsheetID, range_).Do()
	if err != nil {
		return nil, err
	}
	return resp.Values, nil
}

func (s *SheetClient) Append(sheetName string, values [][]interface{}) error {
	_, err := s.srv.Spreadsheets.Values.Append(s.spreadsheetID, sheetName, &sheets.ValueRange{
		Values: values,
	}).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Do()
	if err != nil {
		return err
	}
	return nil
}

func (s *SheetClient) Update(range_ string, values [][]interface{}) error {
	_, err := s.srv.Spreadsheets.Values.Update(s.spreadsheetID, range_, &sheets.ValueRange{
		Values: values,
	}).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		return err
	}
	return nil
}
