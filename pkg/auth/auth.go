package auth

import (
	"github.com/gobuffalo/buffalo"
)

// Endpoints
// /auth/register/initiate
// /auth/register/verify

// /auth/login/initiate
// /auth/login/verify

// /auth/am-i
// /auth/can-i

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Auth)
// DB Table: Plural (auths)
// Resource: Plural (Auths)
// Path: Plural (/auths)
// View Template Folder: Plural (/templates/auths/)

// AuthsResource is the resource for the Auth model
type AuthsResource struct {
	buffalo.Resource
}
