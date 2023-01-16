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

func easyjson30bd9105DecodeDbPerformanceProjectInternalUserDeliveryModels(in *jlexer.Lexer, out *UserCreateResponse) {
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
		case "nickname":
			out.Nickname = string(in.String())
		case "fullname":
			out.FullName = string(in.String())
		case "about":
			out.About = string(in.String())
		case "email":
			out.Email = string(in.String())
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
func easyjson30bd9105EncodeDbPerformanceProjectInternalUserDeliveryModels(out *jwriter.Writer, in UserCreateResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix[1:])
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"fullname\":"
		out.RawString(prefix)
		out.String(string(in.FullName))
	}
	{
		const prefix string = ",\"about\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserCreateResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson30bd9105EncodeDbPerformanceProjectInternalUserDeliveryModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserCreateResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson30bd9105EncodeDbPerformanceProjectInternalUserDeliveryModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserCreateResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson30bd9105DecodeDbPerformanceProjectInternalUserDeliveryModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserCreateResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson30bd9105DecodeDbPerformanceProjectInternalUserDeliveryModels(l, v)
}
func easyjson30bd9105DecodeDbPerformanceProjectInternalUserDeliveryModels1(in *jlexer.Lexer, out *UserCreateRequest) {
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
		case "Nickname":
			out.Nickname = string(in.String())
		case "fullname":
			out.FullName = string(in.String())
		case "about":
			out.About = string(in.String())
		case "email":
			out.Email = string(in.String())
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
func easyjson30bd9105EncodeDbPerformanceProjectInternalUserDeliveryModels1(out *jwriter.Writer, in UserCreateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Nickname\":"
		out.RawString(prefix[1:])
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"fullname\":"
		out.RawString(prefix)
		out.String(string(in.FullName))
	}
	{
		const prefix string = ",\"about\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserCreateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson30bd9105EncodeDbPerformanceProjectInternalUserDeliveryModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserCreateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson30bd9105EncodeDbPerformanceProjectInternalUserDeliveryModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserCreateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson30bd9105DecodeDbPerformanceProjectInternalUserDeliveryModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserCreateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson30bd9105DecodeDbPerformanceProjectInternalUserDeliveryModels1(l, v)
}
