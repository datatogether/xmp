package xmp

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

func XmpMetaFromObject(stream []byte) (map[string]interface{}, error) {
	data := &XmpMeta{}

	err := xml.Unmarshal(stream, &data)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(jsonBytes))
	out := map[string]interface{}{}
	err = json.Unmarshal(jsonBytes, &out)
	return out, err
}

type XmpDocument struct {
	Meta XmpMeta `xml:"xmpmeta"`
}

type XmpMeta struct {
	Rdf Rdf `xml:"RDF"`
}

type Rdf struct {
	Description RdfDescription `xml:"Description"`
}

type RdfDescription struct {
	Title        XmpTitle `xml:"title"`
	Format       string   `xml:"format"`
	Creator      string   `xml:"creator"`
	Subject      []string `xml:"subject"`
	InstanceId   string   `xml:"InstanceID"`
	CreateDate   *time.Time
	MetadataDate *time.Time
	ModifyDate   *time.Time
	Keywords     string
}

type XmpTitle struct {
	Title RdfAlt `xml:"alt"`
}

type RdfAlt struct {
	Items []string `xml:"li"`
}
