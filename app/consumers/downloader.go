package consumers
import (
	"encoding/json"
	"net/http"
	"io"
	"reflect"
	"github.com/revel/revel"
)

type UserServiceAuth struct {
	Username string
	Password string
}

type Downloader struct {
	Auth *UserServiceAuth
	Url string
	Response Response
}

// Used to return the response body
type Response interface{}

// Returns the json decoded body
func (d *Downloader) JsonResponse(body io.Reader) (error) {
	revel.TRACE.Printf("Putting %s into %s\n", body, reflect.TypeOf(d.Response))
	return json.NewDecoder(body).Decode(&d.Response)
}

func (d *Downloader) GetPage() (error) {
	revel.TRACE.Printf("pg: %s\n", d.Url)
	revel.TRACE.Println("Before settings json ", d.Response)
	// Already set, bounce out. Used for testing
	if d.Response != nil {
		return nil
	}
	resp, err := http.Get(d.Url)
	perror(err)
	defer resp.Body.Close()
	return d.JsonResponse(resp.Body)
}

// Create the new Downloader and assign the Response type
// Response is where we will stuff the JSON response
func NewDownloader(url string, responseHolder Response) *Downloader {
	rt := reflect.ValueOf(responseHolder)
	revel.TRACE.Println("Response type: ", rt)
	return &Downloader{Url: url, Response: responseHolder }
}

