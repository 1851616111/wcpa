package api

import (
	"fmt"
	"testing"
)

func TestGetUser(t *testing.T) {
	accToken, err := RequestAccessToken("wxd09c7682905819e6", "b9938ddfec045280eba89fab597a0c41")
	if err != nil {
		t.Fatal(err)
	}

	user, err := GetUser(accToken.Token, "oH4HtwGsY-0JSjhNhJLA7jYYOMsQ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", *user)

	user, err = GetUser(accToken.Token, "oH4HtwGsY-0JSjhNhJLA7jYYOMsQxxx")
	if err == nil {
		t.Fatal()
	}
	if user != nil {
		t.Fatal()
	}
}
