/*
Used to consume external authorization.
Given a key parameter, return a user's ability to download the project refs.
Uses https://github.com/bndr/gopencils
 */
package consumers
import (
	"fmt"
	"io"
)

// declare the type of UserAccessGetter
// The url must conform to consumers_spec user_service spec
//type UserAccessGetter func(url string) string

type UserAccessResponse struct {
	Access bool `json:"access"`
	Status string `json:"status"`
	Message string `json:"message"`
	RawResponse io.Reader
}

type UserAccessRequest struct {
	Username string
	Project string
	Action string
}

type UserService struct {
	Downloader *Downloader
}
const BASE_URL = "http://localhost:3001/get_page"

func NewUserService(base string, uareq UserAccessRequest) *Downloader {
//	rr := new(UserAccessResponse)
	return NewDownloader(fmt.Sprintf("%s?username=%s&project=%s&action=%s", BASE_URL, uareq.Username, uareq.Project, uareq.Action), new(UserAccessResponse))
}

func (downloader *Downloader) Can() bool {
	downloader.GetPage()
	myr := downloader.Response.(*UserAccessResponse)
	return myr.Access
}
