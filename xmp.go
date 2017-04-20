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

// XMP is itself a subset of the Resource Description Framework.
// RDF is an OG (Original Gangter) web (W3) standard for data interchange.
// https://www.w3.org/TR/rdf11-concepts/
// All XMPPackets are a single instance of an RDF document, with limitations
// defined by the XMP spec
type RDFMeta struct {
	Description DublinCore `xml:"Description"`
}

// DublinCore is the "dc:" namespace often seen in xmp meta. It's pretty sweet.
// https://en.wikipedia.org/wiki/Dublin_Core
// http://dublincore.org
type DublinCore struct {
	Title        string   `xml:"title"`
	Description  string   `xml:"description"`
	Format       string   `xml:"format"`
	Creator      string   `xml:"creator"`
	Subject      []string `xml:"subject"`
	InstanceId   string   `xml:"InstanceID"`
	CreateDate   *time.Time
	MetadataDate *time.Time
	ModifyDate   *time.Time
	Keywords     string
}
