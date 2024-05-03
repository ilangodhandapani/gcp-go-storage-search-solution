# Purpose

A storage and search solution in google cloud using bucket and firestore. Create, Read, Search and Delete files using the apis.


# Instructions

1. Set up gcloud locally vy downloading google cli https://cloud.google.com/sdk/docs/install

2. Initiate gcloud

./google-cloud-sdk/bin/gcloud init

3. Authenticate to google cloud
./google-cloud-sdk/bin/gcloud auth login
./google-cloud-sdk/bin/config set project PROJECT_ID
./google-cloud-sdk/bin/gcloud auth application-default login

4. Create bucket and Firestore collection, use same name in both places.

5. Download this repo to your local and initialize go module
  go mod init
  go mod tidy

6. Export env variables present in exportlocalenv.sh and start the app locally. 
  go run routes.go

7. Check if app is up. "http://localhost:8080"

8. Use the below APIs in postman to create, read, search and delete functions.

a. Create/Upload file along with metadata.

    curl --location --request POST 'http://localhost:8080/<your-project-id>/<your-bucket-collection-name>/createobjectmetadata/1.pdf' \
--header 'metadata: {"FileName":"1.pdf","FileSize":"123456","Location":"GA","FilePath":"/","ObjectType":"Test"}' \
--form 'file=@"/**/**/1.pdf"'

b. Read metadata of a file.

    curl --location --request GET 'http://localhost:8080/<your-project-id>/<your-bucket-collection-name>/readmetadata/abc1234'

c. Read/Download file by file name

    curl --location --request GET 'http://localhost:8080/<your-project-id>/<your-bucket-collection-name>/readobjectbyname/1.pdf'

d. Search by metadata

    curl --location --request GET 'http://localhost:8080/<your-project-id>/<your-bucket-collection-name>/searchobjectbymetadata' \
--header 'search: {"ObjectType":"Test","Location":"GA"}'

e. Delete file along with metadata

    curl --location --request DELETE 'http://localhost:8080/<your-project-id>/<your-bucket-collection-name>/deleteobject/1.pdf'

9. You can test this by pushing this code to cloud run also.
    https://cloud.google.com/run/docs/quickstarts/build-and-deploy/deploy-go-service

10. You can add Authentication under security as per your need, modify controllers to secure.
