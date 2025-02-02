package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*Profile Profile

swagger:model Profile
*/
type Profile struct {

	/* Account account
	 */
	Account *Account `json:"account,omitempty"`

	/* { "description" : "", "verbs":["GET"] }

	Required: true
	*/
	AccountID string `json:"accountID,omitempty"`

	/* { "description" : "Any additional information", "verbs":["POST","PUT","GET"] }
	 */
	AdditionalInformation *string `json:"additionalInformation,omitempty"`

	/* { "description" : "Address associated with the profile", "verbs":["POST","PUT","GET"] }
	 */
	Addresses []*Address `json:"addresses,omitempty"`

	/* BillingEntity billing entity
	 */
	BillingEntity *string `json:"billingEntity,omitempty"`

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy *string `json:"changedBy,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	CompanyName *string `json:"companyName,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* { "description" : "Date of birth in YYYY-MM-DD format", "verbs":["POST","PUT","GET"] }
	 */
	Dob strfmt.DateTime `json:"dob,omitempty"`

	/* { "description" : "E-mail address", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Email string `json:"email,omitempty"`

	/* { "description" : "Fax number", "verbs":["POST","PUT","GET"] }
	 */
	Fax *string `json:"fax,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	FirstName string `json:"firstName,omitempty"`

	/* { "description" : "ID of the profile.", "verbs":["POST","PUT","GET"] }
	 */
	ID *string `json:"id,omitempty"`

	/* { "description" : "Home telephone number", "verbs":["POST","PUT","GET"] }
	 */
	Landline *string `json:"landline,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	LastName string `json:"lastName,omitempty"`

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	LogoURL *string `json:"logoURL,omitempty"`

	/* { "description" : "Mobile telephone number", "verbs":["POST","PUT","GET"] }
	 */
	Mobile *string `json:"mobile,omitempty"`

	/* NotificationObjectGraph notification object graph
	 */
	NotificationObjectGraph *string `json:"notificationObjectGraph,omitempty"`

	/* { "description" : "", "verbs":[] }

	Required: true
	*/
	OrganizationID string `json:"organizationID,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated. ", "verbs":[] }
	 */
	Updated strfmt.DateTime `json:"updated,omitempty"`

	/* { "description" : "VAT number", "verbs":["POST","PUT","GET"] }
	 */
	VatNumber *string `json:"vatNumber,omitempty"`
}

// Validate validates this profile
func (m *Profile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAddresses(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBillingEntity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Profile) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountID", "body", string(m.AccountID)); err != nil {
		return err
	}

	return nil
}

func (m *Profile) validateAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.Addresses); i++ {

		if m.Addresses[i] != nil {

			if err := m.Addresses[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

var profileBillingEntityEnum []interface{}

func (m *Profile) validateBillingEntityEnum(path, location string, value string) error {
	if profileBillingEntityEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Notification","Organization","OrganizationGateway","Product","User","Subscription","Profile","ProductRatePlan","Client","Invoice","PricingComponentValue","Account","PricingComponentValueChange","PricingComponentTier","PricingComponent","PricingCalculation","CouponDefinition","CouponInstance","CouponModifier","CouponRule","CouponBookDefinition","CouponBook","InvoiceLine","Webhook","SubscriptionCancellation","NotificationSnapshot","InvoicePayment","InvoiceLinePayment","Payment","PaymentMethod","PaymentMethodSubscriptionLink","DunningLine","CybersourceToken","Card","Alias","PaypalSimplePaymentReconciliation","FreePaymentReconciliation","LocustworldPaymentReconciliation","CouponInstanceExistingValue","TaxLine","TaxationStrategy","TaxationLink","Address","AmendmentPriceNTime","Authority","UnitOfMeasure","SearchResult","Amendment","AuditLog","Password","Username","FixedTermDefinition","FixedTerm","Refund","CreditNote","Receipt","AmendmentCompoundConstituent","APIConfiguration","StripeToken","BraintreeToken","BalancedToken","PaypalToken","AuthorizeNetToken","SpreedlyToken","GatewayRevenue","AmendmentDiscardAmendment","CancellationAmendment","CompoundAmendment","CompoundAmendmentConstituent","FixedTermExpiryAmendment","InvoiceNextExecutionAttemptAmendment","PricingComponentValueAmendment","BraintreeMerchantAccount","WebhookSubscription","Migration","CassResult","CassPaymentResult","CassProductRatePlanResult","CassChurnResult","CassUpgradeResult","SubscriptionCharge","CassPaymentPProductResult","ProductPaymentsArgs","StripeACHToken","UsageAmount","UsageSession","Usage","UsagePeriod","Period","OfflinePayment","CreditNotePayment","CardVaultPayment","FreePayment","BraintreePayment","BalancedPayment","CybersourcePayment","PaypalPayment","PaypalSimplePayment","LocustWorldPayment","StripeOnlyPayment","ProductPaymentsResult","StripeACHPayment","AuthorizeNetPayment","CompoundUsageSession","CompoundUsage","UsageRoundingStrategies","BillforwardManagedPaymentsResult","PricingComponentValueMigrationChargeAmendmentMapping","SubscriptionLTVResult","AccountLTVResult","ProductRatePlanPaymentsResult","DebtsResult","AccountPaymentsResult","ComponentChange","QuoteRequest","Quote","CouponCharge","CouponInstanceInvoiceLink","Coupon","CouponDiscount","CouponUniqueCodesRequest","CouponUniqueCodesResponse","GetCouponsResponse","AddCouponCodeRequest","AddCouponCodeResponse","RemoveCouponFromSubscriptionRequest","TokenizationPreAuth","StripeTokenizationPreAuth","BraintreeTokenizationPreAuth","SpreedlyTokenizationPreAuth","SagePayTokenizationPreAuth","PayVisionTokenizationPreAuth","TokenizationPreAuthRequest","AuthCaptureRequest","StripeACHBankAccountVerification","PasswordReset","PricingRequest","AddTaxationStrategyRequest","AddPaymentMethodRequest","APIRequest","SagePayToken","SagePayNotificationRequest","SagePayNotificationResponse","SagePayOutstandingTransaction","SagePayEnabledCardType","TrustCommerceToken","SagePayTransaction","PricingComponentValueResponse","MigrationResponse","TimeResponse","EntityTime","Email","AggregationLink","BFPermission","Role","PermissionLink","PayVisionToken","PayVisionTransaction","KashToken","EmailProvider","DataSynchronizationJob","DataSynchronizationJobError","DataSynchronizationConfiguration","DataSynchronizationAppConfiguration","AggregationChildrenResponse","MetadataKeyValue","Metadata","AggregatingComponent","PricingComponentMigrationValue","InvoiceRecalculationAmendment","IssueInvoiceAmendment","EmailSubscription","RevenueAttribution"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			profileBillingEntityEnum = append(profileBillingEntityEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, profileBillingEntityEnum); err != nil {
		return err
	}
	return nil
}

func (m *Profile) validateBillingEntity(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingEntity) { // not required
		return nil
	}

	if err := m.validateBillingEntityEnum("billingEntity", "body", *m.BillingEntity); err != nil {
		return err
	}

	return nil
}

func (m *Profile) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", string(m.Email)); err != nil {
		return err
	}

	return nil
}

func (m *Profile) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", string(m.FirstName)); err != nil {
		return err
	}

	return nil
}

func (m *Profile) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", string(m.LastName)); err != nil {
		return err
	}

	return nil
}

func (m *Profile) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationID", "body", string(m.OrganizationID)); err != nil {
		return err
	}

	return nil
}
