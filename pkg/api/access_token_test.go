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

	// beta acc_token 12_Pq3iowvUwpuT_8XSzzPf9BHSmd2bjhquzxfthOwGCqB2WZpdOqT3GZLnOW1f0eTPkVVaiYJ06p2fFInS4H4XUcYo610GnRFajhP30xYYIIX1s-yXQ66IXk7K3urLMN1HSNIhAlXunhJGfMD1KLOhAEAYZC
//https://api.weixin.qq.com/cgi-bin/menu/get?access_token=12_5UEUYTZ3eahTBALK7nNQ0PeNqTmZVECa7uOoBJ1miTNT7eiEORvO3jv2fGbyhrU7PfPW43kH8sA6J-jbr2kMQIhYs_fD9oT1wHP7fwXaG7GB-ADynTG3nVobWBJ18fRn6vM8GlX4pcfrWgAaIHFgAIASRP
//https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=12_5UEUYTZ3eahTBALK7nNQ0PeNqTmZVECa7uOoBJ1miTNT7eiEORvO3jv2fGbyhrU7PfPW43kH8sA6J-jbr2kMQIhYs_fD9oT1wHP7fwXaG7GB-ADynTG3nVobWBJ18fRn6vM8GlX4pcfrWgAaIHFgAIASRP

//curl -XPOST  https://api.weixin.qq.com/cgi-bin/menu/create?access_token=12_Pq3iowvUwpuT_8XSzzPf9BHSmd2bjhquzxfthOwGCqB2WZpdOqT3GZLnOW1f0eTPkVVaiYJ06p2fFInS4H4XUcYo610GnRFajhP30xYYIIX1s-yXQ66IXk7K3urLMN1HSNIhAlXunhJGfMD1KLOhAEAYZC -d ''

//curl -XPOST -F media=@timg.jpeg -F description='{"title":"123", "introduction":"123"}' "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=12_qN8Phi2MpfudH2gDlqypzgzrcpBa986fs0Enxgp3V_3Uk1YnxYV4TX1eQbM4B3C30ec0VtcFACUNbN-kBThpvkDbRtsZlUK3YPJxMKPgvyaOYvPIRWm3xCptgpC9MB4KuZIL1rlte6G3C-udCNEjAJASDU&type=image"
//	{"media_id":"ImRr0414MiFv7qOJ3erDXtL-D7ebV3eZtWo3PZr2log","url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/1KicZ7bpfTIdHmBFLic5n2NzK6S4lE5yp6K5rU8sN83TlkR3nsKO1QWuguEWFxj9hkHf2nmAeWibCHsOlkPcuS9MQ\/0?wx_fmt=jpeg"}
fmt.Printf("%#v\n", accToken)
}
