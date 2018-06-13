package controllers

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

// LocationGroup alias
type LocationGroup db.LocationGroup

// ShowAll Location groups
func (lg *LocationGroup) ShowAll(c buffalo.Context) error {
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

// Insert a new locationgroup to the database
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

	_, err = re.Table("locationsgroup").Insert(newLocationGroup).RunWrite(db.Session)
	if err != nil {
		panic(err)
	}

	return c.Render(200, r.JSON(newLocationGroup))
}

// Delete locationGroup with specific id
func (lg *LocationGroup) Delete(c buffalo.Context) error {
	id := c.Param("id")
	res, err := re.Table("locationsgroup").Get(id).Run(db.Session)

	if res.IsNil() || err != nil {
		return c.Render(400, r.JSON("locationgroup "+id+" not found"))
	}
	re.Table("locationsgroup").Get(id).Delete().RunWrite(db.Session)
	fmt.Println(res)
	return c.Render(200, r.JSON("location group "+id+" deleted"))
}
