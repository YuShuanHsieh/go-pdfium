package implementation

// #cgo pkg-config: pdfium
// #include "fpdf_doc.h"
// #include <stdlib.h>
import "C"

import (
	"errors"
	"unsafe"

	"github.com/klippa-app/go-pdfium/requests"
	"github.com/klippa-app/go-pdfium/responses"
)

// GetMetaData returns the metadata values of the document.
func (p *PdfiumImplementation) GetMetaData(request *requests.GetMetaData) (*responses.GetMetaData, error) {
	p.Lock()
	defer p.Unlock()

	nativeDoc, err := p.getNativeDocument(request.Document)
	if err != nil {
		return nil, err
	}

	if nativeDoc.currentDoc == nil {
		return nil, errors.New("no current document")
	}

	getMetaText := func(tag string) (string, error) {
		cstr := C.CString(tag)
		defer C.free(unsafe.Pointer(cstr))

		// First get the metadata length.
		metaSize := C.FPDF_GetMetaText(nativeDoc.currentDoc, cstr, C.NULL, 0)
		if metaSize == 0 {
			return "", errors.New("Could not get metadata")
		}

		charData := make([]byte, metaSize)
		C.FPDF_GetMetaText(nativeDoc.currentDoc, cstr, unsafe.Pointer(&charData[0]), C.ulong(len(charData)))

		transformedText, err := p.transformUTF16LEText(charData)
		if err != nil {
			return "", err
		}

		return transformedText, nil
	}

	if request.Tags == nil {
		request.Tags = &[]string{"Title", "Author", "Subject", "Keywords", "Creator", "Producer", "CreationDate", "ModDate"}
	}

	results := []responses.GetMetaDataTag{}

	tags := *request.Tags
	for i := range tags {
		result, err := getMetaText(tags[i])
		if err != nil {
			return nil, err
		}

		results = append(results, responses.GetMetaDataTag{
			Tag:   tags[i],
			Value: result,
		})
	}

	return &responses.GetMetaData{
		Tags: results,
	}, nil
}
