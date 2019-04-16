package gocudnn

import (
	"testing"
)

func TestCreateActivationDescriptor(t *testing.T) {
	var amflg ActivationMode
	var nanflg NANProp
	coef := 5.0
	ad, err := CreateActivationDescriptor()
	if err != nil {
		t.Error(err)
	}
	err = ad.Set(amflg.ClippedRelu(), nanflg.NotPropigate(), coef)
	if err != nil {
		t.Error(err)
	}
	mode, nanprop, coefreturned, err := ad.Get()
	if err != nil {
		t.Error(err)
	}
	if amflg != mode {
		t.Error("Activation Set dooesn't match returned from Get")
	}
	if nanflg != nanprop {
		t.Error("NanPropigation Set dooesn't match returned from Get")
	}
	if coef != coefreturned {
		t.Error("coef Set dooesn't match returned from Get: ", coef, coefreturned)
	}
}