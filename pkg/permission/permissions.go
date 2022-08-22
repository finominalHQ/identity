package permission

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
// Model: Singular (Permission)
// DB Table: Plural (permissions)
// Resource: Plural (Permissions)
// Path: Plural (/permissions)
// View Template Folder: Plural (/templates/permissions/)

// PermissionsResource is the resource for the Permission model
type PermissionsResource struct {
	buffalo.Resource
}

// List gets all Permissions. This function is mapped to the path
// GET /permissions
func (v PermissionsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	permissions := &Permissions{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Permissions from the DB
	if err := q.All(permissions); err != nil {
		return err
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(200, actions.R.JSON(permissions))
	}).Respond(c)
}

// Show gets the data for one Permission. This function is mapped to
// the path GET /permissions/{permission_id}
func (v PermissionsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Permission
	permission := &Permission{}

	// To find the Permission the parameter permission_id is used.
	if err := tx.Find(permission, c.Param("permission_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(200, actions.R.JSON(permission))
	}).Respond(c)
}

// Create adds a Permission to the DB. This function is mapped to the
// path POST /permissions
func (v PermissionsResource) Create(c buffalo.Context) error {
	// Allocate an empty Permission
	permission := &Permission{}

	// Bind permission to the html form elements
	if err := c.Bind(permission); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(permission)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, actions.R.JSON(verrs))
		}).Respond(c)
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, actions.R.JSON(permission))
	}).Respond(c)
}

// Update changes a Permission in the DB. This function is mapped to
// the path PUT /permissions/{permission_id}
func (v PermissionsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Permission
	permission := &Permission{}

	if err := tx.Find(permission, c.Param("permission_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Permission to the html form elements
	if err := c.Bind(permission); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(permission)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, actions.R.JSON(verrs))
		}).Respond(c)
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, actions.R.JSON(permission))
	}).Respond(c)
}

// Destroy deletes a Permission from the DB. This function is mapped
// to the path DELETE /permissions/{permission_id}
func (v PermissionsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Permission
	permission := &Permission{}

	// To find the Permission the parameter permission_id is used.
	if err := tx.Find(permission, c.Param("permission_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(permission); err != nil {
		return err
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, actions.R.JSON(permission))
	}).Respond(c)
}