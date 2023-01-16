// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson8ce2cb00DecodeDbPerformancEprojectInternalForumDeliveryModels(in *jlexer.Lexer, out *ThreadsList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(ThreadsList, 0, 0)
			} else {
				*out = ThreadsList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 ForumGetSlugThreadsResponse
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8ce2cb00EncodeDbPerformancEprojectInternalForumDeliveryModels(out *jwriter.Writer, in ThreadsList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v ThreadsList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8ce2cb00EncodeDbPerformancEprojectInternalForumDeliveryModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ThreadsList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8ce2cb00EncodeDbPerformancEprojectInternalForumDeliveryModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ThreadsList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8ce2cb00DecodeDbPerformancEprojectInternalForumDeliveryModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ThreadsList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8ce2cb00DecodeDbPerformancEprojectInternalForumDeliveryModels(l, v)
}
func easyjson8ce2cb00DecodeDbPerformancEprojectInternalForumDeliveryModels1(in *jlexer.Lexer, out *ForumGetSlugThreadsResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint32(in.Uint32())
		case "title":
			out.Title = string(in.String())
		case "author":
			out.Author = string(in.String())
		case "forum":
			out.Forum = string(in.String())
		case "slug":
			out.Slug = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "created":
			out.Created = string(in.String())
		case "votes":
			out.Votes = int32(in.Int32())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8ce2cb00EncodeDbPerformancEprojectInternalForumDeliveryModels1(out *jwriter.Writer, in ForumGetSlugThreadsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint32(uint32(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"forum\":"
		out.RawString(prefix)
		out.String(string(in.Forum))
	}
	{
		const prefix string = ",\"slug\":"
		out.RawString(prefix)
		out.String(string(in.Slug))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"created\":"
		out.RawString(prefix)
		out.String(string(in.Created))
	}
	{
		const prefix string = ",\"votes\":"
		out.RawString(prefix)
		out.Int32(int32(in.Votes))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ForumGetSlugThreadsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8ce2cb00EncodeDbPerformancEprojectInternalForumDeliveryModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ForumGetSlugThreadsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8ce2cb00EncodeDbPerformancEprojectInternalForumDeliveryModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ForumGetSlugThreadsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8ce2cb00DecodeDbPerformancEprojectInternalForumDeliveryModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ForumGetSlugThreadsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8ce2cb00DecodeDbPerformancEprojectInternalForumDeliveryModels1(l, v)
}
