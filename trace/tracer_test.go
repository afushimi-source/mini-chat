package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Newからの戻り地がnilです")
	} else {
		tracer.Trace("こんにちは, trace package")
		if buf.String() != "こんにちは, trace package\n" {
			t.Errorf("'%s' という誤った文字列が出力されました", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("data")
}
