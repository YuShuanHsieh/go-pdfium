package responses

import (
	"github.com/YuShuanHsieh/go-pdfium/references"
)

type GetBookmarksBookmark struct {
	Title      string
	Reference  references.FPDF_BOOKMARK
	ActionInfo *ActionInfo
	DestInfo   *DestInfo
	Children   []GetBookmarksBookmark
}

type GetBookmarks struct {
	Bookmarks []GetBookmarksBookmark
}
