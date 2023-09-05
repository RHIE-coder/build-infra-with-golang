package demo

type AtomicSwapContract struct {
	pointContractAddr string
	tokenContractAddr string
	swapPointContract string
	swapTokenContract string
}

func NewAtomicSwapController(
	pointContractAddr string,
	tokenContractAddr string,
	swapPointContract string,
	swapTokenContract string,
) *AtomicSwapContract {
	return &AtomicSwapContract{
		pointContractAddr: pointContractAddr,
		tokenContractAddr: tokenContractAddr,
		swapPointContract: swapPointContract,
		swapTokenContract: swapTokenContract,
	}
}

func (atomicSwap *AtomicSwapContract) Addresses() map[string]string {
	return map[string]string{
		"point_contract": atomicSwap.pointContractAddr,
		"token_contract": atomicSwap.tokenContractAddr,
		"swap_point":     atomicSwap.swapPointContract,
		"swap_token":     atomicSwap.swapTokenContract,
	}
}

type Swap struct {
	PoolInitiatorAddress string // sender address
	ReceiverAddress      string // receiver address
	SecretHash           string // keccack256 algorithm
	Amount               string // Amount = amount * decimals
}

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

type AtomicSwap interface {
	Name() (string, error)
	Symbol() (string, error)
	Decimals() (string, error)
	AddressOfContract() (string, error)
	CreateSwap(swap Swap) error
	Redeem(secret []byte, secretHash string) error
	Refund(secretHash string) error
	GetSwap(secretHash string) Swap
	GetSwapStatus(secretHash string) Stage
	IsRedeemed(secretHash string)
	IsRefunded(secretHash string)
}
