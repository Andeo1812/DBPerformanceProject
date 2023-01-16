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

func easyjson8ebeb8efDecodeDbPerformanceProjectInternalModels(in *jlexer.Lexer, out *StatusService) {
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
		case "user":
			out.User = uint32(in.Uint32())
		case "forum":
			out.Forum = uint32(in.Uint32())
		case "thread":
			out.Thread = uint32(in.Uint32())
		case "post":
			out.Post = uint32(in.Uint32())
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
func easyjson8ebeb8efEncodeDbPerformanceProjectInternalModels(out *jwriter.Writer, in StatusService) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix[1:])
		out.Uint32(uint32(in.User))
	}
	{
		const prefix string = ",\"forum\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Forum))
	}
	{
		const prefix string = ",\"thread\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Thread))
	}
	{
		const prefix string = ",\"post\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Post))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StatusService) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8ebeb8efEncodeDbPerformanceProjectInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StatusService) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8ebeb8efEncodeDbPerformanceProjectInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StatusService) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8ebeb8efDecodeDbPerformanceProjectInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StatusService) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8ebeb8efDecodeDbPerformanceProjectInternalModels(l, v)
}
func easyjson8ebeb8efDecodeDbPerformanceProjectInternalModels1(in *jlexer.Lexer, out *PostDetails) {
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
		case "post":
			easyjson8ebeb8efDecodeDbPerformanceProjectInternalModels2(in, &out.Post)
		case "author":
			easyjson8ebeb8efDecodeDbPerformanceProjectInternalModels3(in, &out.Author)
		case "thread":
			easyjson8ebeb8efDecodeDbPerformanceProjectInternalModels4(in, &out.Thread)
		case "forum":
			easyjson8ebeb8efDecodeDbPerformanceProjectInternalModels5(in, &out.Forum)
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
func easyjson8ebeb8efEncodeDbPerformanceProjectInternalModels1(out *jwriter.Writer, in PostDetails) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"post\":"
		out.RawString(prefix[1:])
		easyjson8ebeb8efEncodeDbPerformanceProjectInternalModels2(out, in.Post)
	}
	if true {
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		easyjson8ebeb8efEncodeDbPerformanceProjectInternalModels3(out, in.Author)
	}
	if true {
		const prefix string = ",\"thread\":"
		out.RawString(prefix)
		easyjson8ebeb8efEncodeDbPerformanceProjectInternalModels4(out, in.Thread)
	}
	if true {
		const prefix string = ",\"forum\":"
		out.RawString(prefix)
		easyjson8ebeb8efEncodeDbPerformanceProjectInternalModels5(out, in.Forum)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostDetails) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8ebeb8efEncodeDbPerformanceProjectInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostDetails) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8ebeb8efEncodeDbPerformanceProjectInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostDetails) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8ebeb8efDecodeDbPerformanceProjectInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostDetails) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8ebeb8efDecodeDbPerformanceProjectInternalModels1(l, v)
}
func easyjson8ebeb8efDecodeDbPerformanceProjectInternalModels5(in *jlexer.Lexer, out *Forum) {
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
			out.ID = uint32(in.Uint32())
		case "Title":
			out.Title = string(in.String())
		case "User":
			out.User = string(in.String())
		case "Slug":
			out.Slug = string(in.String())
		case "Posts":
			out.Posts = uint32(in.Uint32())
		case "Threads":
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
func easyjson8ebeb8efEncodeDbPerformanceProjectInternalModels5(out *jwriter.Writer, in Forum) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Uint32(uint32(in.ID))
	}
	{
		const prefix string = ",\"Title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"User\":"
		out.RawString(prefix)
		out.String(string(in.User))
	}
	{
		const prefix string = ",\"Slug\":"
		out.RawString(prefix)
		out.String(string(in.Slug))
	}
	{
		const prefix string = ",\"Posts\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Posts))
	}
	{
		const prefix string = ",\"Threads\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Threads))
	}
	out.RawByte('}')
}
func easyjson8ebeb8efDecodeDbPerformanceProjectInternalModels4(in *jlexer.Lexer, out *Thread) {
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
			out.ID = uint32(in.Uint32())
		case "Title":
			out.Title = string(in.String())
		case "Author":
			out.Author = string(in.String())
		case "Forum":
			out.Forum = string(in.String())
		case "Slug":
			out.Slug = string(in.String())
		case "Message":
			out.Message = string(in.String())
		case "Created":
			out.Created = string(in.String())
		case "Votes":
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
func easyjson8ebeb8efEncodeDbPerformanceProjectInternalModels4(out *jwriter.Writer, in Thread) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Uint32(uint32(in.ID))
	}
	{
		const prefix string = ",\"Title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"Author\":"
		out.RawString(prefix)
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"Forum\":"
		out.RawString(prefix)
		out.String(string(in.Forum))
	}
	{
		const prefix string = ",\"Slug\":"
		out.RawString(prefix)
		out.String(string(in.Slug))
	}
	{
		const prefix string = ",\"Message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"Created\":"
		out.RawString(prefix)
		out.String(string(in.Created))
	}
	{
		const prefix string = ",\"Votes\":"
		out.RawString(prefix)
		out.Int32(int32(in.Votes))
	}
	out.RawByte('}')
}
func easyjson8ebeb8efDecodeDbPerformanceProjectInternalModels3(in *jlexer.Lexer, out *User) {
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
			out.ID = uint32(in.Uint32())
		case "Nickname":
			out.Nickname = string(in.String())
		case "FullName":
			out.FullName = string(in.String())
		case "About":
			out.About = string(in.String())
		case "Email":
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
func easyjson8ebeb8efEncodeDbPerformanceProjectInternalModels3(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Uint32(uint32(in.ID))
	}
	{
		const prefix string = ",\"Nickname\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"FullName\":"
		out.RawString(prefix)
		out.String(string(in.FullName))
	}
	{
		const prefix string = ",\"About\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	{
		const prefix string = ",\"Email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	out.RawByte('}')
}
func easyjson8ebeb8efDecodeDbPerformanceProjectInternalModels2(in *jlexer.Lexer, out *Post) {
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
			out.ID = uint32(in.Uint32())
		case "Parent":
			out.Parent = uint32(in.Uint32())
		case "Author":
			out.Author = string(in.String())
		case "Message":
			out.Message = string(in.String())
		case "IsEdited":
			out.IsEdited = bool(in.Bool())
		case "Forum":
			out.Forum = string(in.String())
		case "Thread":
			out.Thread = uint32(in.Uint32())
		case "Created":
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
func easyjson8ebeb8efEncodeDbPerformanceProjectInternalModels2(out *jwriter.Writer, in Post) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Uint32(uint32(in.ID))
	}
	{
		const prefix string = ",\"Parent\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Parent))
	}
	{
		const prefix string = ",\"Author\":"
		out.RawString(prefix)
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"Message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"IsEdited\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsEdited))
	}
	{
		const prefix string = ",\"Forum\":"
		out.RawString(prefix)
		out.String(string(in.Forum))
	}
	{
		const prefix string = ",\"Thread\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Thread))
	}
	{
		const prefix string = ",\"Created\":"
		out.RawString(prefix)
		out.String(string(in.Created))
	}
	out.RawByte('}')
}
