package qr

import (
	"testing"
)

func TestModuler(t *testing.T) {
	input := "linkedin.com/in/andreibuiciuc"

	v := newVersioner()
	string, _ := v.getMode(input)
	vrs, _ := v.getVersion(input, string, ec_MEDIUM)

	e := newEncoder()
	encoded, _ := e.encode(input, ec_MEDIUM)
	encoded = e.augmentEncodedInput(encoded, vrs, ec_MEDIUM)

	i := newInterleaver()
	data := i.getFinalMessage(encoded, vrs, ec_MEDIUM)

	m := newModuler(vrs, ec_MEDIUM)
	matrix, _ := m.createModuleMatrix(data)

	qi := NewImage()
	qi.CreateImage("best.png", matrix.GetMatrix())
}
