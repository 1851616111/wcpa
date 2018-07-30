package api

import (
	"testing"
	"encoding/xml"
	"fmt"
)

func TestType_Msg(t *testing.T) {
	s := `<xml>
		<ToUserName><![CDATA[gh_b4ec5147c722]]></ToUserName>
		<FromUserName><![CDATA[oH4HtwGsY-0JSjhNhJLA7jYYOMsQ]]></FromUserName>
		<CreateTime>1532929824</CreateTime>
		<MsgType><![CDATA[text]]></MsgType>
		<Content><![CDATA[xxx]]></Content>
		<MsgId>6583883461605926895</MsgId>
	</xml>`

	var e Event
	if err := xml.Unmarshal(([]byte)(s), &e); err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", e)
}

func TestType_Event(t *testing.T) {
	s := `<xml>
		<ToUserName><![CDATA[gh_b4ec5147c722]]></ToUserName>
		<FromUserName><![CDATA[oH4HtwGsY-0JSjhNhJLA7jYYOMsQ]]></FromUserName>
		<CreateTime>1532928236</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[SCAN]]></Event>
		<EventKey><![CDATA[qr_scene_1532927897972126214]]></EventKey>
		<Ticket><![CDATA[gQEi8DwAAAAAAAAAAS5odHRwOi8vd2VpeGluLnFxLmNvbS9xLzAyWkpCZ3dudjFmSjAxcWV4dWhyYzYAAgSan15bAwT0AQAA]]></Ticket>
	</xml>`

	var e Event
	if err := xml.Unmarshal(([]byte)(s), &e); err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", e)
}