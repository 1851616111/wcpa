package api

import (
	xmlutil "github.com/1851616111/util/xml"
	"encoding/xml"
)

type Message struct {
	XMLName xml.Name `xml:"xml"`

	To         xmlutil.CDATA `xml:"ToUserName"`
	From       xmlutil.CDATA `xml:"FromUserName"`
	CreateTime int64         `xml:"CreateTime"`
	MsgType    xmlutil.CDATA `xml:"MsgType"`

	Content xmlutil.CDATA `xml:"Content"`
	MsgId   string        `xml:"MsgId"`

	Event    xmlutil.CDATA `xml:"Event"`
	EventKey xmlutil.CDATA `xml:"EventKey"`
	Ticket   xmlutil.CDATA `xml:"Ticket"`

	//ArticleCount int     `xml:"ArticleCount,omitempty"`
	//articles     string  `xml:"Articles,omitempty"`
	//Items        *[]item `xml:"Articles>item,omitempty"`
}

func (e Message) IsEvent() bool {
	return string(e.MsgType) == "event"
}
func (e Message) IsScanEvent() bool {
	return string(e.Event) == "SCAN"
}

//扫码事件
//<xml>
//	<ToUserName><![CDATA[gh_b4ec5147c722]]></ToUserName>\n
//	<FromUserName><![CDATA[oH4HtwGsY-0JSjhNhJLA7jYYOMsQ]]></FromUserName>\n
//	<CreateTime>1532928236</CreateTime>\n
//	<MsgType><![CDATA[event]]></MsgType>\n
//	<Event><![CDATA[SCAN]]></Event>\n
//	<EventKey><![CDATA[qr_scene_1532927897972126214]]></EventKey>\n
//	<Ticket><![CDATA[gQEi8DwAAAAAAAAAAS5odHRwOi8vd2VpeGluLnFxLmNvbS9xLzAyWkpCZ3dudjFmSjAxcWV4dWhyYzYAAgSan15bAwT0AQAA]]></Ticket>\n
//</xml>

//<xml>
//	<ToUserName><![CDATA[gh_b4ec5147c722]]></ToUserName>
//	<FromUserName><![CDATA[oH4HtwGsY-0JSjhNhJLA7jYYOMsQ]]></FromUserName>
//	<CreateTime>1532929824</CreateTime>
//	<MsgType><![CDATA[text]]></MsgType>
//	<Content><![CDATA[xxx]]></Content>
//	<MsgId>6583883461605926895</MsgId>
//</xml>
