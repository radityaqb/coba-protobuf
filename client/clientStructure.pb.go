// Code generated by protoc-gen-go.
// source: clientStructure.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	clientStructure.proto

It has these top-level messages:
	Client
*/
package main

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Client struct {
	Id               *int32         `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name             *string        `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email            *string        `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Country          *string        `protobuf:"bytes,4,opt,name=country" json:"country,omitempty"`
	Inbox            []*Client_Mail `protobuf:"bytes,5,rep,name=inbox" json:"inbox,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}

func (m *Client) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Client) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Client) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *Client) GetCountry() string {
	if m != nil && m.Country != nil {
		return *m.Country
	}
	return ""
}

func (m *Client) GetInbox() []*Client_Mail {
	if m != nil {
		return m.Inbox
	}
	return nil
}

type Client_Mail struct {
	RemoteEmail      *string `protobuf:"bytes,1,opt,name=remoteEmail" json:"remoteEmail,omitempty"`
	Body             *string `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Client_Mail) Reset()         { *m = Client_Mail{} }
func (m *Client_Mail) String() string { return proto.CompactTextString(m) }
func (*Client_Mail) ProtoMessage()    {}

func (m *Client_Mail) GetRemoteEmail() string {
	if m != nil && m.RemoteEmail != nil {
		return *m.RemoteEmail
	}
	return ""
}

func (m *Client_Mail) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func init() {
}
