package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestVeIDFromUint64(t *testing.T) {
	id := uint64(10000)
	veID := VeIDFromUint64(id)
	require.Equal(t, "ve-10000", veID)
}

func TestUint64FromVeID(t *testing.T) {
	veID := "xxx"
	val := Uint64FromVeID(veID)
	require.Equal(t, uint64(EmptyVeID), val)

	veID = "ve-xxx"
	val = Uint64FromVeID(veID)
	require.Equal(t, uint64(EmptyVeID), val)

	veID = "ve-10000"
	val = Uint64FromVeID(veID)
	require.Equal(t, uint64(10000), val)
}

func TestVeNftUri(t *testing.T) {
	var (
		nftID     = "10000"
		balance   = sdk.NewInt(10000)
		lockedEnd = uint64(10000)
		value     = sdk.NewInt(10000)
		uri       = VeNftUri(nftID, balance, lockedEnd, value)
	)
	require.Equal(t, "data:application/json;base64,eyJuYW1lIjoibG9jayAjMTAwMDAiLCJkZXNjcmlwdGlvbiI6IkJsYWNrZnVyeSBsb2NrcywgY2FuIGJlIHVzZWQgdG8gYm9vc3QgZ2F1Z2UgeWllbGRzLCB2b3RlIG9uIHRva2VuIGVtaXNzaW9uLCBhbmQgcmVjZWl2ZSBicmliZXMiLCJpbWFnZSI6ImRhdGE6aW1hZ2Uvc3ZnK3htbDtiYXNlNjQsUEhOMlp5QjRiV3h1Y3owaWFIUjBjRG92TDNkM2R5NTNNeTV2Y21jdk1qQXdNQzl6ZG1jaUlIQnlaWE5sY25abFFYTndaV04wVW1GMGFXODlJbmhOYVc1WlRXbHVJRzFsWlhRaUlIWnBaWGRDYjNnOUlqQWdNQ0F6TlRBZ016VXdJajQ4YzNSNWJHVS1MbUpoYzJVZ2V5Qm1hV3hzT2lCM2FHbDBaVHNnWm05dWRDMW1ZVzFwYkhrNklITmxjbWxtT3lCbWIyNTBMWE5wZW1VNklERTBjSGc3SUgwOEwzTjBlV3hsUGp4eVpXTjBJSGRwWkhSb1BTSXhNREFsSWlCb1pXbG5hSFE5SWpFd01DVWlJR1pwYkd3OUltSnNZV05ySWlBdlBqeDBaWGgwSUhnOUlqRXdJaUI1UFNJeU1DSWdZMnhoYzNNOUltSmhjMlVpUG5SdmEyVnVJREV3TURBd1BDOTBaWGgwUGp4MFpYaDBJSGc5SWpFd0lpQjVQU0kwTUNJZ1kyeGhjM005SW1KaGMyVWlQbUpoYkdGdVkyVlBaaUF4TURBd01Ed3ZkR1Y0ZEQ0OGRHVjRkQ0I0UFNJeE1DSWdlVDBpTmpBaUlHTnNZWE56UFNKaVlYTmxJajVzYjJOclpXUmZaVzVrSURFd01EQXdQQzkwWlhoMFBqeDBaWGgwSUhnOUlqRXdJaUI1UFNJNE1DSWdZMnhoYzNNOUltSmhjMlVpUG5aaGJIVmxJREV3TURBd1BDOTBaWGgwUGp3dmMzWm5QZz09In0=", uri)
}
