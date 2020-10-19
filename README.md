### GCP Create cloud function
in GCP console, navigate to your function that you would like to integrate with Moesif.
Add the following "Runtime environment variable"

```MOESIF_APPLICATION_ID``` : ```<enter the moesif application id here>```

Now, modify your existing GCP function.go file:
1. Import middleware 'github.com/moesif/moesifmiddleware-go'
Create a new func init() (or modify existing func) add:
2. Load environment variable MOESIF_APPLICATION_ID and set it in moesifmiddleware StartCaptureOutgoing
The init() function is called once when the container is (re)loaded, and not for any subsequent invocations.

All future

Example: see ```news-bot-example.go```

### Local testing and testing in Google cloud shell
