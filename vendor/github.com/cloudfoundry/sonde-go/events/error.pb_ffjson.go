// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: ./vendor/github.com/cloudfoundry/sonde-go/events/error.pb.go
// DO NOT EDIT!

package events

import (
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Error) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Error) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if mj.Source != nil {
		if true {
			buf.WriteString(`"source":`)
			fflib.WriteJsonString(buf, string(*mj.Source))
			buf.WriteByte(',')
		}
	}
	if mj.Code != nil {
		if true {
			buf.WriteString(`"code":`)
			fflib.FormatBits2(buf, uint64(*mj.Code), 10, *mj.Code < 0)
			buf.WriteByte(',')
		}
	}
	if mj.Message != nil {
		if true {
			buf.WriteString(`"message":`)
			fflib.WriteJsonString(buf, string(*mj.Message))
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}