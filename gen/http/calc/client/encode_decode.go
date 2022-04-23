// Code generated by goa v3.7.2, DO NOT EDIT.
//
// calc HTTP client encoders and decoders
//
// Command:
// $ goa gen calcsvc/design

package client

import (
	"bytes"
	calc "calcsvc/gen/calc"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildMultiplyRequest instantiates a HTTP request object with method and path
// set to call the "calc" service "multiply" endpoint
func (c *Client) BuildMultiplyRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: MultiplyCalcPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("calc", "multiply", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeMultiplyRequest returns an encoder for requests sent to the calc
// multiply server.
func EncodeMultiplyRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*calc.MultiplyPayload)
		if !ok {
			return goahttp.ErrInvalidType("calc", "multiply", "*calc.MultiplyPayload", v)
		}
		body := NewMultiplyRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("calc", "multiply", err)
		}
		return nil
	}
}

// DecodeMultiplyResponse returns a decoder for responses returned by the calc
// multiply endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeMultiplyResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body int
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "multiply", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("calc", "multiply", resp.StatusCode, string(body))
		}
	}
}