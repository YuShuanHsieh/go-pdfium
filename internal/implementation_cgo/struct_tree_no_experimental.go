//go:build !pdfium_experimental
// +build !pdfium_experimental

package implementation_cgo

// #cgo pkg-config: pdfium
// #include "fpdf_structtree.h"
import "C"
import (
	"github.com/YuShuanHsieh/go-pdfium/references"

	"github.com/google/uuid"
)

type StructElementAttributeHandle struct {
	nativeRef references.FPDF_STRUCTELEMENT_ATTR // A string that is our reference inside the process. We need this to close the references in DestroyLibrary.
}

func (p *PdfiumImplementation) registerStructElementAttribute(structElementAttribute interface{}) *StructElementAttributeHandle {
	ref := uuid.New()
	handle := &StructElementAttributeHandle{
		nativeRef: references.FPDF_STRUCTELEMENT_ATTR(ref.String()),
	}

	p.structElementAttributeRefs[handle.nativeRef] = handle

	return handle
}

type StructElementAttributeValueHandle struct {
	nativeRef references.FPDF_STRUCTELEMENT_ATTR_VALUE // A string that is our reference inside the process. We need this to close the references in DestroyLibrary.
}

func (p *PdfiumImplementation) registerStructElementAttributeValue(structElementAttributeValue interface{}) *StructElementAttributeValueHandle {
	ref := uuid.New()
	handle := &StructElementAttributeValueHandle{
		nativeRef: references.FPDF_STRUCTELEMENT_ATTR_VALUE(ref.String()),
	}

	p.structElementAttributeValueRefs[handle.nativeRef] = handle

	return handle
}
