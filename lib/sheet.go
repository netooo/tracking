package tracking

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rkoesters/xdg/basedir"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"path/filepath"
)

type SheetClient struct {
	srv           *sheets.Service
	spreadsheetID string
}

func NewSheetClient(ctx context.Context) (*SheetClient, error) {
	secretPath := filepath.Join(basedir.ConfigHome, "tracking", "secret.json")
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

	spreadsheetID, err := GetSheetID()
	return &SheetClient{
		srv:           srv,
		spreadsheetID: spreadsheetID,
	}, nil
}

func GetSheetID() (string, error) {
	configPath := filepath.Join(basedir.ConfigHome, "tracking", "config.json")
	configBlob, err := ioutil.ReadFile(configPath)
	if err != nil {
		return "", errors.New("command failed")
	}

	var configJson interface{}
	err = json.Unmarshal(configBlob, &configJson)
	if err != nil {
		return "", errors.New("command failed")
	}
	sheetId := configJson.(map[string]interface{})["spread_sheet_id"].(string)

	return sheetId, nil
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
