package otp

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/x/responder"

	"identity/actions"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Otp)
// DB Table: Plural (otps)
// Resource: Plural (Otps)
// Path: Plural (/otps)
// View Template Folder: Plural (/templates/otps/)

// OtpsResource is the resource for the Otp model
type OtpsResource struct {
	buffalo.Resource
}

// List gets all Otps. This function is mapped to the path
// GET /otps
func (v OtpsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	otps := &Otps{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Otps from the DB
	if err := q.All(otps); err != nil {
		return err
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(200, actions.R.JSON(otps))
	}).Respond(c)
}

// Show gets the data for one Otp. This function is mapped to
// the path GET /otps/{otp_id}
func (v OtpsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Otp
	otp := &Otp{}

	// To find the Otp the parameter otp_id is used.
	if err := tx.Find(otp, c.Param("otp_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(200, actions.R.JSON(otp))
	}).Respond(c)
}

// Create adds a Otp to the DB. This function is mapped to the
// path POST /otps
func (v OtpsResource) Create(c buffalo.Context) error {
	// Allocate an empty Otp
	otp := &Otp{}

	// Bind otp to the html form elements
	if err := c.Bind(otp); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(otp)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, actions.R.JSON(verrs))
		}).Respond(c)
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, actions.R.JSON(otp))
	}).Respond(c)
}

// Update changes a Otp in the DB. This function is mapped to
// the path PUT /otps/{otp_id}
func (v OtpsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Otp
	otp := &Otp{}

	if err := tx.Find(otp, c.Param("otp_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Otp to the html form elements
	if err := c.Bind(otp); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(otp)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, actions.R.JSON(verrs))
		}).Respond(c)
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, actions.R.JSON(otp))
	}).Respond(c)
}

// Destroy deletes a Otp from the DB. This function is mapped
// to the path DELETE /otps/{otp_id}
func (v OtpsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Otp
	otp := &Otp{}

	// To find the Otp the parameter otp_id is used.
	if err := tx.Find(otp, c.Param("otp_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(otp); err != nil {
		return err
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, actions.R.JSON(otp))
	}).Respond(c)
}
