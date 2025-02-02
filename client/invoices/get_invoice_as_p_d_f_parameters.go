package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewGetInvoiceAsPDFParams creates a new GetInvoiceAsPDFParams object
// with the default values initialized.
func NewGetInvoiceAsPDFParams() *GetInvoiceAsPDFParams {
	var ()
	return &GetInvoiceAsPDFParams{}
}

/*GetInvoiceAsPDFParams contains all the parameters to send to the API endpoint
for the get invoice as p d f operation typically these are written to a http.Request
*/
type GetInvoiceAsPDFParams struct {

	/*ID
	  The ID of the invoice.

	*/
	ID string
	/*Organizations
	  A list of organization-IDs used to restrict the scope of API calls.

	*/
	Organizations []string
	/*TierBreakdown
	  Whether to provide a breakdown of charge tiering.

	*/
	TierBreakdown *bool
}

// WithID adds the id to the get invoice as p d f params
func (o *GetInvoiceAsPDFParams) WithID(id string) *GetInvoiceAsPDFParams {
	o.ID = id
	return o
}

// WithOrganizations adds the organizations to the get invoice as p d f params
func (o *GetInvoiceAsPDFParams) WithOrganizations(organizations []string) *GetInvoiceAsPDFParams {
	o.Organizations = organizations
	return o
}

// WithTierBreakdown adds the tierBreakdown to the get invoice as p d f params
func (o *GetInvoiceAsPDFParams) WithTierBreakdown(tierBreakdown *bool) *GetInvoiceAsPDFParams {
	o.TierBreakdown = tierBreakdown
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceAsPDFParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param ID
	if err := r.SetPathParam("ID", o.ID); err != nil {
		return err
	}

	valuesOrganizations := o.Organizations

	joinedOrganizations := swag.JoinByFormat(valuesOrganizations, "multi")
	// query array param organizations
	if err := r.SetQueryParam("organizations", joinedOrganizations...); err != nil {
		return err
	}

	if o.TierBreakdown != nil {

		// query param tier_breakdown
		var qrTierBreakdown bool
		if o.TierBreakdown != nil {
			qrTierBreakdown = *o.TierBreakdown
		}
		qTierBreakdown := swag.FormatBool(qrTierBreakdown)
		if qTierBreakdown != "" {
			if err := r.SetQueryParam("tier_breakdown", qTierBreakdown); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
