# Moesif GCP "Cloud Functions" example for Go to log outbound HTTP(s) calls
[Moesif](https://www.moesif.com) is an API analytics platform.
[Cloud Functions](https://cloud.google.com/functions) is a serverless Function as a service (Faas) offer from [Google Cloud Platform (GCP)](https://cloud.google.com)

### Scenario:
A Moesif client who utilizes Cloud Functions, written in Go Lang, would like to log all "outbound" calls originating from their cloud functions to be logged into Moesif. 

### Solution:
This is a simple "Cloud function" that shows how to integrate Moesif outbound call logging.

### Overview:
1. Create a Cloud Function in `go` by following [GCP Quickstart](https://cloud.google.com/functions/docs/quickstart-go). Deploy and verify it works.
2. Copy `moesif-init.go` from this project into your Cloud Function project. Update the `package` name to match yours.
3. In your Cloud Functions file (eg: `function.go`) add/update the init function to invoke `moesifInit`:
```go
func init() {
	moesifInit()
}
```
4. In GCP console, add the following `Runtime environment variable`

```MOESIF_APPLICATION_ID``` : ```<enter the moesif application id here>```

---

# Try it in 2 minutes!
Follow 3 simple steps to create sample function, invoke it, and see the results in Moesif dashboard:

[HOWTO quickly run sample in GCP using ```gcloud``` tool](README-RUN-SAMPLE.md)

