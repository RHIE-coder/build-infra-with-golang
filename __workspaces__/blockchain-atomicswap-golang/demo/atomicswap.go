package demo

type AtomicSwap interface {
	Name(SwapTargetERC20) (string, error)
	Symbol(SwapTargetERC20) (string, error)
	Decimals(SwapTargetERC20) (string, error)
	AddressOfContract() (string, error)
	CreateSwap(swap Swap) error
	Redeem(secret []byte, secretHash string) error
	Refund(secretHash string) error
	GetSwap(secretHash string) Swap
	GetSwapStatus(secretHash string) Stage
	IsRedeemed(secretHash string)
	IsRefunded(secretHash string)
}

/* Swap */
type Swap struct {
	PoolInitiatorAddress string // sender address
	ReceiverAddress      string // receiver address
	SecretHash           string // keccack256 algorithm
	Amount               string // Amount = amount * decimals
}

/* enum Stage */
type Stage int64

const (
	INVALID Stage = iota
	PENDING
	COMPLETED
	CANCELED
)

func (s Stage) String() string {
	switch s {

	case INVALID:
		return "invalid"
	case PENDING:
		return "pending"
	case COMPLETED:
		return "completed"
	case CANCELED:
		return "canceled"
	}

	return "unknown"
}

/* Dispatcher for Atomic Swap */
type SwapTargetERC20 int64

const (
	POINT SwapTargetERC20 = iota
	TOKEN
)
