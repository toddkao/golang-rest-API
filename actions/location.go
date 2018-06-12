package actions

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gobuffalo/buffalo"
	db "github.com/toddkao/ecomm2/models"
	re "gopkg.in/gorethink/gorethink.v4"
	"gopkg.in/validator.v2"
)

// LocationGroup struct describes fields in Location
type LocationGroup struct {
	ID          string    `json:"id,omitempty" gorethink:"id,omitempty"`
	Name        string    `validate:"nonzero" json:"name" gorethink:"name"`
	Description string    `validate:"nonzero" json:"description" gorethink:"description"`
	Status      byte      `validate:"nonzero" json:"status" gorethink:"status"`
	Created     time.Time `json:"created_at" gorethink:"created_at"`
	Updated     time.Time `json:"updated_at" gorethink:"updated_at"`
}

// GetAllLocationGroupHandler returns all Location groups
func GetAllLocationGroupHandler(c buffalo.Context) error {
	// Fetch all LocationGroups from the database
	list := make([]LocationGroup, 0)
	query := re.Table("locationsgroup")
	res, err := query.Run(db.Session)
	err = res.All(&list)
	if err != nil {
		err = errors.New("invalid credentials")
	}
	fmt.Println(list)
	return c.Render(200, r.JSON(&list))
}

// InsertLocationGroupHandler displays posted data
func InsertLocationGroupHandler(c buffalo.Context) error {
	request := c.Request()
	// fmt.Println(request)
	var newLocationGroup LocationGroup
	newLocationGroup.Created = time.Now()
	newLocationGroup.Updated = time.Now()
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&newLocationGroup)
	if err != nil {
		err = errors.New("missing fields")
	}
	if errs := validator.Validate(newLocationGroup); errs != nil {
		return errs
	}
	fmt.Println(newLocationGroup)

	_, err = re.Table("locationsgroup").Insert(newLocationGroup).RunWrite(db.Session)
	if err != nil {
		panic(err)
	}

	return c.Render(200, r.JSON(newLocationGroup))
}

// DeleteLocationGroupHandler deletes locationGroup with specific id
func DeleteLocationGroupHandler(c buffalo.Context) error {
	// Fetch all the items from the database
	// list := make([]LocationGroup, 0)
	// query := re.Table("locationsgroup")
	// res, err := query.Run(db.Session)
	// err = res.All(&list)
	// if err != nil {
	// 	panic(err)
	// }
	id := c.Param("id")
	res, err := re.Table("locationsgroup").Get(id).Run(db.Session)
	if err != nil {
		return c.Render(200, r.JSON("locationgroup"+id+" not found"))
	}

	if res.IsNil() {
		return c.Render(200, r.JSON("locationgroup"+id+" not found"))
	}
	_, err = re.Table("locationsgroup").Get(id).Delete().RunWrite(db.Session)
	if err != nil {
		return c.Render(200, r.JSON("could not delete locationgroup"+id))
	}
	fmt.Println(res)
	return c.Render(200, r.JSON(map[string]string{"success": "location group " + id + " deleted"}))
}
