package stockchangehistorymodel

type StockChangeType string

const (
	Sell   StockChangeType = "Sell"
	Import StockChangeType = "Import"
	Export StockChangeType = "Export"
	Modify StockChangeType = "Modify"
)
