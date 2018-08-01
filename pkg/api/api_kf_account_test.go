package api

import (
	"testing"
	"fmt"
)

func TestListAllKFAccounts(t *testing.T) {
	l, err := ListAllKFAccounts("12_v6U3E7KQTd-RuJOc86PdwDCpaI9MrovAqEdy33Hsf659aEUQPxbacWm5zoRl98CVVAdYuh2GVD_62mdC5kupB_MLYfXko3tT-oPEmRc_1gA5ThFbX7fgnAlDHZQDMNhAIAVEF")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(l)
}

func TestCreateKFAccount(t *testing.T) {
	acc := NewKFAccount{
		Account:  "",
		NickName: "test",
		Password: "123456",
	}
	err := CreateKFAccount("12_wsH_9Rl3BoC1RkIPJafcXTCc19TWy9oAAmLKmj8lyjMc3PRRYemCNsh9mvHa838Dzmn3HMN25Jmf46l2wME5RiAc8oxUFBepDFCNYriwpt3wMStJ--dC_72ZujwtXsvWZo2HHDalVIjRZHxbLDNaAIARJT", acc)
	if err != nil {
		t.Fatal(err)
	}
}
