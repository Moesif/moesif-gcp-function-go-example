package p

import (
	"os"

	moesifmiddleware "github.com/moesif/moesifmiddleware-go"
)

/* Retrieve "MOESIF_APPLICATION_ID" from GCP environment varilable
and set it for 'moesifmiddleware-go'
*/
func moesifInit() {
	var moesifOptions = map[string]interface{}{
		"Application_Id": os.Getenv("MOESIF_APPLICATION_ID"),
	}
	moesifmiddleware.StartCaptureOutgoing(moesifOptions)
}
