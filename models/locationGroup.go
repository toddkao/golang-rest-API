package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gobuffalo/buffalo"
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

// ShowAll returns all Location groups
func (lg *LocationGroup) ShowAll(c buffalo.Context) error {
	list := make([]LocationGroup, 0)
	query := re.Table("locationsgroup")
	res, err := query.Run(Session)
	err = res.All(&list)
	if err != nil {
		err = errors.New("invalid credentials")
	}
	fmt.Println(list)
	return c.Render(200, r.JSON(&list))
}

// Insert inserts a new location group to the table
func (lg *LocationGroup) Insert(c buffalo.Context) error {
	var err error
	request := c.Request()
	var newLocationGroup LocationGroup
	newLocationGroup.Created = time.Now()
	newLocationGroup.Updated = time.Now()
	decoder := json.NewDecoder(request.Body)
	if err = decoder.Decode(&newLocationGroup); err != nil {
		err = errors.New("missing fields")
	}
	if err = validator.Validate(newLocationGroup); err != nil {
		return err
	}
	fmt.Println(newLocationGroup)

	_, err = re.Table("locationsgroup").Insert(newLocationGroup).RunWrite(Session)
	if err != nil {
		panic(err)
	}

	return c.Render(200, r.JSON(newLocationGroup))
}

// Delete deletes locationGroup with specific id
func (lg *LocationGroup) Delete(c buffalo.Context) error {
	id := c.Param("id")
	res, err := re.Table("locationsgroup").Get(id).Run(Session)

	if res.IsNil() || err != nil {
		return c.Render(400, r.JSON("locationgroup "+id+" not found"))
	}
	re.Table("locationsgroup").Get(id).Delete().RunWrite(Session)
	fmt.Println(res)
	return c.Render(200, r.JSON("location group "+id+" deleted"))
}
