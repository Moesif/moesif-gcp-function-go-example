# How to run this sample with `gcloud`

### In Google cloud shell or locally

Disclaimer: Intended to be run in a sandboxed test project. 

_Pre-requisites: Ensure ```gcloud``` is installed and correct project id, region and auth is set._

1. Set the Moesif Application ID key
```sh
export MY_MOE_APP_ID="ENTER YOUR MOESIF APP ID HERE"
```

2. Download sample code
```sh
TEMP_FLDR=/tmp/gcp-func-moesif/
rm -rf ${TEMP_FLDR}
mkdir -p ${TEMP_FLDR}
cd ${TEMP_FLDR}

git clone https://github.com/moesif/moesif-gcp-function-go-example.git

cd moesif-gcp-function-go-example
```

(Optional) To test it locally, run the following command. Requires ```go lang```
```sh
export MOESIF_APPLICATION_ID=${MY_MOE_APP_ID}
go run cmd/local-main.go
```
And browse to https://localhost:8080. _Tip: In Google Cloud Shell, use [Web Preview](https://cloud.google.com/shell/docs/using-web-preview)_



3. Deploy the function using ```gcloud``` (first ensure correct project id, region and auth is set)
```sh
gcloud functions deploy FetchNews \
--runtime go113 --trigger-http --allow-unauthenticated \
--set-env-vars="MOESIF_APPLICATION_ID=${MY_MOE_APP_ID}"
```

Invoke the function using gcloud 
```sh
gcloud functions call FetchNews
```
---
### Call/Invoke the function
Invoke the using browser: The URL of the function can be obtained either via console
or using command
```sh
gcloud functions describe FetchNews | grep "url:"
```

----
### Cleanup/Delete
Delete/cleanup the function (WARNING this will destroy the function)
```sh
gcloud functions delete FetchNews
```