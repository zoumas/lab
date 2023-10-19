// package bitcoin contains the Bitcoin type definition and interface implementations
package bitcoin

import "fmt"

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
