package financialvalueobjects

type AccountType int

const (
	CASH AccountType = iota
	BANK
	ACCOUNTS_RECIVEABLE
	INVENTORY
	OTHER_CURRENT_ASSET
	FIXED_ASSET
	NON_CURRENT_ASSET
	ACCOUNTS_PAYABLE
	CREDIT_CARD
	TAX_PAYABLE
	OTHER_CURRENT_LIABILITY
	LONG_TERM_LIABILITY
	NON_CURRENT_LIABILITY
	EQUITY
	INCOME
	OTHER_INCOME
	COSTS_OF_GOODS_SOLD
	EXPENSE
	OTHER_EXPENSE
)
