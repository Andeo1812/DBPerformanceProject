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

func easyjson5cae51c1DecodeDbPerformanceProjectInternalPostDeliveryModels(in *jlexer.Lexer, out *PostUpdateResponse) {
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
			out.ID = int64(in.Int64())
		case "parent":
			out.Parent = int64(in.Int64())
		case "author":
			out.Author = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "isEdited":
			out.IsEdited = bool(in.Bool())
		case "forum":
			out.Forum = string(in.String())
		case "thread":
			out.Thread = int64(in.Int64())
		case "created":
			out.Created = string(in.String())
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
func easyjson5cae51c1EncodeDbPerformanceProjectInternalPostDeliveryModels(out *jwriter.Writer, in PostUpdateResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"parent\":"
		out.RawString(prefix)
		out.Int64(int64(in.Parent))
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"isEdited\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsEdited))
	}
	{
		const prefix string = ",\"forum\":"
		out.RawString(prefix)
		out.String(string(in.Forum))
	}
	{
		const prefix string = ",\"thread\":"
		out.RawString(prefix)
		out.Int64(int64(in.Thread))
	}
	{
		const prefix string = ",\"created\":"
		out.RawString(prefix)
		out.String(string(in.Created))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostUpdateResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5cae51c1EncodeDbPerformanceProjectInternalPostDeliveryModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostUpdateResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5cae51c1EncodeDbPerformanceProjectInternalPostDeliveryModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostUpdateResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5cae51c1DecodeDbPerformanceProjectInternalPostDeliveryModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostUpdateResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5cae51c1DecodeDbPerformanceProjectInternalPostDeliveryModels(l, v)
}
func easyjson5cae51c1DecodeDbPerformanceProjectInternalPostDeliveryModels1(in *jlexer.Lexer, out *PostUpdateRequest) {
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
		case "ID":
			out.ID = int64(in.Int64())
		case "message":
			out.Message = string(in.String())
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
func easyjson5cae51c1EncodeDbPerformanceProjectInternalPostDeliveryModels1(out *jwriter.Writer, in PostUpdateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostUpdateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5cae51c1EncodeDbPerformanceProjectInternalPostDeliveryModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostUpdateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5cae51c1EncodeDbPerformanceProjectInternalPostDeliveryModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostUpdateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5cae51c1DecodeDbPerformanceProjectInternalPostDeliveryModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostUpdateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5cae51c1DecodeDbPerformanceProjectInternalPostDeliveryModels1(l, v)
}
