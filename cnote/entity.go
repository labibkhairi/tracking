package cnote

import (
	"time"
)

type Cnote struct {
	CnoteNo           string
	Cdate             time.Time
	CnoteOrigin       string
	CnoteDestination  string
	CnoteServicesCode string
	ServiceType       string
	CnoteCustNo       string
	CnoteDate         time.Time
	CnotePodReceiver  string
	CnoteReceiverName string
	CityName          string
	CnotePodDate      string
	PodStatus         string
	CustType          string
	CnoteAmount       int
	CnoteWeight       int
	PodCode           string
	Keterangan        string
	CnoteGoodsDescr   string
	FreightCharge     int
	ShippingCost      int
	InsuranceAmount   int
}
