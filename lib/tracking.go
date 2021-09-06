package tracking

import (
	"context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
)

type SheetClient struct {
	srv           *sheets.Service
	spreadsheetID string
}

func NewSheetClient(ctx context.Context, spreadsheetID string) (*SheetClient, error) {
	b, err := ioutil.ReadFile("secret.json")
	if err != nil {
		return nil, err
	}
	// read & write permission
	jwt, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, err
	}
	srv, err := sheets.New(jwt.Client(ctx))
	if err != nil {
		return nil, err
	}
	return &SheetClient{
		srv:           srv,
		spreadsheetID: spreadsheetID,
	}, nil
}
