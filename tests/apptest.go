package tests

import (
	"github.com/revel/revel/testing"
	"github.com/memikequinn/lfs-server-go/app/consumers"
	"fmt"
//	"encoding/json"
	"bytes"
	"io"
//	"reflect"
)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) TestThatObjectsPageWorks() {
	t.Get("/objects")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func mock_get_page(url string, access bool) io.Reader {
//	fmt.Printf("Url is %s\n", url)
	return bytes.NewBufferString(fmt.Sprintf("{\"access\": %t, \"status\": \"yay\", \"message\": \"Some Message\"}", access))
}

func (t *AppTest) TestThatConsumerUserServiceRespondsWhenTrue() {
	uar := &consumers.UserAccessRequest{Username:"testuser", Project:"testproject", Action:"download"}
	downloader := consumers.NewUserService("http://somewhere.net", *uar)
	downloader.JsonResponse(mock_get_page("http://somewhere.net", true))
//	fmt.Printf("Can: %s\n", downloader.Can())
	t.AssertEqual(true, downloader.Can())
}

func (t *AppTest) TestThatConsumerUserServiceRespondsWhenFalse() {
	uar := &consumers.UserAccessRequest{Username:"testuser", Project:"testproject", Action:"download"}
	downloader := consumers.NewUserService("http://somewhere.net", *uar)
//	fmt.Printf("Can: %s\n", downloader.Can())
	t.AssertEqual(false, downloader.Can())
}

func (t *AppTest) TestThatConsumerUserServiceSetsMessageWhenInvalidAction() {
	uar := &consumers.UserAccessRequest{Username:"testuser", Project:"testproject", Action:"poo"}
	downloader := consumers.NewUserService("http://somewhere.net", *uar)
	t.AssertEqual(false, downloader.Can())
	myr := downloader.Response.(*consumers.UserAccessResponse)
	t.AssertEqual("poo is not in AllowedActions", myr.Message)
}

func (t *AppTest) TestConsumerUserServiceResponds_Can() {
	uar := &consumers.UserAccessRequest{Username:"testuser", Project:"testproject", Action:"download"}
	downloader := consumers.NewUserService("http://somewhere.net", *uar)
	downloader.JsonResponse(mock_get_page("http://somewhere.net", true))
//	fmt.Printf("Can: %s\n", downloader.Can())
	t.AssertEqual(true, downloader.Can())
}

func (t *AppTest) TestDownloaderAccessWhenTrue() {
	r := new(consumers.UserAccessResponse)
	d := consumers.NewDownloader("somewhere.in", r)
	d.JsonResponse(mock_get_page("http://somewhere.net", true))
	//Coerce to the type we want
	myr := d.Response.(*consumers.UserAccessResponse)
	t.AssertEqual(true, myr.Access)
	t.AssertEqual("yay", myr.Status)
	t.AssertEqual("Some Message", myr.Message)
}

func (t *AppTest) TestDownloaderAccessWhenFalse() {
	r := new(consumers.UserAccessResponse)
	d := consumers.NewDownloader("somewhere.in", r)
	d.JsonResponse(mock_get_page("http://somewhere.net", false))
	//Coerce to the type we want
	myr := d.Response.(*consumers.UserAccessResponse)
	t.AssertEqual(false, myr.Access)
	t.AssertEqual("yay", myr.Status)
	t.AssertEqual("Some Message", myr.Message)
}

func (t *AppTest) After() {
	println("Tear down")
}
