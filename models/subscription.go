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

/*Subscription Subscription

swagger:model Subscription
*/
type Subscription struct {

	/* {"description":"","verbs":["POST","GET"]}

	Required: true
	*/
	AccountID string `json:"accountID,omitempty"`

	/* {  "default":"true", "description":"Whether the subscription will aggregate all other subscriptions on the account.","verbs":["GET", "PUT", "POST"]}
	 */
	AggregateAllSubscriptionsOnAccount *bool `json:"aggregateAllSubscriptionsOnAccount,omitempty"`

	/* BillingEntity billing entity
	 */
	BillingEntity *string `json:"billingEntity,omitempty"`

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy *string `json:"changedBy,omitempty"`

	/* {"description":"start of the contracted period.  This will be after a trial, if one exists","verbs":["GET"]}
	 */
	ContractStart strfmt.DateTime `json:"contractStart,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* {  "default":"true", "description":"Can credit-notes be used to pay outstanding invoices for this subscription.","verbs":["GET", "PUT", "POST"]}
	 */
	CreditEnabled *bool `json:"creditEnabled,omitempty"`

	/* Crmid crmid
	 */
	Crmid *string `json:"crmid,omitempty"`

	/* CurrentFixedTerm current fixed term
	 */
	CurrentFixedTerm *FixedTerm `json:"currentFixedTerm,omitempty"`

	/* {"description":"End of the current period invoiced for. This can be manually updated to extend trials or delay invoice generation.","verbs":["PUT","GET"]}
	 */
	CurrentPeriodEnd strfmt.DateTime `json:"currentPeriodEnd,omitempty"`

	/* {"description":"Override for the initial subscription period. Allows periods to align to a date or time regardless of purchase date/time.","verbs":["POST","PUT","GET"]}
	 */
	CurrentPeriodEndExplicit strfmt.DateTime `json:"currentPeriodEndExplicit,omitempty"`

	/* {"description":"Start of the current invoice period. At the end of this period, a new new invoice will be generated","verbs":["POST","GET"]}
	 */
	CurrentPeriodStart strfmt.DateTime `json:"currentPeriodStart,omitempty"`

	/* {"description":"The current time &mdash; from the point of view of the subscription.","verbs":["GET"]}

	Required: true
	*/
	CurrentTime strfmt.DateTime `json:"currentTime,omitempty"`

	/* DaysElapsed days elapsed
	 */
	DaysElapsed *Days `json:"daysElapsed,omitempty"`

	/* DaysRemaining days remaining
	 */
	DaysRemaining *Days `json:"daysRemaining,omitempty"`

	/* {"description":"","verbs":["POST","PUT","GET"]}
	 */
	Description *string `json:"description,omitempty"`

	/* {  "default":"false", "description":"Are there outstanding invoices which are currently in dunning.","verbs":["GET", "PUT", "POST"]}
	 */
	Dunning *bool `json:"dunning,omitempty"`

	/* Fields fields
	 */
	Fields map[string]string `json:"fields,omitempty"`

	/* FixedTerm fixed term
	 */
	FixedTerm *bool `json:"fixedTerm,omitempty"`

	/* List of fixed terms that have been or are applied to the subscription

	Required: true
	*/
	FixedTerms []*FixedTerm `json:"fixedTerms,omitempty"`

	/* {"description":"","verbs":["GET"]}

	Required: true
	*/
	ID string `json:"id,omitempty"`

	/* {"description":"Start of the first successful period","verbs":["GET"]}

	Required: true
	*/
	InitialPeriodStart strfmt.DateTime `json:"initialPeriodStart,omitempty"`

	/* {"description":"If the subscription is locked, it will not be processed by the system","verbs":[]}
	 */
	Locked *string `json:"locked,omitempty"`

	/* {"description":"Which system is responsible for managing the subscription.","verbs":[]}
	 */
	ManagedBy *string `json:"managedBy,omitempty"`

	/* MinutesElapsed minutes elapsed
	 */
	MinutesElapsed *Minutes `json:"minutesElapsed,omitempty"`

	/* MinutesRemaining minutes remaining
	 */
	MinutesRemaining *Minutes `json:"minutesRemaining,omitempty"`

	/* {"description":"User definable friendly name for the subscription.","verbs":["POST","PUT","GET"]}

	Required: true
	*/
	Name string `json:"name,omitempty"`

	/* NewSubscription new subscription
	 */
	NewSubscription *bool `json:"newSubscription,omitempty"`

	/* NotificationObjectGraph notification object graph
	 */
	NotificationObjectGraph *string `json:"notificationObjectGraph,omitempty"`

	/* {"description":"Organization associated with the subscription.","verbs":[]}

	Required: true
	*/
	OrganizationID string `json:"organizationID,omitempty"`

	/* {"description":"","verbs":["GET"]}
	 */
	ParentID *string `json:"parentID,omitempty"`

	/* PaymentMethodSubscriptionLinks payment method subscription links
	 */
	PaymentMethodSubscriptionLinks []*PaymentMethodSubscriptionLink `json:"paymentMethodSubscriptionLinks,omitempty"`

	/* {"description":"The quantities for each pricing component of the rate-plan. Values should be set for all pricing components of the rate-plan apart from the usage components. Usage components should be added when the usage is known, this is often after the end of the current billing cycle.","verbs":["GET", "POST"]}
	 */
	PricingComponentValues []*PricingComponentValue `json:"pricingComponentValues,omitempty"`

	/* {"description":"","verbs":["GET"]}

	Required: true
	*/
	ProductID string `json:"productID,omitempty"`

	/* {"description":"Identifier of the rate-plan being billed for.","verbs":["POST","PUT","GET"]}

	Required: true
	*/
	ProductRatePlanID string `json:"productRatePlanID,omitempty"`

	/* SecondsRemaining seconds remaining
	 */
	SecondsRemaining *Seconds `json:"secondsRemaining,omitempty"`

	/* ShortID short ID
	 */
	ShortID *string `json:"shortID,omitempty"`

	/* {"PUT_description":"A <span class=\"label label-default\">Provisioned</span> subscription can be updated to either <span class=\"label label-default\">Trial</span> or <span class=\"label label-default\">AwaitingPayment</span>, this will start the subscription. Any updates to the state of a non-<span class=\"label label-default\">Provisioned</span> will be ignored. To cancel or otherwise amend a subscription please use the explict amendment calls.", "description":"A <span class=\"label label-default\">Provisioned</span> subscription will not begin until marked as <span class=\"label label-default\">Trial</span> or <span class=\"label label-default\">AwaitingPayment</span>. Trial subscriptions transition to <span class=\"label label-default\">AwaitingPayment</span> when the trial period is over. On subscription renewal the state becomes <span class=\"label label-default\">AwaitingPayment</span>. Once outstanding invoices are paid the state changes to <span class=\"label label-default\">Paid</span>. A subscription is set as either <span class=\"label label-default\">Failed</span> or left as <span class=\"label label-default\">AwaitingPayment</span>, depending on the rate-plan configuration. If a subscription is non-recurring or fixed-term and ends naturally, it will be marked as <span class=\"label label-default\">Expired</span>. If all payment attempts have failed a subscription is marked as <span class=\"label label-default\">Cancelled</span> if it has been manually ended. Once a subscription is marked as <span class=\"label label-default\">Failed</span>, <span class=\"label label-default\">Expired</span>, or <span class=\"label label-default\">Cancelled</span> no invoices other than a final invoice will be issued. Note: Updating account card details will not lead to BillForward automatically retrying payment, manual payment attempts can be made.","verbs":["POST","PUT","GET"]}

	Required: true
	*/
	State string `json:"state,omitempty"`

	/* {"description":"When a subscription will end. This may be in the future if the cancellation is at the end of the current period.","verbs":["GET"]}
	 */
	SubscriptionEnd strfmt.DateTime `json:"subscriptionEnd,omitempty"`

	/* {"description":"Number of paid-for periods billing, excluding trials, since the subscription started.","verbs":["GET"]}
	 */
	SuccessfulPeriods int32 `json:"successfulPeriods,omitempty"`

	/* {"description":"How far in the future is the entity (in seconds) compared to the BillForward server's time.","verbs":["GET"]}
	 */
	TimeOffset int64 `json:"timeOffset,omitempty"`

	/* {"description":"Total number of subscription periods.","verbs":["GET"]}
	 */
	TotalPeriods int32 `json:"totalPeriods,omitempty"`

	/* {"default":"dependent on product", "description":"","verbs":["POST","GET"]}
	 */
	Type *string `json:"type,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated. ", "verbs":[] }
	 */
	Updated strfmt.DateTime `json:"updated,omitempty"`

	/* {"description":"When the current version of the subscription ended, null indicates current version.","verbs":["GET"]}
	 */
	VersionEnd strfmt.DateTime `json:"versionEnd,omitempty"`

	/* {"description":"","verbs":["GET"]}
	 */
	VersionID *string `json:"versionID,omitempty"`

	/* {"description":"Incremental version number of the subscription, starts at 1.","verbs":["GET"]}

	Required: true
	*/
	VersionNumber int32 `json:"versionNumber,omitempty"`

	/* {"description":"When the current version of the subscription started.","verbs":["GET"]}

	Required: true
	*/
	VersionStart strfmt.DateTime `json:"versionStart,omitempty"`
}

