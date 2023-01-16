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

func easyjson315bdd5DecodeDbPerformancEprojectInternalForumDeliveryModels(in *jlexer.Lexer, out *ForumCreateResponse) {
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
		case "title":
			out.Title = string(in.String())
		case "user":
			out.User = string(in.String())
		case "slug":
			out.Slug = string(in.String())
		case "posts":
			out.Posts = uint32(in.Uint32())
		case "threads":
			out.Threads = uint32(in.Uint32())
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
func easyjson315bdd5EncodeDbPerformancEprojectInternalForumDeliveryModels(out *jwriter.Writer, in ForumCreateResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		out.String(string(in.User))
	}
	{
		const prefix string = ",\"slug\":"
		out.RawString(prefix)
		out.String(string(in.Slug))
	}
	{
		const prefix string = ",\"posts\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Posts))
	}
	{
		const prefix string = ",\"threads\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Threads))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ForumCreateResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson315bdd5EncodeDbPerformancEprojectInternalForumDeliveryModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ForumCreateResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson315bdd5EncodeDbPerformancEprojectInternalForumDeliveryModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ForumCreateResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson315bdd5DecodeDbPerformancEprojectInternalForumDeliveryModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ForumCreateResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson315bdd5DecodeDbPerformancEprojectInternalForumDeliveryModels(l, v)
}
func easyjson315bdd5DecodeDbPerformancEprojectInternalForumDeliveryModels1(in *jlexer.Lexer, out *ForumCreateRequest) {
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
		case "title":
			out.Title = string(in.String())
		case "user":
			out.User = string(in.String())
		case "slug":
			out.Slug = string(in.String())
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
func easyjson315bdd5EncodeDbPerformancEprojectInternalForumDeliveryModels1(out *jwriter.Writer, in ForumCreateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		out.String(string(in.User))
	}
	{
		const prefix string = ",\"slug\":"
		out.RawString(prefix)
		out.String(string(in.Slug))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ForumCreateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson315bdd5EncodeDbPerformancEprojectInternalForumDeliveryModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ForumCreateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson315bdd5EncodeDbPerformancEprojectInternalForumDeliveryModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ForumCreateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson315bdd5DecodeDbPerformancEprojectInternalForumDeliveryModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ForumCreateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson315bdd5DecodeDbPerformancEprojectInternalForumDeliveryModels1(l, v)
}
