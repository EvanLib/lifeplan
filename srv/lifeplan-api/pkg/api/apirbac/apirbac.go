package apirbac

import (
	"fmt"
	"net/http"

	"github.com/ProtocolONE/rbac"
	"github.com/labstack/echo/v4"
)

// ResourceRouter is the interface that wraps the basic GetOwner Method
type ResourceRouter interface {
	GetOwner(AppContext) (string, error)
}

// AppContext
type AppContext struct {
	echo.Context
	enf *rbac.Enforcer
}

// NewAppContextMiddleware returns a echo Middleware
func NewAppContextMiddleware(enf *rbac.Enforcer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			context := AppContext{
				enf:     enf,
				Context: c,
			}
			return next(context)
		}
	}
}

func (c *AppContext) CheckPermissions(userId, domain, resource, resourceId, owner, action string) error {
	ctx := rbac.Context{
		Domain:        domain,
		User:          userId,
		ResourceId:    resourceId,
		Resource:      resource,
		ResourceOwner: owner,
		Action:        action,
	}
	if c.enf.Enforce(ctx) == false {
		return NewServiceErrorf(http.StatusForbidden, "Enforce failed for user: `%s`, resource `%s` with id `%s` and action `%s` in domain `%s`", userId, resource, resourceId, action, domain)
	}
	return nil
}

func CheckPermissions(group *RbacGroup, router ResourceRouter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//type assertion
			appContext := c.(AppContext)

			paths := group.paths
			path := c.Path()
			perm, ok := paths[path]
			if !ok {
				return NewServiceErrorf(http.StatusForbidden, "Could not map `%s` in paths", path)
			}

			owner, err := router.GetOwner(appContext)
			if err != nil {
				return err
			}

			// TODO: Get userID from JWT.
			userId := appContext.QueryParam("userId")

			resourceID := "*"
			if perm[0] != "*" {
				resourceID = c.Param(perm[0])
				fmt.Printf("ResourceID %s \n", resourceID)
			}

			action := "any"
			method := c.Request().Method
			switch method {
			case echo.GET:
				action = "read"
			case echo.PUT:
			case echo.POST:
			case echo.PATCH:
			case echo.DELETE:
				action = "write"
			}

			err = appContext.CheckPermissions(userId, perm[2], perm[1], resourceID, owner, action)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}

			return next(c)
		}
	}
}
