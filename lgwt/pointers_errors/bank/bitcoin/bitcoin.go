// package bitcoin contains the Bitcoin type definition and interface implementations
package bitcoin

import "fmt"

// Bitcoin represents the Bitcoin cryptocurrency
type Bitcoin uint

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
