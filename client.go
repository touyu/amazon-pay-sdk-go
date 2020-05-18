package amazonpay

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetProfile get user profile
func (amazonPay *AmazonPay) GetProfile(token string) (profile Profile, err error) {
	fmt.Println(amazonPay.OAuthEndpoint + "/user/profile")
	req, _ := http.NewRequest("GET", amazonPay.OAuthEndpoint+"/user/profile", nil)
	req.Header.Add("Authorization", "bearer "+token)
	resp, err := http.DefaultClient.Do(req)

	if err == nil {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)

		respErr := OAuthResponseError{}
		if err = json.Unmarshal(contents, &respErr); err == nil {
			if respErr.Error != "" {
				return profile, errors.New(respErr.ErrorDescription)
			}
		}

		if err == nil {
			err = json.Unmarshal(contents, &profile)
		}
		return profile, err
	}
	return profile, err
}

// SetOrderReferenceDetails set order details such as order amount and explanation in Order Reference
func (amazonPay *AmazonPay) SetOrderReferenceDetails(orderReferenceID string, attrs OrderReferenceAttributes) (result SetOrderReferenceDetailsResult, err error) {
	var params = Params{
		"Action":                 "SetOrderReferenceDetails",
		"AmazonOrderReferenceId": orderReferenceID,
	}

	err = updateParams(&params, "OrderReferenceAttributes", attrs)

	if err == nil {
		err = amazonPay.Post(params, &result)
	}

	return result, err
}

// ConfirmOrderReference confirm order details
func (amazonPay *AmazonPay) ConfirmOrderReference(orderReferenceID string) error {
	var params = Params{
		"Action":                 "ConfirmOrderReference",
		"AmazonOrderReferenceId": orderReferenceID,
	}

	return amazonPay.Post(params, nil)
}

// GetOrderReferenceDetails Returns the details and current state of the Order Reference object.
func (amazonPay *AmazonPay) GetOrderReferenceDetails(orderReferenceID string, addressToken string) (result GetOrderReferenceDetailsResponse, err error) {
	var params = Params{
		"Action":                 "GetOrderReferenceDetails",
		"AmazonOrderReferenceId": orderReferenceID,
		"AddressToken":           addressToken,
	}

	err = amazonPay.Post(params, &result)

	return result, err
}

// AuthorizeInput authorize input struct
type AuthorizeInput struct {
	SellerAuthorizationNote string
	TransactionTimedOut     uint
	CaptureNow              bool
	SoftDecriptor           string
}

// Authorize process secures the funds specified for the payment method stored in the Order Reference.
func (amazonPay *AmazonPay) Authorize(orderReferenceID string, authorizationReferenceID string, amount Price, input AuthorizeInput) (result AuthorizeResponse, err error) {
	var params = Params{
		"Action":                   "Authorize",
		"AmazonOrderReferenceId":   orderReferenceID,
		"AuthorizationReferenceId": authorizationReferenceID,
	}

	updateParams(&params, "AuthorizationAmount", amount)
	updateParams(&params, "", input)

	err = amazonPay.Post(params, &result)

	return result, err
}

// GetAuthorizationDetails returns the total authorized amount for authorization status and authorization.
func (amazonPay *AmazonPay) GetAuthorizationDetails(authorizationID string) (result GetAuthorizationDetailsResponse, err error) {
	var params = Params{
		"Action":                "GetAuthorizationDetails",
		"AmazonAuthorizationId": authorizationID,
	}

	err = amazonPay.Post(params, &result)

	return result, err
}

// CloseAuthorization Close authorization
func (amazonPay *AmazonPay) CloseAuthorization(authorizationID string, closureReason string) error {
	var params = Params{
		"Action":                "CloseAuthorization",
		"AmazonAuthorizationId": authorizationID,
		"ClosureReason":         closureReason,
	}

	return amazonPay.Post(params, nil)
}

// CaptureInput capture input struct
type CaptureInput struct {
	SellerCaptureNote string
	SoftDecriptor     string
}

// Capture request funds from the authorized payment method.
func (amazonPay *AmazonPay) Capture(authorizationID string, captureReferenceID string, captureAmount Price, input CaptureInput) (result CaptureResponse, err error) {
	var params = Params{
		"Action":                "Capture",
		"AmazonAuthorizationId": authorizationID,
		"CaptureReferenceId":    captureReferenceID,
	}

	updateParams(&params, "CaptureAmount", captureAmount)
	updateParams(&params, "", input)

	err = amazonPay.Post(params, &result)

	return result, err
}

// GetCaptureDetails returns the detailed sales request status and the total amount refunded by sales request.
func (amazonPay *AmazonPay) GetCaptureDetails(captureID string) (result GetCaptureDetailsResponse, err error) {
	var params = Params{
		"Action":          "GetCaptureDetails",
		"AmazonCaptureId": captureID,
	}

	err = amazonPay.Post(params, &result)
	return result, err
}

// CloseOrderReference complete order reference and will not be able to generate a new authorization from this Order Reference.
func (amazonPay *AmazonPay) CloseOrderReference(orderReferenceID string, closureReason string) error {
	var params = Params{
		"Action":                 "CloseOrderReference",
		"AmazonOrderReferenceId": orderReferenceID,
		"ClosureReason":          closureReason,
	}

	return amazonPay.Post(params, nil)
}

