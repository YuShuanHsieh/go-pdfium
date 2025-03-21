//go:build pdfium_experimental
// +build pdfium_experimental

package shared_tests

import (
	"io/ioutil"

	"github.com/YuShuanHsieh/go-pdfium/enums"
	"github.com/YuShuanHsieh/go-pdfium/references"
	"github.com/YuShuanHsieh/go-pdfium/requests"
	"github.com/YuShuanHsieh/go-pdfium/responses"
	"github.com/YuShuanHsieh/go-pdfium/structs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("fpdf_text", func() {
	BeforeEach(func() {
		Locker.Lock()
	})

	AfterEach(func() {
		Locker.Unlock()
	})

	Context("no page is given", func() {
		It("returns an error when calling FPDFText_LoadPage", func() {
			FPDFText_LoadPage, err := PdfiumInstance.FPDFText_LoadPage(&requests.FPDFText_LoadPage{})
			Expect(err).To(MatchError("either page reference or index should be given"))
			Expect(FPDFText_LoadPage).To(BeNil())
		})
	})

	Context("no text page is given", func() {
		It("returns an error when calling FPDFText_LoadPage", func() {
			FPDFText_ClosePage, err := PdfiumInstance.FPDFText_ClosePage(&requests.FPDFText_ClosePage{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_ClosePage).To(BeNil())
		})

		It("returns an error when calling FPDFText_CountChars", func() {
			FPDFText_CountChars, err := PdfiumInstance.FPDFText_CountChars(&requests.FPDFText_CountChars{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_CountChars).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetUnicode", func() {
			FPDFText_GetUnicode, err := PdfiumInstance.FPDFText_ClosePage(&requests.FPDFText_ClosePage{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_GetUnicode).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetFontSize", func() {
			FPDFText_GetFontSize, err := PdfiumInstance.FPDFText_GetFontSize(&requests.FPDFText_GetFontSize{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_GetFontSize).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetFontInfo", func() {
			FPDFText_GetFontInfo, err := PdfiumInstance.FPDFText_GetFontInfo(&requests.FPDFText_GetFontInfo{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_GetFontInfo).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetFontWeight", func() {
			FPDFText_GetFontWeight, err := PdfiumInstance.FPDFText_GetFontWeight(&requests.FPDFText_GetFontWeight{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_GetFontWeight).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetFillColor", func() {
			FPDFText_GetFillColor, err := PdfiumInstance.FPDFText_GetFillColor(&requests.FPDFText_GetFillColor{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_GetFillColor).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetStrokeColor", func() {
			FPDFText_GetStrokeColor, err := PdfiumInstance.FPDFText_GetStrokeColor(&requests.FPDFText_GetStrokeColor{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_GetStrokeColor).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetCharAngle", func() {
			FPDFText_GetCharAngle, err := PdfiumInstance.FPDFText_GetCharAngle(&requests.FPDFText_GetCharAngle{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_GetCharAngle).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetCharBox", func() {
			FPDFText_GetCharBox, err := PdfiumInstance.FPDFText_GetCharBox(&requests.FPDFText_GetCharBox{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_GetCharBox).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetLooseCharBox", func() {
			FPDFText_GetLooseCharBox, err := PdfiumInstance.FPDFText_GetLooseCharBox(&requests.FPDFText_GetLooseCharBox{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_GetLooseCharBox).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetMatrix", func() {
			FPDFText_GetMatrix, err := PdfiumInstance.FPDFText_GetMatrix(&requests.FPDFText_GetMatrix{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_GetMatrix).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetCharOrigin", func() {
			FPDFText_GetCharOrigin, err := PdfiumInstance.FPDFText_GetCharOrigin(&requests.FPDFText_GetCharOrigin{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_GetCharOrigin).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetCharIndexAtPos", func() {
			FPDFText_GetCharIndexAtPos, err := PdfiumInstance.FPDFText_GetCharIndexAtPos(&requests.FPDFText_GetCharIndexAtPos{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_GetCharIndexAtPos).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetText", func() {
			FPDFText_GetText, err := PdfiumInstance.FPDFText_GetText(&requests.FPDFText_GetText{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_GetText).To(BeNil())
		})

		It("returns an error when calling FPDFText_CountRects", func() {
			FPDFText_CountRects, err := PdfiumInstance.FPDFText_CountRects(&requests.FPDFText_CountRects{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_CountRects).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetRect", func() {
			FPDFText_GetRect, err := PdfiumInstance.FPDFText_GetRect(&requests.FPDFText_GetRect{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_GetRect).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetBoundedText", func() {
			FPDFText_GetBoundedText, err := PdfiumInstance.FPDFText_GetBoundedText(&requests.FPDFText_GetBoundedText{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_GetBoundedText).To(BeNil())
		})

		It("returns an error when calling FPDFText_FindStart", func() {
			FPDFText_FindStart, err := PdfiumInstance.FPDFText_FindStart(&requests.FPDFText_FindStart{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_FindStart).To(BeNil())
		})

		It("returns an error when calling FPDFLink_LoadWebLinks", func() {
			FPDFLink_LoadWebLinks, err := PdfiumInstance.FPDFLink_LoadWebLinks(&requests.FPDFLink_LoadWebLinks{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFLink_LoadWebLinks).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetTextObject", func() {
			FPDFText_GetTextObject, err := PdfiumInstance.FPDFText_GetTextObject(&requests.FPDFText_GetTextObject{})
			Expect(err).To(MatchError("textPage not given"))
			Expect(FPDFText_GetTextObject).To(BeNil())
		})
	})

	Context("no search handle is given", func() {
		It("returns an error when calling FPDFText_FindNext", func() {
			FPDFText_FindNext, err := PdfiumInstance.FPDFText_FindNext(&requests.FPDFText_FindNext{})
			Expect(err).To(MatchError("search not given"))
			Expect(FPDFText_FindNext).To(BeNil())
		})

		It("returns an error when calling FPDFText_FindPrev", func() {
			FPDFText_FindPrev, err := PdfiumInstance.FPDFText_FindPrev(&requests.FPDFText_FindPrev{})
			Expect(err).To(MatchError("search not given"))
			Expect(FPDFText_FindPrev).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetSchResultIndex", func() {
			FPDFText_GetSchResultIndex, err := PdfiumInstance.FPDFText_GetSchResultIndex(&requests.FPDFText_GetSchResultIndex{})
			Expect(err).To(MatchError("search not given"))
			Expect(FPDFText_GetSchResultIndex).To(BeNil())
		})

		It("returns an error when calling FPDFText_GetSchCount", func() {
			FPDFText_GetSchCount, err := PdfiumInstance.FPDFText_GetSchCount(&requests.FPDFText_GetSchCount{})
			Expect(err).To(MatchError("search not given"))
			Expect(FPDFText_GetSchCount).To(BeNil())
		})

		It("returns an error when calling FPDFText_FindClose", func() {
			FPDFText_FindClose, err := PdfiumInstance.FPDFText_FindClose(&requests.FPDFText_FindClose{})
			Expect(err).To(MatchError("search not given"))
			Expect(FPDFText_FindClose).To(BeNil())
		})
	})

	Context("no page link is given", func() {
		It("returns an error when calling FPDFLink_CountWebLinks", func() {
			FPDFLink_CountWebLinks, err := PdfiumInstance.FPDFLink_CountWebLinks(&requests.FPDFLink_CountWebLinks{})
			Expect(err).To(MatchError("pageLink not given"))
			Expect(FPDFLink_CountWebLinks).To(BeNil())
		})

		It("returns an error when calling FPDFLink_GetURL", func() {
			FPDFLink_GetURL, err := PdfiumInstance.FPDFLink_GetURL(&requests.FPDFLink_GetURL{})
			Expect(err).To(MatchError("pageLink not given"))
			Expect(FPDFLink_GetURL).To(BeNil())
		})

		It("returns an error when calling FPDFLink_CountRects", func() {
			FPDFLink_CountRects, err := PdfiumInstance.FPDFLink_CountRects(&requests.FPDFLink_CountRects{})
			Expect(err).To(MatchError("pageLink not given"))
			Expect(FPDFLink_CountRects).To(BeNil())
		})

		It("returns an error when calling FPDFLink_GetRect", func() {
			FPDFLink_GetRect, err := PdfiumInstance.FPDFLink_GetRect(&requests.FPDFLink_GetRect{})
			Expect(err).To(MatchError("pageLink not given"))
			Expect(FPDFLink_GetRect).To(BeNil())
		})

		It("returns an error when calling FPDFLink_GetTextRange", func() {
			FPDFLink_GetTextRange, err := PdfiumInstance.FPDFLink_GetTextRange(&requests.FPDFLink_GetTextRange{})
			Expect(err).To(MatchError("pageLink not given"))
			Expect(FPDFLink_GetTextRange).To(BeNil())
		})

		It("returns an error when calling FPDFLink_CloseWebLinks", func() {
			FPDFLink_CloseWebLinks, err := PdfiumInstance.FPDFLink_CloseWebLinks(&requests.FPDFLink_CloseWebLinks{})
			Expect(err).To(MatchError("pageLink not given"))
			Expect(FPDFLink_CloseWebLinks).To(BeNil())
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

		When("a text page is opened", func() {
			var textPage references.FPDF_TEXTPAGE

			BeforeEach(func() {
				textPageResp, err := PdfiumInstance.FPDFText_LoadPage(&requests.FPDFText_LoadPage{
					Page: requests.Page{
						ByIndex: &requests.PageByIndex{
							Document: doc,
							Index:    0,
						},
					},
				})
				Expect(err).To(BeNil())
				Expect(textPageResp).To(Not(BeNil()))

				textPage = textPageResp.TextPage
			})

			AfterEach(func() {
				resp, err := PdfiumInstance.FPDFText_ClosePage(&requests.FPDFText_ClosePage{
					TextPage: textPage,
				})
				Expect(err).To(BeNil())
				Expect(resp).To(Not(BeNil()))
			})

			It("returns the correct character count", func() {
				FPDFText_CountChars, err := PdfiumInstance.FPDFText_CountChars(&requests.FPDFText_CountChars{
					TextPage: textPage,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_CountChars).To(Equal(&responses.FPDFText_CountChars{
					Count: 57,
				}))
			})

			It("returns the correct unicode for char 0", func() {
				FPDFText_GetUnicode, err := PdfiumInstance.FPDFText_GetUnicode(&requests.FPDFText_GetUnicode{
					TextPage: textPage,
					Index:    0,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetUnicode).To(Equal(&responses.FPDFText_GetUnicode{
					Index:   0,
					Unicode: 70,
				}))
			})

			It("returns the correct font size for char 0", func() {
				FPDFText_GetFontSize, err := PdfiumInstance.FPDFText_GetFontSize(&requests.FPDFText_GetFontSize{
					TextPage: textPage,
					Index:    0,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetFontSize).To(Equal(&responses.FPDFText_GetFontSize{
					Index:    0,
					FontSize: 1,
				}))
			})

			It("returns the correct font info for char 0", func() {
				FPDFText_GetFontInfo, err := PdfiumInstance.FPDFText_GetFontInfo(&requests.FPDFText_GetFontInfo{
					TextPage: textPage,
					Index:    0,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetFontInfo).To(Equal(&responses.FPDFText_GetFontInfo{
					Index:    0,
					FontName: "DejaVuSans",
					Flags:    524320,
				}))
			})

			It("returns an error when getting the font info for an invalid char", func() {
				FPDFText_GetFontInfo, err := PdfiumInstance.FPDFText_GetFontInfo(&requests.FPDFText_GetFontInfo{
					TextPage: textPage,
					Index:    -1,
				})
				Expect(err).To(MatchError("could not get font name"))
				Expect(FPDFText_GetFontInfo).To(BeNil())
			})

			It("returns the correct font weight for char 0", func() {
				FPDFText_GetFontWeight, err := PdfiumInstance.FPDFText_GetFontWeight(&requests.FPDFText_GetFontWeight{
					TextPage: textPage,
					Index:    0,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetFontWeight).To(Equal(&responses.FPDFText_GetFontWeight{
					Index:      0,
					FontWeight: 400,
				}))
			})

			It("returns an error when getting the font weight for an invalid char", func() {
				FPDFText_GetFontWeight, err := PdfiumInstance.FPDFText_GetFontWeight(&requests.FPDFText_GetFontWeight{
					TextPage: textPage,
					Index:    -1,
				})
				Expect(err).To(MatchError("could not get font weight"))
				Expect(FPDFText_GetFontWeight).To(BeNil())
			})

			It("returns the correct text fill color for char 0", func() {
				FPDFText_GetFillColor, err := PdfiumInstance.FPDFText_GetFillColor(&requests.FPDFText_GetFillColor{
					TextPage: textPage,
					Index:    0,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetFillColor).To(Equal(&responses.FPDFText_GetFillColor{
					Index: 0,
					R:     0,
					G:     0,
					B:     0,
					A:     255,
				}))
			})

			It("returns an error when getting the text fill color for an invalid char", func() {
				FPDFText_GetFillColor, err := PdfiumInstance.FPDFText_GetFillColor(&requests.FPDFText_GetFillColor{
					TextPage: textPage,
					Index:    -1,
				})
				Expect(err).To(MatchError("could not get fill color"))
				Expect(FPDFText_GetFillColor).To(BeNil())
			})

			It("returns the correct text stroke color for char 0", func() {
				FPDFText_GetStrokeColor, err := PdfiumInstance.FPDFText_GetStrokeColor(&requests.FPDFText_GetStrokeColor{
					TextPage: textPage,
					Index:    0,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetStrokeColor).To(Equal(&responses.FPDFText_GetStrokeColor{
					Index: 0,
					R:     0,
					G:     0,
					B:     0,
					A:     255,
				}))
			})

			It("returns an error when getting the text stroke color for an invalid char", func() {
				FPDFText_GetStrokeColor, err := PdfiumInstance.FPDFText_GetStrokeColor(&requests.FPDFText_GetStrokeColor{
					TextPage: textPage,
					Index:    -1,
				})
				Expect(err).To(MatchError("could not get stroke color"))
				Expect(FPDFText_GetStrokeColor).To(BeNil())
			})

			It("returns the correct char angle for char 0", func() {
				FPDFText_GetCharAngle, err := PdfiumInstance.FPDFText_GetCharAngle(&requests.FPDFText_GetCharAngle{
					TextPage: textPage,
					Index:    0,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetCharAngle).To(Equal(&responses.FPDFText_GetCharAngle{
					Index:     0,
					CharAngle: 0,
				}))
			})

			It("returns an error when getting the char angle for an invalid char", func() {
				FPDFText_GetCharAngle, err := PdfiumInstance.FPDFText_GetCharAngle(&requests.FPDFText_GetCharAngle{
					TextPage: textPage,
					Index:    -1,
				})
				Expect(err).To(MatchError("could not get char angle"))
				Expect(FPDFText_GetCharAngle).To(BeNil())
			})

			It("returns the correct char box for char 0", func() {
				FPDFText_GetCharBox, err := PdfiumInstance.FPDFText_GetCharBox(&requests.FPDFText_GetCharBox{
					TextPage: textPage,
					Index:    0,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetCharBox).To(Or(
					Equal(&responses.FPDFText_GetCharBox{
						Index:  0,
						Left:   71.9451904296875,
						Right:  76.55418395996094,
						Bottom: 789.1592407226562,
						Top:    797.17822265625,
					}),
					Equal(&responses.FPDFText_GetCharBox{
						Index:  0,
						Left:   71.9451904296875,
						Right:  76.55418395996094,
						Bottom: 789.1593017578125,
						Top:    797.1782836914062,
					}),
				))
			})

			It("returns an error when getting the char box for an invalid char", func() {
				FPDFText_GetCharBox, err := PdfiumInstance.FPDFText_GetCharBox(&requests.FPDFText_GetCharBox{
					TextPage: textPage,
					Index:    -1,
				})
				Expect(err).To(MatchError("could not get char box"))
				Expect(FPDFText_GetCharBox).To(BeNil())
			})

			It("returns the correct loose char box for char 0", func() {
				FPDFText_GetLooseCharBox, err := PdfiumInstance.FPDFText_GetLooseCharBox(&requests.FPDFText_GetLooseCharBox{
					TextPage: textPage,
					Index:    0,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetLooseCharBox).To(Or(
					Equal(&responses.FPDFText_GetLooseCharBox{
						Index: 0,
						Rect: structs.FPDF_FS_RECTF{
							Left:   70.8671875,
							Top:    797.9365234375,
							Right:  77.19218444824219,
							Bottom: 786.9365234375,
						},
					}),
					Equal(&responses.FPDFText_GetLooseCharBox{
						Index: 0,
						Rect: structs.FPDF_FS_RECTF{
							Left:   70.86719512939453,
							Top:    802.7113037109375,
							Right:  77.19219207763672,
							Bottom: 784.0773315429688,
						},
					}),
				))
			})

			It("returns an error when getting the loose char box for an invalid char", func() {
				FPDFText_GetLooseCharBox, err := PdfiumInstance.FPDFText_GetLooseCharBox(&requests.FPDFText_GetLooseCharBox{
					TextPage: textPage,
					Index:    -1,
				})
				Expect(err).To(MatchError("could not get loose char box"))
				Expect(FPDFText_GetLooseCharBox).To(BeNil())
			})

			It("returns the correct char matrix for char 0", func() {
				FPDFText_GetMatrix, err := PdfiumInstance.FPDFText_GetMatrix(&requests.FPDFText_GetMatrix{
					TextPage: textPage,
					Index:    0,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetMatrix).To(Or(
					Equal(&responses.FPDFText_GetMatrix{
						Index: 0,
						Matrix: structs.FPDF_FS_MATRIX{
							A: 11,
							B: 0,
							C: 0,
							D: 11,
							E: 70.8671875,
							F: 789.1592407226562,
						},
					}),
					Equal(&responses.FPDFText_GetMatrix{
						Index: 0,
						Matrix: structs.FPDF_FS_MATRIX{
							A: 11,
							B: 0,
							C: 0,
							D: 11,
							E: 70.8671875,
							F: 789.1593017578125,
						},
					}),
				))
			})

			It("returns an error when getting the char matrix for an invalid char", func() {
				FPDFText_GetMatrix, err := PdfiumInstance.FPDFText_GetMatrix(&requests.FPDFText_GetMatrix{
					TextPage: textPage,
					Index:    -1,
				})
				Expect(err).To(MatchError("could not get char matrix"))
				Expect(FPDFText_GetMatrix).To(BeNil())
			})

			It("returns the correct char origin for char 0", func() {
				FPDFText_GetCharOrigin, err := PdfiumInstance.FPDFText_GetCharOrigin(&requests.FPDFText_GetCharOrigin{
					TextPage: textPage,
					Index:    0,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetCharOrigin).To(Or(
					Equal(&responses.FPDFText_GetCharOrigin{
						Index: 0,
						X:     70.8671875,
						Y:     789.1592407226562,
					}),
					Equal(&responses.FPDFText_GetCharOrigin{
						Index: 0,
						X:     70.8671875,
						Y:     789.1593017578125,
					}),
				))
			})

			It("returns an error when getting the char origin for an invalid char", func() {
				FPDFText_GetCharOrigin, err := PdfiumInstance.FPDFText_GetCharOrigin(&requests.FPDFText_GetCharOrigin{
					TextPage: textPage,
					Index:    -1,
				})
				Expect(err).To(MatchError("could not get char origin"))
				Expect(FPDFText_GetCharOrigin).To(BeNil())
			})

			It("returns the correct char index for the given position", func() {
				FPDFText_GetCharIndexAtPos, err := PdfiumInstance.FPDFText_GetCharIndexAtPos(&requests.FPDFText_GetCharIndexAtPos{
					TextPage: textPage,
					X:        73,
					Y:        793,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetCharIndexAtPos).To(Equal(&responses.FPDFText_GetCharIndexAtPos{
					CharIndex: 0,
				}))
			})

			It("returns the correct char index for no text at position", func() {
				FPDFText_GetCharIndexAtPos, err := PdfiumInstance.FPDFText_GetCharIndexAtPos(&requests.FPDFText_GetCharIndexAtPos{
					TextPage: textPage,
					X:        2,
					Y:        2,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetCharIndexAtPos).To(Equal(&responses.FPDFText_GetCharIndexAtPos{
					CharIndex: -1,
				}))
			})

			It("returns the correct page text", func() {
				FPDFText_GetText, err := PdfiumInstance.FPDFText_GetText(&requests.FPDFText_GetText{
					TextPage:   textPage,
					StartIndex: 0,
					Count:      57,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetText).To(Equal(&responses.FPDFText_GetText{
					Text: "File: Untitled Document 2 Page 1 of 1\r\nThis is a test PDF",
				}))
			})

			It("returns the correct rect count", func() {
				FPDFText_CountRects, err := PdfiumInstance.FPDFText_CountRects(&requests.FPDFText_CountRects{
					TextPage:   textPage,
					StartIndex: 0,
					Count:      57,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_CountRects).To(Equal(&responses.FPDFText_CountRects{
					Count: 3,
				}))
			})

			It("returns an error when getting a rect without calling FPDFText_CountRects first", func() {
				FPDFText_GetRect, err := PdfiumInstance.FPDFText_GetRect(&requests.FPDFText_GetRect{
					TextPage: textPage,
					Index:    0,
				})
				Expect(err).To(MatchError("could not get rect"))
				Expect(FPDFText_GetRect).To(BeNil())
			})

			It("returns the correct position for rect 0", func() {
				FPDFText_CountRects, err := PdfiumInstance.FPDFText_CountRects(&requests.FPDFText_CountRects{
					TextPage:   textPage,
					StartIndex: 0,
					Count:      57,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_CountRects).To(Equal(&responses.FPDFText_CountRects{
					Count: 3,
				}))

				// This only works if FPDFText_CountRects is called first.
				FPDFText_GetRect, err := PdfiumInstance.FPDFText_GetRect(&requests.FPDFText_GetRect{
					TextPage: textPage,
					Index:    0,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetRect).To(Or(
					Equal(&responses.FPDFText_GetRect{
						Left:   71.9451904296875,
						Top:    797.5192260742188,
						Right:  208.60919189453125,
						Bottom: 789.0162353515625,
					}),
					Equal(&responses.FPDFText_GetRect{
						Left:   71.9451904296875,
						Top:    797.519287109375,
						Right:  208.60919189453125,
						Bottom: 789.0162963867188,
					}),
				))
			})

			It("returns the correct text within bounds", func() {
				FPDFText_GetBoundedText, err := PdfiumInstance.FPDFText_GetBoundedText(&requests.FPDFText_GetBoundedText{
					TextPage: textPage,
					Left:     71.9451904296875,
					Top:      797.5192260742188,
					Right:    208.60919189453125,
					Bottom:   789.0162353515625,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetBoundedText).To(Equal(&responses.FPDFText_GetBoundedText{
					Text: "File: Untitled Document 2 ",
				}))
			})

			When("a search handle is opened", func() {
				When("moving forwards", func() {
					var searchHandle references.FPDF_SCHHANDLE

					BeforeEach(func() {
						FPDFText_FindStart, err := PdfiumInstance.FPDFText_FindStart(&requests.FPDFText_FindStart{
							TextPage: textPage,
							Find:     "e",
						})
						Expect(err).To(BeNil())
						Expect(FPDFText_FindStart).To(Not(BeNil()))

						searchHandle = FPDFText_FindStart.Search
					})

					AfterEach(func() {
						resp, err := PdfiumInstance.FPDFText_FindClose(&requests.FPDFText_FindClose{
							Search: searchHandle,
						})
						Expect(err).To(BeNil())
						Expect(resp).To(Not(BeNil()))
					})

					It("returns the correct amount of characters matches", func() {
						// We need to call FindNext first to move to the first result.
						FPDFText_FindNext, err := PdfiumInstance.FPDFText_FindNext(&requests.FPDFText_FindNext{
							Search: searchHandle,
						})

						Expect(err).To(BeNil())
						Expect(FPDFText_FindNext).To(Equal(&responses.FPDFText_FindNext{
							GotMatch: true,
						}))

						FPDFText_GetSchCount, err := PdfiumInstance.FPDFText_GetSchCount(&requests.FPDFText_GetSchCount{
							Search: searchHandle,
						})
						Expect(err).To(BeNil())
						Expect(FPDFText_GetSchCount).To(Equal(&responses.FPDFText_GetSchCount{
							Count: 1,
						}))
					})

					It("returns the correct char position of the first match", func() {
						// We need to call FindNext first to move to the first result.
						FPDFText_FindNext, err := PdfiumInstance.FPDFText_FindNext(&requests.FPDFText_FindNext{
							Search: searchHandle,
						})

						Expect(err).To(BeNil())
						Expect(FPDFText_FindNext).To(Equal(&responses.FPDFText_FindNext{
							GotMatch: true,
						}))

						FPDFText_GetSchResultIndex, err := PdfiumInstance.FPDFText_GetSchResultIndex(&requests.FPDFText_GetSchResultIndex{
							Search: searchHandle,
						})
						Expect(err).To(BeNil())
						Expect(FPDFText_GetSchResultIndex).To(Equal(&responses.FPDFText_GetSchResultIndex{
							Index: 3,
						}))
					})
				})

				When("moving backwards", func() {
					var searchHandle references.FPDF_SCHHANDLE

					BeforeEach(func() {
						FPDFText_FindStart, err := PdfiumInstance.FPDFText_FindStart(&requests.FPDFText_FindStart{
							TextPage:   textPage,
							Find:       "e",
							StartIndex: -1,
						})
						Expect(err).To(BeNil())
						Expect(FPDFText_FindStart).To(Not(BeNil()))

						searchHandle = FPDFText_FindStart.Search
					})

					AfterEach(func() {
						resp, err := PdfiumInstance.FPDFText_FindClose(&requests.FPDFText_FindClose{
							Search: searchHandle,
						})
						Expect(err).To(BeNil())
						Expect(resp).To(Not(BeNil()))
					})

					It("returns the correct amount of characters matches", func() {
						// We need to call FindNext first to move to the first result.
						FPDFText_FindNext, err := PdfiumInstance.FPDFText_FindPrev(&requests.FPDFText_FindPrev{
							Search: searchHandle,
						})

						Expect(err).To(BeNil())
						Expect(FPDFText_FindNext).To(Equal(&responses.FPDFText_FindPrev{
							GotMatch: true,
						}))

						FPDFText_GetSchCount, err := PdfiumInstance.FPDFText_GetSchCount(&requests.FPDFText_GetSchCount{
							Search: searchHandle,
						})
						Expect(err).To(BeNil())
						Expect(FPDFText_GetSchCount).To(Equal(&responses.FPDFText_GetSchCount{
							Count: 1,
						}))
					})

					It("returns the correct char position of the first match", func() {
						// We need to call FindNext first to move to the first result.
						FPDFText_FindNext, err := PdfiumInstance.FPDFText_FindPrev(&requests.FPDFText_FindPrev{
							Search: searchHandle,
						})

						Expect(err).To(BeNil())
						Expect(FPDFText_FindNext).To(Equal(&responses.FPDFText_FindPrev{
							GotMatch: true,
						}))

						FPDFText_GetSchResultIndex, err := PdfiumInstance.FPDFText_GetSchResultIndex(&requests.FPDFText_GetSchResultIndex{
							Search: searchHandle,
						})
						Expect(err).To(BeNil())
						Expect(FPDFText_GetSchResultIndex).To(Equal(&responses.FPDFText_GetSchResultIndex{
							Index: 50,
						}))
					})
				})
			})
		})
	})

	Context("a PDF file with weblinks", func() {
		var doc references.FPDF_DOCUMENT

		BeforeEach(func() {
			pdfData, err := ioutil.ReadFile(TestDataPath + "/testdata/weblinks.pdf")
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

		When("a text page is opened", func() {
			var textPage references.FPDF_TEXTPAGE

			BeforeEach(func() {
				textPageResp, err := PdfiumInstance.FPDFText_LoadPage(&requests.FPDFText_LoadPage{
					Page: requests.Page{
						ByIndex: &requests.PageByIndex{
							Document: doc,
							Index:    0,
						},
					},
				})
				Expect(err).To(BeNil())
				Expect(textPageResp).To(Not(BeNil()))

				textPage = textPageResp.TextPage
			})

			AfterEach(func() {
				resp, err := PdfiumInstance.FPDFText_ClosePage(&requests.FPDFText_ClosePage{
					TextPage: textPage,
				})
				Expect(err).To(BeNil())
				Expect(resp).To(Not(BeNil()))
			})

			When("a web links are loaded", func() {
				var pageLink references.FPDF_PAGELINK

				BeforeEach(func() {
					FPDFLink_LoadWebLinks, err := PdfiumInstance.FPDFLink_LoadWebLinks(&requests.FPDFLink_LoadWebLinks{
						TextPage: textPage,
					})
					Expect(err).To(BeNil())
					Expect(FPDFLink_LoadWebLinks).To(Not(BeNil()))

					pageLink = FPDFLink_LoadWebLinks.PageLink
				})

				AfterEach(func() {
					resp, err := PdfiumInstance.FPDFLink_CloseWebLinks(&requests.FPDFLink_CloseWebLinks{
						PageLink: pageLink,
					})
					Expect(err).To(BeNil())
					Expect(resp).To(Not(BeNil()))
				})

				It("returns the correct web link count", func() {
					FPDFLink_CountWebLinks, err := PdfiumInstance.FPDFLink_CountWebLinks(&requests.FPDFLink_CountWebLinks{
						PageLink: pageLink,
					})
					Expect(err).To(BeNil())
					Expect(FPDFLink_CountWebLinks).To(Equal(&responses.FPDFLink_CountWebLinks{
						Count: 2,
					}))
				})

				It("returns the correct URL for link 0", func() {
					FPDFLink_GetURL, err := PdfiumInstance.FPDFLink_GetURL(&requests.FPDFLink_GetURL{
						PageLink: pageLink,
						Index:    0,
					})
					Expect(err).To(BeNil())
					Expect(FPDFLink_GetURL).To(Equal(&responses.FPDFLink_GetURL{
						Index: 0,
						URL:   "http://example.com?q=foo",
					}))
				})

				It("returns the correct rect count for link 0", func() {
					FPDFLink_CountRects, err := PdfiumInstance.FPDFLink_CountRects(&requests.FPDFLink_CountRects{
						PageLink: pageLink,
						Index:    0,
					})
					Expect(err).To(BeNil())
					Expect(FPDFLink_CountRects).To(Equal(&responses.FPDFLink_CountRects{
						Index: 0,
						Count: 1,
					}))
				})

				It("returns the correct rect for link 0 and rect index 0", func() {
					FPDFLink_GetRect, err := PdfiumInstance.FPDFLink_GetRect(&requests.FPDFLink_GetRect{
						PageLink:  pageLink,
						Index:     0,
						RectIndex: 0,
					})
					Expect(err).To(BeNil())
					Expect(FPDFLink_GetRect).To(Not(BeNil()))
					Expect(FPDFLink_GetRect.Index).To(Equal(0))
					Expect(FPDFLink_GetRect.RectIndex).To(Equal(0))

					// Rect can be a little different depending on the platform.
					Expect(FPDFLink_GetRect.Left).To(BeNumerically("~", 50, 1))
					Expect(FPDFLink_GetRect.Top).To(BeNumerically("~", 108, 1))
					Expect(FPDFLink_GetRect.Right).To(BeNumerically("~", 187, 1))
					Expect(FPDFLink_GetRect.Bottom).To(BeNumerically("~", 97, 1))
				})

				It("returns the correct text range for link 0 and gets the correct text for it", func() {
					FPDFLink_GetTextRange, err := PdfiumInstance.FPDFLink_GetTextRange(&requests.FPDFLink_GetTextRange{
						PageLink: pageLink,
						Index:    0,
					})
					Expect(err).To(BeNil())
					Expect(FPDFLink_GetTextRange).To(Equal(&responses.FPDFLink_GetTextRange{
						Index:          0,
						StartCharIndex: 35,
						CharCount:      24,
					}))

					FPDFText_GetText, err := PdfiumInstance.FPDFText_GetText(&requests.FPDFText_GetText{
						TextPage:   textPage,
						StartIndex: FPDFLink_GetTextRange.StartCharIndex,
						Count:      FPDFLink_GetTextRange.CharCount,
					})
					Expect(err).To(BeNil())
					Expect(FPDFText_GetText).To(Equal(&responses.FPDFText_GetText{
						Text: "http://example.com?q=foo",
					}))
				})
			})
		})
	})

	Context("a PDF file with generated text", func() {
		var doc references.FPDF_DOCUMENT

		BeforeEach(func() {
			pdfData, err := ioutil.ReadFile(TestDataPath + "/testdata/hello_world.pdf")
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

		When("a text page is opened", func() {
			var textPage references.FPDF_TEXTPAGE

			BeforeEach(func() {
				textPageResp, err := PdfiumInstance.FPDFText_LoadPage(&requests.FPDFText_LoadPage{
					Page: requests.Page{
						ByIndex: &requests.PageByIndex{
							Document: doc,
							Index:    0,
						},
					},
				})
				Expect(err).To(BeNil())
				Expect(textPageResp).To(Not(BeNil()))

				textPage = textPageResp.TextPage
			})

			AfterEach(func() {
				resp, err := PdfiumInstance.FPDFText_ClosePage(&requests.FPDFText_ClosePage{
					TextPage: textPage,
				})
				Expect(err).To(BeNil())
				Expect(resp).To(Not(BeNil()))
			})

			It("returns correctly whether character at index 0 is generated", func() {
				FPDFText_IsGenerated, err := PdfiumInstance.FPDFText_IsGenerated(&requests.FPDFText_IsGenerated{
					TextPage: textPage,
					Index:    0,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_IsGenerated).To(Equal(&responses.FPDFText_IsGenerated{
					Index:       0,
					IsGenerated: false,
				}))
			})

			It("returns correctly whether character at index 6 is generated", func() {
				FPDFText_IsGenerated, err := PdfiumInstance.FPDFText_IsGenerated(&requests.FPDFText_IsGenerated{
					TextPage: textPage,
					Index:    6,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_IsGenerated).To(Equal(&responses.FPDFText_IsGenerated{
					Index:       6,
					IsGenerated: false,
				}))
			})

			It("returns correctly whether character at index 13 is generated", func() {
				FPDFText_IsGenerated, err := PdfiumInstance.FPDFText_IsGenerated(&requests.FPDFText_IsGenerated{
					TextPage: textPage,
					Index:    13,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_IsGenerated).To(Equal(&responses.FPDFText_IsGenerated{
					Index:       13,
					IsGenerated: true,
				}))
			})

			It("returns correctly whether character at index 14 is generated", func() {
				FPDFText_IsGenerated, err := PdfiumInstance.FPDFText_IsGenerated(&requests.FPDFText_IsGenerated{
					TextPage: textPage,
					Index:    14,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_IsGenerated).To(Equal(&responses.FPDFText_IsGenerated{
					Index:       14,
					IsGenerated: true,
				}))
			})

			It("returns an error when giving an invalid character", func() {
				FPDFText_IsGenerated, err := PdfiumInstance.FPDFText_IsGenerated(&requests.FPDFText_IsGenerated{
					TextPage: textPage,
					Index:    9999999,
				})
				Expect(err).To(MatchError("could not get whether text is generated"))
				Expect(FPDFText_IsGenerated).To(BeNil())
			})

			It("returns an error when giving an invalid text page", func() {
				FPDFText_IsGenerated, err := PdfiumInstance.FPDFText_IsGenerated(&requests.FPDFText_IsGenerated{
					TextPage: "dffddf",
					Index:    9999999,
				})
				Expect(err).To(MatchError("could not find textPage handle, perhaps the textPage was already closed or you tried to share textPages between instances"))
				Expect(FPDFText_IsGenerated).To(BeNil())
			})
		})
	})

	Context("a PDF file with hypthen text", func() {
		var doc references.FPDF_DOCUMENT

		BeforeEach(func() {
			pdfData, err := ioutil.ReadFile(TestDataPath + "/testdata/bug_781804.pdf")
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

		When("a text page is opened", func() {
			var textPage references.FPDF_TEXTPAGE

			BeforeEach(func() {
				textPageResp, err := PdfiumInstance.FPDFText_LoadPage(&requests.FPDFText_LoadPage{
					Page: requests.Page{
						ByIndex: &requests.PageByIndex{
							Document: doc,
							Index:    0,
						},
					},
				})
				Expect(err).To(BeNil())
				Expect(textPageResp).To(Not(BeNil()))

				textPage = textPageResp.TextPage
			})

			AfterEach(func() {
				resp, err := PdfiumInstance.FPDFText_ClosePage(&requests.FPDFText_ClosePage{
					TextPage: textPage,
				})
				Expect(err).To(BeNil())
				Expect(resp).To(Not(BeNil()))
			})

			It("returns correctly whether character at index 0 is a hyphen", func() {
				FPDFText_IsHyphen, err := PdfiumInstance.FPDFText_IsHyphen(&requests.FPDFText_IsHyphen{
					TextPage: textPage,
					Index:    0,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_IsHyphen).To(Equal(&responses.FPDFText_IsHyphen{
					Index:    0,
					IsHyphen: false,
				}))
			})

			It("returns correctly whether character at index 6 is a hyphen", func() {
				FPDFText_IsHyphen, err := PdfiumInstance.FPDFText_IsHyphen(&requests.FPDFText_IsHyphen{
					TextPage: textPage,
					Index:    6,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_IsHyphen).To(Equal(&responses.FPDFText_IsHyphen{
					Index:    6,
					IsHyphen: true,
				}))
			})

			It("returns correctly whether character at index 14 is a hyphen", func() {
				FPDFText_IsHyphen, err := PdfiumInstance.FPDFText_IsHyphen(&requests.FPDFText_IsHyphen{
					TextPage: textPage,
					Index:    14,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_IsHyphen).To(Equal(&responses.FPDFText_IsHyphen{
					Index:    14,
					IsHyphen: false,
				}))
			})

			It("returns correctly whether character at index 18 is a hyphen", func() {
				FPDFText_IsHyphen, err := PdfiumInstance.FPDFText_IsHyphen(&requests.FPDFText_IsHyphen{
					TextPage: textPage,
					Index:    18,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_IsHyphen).To(Equal(&responses.FPDFText_IsHyphen{
					Index:    18,
					IsHyphen: false,
				}))
			})

			It("returns an error when giving an invalid character", func() {
				FPDFText_IsHyphen, err := PdfiumInstance.FPDFText_IsHyphen(&requests.FPDFText_IsHyphen{
					TextPage: textPage,
					Index:    9999999,
				})
				Expect(err).To(MatchError("could not get whether text is a hyphen"))
				Expect(FPDFText_IsHyphen).To(BeNil())
			})

			It("returns an error when giving an invalid text page", func() {
				FPDFText_IsHyphen, err := PdfiumInstance.FPDFText_IsHyphen(&requests.FPDFText_IsHyphen{
					TextPage: "dffddf",
					Index:    9999999,
				})
				Expect(err).To(MatchError("could not find textPage handle, perhaps the textPage was already closed or you tried to share textPages between instances"))
				Expect(FPDFText_IsHyphen).To(BeNil())
			})
		})
	})

	Context("a PDF file with invalid unicode", func() {
		var doc references.FPDF_DOCUMENT

		BeforeEach(func() {
			pdfData, err := ioutil.ReadFile(TestDataPath + "/testdata/bug_1388_2.pdf")
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

		When("a text page is opened", func() {
			var textPage references.FPDF_TEXTPAGE

			BeforeEach(func() {
				textPageResp, err := PdfiumInstance.FPDFText_LoadPage(&requests.FPDFText_LoadPage{
					Page: requests.Page{
						ByIndex: &requests.PageByIndex{
							Document: doc,
							Index:    0,
						},
					},
				})
				Expect(err).To(BeNil())
				Expect(textPageResp).To(Not(BeNil()))

				textPage = textPageResp.TextPage
			})

			AfterEach(func() {
				resp, err := PdfiumInstance.FPDFText_ClosePage(&requests.FPDFText_ClosePage{
					TextPage: textPage,
				})
				Expect(err).To(BeNil())
				Expect(resp).To(Not(BeNil()))
			})

			It("returns the correct amount of chars", func() {
				FPDFText_CountChars, err := PdfiumInstance.FPDFText_CountChars(&requests.FPDFText_CountChars{
					TextPage: textPage,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_CountChars).To(Equal(&responses.FPDFText_CountChars{
					Count: 5,
				}))
			})

			It("returns an error when checking for a unicode map error on an invalid char index", func() {
				FPDFText_HasUnicodeMapError, err := PdfiumInstance.FPDFText_HasUnicodeMapError(&requests.FPDFText_HasUnicodeMapError{
					TextPage: textPage,
					Index:    -1,
				})
				Expect(err).To(MatchError("could not get whether text has a unicode map error"))
				Expect(FPDFText_HasUnicodeMapError).To(BeNil())
			})

			It("returns the correct char at position 0 without a unicode map error", func() {
				FPDFText_GetUnicode, err := PdfiumInstance.FPDFText_GetUnicode(&requests.FPDFText_GetUnicode{
					TextPage: textPage,
					Index:    0,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetUnicode).To(Equal(&responses.FPDFText_GetUnicode{
					Index:   0,
					Unicode: 'X',
				}))

				FPDFText_HasUnicodeMapError, err := PdfiumInstance.FPDFText_HasUnicodeMapError(&requests.FPDFText_HasUnicodeMapError{
					TextPage: textPage,
					Index:    0,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_HasUnicodeMapError).To(Equal(&responses.FPDFText_HasUnicodeMapError{
					Index:              0,
					HasUnicodeMapError: false,
				}))
			})

			It("returns the correct char at position 1 without a unicode map error", func() {
				FPDFText_GetUnicode, err := PdfiumInstance.FPDFText_GetUnicode(&requests.FPDFText_GetUnicode{
					TextPage: textPage,
					Index:    1,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetUnicode).To(Equal(&responses.FPDFText_GetUnicode{
					Index:   1,
					Unicode: ' ',
				}))

				FPDFText_HasUnicodeMapError, err := PdfiumInstance.FPDFText_HasUnicodeMapError(&requests.FPDFText_HasUnicodeMapError{
					TextPage: textPage,
					Index:    1,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_HasUnicodeMapError).To(Equal(&responses.FPDFText_HasUnicodeMapError{
					Index:              1,
					HasUnicodeMapError: false,
				}))
			})

			It("returns the correct char at position 2 with a unicode map error", func() {
				FPDFText_GetUnicode, err := PdfiumInstance.FPDFText_GetUnicode(&requests.FPDFText_GetUnicode{
					TextPage: textPage,
					Index:    2,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_GetUnicode).To(Equal(&responses.FPDFText_GetUnicode{
					Index:   2,
					Unicode: 31,
				}))

				FPDFText_HasUnicodeMapError, err := PdfiumInstance.FPDFText_HasUnicodeMapError(&requests.FPDFText_HasUnicodeMapError{
					TextPage: textPage,
					Index:    2,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_HasUnicodeMapError).To(Equal(&responses.FPDFText_HasUnicodeMapError{
					Index:              2,
					HasUnicodeMapError: true,
				}))
			})
		})
	})

	Context("a PDF file with text render mode", func() {
		var doc references.FPDF_DOCUMENT

		BeforeEach(func() {
			pdfData, err := ioutil.ReadFile(TestDataPath + "/testdata/text_render_mode.pdf")
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

		When("a text page is opened", func() {
			var textPage references.FPDF_TEXTPAGE

			BeforeEach(func() {
				textPageResp, err := PdfiumInstance.FPDFText_LoadPage(&requests.FPDFText_LoadPage{
					Page: requests.Page{
						ByIndex: &requests.PageByIndex{
							Document: doc,
							Index:    0,
						},
					},
				})
				Expect(err).To(BeNil())
				Expect(textPageResp).To(Not(BeNil()))

				textPage = textPageResp.TextPage
			})

			AfterEach(func() {
				resp, err := PdfiumInstance.FPDFText_ClosePage(&requests.FPDFText_ClosePage{
					TextPage: textPage,
				})
				Expect(err).To(BeNil())
				Expect(resp).To(Not(BeNil()))
			})

			It("returns the correct character count", func() {
				FPDFText_CountChars, err := PdfiumInstance.FPDFText_CountChars(&requests.FPDFText_CountChars{
					TextPage: textPage,
				})
				Expect(err).To(BeNil())
				Expect(FPDFText_CountChars).To(Equal(&responses.FPDFText_CountChars{
					Count: 12,
				}))
			})

			It("returns an error on invalid text object indices", func() {
				FPDFText_GetTextObject, err := PdfiumInstance.FPDFText_GetTextObject(&requests.FPDFText_GetTextObject{
					TextPage: textPage,
					Index:    -1,
				})
				Expect(err).To(MatchError("could not get object"))
				Expect(FPDFText_GetTextObject).To(BeNil())

				FPDFText_GetTextObject, err = PdfiumInstance.FPDFText_GetTextObject(&requests.FPDFText_GetTextObject{
					TextPage: textPage,
					Index:    314,
				})
				Expect(err).To(MatchError("could not get object"))
				Expect(FPDFText_GetTextObject).To(BeNil())
			})

			When("a text object is opened", func() {
				var textObject references.FPDF_PAGEOBJECT

				BeforeEach(func() {
					FPDFText_GetTextObject, err := PdfiumInstance.FPDFText_GetTextObject(&requests.FPDFText_GetTextObject{
						TextPage: textPage,
						Index:    0,
					})
					Expect(err).To(BeNil())
					Expect(FPDFText_GetTextObject).To(Not(BeNil()))

					textObject = FPDFText_GetTextObject.TextObject
				})

				It("returns the correct object type", func() {
					FPDFPageObj_GetType, err := PdfiumInstance.FPDFPageObj_GetType(&requests.FPDFPageObj_GetType{
						PageObject: textObject,
					})
					Expect(err).To(BeNil())
					Expect(FPDFPageObj_GetType).To(Equal(&responses.FPDFPageObj_GetType{
						Type: enums.FPDF_PAGEOBJ_TEXT,
					}))
				})
			})

			When("another text object is opened", func() {
				var textObject references.FPDF_PAGEOBJECT

				BeforeEach(func() {
					FPDFText_GetTextObject, err := PdfiumInstance.FPDFText_GetTextObject(&requests.FPDFText_GetTextObject{
						TextPage: textPage,
						Index:    7,
					})
					Expect(err).To(BeNil())
					Expect(FPDFText_GetTextObject).To(Not(BeNil()))

					textObject = FPDFText_GetTextObject.TextObject
				})

				It("returns the correct object type", func() {
					FPDFPageObj_GetType, err := PdfiumInstance.FPDFPageObj_GetType(&requests.FPDFPageObj_GetType{
						PageObject: textObject,
					})
					Expect(err).To(BeNil())
					Expect(FPDFPageObj_GetType).To(Equal(&responses.FPDFPageObj_GetType{
						Type: enums.FPDF_PAGEOBJ_TEXT,
					}))
				})

				It("returns the correct text render mode", func() {
					FPDFTextObj_GetTextRenderMode, err := PdfiumInstance.FPDFTextObj_GetTextRenderMode(&requests.FPDFTextObj_GetTextRenderMode{
						PageObject: textObject,
					})
					Expect(err).To(BeNil())
					Expect(FPDFTextObj_GetTextRenderMode).To(Equal(&responses.FPDFTextObj_GetTextRenderMode{
						TextRenderMode: enums.FPDF_TEXTRENDERMODE_STROKE,
					}))
				})
			})
		})
	})
})
