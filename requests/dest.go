package requests

import "github.com/YuShuanHsieh/go-pdfium/references"

type GetDestInfo struct {
	Document references.FPDF_DOCUMENT
	Dest     references.FPDF_DEST
}
