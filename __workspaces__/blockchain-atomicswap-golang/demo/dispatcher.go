package demo

type SwapDispatcher struct {
	PointERC20 ERC20
	TokenERC20 ERC20

	PointSwap AtomicSwap
	TokenSwap AtomicSwap
}
