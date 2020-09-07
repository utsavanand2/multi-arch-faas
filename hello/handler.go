package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("Hello from Utsav's OpenFaas function running on Raspberry Pi ARM")
}
