package echotron

import (
	"testing"
	"time"
)

type test struct{}

func (t test) Update(_ *Update) {}

var dsp *Dispatcher

func TestNewDispatcher(t *testing.T) {
	if dsp = NewDispatcher("token", func(_ int64) Bot { return test{} }); dsp == nil {
		t.Fatal("dispatcher is nil")
	}
}

func TestAddSession(t *testing.T) {
	dsp.AddSession(0)

	if len(dsp.sessionMap) == 0 {
		t.Fatal("could not add session")
	}
}

func TestDelSession(t *testing.T) {
	dsp.DelSession(0)

	if len(dsp.sessionMap) != 0 {
		t.Fatal("could not delete session")
	}
}

func TestListenWebhook(_ *testing.T) {
	go dsp.ListenWebhook("http://example.com:8443/test")
	time.Sleep(time.Second)
}

func TestPoll(_ *testing.T) {
	go dsp.Poll()
	dsp.updates <- &Update{}
	dsp.updates <- &Update{Message: &Message{Chat: &Chat{ID: 0}}}
	time.Sleep(time.Second)
}
