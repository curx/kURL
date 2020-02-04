// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package phishingprotection

import (
	"context"
	"fmt"
	"math"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	phishingprotectionpb "google.golang.org/genproto/googleapis/cloud/phishingprotection/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// PhishingProtectionServiceV1Beta1CallOptions contains the retry settings for each method of PhishingProtectionServiceV1Beta1Client.
type PhishingProtectionServiceV1Beta1CallOptions struct {
	ReportPhishing []gax.CallOption
}

func defaultPhishingProtectionServiceV1Beta1ClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("phishingprotection.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPhishingProtectionServiceV1Beta1CallOptions() *PhishingProtectionServiceV1Beta1CallOptions {
	return &PhishingProtectionServiceV1Beta1CallOptions{
		ReportPhishing: []gax.CallOption{},
	}
}

// PhishingProtectionServiceV1Beta1Client is a client for interacting with Phishing Protection API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type PhishingProtectionServiceV1Beta1Client struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	phishingProtectionServiceV1Beta1Client phishingprotectionpb.PhishingProtectionServiceV1Beta1Client

	// The call options for this service.
	CallOptions *PhishingProtectionServiceV1Beta1CallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewPhishingProtectionServiceV1Beta1Client creates a new phishing protection service v1 beta1 client.
//
// Service to report phishing URIs.
func NewPhishingProtectionServiceV1Beta1Client(ctx context.Context, opts ...option.ClientOption) (*PhishingProtectionServiceV1Beta1Client, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultPhishingProtectionServiceV1Beta1ClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &PhishingProtectionServiceV1Beta1Client{
		conn:        conn,
		CallOptions: defaultPhishingProtectionServiceV1Beta1CallOptions(),

		phishingProtectionServiceV1Beta1Client: phishingprotectionpb.NewPhishingProtectionServiceV1Beta1Client(conn),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *PhishingProtectionServiceV1Beta1Client) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PhishingProtectionServiceV1Beta1Client) Close() error {
	return c.conn.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PhishingProtectionServiceV1Beta1Client) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ReportPhishing reports a URI suspected of containing phishing content to be reviewed. Once
// the report review is complete, its result can be found in the Cloud
// Security Command Center findings dashboard for Phishing Protection. If the
// result verifies the existence of malicious phishing content, the site will
// be added the to Google’s Social Engineering
// lists (at https://support.google.com/webmasters/answer/6350487/) in order to
// protect users that could get exposed to this threat in the future.
func (c *PhishingProtectionServiceV1Beta1Client) ReportPhishing(ctx context.Context, req *phishingprotectionpb.ReportPhishingRequest, opts ...gax.CallOption) (*phishingprotectionpb.ReportPhishingResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ReportPhishing[0:len(c.CallOptions.ReportPhishing):len(c.CallOptions.ReportPhishing)], opts...)
	var resp *phishingprotectionpb.ReportPhishingResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.phishingProtectionServiceV1Beta1Client.ReportPhishing(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
