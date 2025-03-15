//go:build !pdfium_experimental
// +build !pdfium_experimental

package implementation_cgo

import (
	pdfium_errors "github.com/YuShuanHsieh/go-pdfium/errors"
	"github.com/YuShuanHsieh/go-pdfium/requests"
	"github.com/YuShuanHsieh/go-pdfium/responses"
)

// FPDFCatalog_IsTagged determines if the given document represents a tagged PDF.
// For the definition of tagged PDF, See (see 10.7 "Tagged PDF" in PDF Reference 1.7).
// Experimental API.
func (p *PdfiumImplementation) FPDFCatalog_IsTagged(request *requests.FPDFCatalog_IsTagged) (*responses.FPDFCatalog_IsTagged, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}

// FPDFCatalog_SetLanguage sets the language of a document.
// Experimental API.
func (p *PdfiumImplementation) FPDFCatalog_SetLanguage(request *requests.FPDFCatalog_SetLanguage) (*responses.FPDFCatalog_SetLanguage, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}
