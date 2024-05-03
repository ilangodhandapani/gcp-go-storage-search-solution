package services

import (
	"APP-GO-GCP/logging"
	"context"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
)

// Permanently delete object from bucket and document from firestore
func DeleteObject(bucket, GCPobjectName, collection, projectID string) error {
	logging.Logging()
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		logging.ErrLog.Println(err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	o := client.Bucket(bucket).Object(GCPobjectName)
	_, err = o.Attrs(ctx)
	if err != nil {
		logging.ErrLog.Println(err)
	}

	o = client.Bucket(bucket).Object(GCPobjectName)

	// Optional: set a generation-match precondition to avoid potential race
	// conditions and data corruptions. The request to delete the file is aborted
	// if the object's generation number does not match your precondition.
	attrs, err := o.Attrs(ctx)
	if err != nil {
		logging.ErrLog.Println(err)
	}
	o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	if err := o.Delete(ctx); err != nil {
		logging.ErrLog.Println(err)
	}
	logging.InfoLog.Println("Hard deleted object ", GCPobjectName, " in bucket", bucket)
	return err
}
func DeleteDocFirestore(projectID, collection, GCPobjectName string) error {
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
	iter := client.Collection(collection).Where("FileName", "==", GCPobjectName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		docID := (doc.Ref.ID)

		_, err = client.Collection(collection).Doc(docID).Delete(ctx)
		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			logging.ErrLog.Println(err)
		}
		logging.InfoLog.Println("Hard Deleted document ", docID, " in collection", collection)
	}
	return nil
}
