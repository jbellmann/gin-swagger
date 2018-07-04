package config_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"
	"github.com/jbellmann/gin-swagger/api"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jbellmann/gin-swagger/example/models"
)

// AddOrUpdateConfigItemEndpoint executes the core logic of the related
// route endpoint.
func AddOrUpdateConfigItemEndpoint(handler func(ctx *gin.Context, params *AddOrUpdateConfigItemParams) *api.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// generate params from request
		params := NewAddOrUpdateConfigItemParams()
		err := params.readRequest(ctx)
		if err != nil {
			errObj := err.(*errors.CompositeError)
			problem := api.Problem{
				Title:  "Unprocessable Entity.",
				Status: int(errObj.Code()),
				Detail: errObj.Error(),
			}
			ctx.Writer.Header().Set("Content-Type", "application/problem+json")
			ctx.JSON(problem.Status, problem)
			return
		}

		resp := handler(ctx, params)
		switch resp.Code {
		case http.StatusNoContent:
			ctx.AbortWithStatus(resp.Code)
		default:
			ctx.JSON(resp.Code, resp.Body)
		}
	}
}

// NewAddOrUpdateConfigItemParams creates a new AddOrUpdateConfigItemParams object
// with the default values initialized.
func NewAddOrUpdateConfigItemParams() *AddOrUpdateConfigItemParams {
	var ()
	return &AddOrUpdateConfigItemParams{}
}

// AddOrUpdateConfigItemParams contains all the bound params for the add or update config item operation
// typically these are obtained from a http.Request
//
// swagger:parameters addOrUpdateConfigItem
type AddOrUpdateConfigItemParams struct {

	/*ID of the cluster.
	  Required: true
	  Pattern: ^[a-z][a-z0-9-:]*[a-z0-9]$
	  In: path
	*/
	ClusterID string
	/*Key for the config value.
	  Required: true
	  Pattern: ^[a-z][a-z0-9_]*[a-z0-9]$
	  In: path
	*/
	ConfigKey string
	/*Config value.
	  Required: true
	  In: body
	*/
	Value *models.ConfigValue
}

// readRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *AddOrUpdateConfigItemParams) readRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	rClusterID := []string{ctx.Param("cluster_id")}
	if err := o.bindClusterID(rClusterID, true, formats); err != nil {
		res = append(res, err)
	}

	rConfigKey := []string{ctx.Param("config_key")}
	if err := o.bindConfigKey(rConfigKey, true, formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(ctx.Request) {
		var body models.ConfigValue
		if err := ctx.BindJSON(&body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("value", "body"))
			} else {
				res = append(res, errors.NewParseError("value", "body", "", err))
			}

		} else {
			if err := body.Validate(formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Value = &body
			}
		}

	} else {
		res = append(res, errors.Required("value", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddOrUpdateConfigItemParams) bindClusterID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ClusterID = raw

	if err := o.validateClusterID(formats); err != nil {
		return err
	}

	return nil
}

func (o *AddOrUpdateConfigItemParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Pattern("cluster_id", "path", o.ClusterID, `^[a-z][a-z0-9-:]*[a-z0-9]$`); err != nil {
		return err
	}

	return nil
}

func (o *AddOrUpdateConfigItemParams) bindConfigKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ConfigKey = raw

	if err := o.validateConfigKey(formats); err != nil {
		return err
	}

	return nil
}

func (o *AddOrUpdateConfigItemParams) validateConfigKey(formats strfmt.Registry) error {

	if err := validate.Pattern("config_key", "path", o.ConfigKey, `^[a-z][a-z0-9_]*[a-z0-9]$`); err != nil {
		return err
	}

	return nil
}

// vim: ft=go
