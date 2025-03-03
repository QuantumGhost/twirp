// Code generated by protoc-gen-twirp v8.1.1, DO NOT EDIT.
// source: multiple2.proto

// Multiple proto files in one package

package multiple

import context "context"
import fmt "fmt"
import http "net/http"
import ioutil "io/ioutil"
import json "encoding/json"
import strconv "strconv"
import strings "strings"

import protojson "google.golang.org/protobuf/encoding/protojson"
import proto "google.golang.org/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// Version compatibility assertion.
// If the constant is not defined in the package, that likely means
// the package needs to be updated to work with this generated code.
// See https://twitchtv.github.io/twirp/docs/version_matrix.html
const _ = twirp.TwirpPackageMinVersion_8_1_0

// ==============
// Svc2 Interface
// ==============

type Svc2 interface {
	Send(context.Context, *Msg2) (*Msg2, error)

	SamePackageProtoImport(context.Context, *Msg1) (*Msg1, error)
}

// ====================
// Svc2 Protobuf Client
// ====================

type svc2ProtobufClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewSvc2ProtobufClient creates a Protobuf client that implements the Svc2 interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewSvc2ProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) Svc2 {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwads compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "twirp.internal.twirptest.multiple", "Svc2")
	urls := [2]string{
		serviceURL + "Send",
		serviceURL + "SamePackageProtoImport",
	}

	return &svc2ProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *svc2ProtobufClient) Send(ctx context.Context, in *Msg2) (*Msg2, error) {
	ctx = ctxsetters.WithPackageName(ctx, "twirp.internal.twirptest.multiple")
	ctx = ctxsetters.WithServiceName(ctx, "Svc2")
	ctx = ctxsetters.WithMethodName(ctx, "Send")
	caller := c.callSend
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *Msg2) (*Msg2, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Msg2)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Msg2) when calling interceptor")
					}
					return c.callSend(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*Msg2)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*Msg2) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *svc2ProtobufClient) callSend(ctx context.Context, in *Msg2) (*Msg2, error) {
	out := new(Msg2)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *svc2ProtobufClient) SamePackageProtoImport(ctx context.Context, in *Msg1) (*Msg1, error) {
	ctx = ctxsetters.WithPackageName(ctx, "twirp.internal.twirptest.multiple")
	ctx = ctxsetters.WithServiceName(ctx, "Svc2")
	ctx = ctxsetters.WithMethodName(ctx, "SamePackageProtoImport")
	caller := c.callSamePackageProtoImport
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *Msg1) (*Msg1, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Msg1)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Msg1) when calling interceptor")
					}
					return c.callSamePackageProtoImport(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*Msg1)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*Msg1) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *svc2ProtobufClient) callSamePackageProtoImport(ctx context.Context, in *Msg1) (*Msg1, error) {
	out := new(Msg1)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ================
// Svc2 JSON Client
// ================

type svc2JSONClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewSvc2JSONClient creates a JSON client that implements the Svc2 interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewSvc2JSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) Svc2 {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwads compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "twirp.internal.twirptest.multiple", "Svc2")
	urls := [2]string{
		serviceURL + "Send",
		serviceURL + "SamePackageProtoImport",
	}

	return &svc2JSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *svc2JSONClient) Send(ctx context.Context, in *Msg2) (*Msg2, error) {
	ctx = ctxsetters.WithPackageName(ctx, "twirp.internal.twirptest.multiple")
	ctx = ctxsetters.WithServiceName(ctx, "Svc2")
	ctx = ctxsetters.WithMethodName(ctx, "Send")
	caller := c.callSend
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *Msg2) (*Msg2, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Msg2)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Msg2) when calling interceptor")
					}
					return c.callSend(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*Msg2)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*Msg2) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *svc2JSONClient) callSend(ctx context.Context, in *Msg2) (*Msg2, error) {
	out := new(Msg2)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *svc2JSONClient) SamePackageProtoImport(ctx context.Context, in *Msg1) (*Msg1, error) {
	ctx = ctxsetters.WithPackageName(ctx, "twirp.internal.twirptest.multiple")
	ctx = ctxsetters.WithServiceName(ctx, "Svc2")
	ctx = ctxsetters.WithMethodName(ctx, "SamePackageProtoImport")
	caller := c.callSamePackageProtoImport
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *Msg1) (*Msg1, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Msg1)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Msg1) when calling interceptor")
					}
					return c.callSamePackageProtoImport(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*Msg1)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*Msg1) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *svc2JSONClient) callSamePackageProtoImport(ctx context.Context, in *Msg1) (*Msg1, error) {
	out := new(Msg1)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ===================
// Svc2 Server Handler
// ===================

type svc2Server struct {
	Svc2
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
	jsonCamelCase    bool   // JSON fields are serialized as lowerCamelCase rather than keeping the original proto names
}

// NewSvc2Server builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewSvc2Server(svc Svc2, opts ...interface{}) TwirpServer {
	serverOpts := newServerOpts(opts)

	// Using ReadOpt allows backwards and forwads compatibility with new options in the future
	jsonSkipDefaults := false
	_ = serverOpts.ReadOpt("jsonSkipDefaults", &jsonSkipDefaults)
	jsonCamelCase := false
	_ = serverOpts.ReadOpt("jsonCamelCase", &jsonCamelCase)
	var pathPrefix string
	if ok := serverOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	return &svc2Server{
		Svc2:             svc,
		hooks:            serverOpts.Hooks,
		interceptor:      twirp.ChainInterceptors(serverOpts.Interceptors...),
		pathPrefix:       pathPrefix,
		jsonSkipDefaults: jsonSkipDefaults,
		jsonCamelCase:    jsonCamelCase,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *svc2Server) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *svc2Server) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
	if context.Canceled == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.Canceled, "failed to read request: context canceled"))
		return
	}
	if context.DeadlineExceeded == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.DeadlineExceeded, "failed to read request: deadline exceeded"))
		return
	}
	s.writeError(ctx, resp, twirp.WrapError(malformedRequestError(msg), err))
}

