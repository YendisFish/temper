package temper_test

import (
	"YendisFish/temper"
	"testing"
)

func TestError(t *testing.T) {
	temper.Error("Test Error", temper.LogTime{"Jan/01/2006", temper.Blue}, []string{"SOME INFORMATION", "The actual information"})
}

func TestInfo(t *testing.T) {
	temper.Info("Test Info", temper.LogTime{"Jan/01/2006", temper.Blue}, []string{"SOME INFORMATION", "The actual information"})
}

func TestWarn(t *testing.T) {
	temper.Warn("Test Warning", temper.LogTime{"Jan/01/2006", temper.Blue}, []string{"SOME INFORMATION", "The actual information"})
}

func TestCustom(t *testing.T) {
	temper.Custom([]string{"Custom Message", "Some cool message!"}, temper.Magenta, temper.LogTime{"Jan/01/2006", temper.Blue}, []string{"SOME INFORMATION", "The actual information"})
}
