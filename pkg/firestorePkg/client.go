package firestorePkg

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

const (
	DefaultProjectID = "<YOUR_PROJECT_ID>"
	CredentialPath   = "<YOUR_CREDENTIAL_PATH>"
)

var (
	client *firestore.Client
)

func New() *firestore.Client {
	if client == nil {
		ctx := context.Background()
		data, err := os.ReadFile(CredentialPath)
		options := option.WithCredentialsJSON(data)
		client, err = firestore.NewClient(ctx, DefaultProjectID, options)
		if err != nil {
			log.Fatalf("firebase.NewClient err: %v", err)
		}
	}
	return client
}
