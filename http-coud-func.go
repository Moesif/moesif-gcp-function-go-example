package p

/* This sample shows how subsequent outbound http calls from this function
can be logged into Moesif collector by utilizing moesifmiddleware-go library
*/

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/* init is called once per container launch.
 */
func init() {
	moesifInit() // Initialize Moesif middleware
}

func FetchNews(w http.ResponseWriter, r *http.Request) {
	// The following http outbound call will be logged into moesif api
	var myNewsSiteUrl string = "https://moesif.com/blog"
	var txt string = callWebhook(myNewsSiteUrl)
	fmt.Fprint(w, txt)
}

func callWebhook(webhook string) string {
	resp, _ := http.Get(webhook)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
