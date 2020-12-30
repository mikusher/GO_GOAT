// Code generated by protoc-gen-go.
// source: semanticconcept.proto
// DO NOT EDIT!

package nyt

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SemanticConceptResponse struct {
	Status     string                   `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Copyright  string                   `protobuf:"bytes,2,opt,name=copyright" json:"copyright,omitempty"`
	NumResults uint32                   `protobuf:"varint,3,opt,name=num_results,json=numResults" json:"num_results,omitempty"`
	Results    []*SemanticConceptResult `protobuf:"bytes,4,rep,name=results" json:"results,omitempty"`
}

func (m *SemanticConceptResponse) Reset()                    { *m = SemanticConceptResponse{} }
func (m *SemanticConceptResponse) String() string            { return proto.CompactTextString(m) }
func (*SemanticConceptResponse) ProtoMessage()               {}
func (*SemanticConceptResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *SemanticConceptResponse) GetResults() []*SemanticConceptResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type SemanticConceptResult struct {
	ArticleList *SemanticConceptArticleList `protobuf:"bytes,1,opt,name=article_list,json=articleList" json:"article_list,omitempty"`
}

func (m *SemanticConceptResult) Reset()                    { *m = SemanticConceptResult{} }
func (m *SemanticConceptResult) String() string            { return proto.CompactTextString(m) }
func (*SemanticConceptResult) ProtoMessage()               {}
func (*SemanticConceptResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SemanticConceptResult) GetArticleList() *SemanticConceptArticleList {
	if m != nil {
		return m.ArticleList
	}
	return nil
}

type SemanticConceptArticleList struct {
	Results []*SemanticConceptArticle `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	Total   uint32                    `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
}

func (m *SemanticConceptArticleList) Reset()                    { *m = SemanticConceptArticleList{} }
func (m *SemanticConceptArticleList) String() string            { return proto.CompactTextString(m) }
func (*SemanticConceptArticleList) ProtoMessage()               {}
func (*SemanticConceptArticleList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *SemanticConceptArticleList) GetResults() []*SemanticConceptArticle {
	if m != nil {
		return m.Results
	}
	return nil
}

type SemanticConceptArticle struct {
	Body   string `protobuf:"bytes,1,opt,name=body" json:"body,omitempty"`
	Byline string `protobuf:"bytes,2,opt,name=byline" json:"byline,omitempty"`
	Title  string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Url    string `protobuf:"bytes,4,opt,name=url" json:"url,omitempty"`
}

func (m *SemanticConceptArticle) Reset()                    { *m = SemanticConceptArticle{} }
func (m *SemanticConceptArticle) String() string            { return proto.CompactTextString(m) }
func (*SemanticConceptArticle) ProtoMessage()               {}
func (*SemanticConceptArticle) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func init() {
	proto.RegisterType((*SemanticConceptResponse)(nil), "nyt.SemanticConceptResponse")
	proto.RegisterType((*SemanticConceptResult)(nil), "nyt.SemanticConceptResult")
	proto.RegisterType((*SemanticConceptArticleList)(nil), "nyt.SemanticConceptArticleList")
	proto.RegisterType((*SemanticConceptArticle)(nil), "nyt.SemanticConceptArticle")
}

var fileDescriptor1 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x51, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x55, 0xbe, 0xf4, 0x2b, 0xca, 0x0d, 0x95, 0xd0, 0x15, 0x2d, 0x51, 0x41, 0x6a, 0x95, 0xa9,
	0x53, 0x86, 0x02, 0x0f, 0x00, 0xac, 0x4c, 0x66, 0x64, 0xa8, 0x92, 0x60, 0x81, 0x25, 0xd7, 0x8e,
	0xe2, 0x9b, 0x21, 0x2f, 0xc4, 0x73, 0xe2, 0xd8, 0x2e, 0x19, 0x08, 0x6c, 0x3e, 0x3f, 0x3e, 0x3a,
	0x3a, 0x17, 0x96, 0x86, 0x1f, 0x4b, 0x45, 0xa2, 0xae, 0xb5, 0xaa, 0x79, 0x43, 0x45, 0xd3, 0x6a,
	0xd2, 0x18, 0xab, 0x9e, 0xf2, 0xcf, 0x08, 0xae, 0x5e, 0x82, 0xfc, 0xe4, 0x65, 0xc6, 0x4d, 0xa3,
	0x95, 0xe1, 0xb8, 0x82, 0xb9, 0xa1, 0x92, 0x3a, 0x93, 0x45, 0xdb, 0x68, 0x97, 0xb0, 0x80, 0xf0,
	0x06, 0x92, 0x5a, 0x37, 0x7d, 0x2b, 0xde, 0x3f, 0x28, 0xfb, 0xe7, 0xa4, 0x91, 0xc0, 0x0d, 0xa4,
	0xaa, 0x3b, 0x1e, 0x5a, 0x6e, 0x3a, 0x49, 0x26, 0x8b, 0xad, 0xbe, 0x60, 0x60, 0x29, 0xe6, 0x19,
	0xbc, 0x83, 0xb3, 0x93, 0x38, 0xdb, 0xc6, 0xbb, 0x74, 0xbf, 0x2e, 0x6c, 0x93, 0xe2, 0x67, 0x0b,
	0x6b, 0x61, 0x27, 0x6b, 0xfe, 0x0a, 0xcb, 0x49, 0x07, 0x3e, 0xc2, 0x79, 0xd9, 0x5a, 0x56, 0xf2,
	0x83, 0x14, 0x86, 0x5c, 0xd7, 0x74, 0xbf, 0x99, 0xca, 0x7c, 0xf0, 0xbe, 0x67, 0x6b, 0x63, 0x69,
	0x39, 0x82, 0x5c, 0xc0, 0xfa, 0x77, 0x2b, 0xde, 0x8f, 0x85, 0x23, 0x57, 0xf8, 0xfa, 0x8f, 0xf0,
	0xef, 0xc6, 0x78, 0x09, 0xff, 0x49, 0x53, 0x29, 0xdd, 0x44, 0x0b, 0xe6, 0x41, 0x2e, 0x61, 0x35,
	0xfd, 0x11, 0x11, 0x66, 0x95, 0x7e, 0xeb, 0xc3, 0xd8, 0xee, 0x3d, 0x9c, 0xa0, 0xea, 0xa5, 0x50,
	0x3c, 0xec, 0x1c, 0x90, 0xcb, 0x16, 0x24, 0xb9, 0x9b, 0x37, 0x61, 0x1e, 0xe0, 0x05, 0xc4, 0x5d,
	0x2b, 0xed, 0xaa, 0x03, 0x37, 0x3c, 0xab, 0xb9, 0x3b, 0xf5, 0xed, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x74, 0x45, 0x83, 0xfb, 0x03, 0x02, 0x00, 0x00,
}
