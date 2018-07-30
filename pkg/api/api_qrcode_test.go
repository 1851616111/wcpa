package api

import (
	"testing"
	"fmt"
)

func TestCreateLimitSceneQrCode(t *testing.T) {
	accToken, err := RequestAccessToken("wxd09c7682905819e6", "b9938ddfec045280eba89fab597a0c41")
	if err != nil {
		t.Fatal(err)
	}

	id, tk , err := CreateSceneQrCode(accToken.Token, 500)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(id)
	fmt.Printf("%+8v\n", *tk)
}
