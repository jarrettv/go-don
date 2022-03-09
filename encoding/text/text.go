package json

import (
	"fmt"
	"net/http"
	"strconv"
	"unsafe"

	"github.com/abemedia/go-don"
	"github.com/valyala/bytebufferpool"
)

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func decodeText(r *http.Request, v interface{}) error {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		return err
	}

	if buf.Len() > 0 {
		switch t := v.(type) {
		case *string:
			*t = b2s(buf.Bytes())

		case *[]byte:
			*t = buf.Bytes()

		case *int:
			d, err := strconv.ParseInt(b2s(buf.Bytes()), 10, 64)
			if err != nil {
				return err
			}
			*t = int(d)

		case *int8:
			d, err := strconv.ParseInt(b2s(buf.Bytes()), 10, 8)
			if err != nil {
				return err
			}
			*t = int8(d)

		case *int16:
			d, err := strconv.ParseInt(b2s(buf.Bytes()), 10, 16)
			if err != nil {
				return err
			}
			*t = int16(d)

		case *int32:
			d, err := strconv.ParseInt(b2s(buf.Bytes()), 10, 32)
			if err != nil {
				return err
			}
			*t = int32(d)

		case *int64:
			d, err := strconv.ParseInt(b2s(buf.Bytes()), 10, 64)
			if err != nil {
				return err
			}
			*t = d

		case *uint:
			d, err := strconv.ParseUint(b2s(buf.Bytes()), 10, 64)
			if err != nil {
				return err
			}
			*t = uint(d)

		case *uint8:
			d, err := strconv.ParseUint(b2s(buf.Bytes()), 10, 8)
			if err != nil {
				return err
			}
			*t = uint8(d)

		case *uint16:
			d, err := strconv.ParseUint(b2s(buf.Bytes()), 10, 16)
			if err != nil {
				return err
			}
			*t = uint16(d)

		case *uint32:
			d, err := strconv.ParseUint(b2s(buf.Bytes()), 10, 32)
			if err != nil {
				return err
			}
			*t = uint32(d)

		case *uint64:
			d, err := strconv.ParseUint(b2s(buf.Bytes()), 10, 64)
			if err != nil {
				return err
			}
			*t = d

		case *float32:
			d, err := strconv.ParseFloat(b2s(buf.Bytes()), 32)
			if err != nil {
				return err
			}
			*t = float32(d)

		case *float64:
			d, err := strconv.ParseFloat(b2s(buf.Bytes()), 64)
			if err != nil {
				return err
			}
			*t = d

		case *bool:
			d, err := strconv.ParseBool(b2s(buf.Bytes()))
			if err != nil {
				return err
			}
			*t = d

		default:
			return don.ErrUnsupportedMediaType
		}
	}

	return nil
}

func encodeText(w http.ResponseWriter, v interface{}) error {
	_, err := fmt.Fprint(w, v)
	return err
}

func init() {
	don.RegisterDecoder("text/plain", decodeText)
	don.RegisterEncoder("text/plain", encodeText)
}
