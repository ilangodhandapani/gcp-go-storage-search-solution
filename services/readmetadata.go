package services

import (
	"context"
	"encoding/json"
	"log"

	firebase "firebase.google.com/go"
)

func ReadMetadataByDocumentId(collection, projectID, DocumentId string) (metadata string, err1 error) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	dsnap, err1 := client.Collection(collection).Doc(DocumentId).Get(ctx)
	m := dsnap.Data()
	jsonStr, _ := json.Marshal(m)
	metadata = (string(jsonStr))

	return metadata, err1
}
