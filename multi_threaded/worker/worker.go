package worker

import (
	"github.com/YuShuanHsieh/go-pdfium"
	"github.com/YuShuanHsieh/go-pdfium/internal/implementation_cgo"
)

func StartWorker(config *pdfium.LibraryConfig) {
	implementation_cgo.StartPlugin(config)
}
