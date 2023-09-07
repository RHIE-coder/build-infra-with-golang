package demo

type ERC20 interface {
	Name() (string, error)
	GetName() string
	Symbol() (string, error)
	GetSymbol() string
	Decimals() (string, error)
	GetDecimals() string
	Approve(string, string) (bool, error)
	Allowance(string, string) (string, error)
	ContractAddress() string
}