// Svc2PathPrefix is a convenience constant that may identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// with the default "/twirp" prefix and default CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const Svc2PathPrefix = "/twirp/twirp.internal.twirptest.multiple.Svc2/"

func (s *svc2Server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "twirp.internal.twirptest.multiple")
	ctx = ctxsetters.WithServiceName(ctx, "Svc2")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	// Verify path format: [<prefix>]/<package>.<Service>/<Method>
	prefix, pkgService, method := parseTwirpPath(req.URL.Path)
	if pkgService != "twirp.internal.twirptest.multiple.Svc2" {
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
	if prefix != s.pathPrefix {
		msg := fmt.Sprintf("invalid path prefix %q, expected %q, on path %q", prefix, s.pathPrefix, req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	switch method {
	case "Send":
		s.serveSend(ctx, resp, req)
		return
	case "SamePackageProtoImport":
		s.serveSamePackageProtoImport(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *svc2Server) serveSend(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveSendJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveSendProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *svc2Server) serveSendJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Send")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(Msg2)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.Svc2.Send
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *Msg2) (*Msg2, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Msg2)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Msg2) when calling interceptor")
					}
					return s.Svc2.Send(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*Msg2)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*Msg2) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *Msg2
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Msg2 and nil error while calling Send. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *svc2Server) serveSendProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Send")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(Msg2)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.Svc2.Send
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *Msg2) (*Msg2, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Msg2)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Msg2) when calling interceptor")
					}
					return s.Svc2.Send(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*Msg2)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*Msg2) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *Msg2
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Msg2 and nil error while calling Send. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *svc2Server) serveSamePackageProtoImport(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveSamePackageProtoImportJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveSamePackageProtoImportProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *svc2Server) serveSamePackageProtoImportJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "SamePackageProtoImport")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(Msg1)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.Svc2.SamePackageProtoImport
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *Msg1) (*Msg1, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Msg1)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Msg1) when calling interceptor")
					}
					return s.Svc2.SamePackageProtoImport(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*Msg1)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*Msg1) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *Msg1
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Msg1 and nil error while calling SamePackageProtoImport. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *svc2Server) serveSamePackageProtoImportProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "SamePackageProtoImport")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(Msg1)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.Svc2.SamePackageProtoImport
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *Msg1) (*Msg1, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Msg1)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Msg1) when calling interceptor")
					}
					return s.Svc2.SamePackageProtoImport(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*Msg1)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*Msg1) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *Msg1
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Msg1 and nil error while calling SamePackageProtoImport. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *svc2Server) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor1, 0
}

func (s *svc2Server) ProtocGenTwirpVersion() string {
	return "v8.1.1"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *svc2Server) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "twirp.internal.twirptest.multiple", "Svc2")
}

var twirpFileDescriptor1 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x2d, 0xcd, 0x29,
	0xc9, 0x2c, 0xc8, 0x49, 0x35, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2c, 0x29, 0xcf,
	0x2c, 0x2a, 0xd0, 0xcb, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x03, 0x73, 0x4b, 0x52,
	0x8b, 0x4b, 0xf4, 0x60, 0x2a, 0xa5, 0xe0, 0x7a, 0x0c, 0x21, 0x7a, 0x94, 0xd8, 0xb8, 0x58, 0x7c,
	0x8b, 0xd3, 0x8d, 0x8c, 0xce, 0x30, 0x72, 0xb1, 0x04, 0x97, 0x25, 0x1b, 0x09, 0x45, 0x70, 0xb1,
	0x04, 0xa7, 0xe6, 0xa5, 0x08, 0xa9, 0xeb, 0x11, 0x34, 0x4d, 0x0f, 0xa4, 0x53, 0x8a, 0x58, 0x85,
	0x42, 0x59, 0x5c, 0x62, 0xc1, 0x89, 0xb9, 0xa9, 0x01, 0x89, 0xc9, 0xd9, 0x89, 0xe9, 0xa9, 0x01,
	0x20, 0xeb, 0x3d, 0x73, 0x0b, 0xf2, 0x8b, 0x4a, 0x88, 0xb5, 0xcb, 0x90, 0x58, 0xbb, 0x0c, 0x9d,
	0xb8, 0xa3, 0x38, 0xf5, 0x61, 0x22, 0x49, 0x6c, 0x60, 0xaf, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x7b, 0x20, 0x14, 0xe1, 0x31, 0x01, 0x00, 0x00,
}
