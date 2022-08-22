package main

import (
	"log"

	"identity/actions"
	"identity/pkg/otp"
	"identity/pkg/permission"
	"identity/pkg/role"
	"identity/pkg/role_permission"
	"identity/pkg/user"
	"identity/pkg/user_role"
)

// main is the starting point for your Buffalo application.
// You can feel free and add to this `main` method, change
// what it does, etc...
// All we ask is that, at some point, you make sure to
// call `app.Serve()`, unless you don't want to start your
// application that is. :)
func main() {
	app := actions.App()
	// app.Resource("/auth", auth.)
	app.Resource("/otp", otp.OtpsResource{})
	app.Resource("/permission", permission.PermissionsResource{})
	app.Resource("/role", role.RolesResource{})
	app.Resource("/role-permission", role_permission.RolePermissionsResource{})
	app.Resource("/user", user.UsersResource{})
	app.Resource("/user-role", user_role.UserRolesResource{})
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}

/*
# Notes about `main.go`

## SSL Support

We recommend placing your application behind a proxy, such as
Apache or Nginx and letting them do the SSL heavy lifting
for you. https://gobuffalo.io/en/docs/proxy

## Buffalo Build

When `buffalo build` is run to compile your binary, this `main`
function will be at the heart of that binary. It is expected
that your `main` function will start your application using
the `app.Serve()` method.

*/
