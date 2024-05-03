package services

import (
	"APP-GO-GCP/logging"
	"APP-GO-GCP/models"
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
)

func SearchActiveObjects(collection, projectID string, c models.SearchMetadata) (results []string) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		logging.ErrLog.Println(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		logging.ErrLog.Println(err)
	}
	defer client.Close()
	results = make([]string, 0)

	coll := client.Collection(collection)
	var query firestore.Query
	for key, value := range c.SearchAttributes {
		if value != "" {
			query = coll.Where(key, "==", value)
		}
	}

	iter := query.Limit(10).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		results = append(results, fmt.Sprintf("%v", doc.Data()["FileName"]))
	}
	return results
}
