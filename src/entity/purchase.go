package entity

type Purchase struct {
	CPFCNPJ            *string
	Private            *bool
	Incomplete         *bool
	DateLastCompra     *string
	AverageTicket      *float64
	LastPurchaseTicket *float64
	StoreMoreFrequent  *string
	StoreLastPurchase  *string
}
