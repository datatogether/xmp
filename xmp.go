// XMP is a package for parsing Extensible Metadata Platform documents.
//
// from: https://en.wikipedia.org/wiki/Extensible_Metadata_Platform
// The Extensible Metadata Platform (XMP) is an ISO standard (ISO 16684-1),
// originally created by Adobe Systems Inc., for the creation, processing and
// interchange of standardized and custom metadata for digital documents and
// data sets.
// XMP standardizes a data model, a serialization format and core properties
// for the definition and processing of extensible metadata.
// It also provides guidelines for embedding XMP information into popular image,
// video and document file formats, such as JPEG and PDF, without breaking their
// readability by applications that do not support XMP. Therefore, the non-XMP
// metadata have to be reconciled with the XMP properties.
package xmp

import (
	"time"
)

// XMPPacket is an Instance of the XMP Data Model.
// Also known as an XMP document.
type XMPPacket struct {
	RDF RDFMeta `xml:"RDF"`
}

// DublinCore is the "dc:" namespace often seen in xmp meta. It's pretty sweet.
// https://en.wikipedia.org/wiki/Dublin_Core
// http://dublincore.org
type DublinCore struct {
	Title        RDFAlt `xml:"title"`
	Description  RDFAlt `xml:"description"`
	Format       string `xml:"format"`
	Creator      RDFSeq `xml:"creator"`
	Subject      RDFBag `xml:"subject"`
	InstanceId   string `xml:"InstanceID"`
	CreateDate   *time.Time
	MetadataDate *time.Time
	ModifyDate   *time.Time
	Keywords     string
}

// The XMP basic namespace contains properties that provide basic descriptive information.
// XMP spec Section 8.4
type XMPBasic struct {
	CreateDate   *time.Time
	CreatorTool  string
	Identifier   []string
	Label        string
	MetadataDate *time.Time
}

// XMP spec Section 8.5
type XMPRights struct {
}

// The XMP Media Management namespace contains properties that provide information
// regarding the identification, composition, and history of a resource.
// XMP spec Section 8.6
type XMPMM struct {
}

// XMP is itself a subset of the Resource Description Framework.
// RDF is an OG (Original Gangter) web (W3) standard for data interchange.
// https://www.w3.org/TR/rdf11-concepts/
// All XMPPackets are a single instance of an RDF document, with limitations
// defined by the XMP spec
// TODO - RDF should probably get it's own package
type RDFMeta struct {
	Description DublinCore `xml:"Description"`
}

type RDFAlt struct {
	RDFList RDFList `xml:"Alt"`
}

func (d RDFAlt) Default() string {
	return d.RDFList.Default()
}

type RDFList struct {
	Items []string `xml:"li"`
}

func (a RDFList) Default() string {
	if a.Items == nil {
		return ""
	}
	return a.Items[0]
}

type RDFBag struct {
	Items RDFList `xml:"Bag"`
}

func (b RDFBag) Default() []string {
	return b.Items.Items
}

type RDFSeq struct {
	Items RDFList `xml:"Seq"`
}

func (seq RDFSeq) Default() []string {
	return seq.Items.Items
}

func (seq RDFSeq) String() string {
	s := ""
	for _, str := range seq.Default() {
		s += "; " + str
	}
	return s[len("; "):]
}
