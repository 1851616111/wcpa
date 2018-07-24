package api

import (
	"fmt"
	"testing"
)

func TestRequestAccessToken(t *testing.T) {
	accToken, err := RequestAccessToken("wxd09c7682905819e6", "b9938ddfec045280eba89fab597a0c41")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", *accToken)
}
