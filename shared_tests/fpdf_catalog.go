//go:build pdfium_experimental
// +build pdfium_experimental

package shared_tests

import (
	"io/ioutil"

	"github.com/YuShuanHsieh/go-pdfium/references"
	"github.com/YuShuanHsieh/go-pdfium/requests"
	"github.com/YuShuanHsieh/go-pdfium/responses"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("fpdf_catalog", func() {
	BeforeEach(func() {
		Locker.Lock()
	})

	AfterEach(func() {
		Locker.Unlock()
	})

	Context("no document", func() {
		When("is opened", func() {
			It("returns an error when getting the document tagged status", func() {
				isTagged, err := PdfiumInstance.FPDFCatalog_IsTagged(&requests.FPDFCatalog_IsTagged{})
				Expect(err).To(MatchError("document not given"))
				Expect(isTagged).To(BeNil())
			})

			It("returns an error when setting the document language", func() {
				FPDFCatalog_SetLanguage, err := PdfiumInstance.FPDFCatalog_SetLanguage(&requests.FPDFCatalog_SetLanguage{})
				Expect(err).To(MatchError("document not given"))
				Expect(FPDFCatalog_SetLanguage).To(BeNil())
			})
		})
	})

	Context("a normal PDF file", func() {
		var doc references.FPDF_DOCUMENT

		BeforeEach(func() {
			pdfData, err := ioutil.ReadFile(TestDataPath + "/testdata/test.pdf")
			Expect(err).To(BeNil())

			newDoc, err := PdfiumInstance.FPDF_LoadMemDocument(&requests.FPDF_LoadMemDocument{
				Data: &pdfData,
			})
			Expect(err).To(BeNil())

			doc = newDoc.Document
		})

		AfterEach(func() {
			FPDF_CloseDocument, err := PdfiumInstance.FPDF_CloseDocument(&requests.FPDF_CloseDocument{
				Document: doc,
			})
			Expect(err).To(BeNil())
			Expect(FPDF_CloseDocument).To(Not(BeNil()))
		})

		When("is opened", func() {
			It("returns that its not tagged", func() {
				isTagged, err := PdfiumInstance.FPDFCatalog_IsTagged(&requests.FPDFCatalog_IsTagged{
					Document: doc,
				})
				Expect(err).To(BeNil())
				Expect(isTagged).To(Equal(&responses.FPDFCatalog_IsTagged{
					IsTagged: false,
				}))
			})
		})
	})

	Context("a tagged PDF file", func() {
		var doc references.FPDF_DOCUMENT

		BeforeEach(func() {
			pdfData, err := ioutil.ReadFile(TestDataPath + "/testdata/tagged_table.pdf")
			Expect(err).To(BeNil())

			newDoc, err := PdfiumInstance.FPDF_LoadMemDocument(&requests.FPDF_LoadMemDocument{
				Data: &pdfData,
			})
			Expect(err).To(BeNil())

			doc = newDoc.Document
		})

		AfterEach(func() {
			FPDF_CloseDocument, err := PdfiumInstance.FPDF_CloseDocument(&requests.FPDF_CloseDocument{
				Document: doc,
			})
			Expect(err).To(BeNil())
			Expect(FPDF_CloseDocument).To(Not(BeNil()))
		})

		When("is opened", func() {
			It("returns that it is tagged", func() {
				isTagged, err := PdfiumInstance.FPDFCatalog_IsTagged(&requests.FPDFCatalog_IsTagged{
					Document: doc,
				})
				Expect(err).To(BeNil())
				Expect(isTagged).To(Equal(&responses.FPDFCatalog_IsTagged{
					IsTagged: true,
				}))
			})

			It("can have its language changed", func() {
				FPDFCatalog_SetLanguage, err := PdfiumInstance.FPDFCatalog_SetLanguage(&requests.FPDFCatalog_SetLanguage{
					Document: doc,
					Language: "hu",
				})
				Expect(err).To(BeNil())
				Expect(FPDFCatalog_SetLanguage).To(Equal(&responses.FPDFCatalog_SetLanguage{}))
			})
		})
	})

	Context("a new PDF file", func() {
		var doc references.FPDF_DOCUMENT

		BeforeEach(func() {
			newDoc, err := PdfiumInstance.FPDF_CreateNewDocument(&requests.FPDF_CreateNewDocument{})
			Expect(err).To(BeNil())

			doc = newDoc.Document
		})

		AfterEach(func() {
			FPDF_CloseDocument, err := PdfiumInstance.FPDF_CloseDocument(&requests.FPDF_CloseDocument{
				Document: doc,
			})
			Expect(err).To(BeNil())
			Expect(FPDF_CloseDocument).To(Not(BeNil()))
		})

		When("is opened", func() {
			It("can have its language changed", func() {
				FPDFCatalog_SetLanguage, err := PdfiumInstance.FPDFCatalog_SetLanguage(&requests.FPDFCatalog_SetLanguage{
					Document: doc,
					Language: "hu",
				})
				Expect(err).To(BeNil())
				Expect(FPDFCatalog_SetLanguage).To(Equal(&responses.FPDFCatalog_SetLanguage{}))
			})
		})
	})
})
