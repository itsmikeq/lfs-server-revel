/*
Used to consume external authorization.
Given a key parameter, return a user's ability to download the project refs.
Uses https://github.com/bndr/gopencils
 */
package consumers
import (
	"fmt"
	"io"
	"github.com/revel/revel"

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
var AllowedActions = []string{"download", "push", "force_push", "admin"}
func vetAction(uareq UserAccessRequest) bool {
	for _, b :=  range AllowedActions {
		if b == uareq.Action {
			return true
		}
	}
	return false
}

func NewUserService(base string, uareq UserAccessRequest) *Downloader {
//	rr := new(UserAccessResponse)
	uar := new(UserAccessResponse)
	if vetAction(uareq) != true {
		revel.ERROR.Println(uareq.Action, "is not in AllowedActions")
		uar.Message = fmt.Sprintf("%s is not in AllowedActions", uareq.Action)
	}
	return NewDownloader(fmt.Sprintf("%s?username=%s&project=%s&action=%s", BASE_URL, uareq.Username, uareq.Project, uareq.Action), uar)
}

func (downloader *Downloader) Can() bool {
	downloader.GetPage()
	myr := downloader.Response.(*UserAccessResponse)
	return myr.Access
}
