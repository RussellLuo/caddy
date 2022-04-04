package reverseproxy

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func TestAUpstreams_GetUpstreams(t *testing.T) {
	au := &AUpstreams{
		Name: "caddyserver.com",
	}

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	repl := caddyhttp.NewTestReplacer(r)
	ctx := context.WithValue(r.Context(), caddy.ReplacerCtxKey, repl)

	if err := au.Provision(caddy.Context{Context: ctx}); err != nil {
		t.Fatalf("err: %v", err)
	}

	want := []*Upstream{
		{Dial: "165.227.20.207:80"},
		{Dial: "[2604:a880:2:d0::21b0:6001]:80"},
	}

	got, err := au.GetUpstreams(r)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("GetUpstreams: want(%+v) != got(%+v)", want, got)
	}

	got, err = au.GetUpstreamsNoTimeSince(r)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("GetUpstreamsNoTimeSince: want(%+v) != got(%+v)", want, got)
	}

	got, err = au.GetUpstreamsSyncMapNoTimeSince(r)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("GetUpstreamsSyncMapNoTimeSince: want(%+v) != got(%+v)", want, got)
	}
}

func BenchmarkAUpstreams_GetUpstreams(b *testing.B) {
	au := &AUpstreams{
		Name: "caddyserver.com",
	}

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	repl := caddyhttp.NewTestReplacer(r)
	ctx := context.WithValue(r.Context(), caddy.ReplacerCtxKey, repl)

	_ = au.Provision(caddy.Context{Context: ctx})
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = au.GetUpstreams(r)
		}
	})
}

func BenchmarkAUpstreams_GetUpstreamsNoTimeSince(b *testing.B) {
	au := &AUpstreams{
		Name: "caddyserver.com",
	}

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	repl := caddyhttp.NewTestReplacer(r)
	ctx := context.WithValue(r.Context(), caddy.ReplacerCtxKey, repl)

	_ = au.Provision(caddy.Context{Context: ctx})
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = au.GetUpstreamsNoTimeSince(r)
		}
	})
}

func BenchmarkAUpstreams_GetUpstreamsSyncMapNoTimeSince(b *testing.B) {
	au := &AUpstreams{
		Name: "caddyserver.com",
	}

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	repl := caddyhttp.NewTestReplacer(r)
	ctx := context.WithValue(r.Context(), caddy.ReplacerCtxKey, repl)

	_ = au.Provision(caddy.Context{Context: ctx})
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = au.GetUpstreamsSyncMapNoTimeSince(r)
		}
	})
}
