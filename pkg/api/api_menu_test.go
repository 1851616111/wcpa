package api

import (
	"testing"
	"fmt"
)

func TestCreateMenu(t *testing.T) {
	accToken, err := RequestAccessToken("wxe2d92554e152f3bf", "4628c756dc7736e9e5ebfdc04667878e")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(accToken.Token)
	subViewM1 := map[string]string{
		"儿歌部落":  "https://mp.weixin.qq.com/mp/homepage?__biz=MzU4ODMyMTExNA==&hid=2&sn=17a2bf6adaa552c7ee1d6064b65f4a4c&scene=18",
		"绘本峡谷":  "https://mp.weixin.qq.com/mp/homepage?__biz=MzU4ODMyMTExNA==&hid=1&sn=27db1edb85f0d9767bafb4c946b2946a&scene=18",
		"动画港湾":  "https://mp.weixin.qq.com/s/50EHCRqW39WTM1y1Wgo87g",
		"我不是鸡汤": "https://mp.weixin.qq.com/mp/homepage?__biz=MzU4ODMyMTExNA==&hid=4&sn=988cf27a483461b5e7b28da2642c2c55&scene=18",
	}

	bt1 := NewTopButton("Eva星球")
	for name, url := range subViewM1 {
		bt1.AddSub(NewViewButton(name, url))
	}
	subViewM2 := map[string]string{
		"8天免费训练营": "https://dznwyy.51baibao.com/wx/?MTAwMDQ3NQ==",
		"成长时光精品课": "https://dznwyy.51baibao.com/wx/?MTAwMDQ0Ng==",
		"会员免费课":   "https://dznwyy.51baibao.com/wx/?MTAwMDQ3MQ==",
	}

	bt2 := NewTopButton("见面礼")
	for name, url := range subViewM2 {
		bt2.AddSub(NewViewButton(name, url))
	}
	bt2.AddSub(NewMediaButton("我要报名", "Ho4LM-FX2a1H3G7gccMCj34lN4EXQsWSq2WDqRpuib2eqYGNJ8FRrBpq1XX8tR-V"))

	bt3 := NewTopButton("我的")
	bt3.AddSub(NewMiniProgramButton("wx42b13b304e0b37db", "我的课程", "https://www.youzan.com/v2/showcase/usercenter/index", "pages/lesson/lesson"))
	bt3.AddSub(NewMediaButton("订单开课", "6LtuzGCxGhPmXQ784UZGkHe1krKVyD5ZcDWOi1ovRIcpwrLjUSJ6sg_t3s83C6BU"))
	bt3.AddSub(NewMediaButton("联系客服", "azG95Co0_SD_P94xBcTJzVsiEXDsbpWcPCVnxFrQ-sA2zj9LUZw2noQJxcAEr-rt"))
	bt3.AddSub(NewViewLimitedButton("关于大嘴妞", "XE0SukUeRkIhGbxSQMOVWyGYHmERIaKLu69Xhqj4W9A"))

	//err = CreateMenu(accToken.Token, bt1, bt2, bt3)
	//if err != nil {
	//	t.Fatal(err)
	//}
}
