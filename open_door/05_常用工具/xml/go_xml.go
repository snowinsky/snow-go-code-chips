package main

import (
	"encoding/xml"
	"fmt"
)

//https://tool.hiofd.com/xml-to-go/ 在线工具，可以把xml字符串转成struct

func main() {
	xmlString := `<batchFile>
	<header count="2" totalAmount="1000">
	    <language>english</language>
	    <format>xml</format>
	</header>
	<content>
	    <record id="1">
	        <name>aaa</name>
	        <value>1111</value>
	    </record>
	    <record id="2">
	        <name>aaa</name>
	        <value>1111</value>
	    </record>
	</content>
</batchFile>`

	bf := new(BatchFile)
	//把xml数据解析成bs对象
	err := xml.Unmarshal([]byte(xmlString), bf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bf.Header.Format)
	fmt.Println(bf.Header.TotalAmount)

	//把对象转成xml
	data, _ := xml.MarshalIndent(bf, "", "  ")
	fmt.Println(string(data))

}

type BatchFile struct {
	XMLName xml.Name `xml:"batchFile"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text        string `xml:",chardata"`
		Count       string `xml:"count,attr"`
		TotalAmount string `xml:"totalAmount,attr"`
		Language    string `xml:"language"`
		Format      string `xml:"format"`
	} `xml:"header"`
	Content struct {
		Text   string `xml:",chardata"`
		Record []struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id,attr"`
			Name  string `xml:"name"`
			Value string `xml:"value"`
		} `xml:"record"`
	} `xml:"content"`
}
