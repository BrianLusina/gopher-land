package wallet

type Ledger struct {
	entries []ledgerEntry
}

type ledgerEntry struct {
	accountId string
	txnType   string
	amount    int
}

func (l *Ledger) makeEntry(accountId, txnType string, amount int) {
	entry := ledgerEntry{accountId: accountId, txnType: txnType, amount: amount}
	l.entries = append(l.entries, entry)
	return
}
