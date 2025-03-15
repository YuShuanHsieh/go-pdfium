package requests

import "github.com/YuShuanHsieh/go-pdfium/references"

type GetActionInfo struct {
	Document references.FPDF_DOCUMENT
	Action   references.FPDF_ACTION
}
