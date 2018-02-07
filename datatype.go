package amazonpay

import "time"

// Address postal address information
type Address struct {
	Name                                     string
	AddressLine1, AddressLine2, AddressLine3 string
	City                                     string
	Country                                  string
	District                                 string
	StateOrRegion                            string
	PostalCode                               string
	CountryCode                              string
	Phone                                    string
}

// Buyer buyer info
type Buyer struct {
	Name  string
	Email string
	Phone string
}

// CaptureDetails the details of the sales claim object and its current state.
type CaptureDetails struct {
	AmazonCaptureID    string
	CaptureReferenceID string
	SellerCaptureNote  string
	CaptureAmount      Price
	RefundAmount       Price
	CaptureFee         Price
	IDList             string
	CreationTimestamp  *time.Time
	CaptureStatus      Status
	SoftDescriptor     string
}

// Constraint represents a mistake or error information of a Billing Agreement or Order Reference object.
type Constraint struct {
	ConstraintID string
	Description  string
}

// Destination the address selected by the buyer via the address book widget.
type Destination struct {
	DestinationType     string
	PhysicalDestination Address
}

// OrderReferenceAttributes attribute of the Order Reference object specified by the seller.
type OrderReferenceAttributes struct {
	OrderTotal            OrderTotal
	PlatformID            string
	SellerNote            string
	SellerOrderAttributes SellerOrderAttributes
}

// OrderReferenceDetails details and current state of the Order Reference object.
type OrderReferenceDetails struct {
	AmazonOrderReferenceID string
	Buyer                  Buyer
	OrderTotal             OrderTotal
	SellerNote             string
	PlatformID             string
	Destination            Destination
	ReleaseEnvironment     string
	SellerOrderAttributes  SellerOrderAttributes
	OrderReferenceStatus   OrderReferenceStatus
	Constraints            []Constraint
	CreationTimestamp      *time.Time
	ExpirationTimestamp    *time.Time
	IDList                 string
}

// OrderReferenceStatus the current state of the Order Reference object.
type OrderReferenceStatus struct {
	State               string
	LastUpdateTimestamp *time.Time
	ReasonCode          string
	ReasonDescription   string
}

// OrderTotal total order amount presented in this Order Reference.
type OrderTotal struct {
	CurrencyCode string
	Amount       string
}

// Price currency type and amount.
type Price struct {
	Amount      string
	CurrecyCode string
}

// RefundDetails details and the current state of the refund object.
type RefundDetails struct {
	AmazonRefundID    string
	RefundReferenceID string
	SellerRefundNote  string
	RefundType        string
	RefundAmount      Price
	FeeRefunded       Price
	CreationTimestamp *time.Time
	RefundStatus      Status
	SoftDescriptor    string
}

// SellerOrderAttributes provides detailed information on the Order Reference object
type SellerOrderAttributes struct {
	SellerOrderID     string
	StoreName         string
	CustomInformation string
}

// Status represents the current status of authorization object, sales request object, refund object.
type Status struct {
	State               string
	LastUpdateTimestamp *time.Time
	ReasonCode          string
	ReasonDescription   string
}
