// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPcloudNetworksPortsGetallParams creates a new PcloudNetworksPortsGetallParams object
// with the default values initialized.
func NewPcloudNetworksPortsGetallParams() *PcloudNetworksPortsGetallParams {
	var ()
	return &PcloudNetworksPortsGetallParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudNetworksPortsGetallParamsWithTimeout creates a new PcloudNetworksPortsGetallParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudNetworksPortsGetallParamsWithTimeout(timeout time.Duration) *PcloudNetworksPortsGetallParams {
	var ()
	return &PcloudNetworksPortsGetallParams{

		timeout: timeout,
	}
}

// NewPcloudNetworksPortsGetallParamsWithContext creates a new PcloudNetworksPortsGetallParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudNetworksPortsGetallParamsWithContext(ctx context.Context) *PcloudNetworksPortsGetallParams {
	var ()
	return &PcloudNetworksPortsGetallParams{

		Context: ctx,
	}
}

// NewPcloudNetworksPortsGetallParamsWithHTTPClient creates a new PcloudNetworksPortsGetallParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudNetworksPortsGetallParamsWithHTTPClient(client *http.Client) *PcloudNetworksPortsGetallParams {
	var ()
	return &PcloudNetworksPortsGetallParams{
		HTTPClient: client,
	}
}

/*PcloudNetworksPortsGetallParams contains all the parameters to send to the API endpoint
for the pcloud networks ports getall operation typically these are written to a http.Request
*/
type PcloudNetworksPortsGetallParams struct {

	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud networks ports getall params
func (o *PcloudNetworksPortsGetallParams) WithTimeout(timeout time.Duration) *PcloudNetworksPortsGetallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud networks ports getall params
func (o *PcloudNetworksPortsGetallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud networks ports getall params
func (o *PcloudNetworksPortsGetallParams) WithContext(ctx context.Context) *PcloudNetworksPortsGetallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud networks ports getall params
func (o *PcloudNetworksPortsGetallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud networks ports getall params
func (o *PcloudNetworksPortsGetallParams) WithHTTPClient(client *http.Client) *PcloudNetworksPortsGetallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud networks ports getall params
func (o *PcloudNetworksPortsGetallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud networks ports getall params
func (o *PcloudNetworksPortsGetallParams) WithCloudInstanceID(cloudInstanceID string) *PcloudNetworksPortsGetallParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud networks ports getall params
func (o *PcloudNetworksPortsGetallParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithNetworkID adds the networkID to the pcloud networks ports getall params
func (o *PcloudNetworksPortsGetallParams) WithNetworkID(networkID string) *PcloudNetworksPortsGetallParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the pcloud networks ports getall params
func (o *PcloudNetworksPortsGetallParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudNetworksPortsGetallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
