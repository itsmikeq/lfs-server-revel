package controllers
import (
	"encoding/json"
	"github.com/revel/revel"
	"github.com/memikequinn/lfs-server-go/app/models"
)
type Objects struct {
	GorpController
}

func (c Objects) parseObjectRequest() (models.Object, error) {
	object := models.Object{}
	err := json.NewDecoder(c.Request.Body).Decode(&object)
	return object, err
}

func (c Objects) Add() revel.Result {
	if object, err := c.parseObjectRequest(); err != nil {
		return c.RenderText("Unable to parse the GitObject from JSON.")
	} else {
		// Validate the model
		object.Validate(c.Validation)
		if c.Validation.HasErrors() {
			// Do something better here!
			return c.RenderText("You have error in your GitObject.")
		} else {
			if err := c.Txn.Insert(&object); err != nil {
				return c.RenderText(
					"Error inserting record into database!")
			} else {
				return c.RenderJson(object)
			}
		}
	}
}

func (c Objects) List() revel.Result {
	lastId := parseIntOrDefault(c.Params.Get("lid"), -1)
	limit := parseUintOrDefault(c.Params.Get("limit"), uint64(25))
	objects, err := c.Txn.Select(models.Object{},
		`SELECT * FROM Object WHERE Id > ? LIMIT ?`, lastId, limit)
	if err != nil {
		return c.RenderText(
			"Error trying to get records from DB.")
	}
	return c.RenderJson(objects)
}

func (c Objects) Update(id int64) revel.Result {
	object, err := c.parseObjectRequest()
	if err != nil {
		return c.RenderText("Unable to parse the Objects from JSON.")
	}
	// Ensure the Id is set.
	object.Id = id
	success, err := c.Txn.Update(&object)
	if err != nil || success == 0 {
		return c.RenderText("Unable to update bid item.")
	}
	return c.RenderText("Updated %v", id)
}

func (c Objects) Delete(id int64) revel.Result {
	success, err := c.Txn.Delete(&models.Object{Id: int64(id)})
	if err != nil || success == 0 {
		return c.RenderText("Failed to remove Objects")
	}
	return c.RenderText("Deleted %v", id)
}