// CancelOrderReference Cancels a previously confirmed order reference.
func (amazonPay *AmazonPay) CancelOrderReference(orderReferenceID string, reason string) error {
	var params = Params{
		"Action":                 "CancelOrderReference",
		"AmazonOrderReferenceId": orderReferenceID,
		"CancelationReason":      reason,
	}

	return amazonPay.Post(params, nil)
}

// RefundInput refund input struct
type RefundInput struct {
	SellerRefundNote string
	SoftDescriptor   string
}

// Refund refund the funds requested
func (amazonPay *AmazonPay) Refund(captureID string, refundReferenceID string, refundAmount Price, input RefundInput) (result RefundResponse, err error) {
	var params = Params{
		"Action":            "Refund",
		"AmazonCaptureId":   captureID,
		"RefundReferenceId": refundReferenceID,
	}

	updateParams(&params, "RefundAmount", refundAmount)
	updateParams(&params, "", input)

	err = amazonPay.Post(params, &result)

	return result, err
}

// GetRefundDetails get refund details
func (amazonPay *AmazonPay) GetRefundDetails(refundID string) (result GetRefundDetailsResponse, err error) {
	var params = Params{
		"Action":         "GetRefundDetails",
		"AmazonRefundId": refundID,
	}

	err = amazonPay.Post(params, &result)
	return result, err
}

// AuthorizeOnBillingAgreementInput authorize on billing agreement input struct
type AuthorizeOnBillingAgreementInput struct {
	SellerAuthorizationNote string
	TransactionTimeout      uint
	CaptureNow              bool
	SoftDecriptor           string
	SellerNote              string
	PlatformID              string `json:"PlatformId"`
	SellerOrderAttributes   SellerOrderAttributes
	InheritShippingAddress  bool
}

// AuthorizeOnBillingAgreement process secures the funds specified for the payment method stored in the Billing Agreement.
func (amazonPay *AmazonPay) AuthorizeOnBillingAgreement(billingAgreementID string, authorizationReferenceID string, amount Price, input AuthorizeOnBillingAgreementInput) (result AuthorizeOnBillingAgreementResponse, err error) {
	var params = Params{
		"Action":                   "AuthorizeOnBillingAgreement",
		"AmazonBillingAgreementId": billingAgreementID,
		"AuthorizationReferenceId": authorizationReferenceID,
	}

	updateParams(&params, "AuthorizationAmount", amount)
	updateParams(&params, "", input)

	err = amazonPay.Post(params, &result)

	return result, err
}

// CloseBillingAgreement Close billing agreement
func (amazonPay *AmazonPay) CloseBillingAgreement(billingAgreementID string, closureReason string) error {
	var params = Params{
		"Action":                   "CloseBillingAgreement",
		"AmazonBillingAgreementId": billingAgreementID,
		"ClosureReason":            closureReason,
	}

	return amazonPay.Post(params, nil)
}

// ConfirmBillingAgreement confirm billing agreement
func (amazonPay *AmazonPay) ConfirmBillingAgreement(billingAgreementID string) error {
	var params = Params{
		"Action":                   "ConfirmBillingAgreement",
		"AmazonBillingAgreementId": billingAgreementID,
	}

	return amazonPay.Post(params, nil)
}

// CreateOrderReferenceForIdInput create order reference for id input struct
type CreateOrderReferenceForIdInput struct {
	InheritShippingAddress   bool
	ConfirmNow               bool
	OrderReferenceAttributes OrderReferenceAttributes
}

// CreateOrderReferenceForId creates an order reference.
func (amazonPay *AmazonPay) CreateOrderReferenceForId(id string, idType string, input CreateOrderReferenceForIdInput) (result CreateOrderReferenceForIdResponse, err error) {
	var params = Params{
		"Action": "CreateOrderReferenceForId",
		"Id":     id,
		"IdType": idType,
	}

	updateParams(&params, "", input)

	err = amazonPay.Post(params, &result)

	return result, err
}

// GetBillingAgreementDetails returns the details and current state of the Billing Agreement object.
func (amazonPay *AmazonPay) GetBillingAgreementDetails(billingAgreementID string, accessToken string) (result GetBillingAgreementDetailsResponse, err error) {
	var params = Params{
		"Action":                   "GetBillingAgreementDetails",
		"AmazonBillingAgreementId": billingAgreementID,
		"AddressToken":      accessToken,
	}

	err = amazonPay.Post(params, &result)

	return result, err
}

// SetBillingAgreementDetails set billing agreement such as a description of the agreement and other information about the merchant.
func (amazonPay *AmazonPay) SetBillingAgreementDetails(billingAgreementID string, attrs BillingAgreementAttributes) (result SetBillingAgreementDetailsResponse, err error) {
	var params = Params{
		"Action":                   "SetBillingAgreementDetails",
		"AmazonBillingAgreementId": billingAgreementID,
	}

	err = updateParams(&params, "BillingAgreementAttributes", attrs)

	if err == nil {
		err = amazonPay.Post(params, &result)
	}

	return result, err
}

// ValidateBillingAgreement validates the status of the BillingAgreement object and the payment method associated with it.
func (amazonPay *AmazonPay) ValidateBillingAgreement(billingAgreementID string) (result ValidateBillingAgreementResponse, err error) {
	var params = Params{
		"Action":                   "ValidateBillingAgreement",
		"AmazonBillingAgreementId": billingAgreementID,
	}

	err = amazonPay.Post(params, &result)

	return result, err
}
