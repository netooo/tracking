package tracking

import (
	"context"
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

func NewSheetClient(ctx context.Context, spreadsheetID string) (*SheetClient, error) {
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
	return &SheetClient{
		srv:           srv,
		spreadsheetID: spreadsheetID,
	}, nil
}
