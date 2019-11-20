// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/spkane/todo-for-terraform/models"
)

// NewUpdateOneParams creates a new UpdateOneParams object
// with the default values initialized.
func NewUpdateOneParams() *UpdateOneParams {
	var ()
	return &UpdateOneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOneParamsWithTimeout creates a new UpdateOneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateOneParamsWithTimeout(timeout time.Duration) *UpdateOneParams {
	var ()
	return &UpdateOneParams{

		timeout: timeout,
	}
}

// NewUpdateOneParamsWithContext creates a new UpdateOneParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateOneParamsWithContext(ctx context.Context) *UpdateOneParams {
	var ()
	return &UpdateOneParams{

		Context: ctx,
	}
}

// NewUpdateOneParamsWithHTTPClient creates a new UpdateOneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateOneParamsWithHTTPClient(client *http.Client) *UpdateOneParams {
	var ()
	return &UpdateOneParams{
		HTTPClient: client,
	}
}

/*UpdateOneParams contains all the parameters to send to the API endpoint
for the update one operation typically these are written to a http.Request
*/
type UpdateOneParams struct {

	/*Body*/
	Body *models.Item
	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update one params
func (o *UpdateOneParams) WithTimeout(timeout time.Duration) *UpdateOneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update one params
func (o *UpdateOneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update one params
func (o *UpdateOneParams) WithContext(ctx context.Context) *UpdateOneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update one params
func (o *UpdateOneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update one params
func (o *UpdateOneParams) WithHTTPClient(client *http.Client) *UpdateOneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update one params
func (o *UpdateOneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update one params
func (o *UpdateOneParams) WithBody(body *models.Item) *UpdateOneParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update one params
func (o *UpdateOneParams) SetBody(body *models.Item) {
	o.Body = body
}

// WithID adds the id to the update one params
func (o *UpdateOneParams) WithID(id int64) *UpdateOneParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update one params
func (o *UpdateOneParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
