package belex

type BelexQuote struct {
	stockTicker  string
	currentQuote float64
}

// MARK: Implementation of the SimpleQuote interface
func (quote BelexQuote) Quote() float64 {
	return quote.currentQuote
}

func (quote BelexQuote) Ticker() string {
	return quote.stockTicker
}