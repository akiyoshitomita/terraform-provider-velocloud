package vcoclient

import (
	"time"
)

type Site struct {
	Id             int32     `json:"id,omitempty"`
	Created        time.Time `json:"created,omitempty"`
	Name           string    `json:"name,omitempty"`
	LogicalId      string    `json:"logicalId,omitempty"`
	ContactName    string    `json:"contactName,omitempty"`
	ContactPhone   string    `json:"contactPhone,omitempty"`
	ContactMobile  string    `json:"contactMobile,omitempty"`
	ContactEmail   string    `json:"contactEmail,omitempty"`
	StreetAddress  string    `json:"streetAddress,omitempty"`
	StreetAddress2 string    `json:"streetAddress2,omitempty"`
	City           string    `json:"city,omitempty"`
	State          string    `json:"state,omitempty"`
	PostalCode     string    `json:"postalCode,omitempty"`
	Country        string    `json:"country,omitempty"`
	Lat            float64   `json:"lat,omitempty"`
	Lon            float64   `json:"lon,omitempty"`
	Timezone       string    `json:"timezone,omitempty"`
	Locale         string    `json:"locale,omitempty"`
	//ShippingSameAsLocation *Tinyint  `json:"shippingSameAsLocation,omitempty"`
	ShippingSameAsLocation int       `json:"shippingSameAsLocation,omitempty"`
	ShippingContactName    string    `json:"shippingContactName,omitempty"`
	ShippingAddress        string    `json:"shippingAddress,omitempty"`
	ShippingAddress2       string    `json:"shippingAddress2,omitempty"`
	ShippingCity           string    `json:"shippingCity,omitempty"`
	ShippingState          string    `json:"shippingState,omitempty"`
	ShippingCountry        string    `json:"shippingCountry,omitempty"`
	ShippingPostalCode     string    `json:"shippingPostalCode,omitempty"`
	Modified               time.Time `json:"modified,omitempty"`
}
