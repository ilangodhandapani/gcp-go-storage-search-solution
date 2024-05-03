package services

import (
	"APP-GO-GCP/logging"
	"APP-GO-GCP/models"
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
)

func CreateObjectMetadata(bucket string, file []byte, object string, m models.Metadata, creationTS string) (err error) {
	logging.Logging()
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		logging.ErrLog.Println(err)
	}a
	defer client.Close()

	buf := bytes.NewBuffer(file)
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	o := client.Bucket(bucket).Object(object)
	o = o.If(storage.Conditions{DoesNotExist: true})

	// Upload an object with storage.Writer.
	wc := o.NewWriter(ctx)
	wc.ChunkSize = 0 // note retries are not supported for chunk size 0.

	if _, err = io.Copy(wc, buf); err != nil {
		logging.ErrLog.Println(err)
	}
	// Data can continue to be added to the file until the writer is closed.
	if err := wc.Close(); err != nil {
		logging.ErrLog.Println(err)
	}
	o = client.Bucket(bucket).Object(object)
	attrs, err := o.Attrs(ctx)

	if err != nil {
		logging.ErrLog.Println(err)
	}
	o = o.If(storage.Conditions{DoesNotExist: false})
	o = o.If(storage.Conditions{MetagenerationMatch: attrs.Metageneration})

	objectAttrsToUpdate := storage.ObjectAttrsToUpdate{
		Metadata: map[string]string{
			"FileName":   m.Values[0].FileName,
			"FilePath":   m.Values[0].FilePath,
			"FileSize ":  m.Values[0].FileSize,
			"Location":   m.Values[0].Location,
			"ObjectType": m.Values[0].ObjectType,
			"CreationTS": creationTS,
		},
	}
	if _, err := o.Update(ctx, objectAttrsToUpdate); err != nil {
		logging.ErrLog.Println(err)
	}
	logging.InfoLog.Println("Success: Updated metadata for object", object, "in", bucket)
	return nil
}

// Function addDocFirestore uploads metadata to Firestore
func AddDocFirestore(projectID, collection string, m models.Metadata, creationTS string) (docID, FileName string, err error) {
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

	// Create doc first with the mandatory CM metadata
	_, _, err = client.Collection(collection).Add(ctx, map[string]interface{}{
		"FileName":   m.Values[0].FileName,
		"FilePath":   m.Values[0].FilePath,
		"FileSize ":  m.Values[0].FileSize,
		"Location":   m.Values[0].Location,
		"ObjectType": m.Values[0].ObjectType,
		"CreationTS": creationTS,
	})

	fmt.Println("Success: Add document", m.Values[0].FileName, " in collection", collection)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		logging.ErrLog.Println(err)
	}

	// Get the doc ID for the updated doc using objectName
	iter := client.Collection(collection).Where("FileName", "==", m.Values[0].FileName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		docID = (doc.Ref.ID)
	}
	return docID, FileName, err
}