// Validate validates this subscription
func (m *Subscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBillingEntity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrentTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFields(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFixedTerms(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInitialPeriodStart(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateManagedBy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentMethodSubscriptionLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePricingComponentValues(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProductID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProductRatePlanID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersionNumber(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersionStart(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subscription) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountID", "body", string(m.AccountID)); err != nil {
		return err
	}

	return nil
}

var subscriptionBillingEntityEnum []interface{}

func (m *Subscription) validateBillingEntityEnum(path, location string, value string) error {
	if subscriptionBillingEntityEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Notification","Organization","OrganizationGateway","Product","User","Subscription","Profile","ProductRatePlan","Client","Invoice","PricingComponentValue","Account","PricingComponentValueChange","PricingComponentTier","PricingComponent","PricingCalculation","CouponDefinition","CouponInstance","CouponModifier","CouponRule","CouponBookDefinition","CouponBook","InvoiceLine","Webhook","SubscriptionCancellation","NotificationSnapshot","InvoicePayment","InvoiceLinePayment","Payment","PaymentMethod","PaymentMethodSubscriptionLink","DunningLine","CybersourceToken","Card","Alias","PaypalSimplePaymentReconciliation","FreePaymentReconciliation","LocustworldPaymentReconciliation","CouponInstanceExistingValue","TaxLine","TaxationStrategy","TaxationLink","Address","AmendmentPriceNTime","Authority","UnitOfMeasure","SearchResult","Amendment","AuditLog","Password","Username","FixedTermDefinition","FixedTerm","Refund","CreditNote","Receipt","AmendmentCompoundConstituent","APIConfiguration","StripeToken","BraintreeToken","BalancedToken","PaypalToken","AuthorizeNetToken","SpreedlyToken","GatewayRevenue","AmendmentDiscardAmendment","CancellationAmendment","CompoundAmendment","CompoundAmendmentConstituent","FixedTermExpiryAmendment","InvoiceNextExecutionAttemptAmendment","PricingComponentValueAmendment","BraintreeMerchantAccount","WebhookSubscription","Migration","CassResult","CassPaymentResult","CassProductRatePlanResult","CassChurnResult","CassUpgradeResult","SubscriptionCharge","CassPaymentPProductResult","ProductPaymentsArgs","StripeACHToken","UsageAmount","UsageSession","Usage","UsagePeriod","Period","OfflinePayment","CreditNotePayment","CardVaultPayment","FreePayment","BraintreePayment","BalancedPayment","CybersourcePayment","PaypalPayment","PaypalSimplePayment","LocustWorldPayment","StripeOnlyPayment","ProductPaymentsResult","StripeACHPayment","AuthorizeNetPayment","CompoundUsageSession","CompoundUsage","UsageRoundingStrategies","BillforwardManagedPaymentsResult","PricingComponentValueMigrationChargeAmendmentMapping","SubscriptionLTVResult","AccountLTVResult","ProductRatePlanPaymentsResult","DebtsResult","AccountPaymentsResult","ComponentChange","QuoteRequest","Quote","CouponCharge","CouponInstanceInvoiceLink","Coupon","CouponDiscount","CouponUniqueCodesRequest","CouponUniqueCodesResponse","GetCouponsResponse","AddCouponCodeRequest","AddCouponCodeResponse","RemoveCouponFromSubscriptionRequest","TokenizationPreAuth","StripeTokenizationPreAuth","BraintreeTokenizationPreAuth","SpreedlyTokenizationPreAuth","SagePayTokenizationPreAuth","PayVisionTokenizationPreAuth","TokenizationPreAuthRequest","AuthCaptureRequest","StripeACHBankAccountVerification","PasswordReset","PricingRequest","AddTaxationStrategyRequest","AddPaymentMethodRequest","APIRequest","SagePayToken","SagePayNotificationRequest","SagePayNotificationResponse","SagePayOutstandingTransaction","SagePayEnabledCardType","TrustCommerceToken","SagePayTransaction","PricingComponentValueResponse","MigrationResponse","TimeResponse","EntityTime","Email","AggregationLink","BFPermission","Role","PermissionLink","PayVisionToken","PayVisionTransaction","KashToken","EmailProvider","DataSynchronizationJob","DataSynchronizationJobError","DataSynchronizationConfiguration","DataSynchronizationAppConfiguration","AggregationChildrenResponse","MetadataKeyValue","Metadata","AggregatingComponent","PricingComponentMigrationValue","InvoiceRecalculationAmendment","IssueInvoiceAmendment","EmailSubscription","RevenueAttribution"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			subscriptionBillingEntityEnum = append(subscriptionBillingEntityEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, subscriptionBillingEntityEnum); err != nil {
		return err
	}
	return nil
}

func (m *Subscription) validateBillingEntity(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingEntity) { // not required
		return nil
	}

	if err := m.validateBillingEntityEnum("billingEntity", "body", *m.BillingEntity); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateCurrentTime(formats strfmt.Registry) error {

	if err := validate.Required("currentTime", "body", strfmt.DateTime(m.CurrentTime)); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	if err := validate.Required("fields", "body", m.Fields); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateFixedTerms(formats strfmt.Registry) error {

	if err := validate.Required("fixedTerms", "body", m.FixedTerms); err != nil {
		return err
	}

	for i := 0; i < len(m.FixedTerms); i++ {

		if m.FixedTerms[i] != nil {

			if err := m.FixedTerms[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Subscription) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateInitialPeriodStart(formats strfmt.Registry) error {

	if err := validate.Required("initialPeriodStart", "body", strfmt.DateTime(m.InitialPeriodStart)); err != nil {
		return err
	}

	return nil
}

var subscriptionManagedByEnum []interface{}

func (m *Subscription) validateManagedByEnum(path, location string, value string) error {
	if subscriptionManagedByEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["BillForward","Stripe"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			subscriptionManagedByEnum = append(subscriptionManagedByEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, subscriptionManagedByEnum); err != nil {
		return err
	}
	return nil
}

func (m *Subscription) validateManagedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ManagedBy) { // not required
		return nil
	}

	if err := m.validateManagedByEnum("managedBy", "body", *m.ManagedBy); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationID", "body", string(m.OrganizationID)); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validatePaymentMethodSubscriptionLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMethodSubscriptionLinks) { // not required
		return nil
	}

	for i := 0; i < len(m.PaymentMethodSubscriptionLinks); i++ {

		if m.PaymentMethodSubscriptionLinks[i] != nil {

			if err := m.PaymentMethodSubscriptionLinks[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Subscription) validatePricingComponentValues(formats strfmt.Registry) error {

	if swag.IsZero(m.PricingComponentValues) { // not required
		return nil
	}

	for i := 0; i < len(m.PricingComponentValues); i++ {

		if m.PricingComponentValues[i] != nil {

			if err := m.PricingComponentValues[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Subscription) validateProductID(formats strfmt.Registry) error {

	if err := validate.Required("productID", "body", string(m.ProductID)); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateProductRatePlanID(formats strfmt.Registry) error {

	if err := validate.Required("productRatePlanID", "body", string(m.ProductRatePlanID)); err != nil {
		return err
	}

	return nil
}

var subscriptionStateEnum []interface{}

func (m *Subscription) validateStateEnum(path, location string, value string) error {
	if subscriptionStateEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Trial","Provisioned","Paid","AwaitingPayment","Cancelled","Failed","Expired"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			subscriptionStateEnum = append(subscriptionStateEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, subscriptionStateEnum); err != nil {
		return err
	}
	return nil
}

func (m *Subscription) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", string(m.State)); err != nil {
		return err
	}

	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var subscriptionTypeEnum []interface{}

func (m *Subscription) validateTypeEnum(path, location string, value string) error {
	if subscriptionTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Subscription","FixedTerm","Trial"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			subscriptionTypeEnum = append(subscriptionTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, subscriptionTypeEnum); err != nil {
		return err
	}
	return nil
}

func (m *Subscription) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateVersionNumber(formats strfmt.Registry) error {

	if err := validate.Required("versionNumber", "body", int32(m.VersionNumber)); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateVersionStart(formats strfmt.Registry) error {

	if err := validate.Required("versionStart", "body", strfmt.DateTime(m.VersionStart)); err != nil {
		return err
	}

	return nil
}
