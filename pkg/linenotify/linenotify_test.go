package linenotify_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hsmtkk/line_notify_go/pkg/linenotify"
	"github.com/hsmtkk/line_notify_go/test"
)

func TestStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		want := "Bearer test"
		got := r.Header.Get("Authorization")
		test.AssertEqualString(t, want, got)
	}))
	defer ts.Close()
	notifier := linenotify.NewForTest(ts.Client(), ts.URL)
	err := notifier.Status()
	test.AssertNil(t, err)
}

func TestNotifyMessage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		want := "Bearer test"
		got := r.Header.Get("Authorization")
		test.AssertEqualString(t, want, got)
		want = "test"
		got = r.FormValue("message")
		test.AssertEqualString(t, want, got)
	}))
	defer ts.Close()
	notifier := linenotify.NewForTest(ts.Client(), ts.URL)
	err := notifier.NotifyMessage("test")
	test.AssertNil(t, err)
}
