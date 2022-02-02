package references

// All the references are just an alias for string, but it's useful to keep the alias,
// so that we can change it later. They all should contain a UUID.

// FPDF_DOCUMENT is an internal reference to a C.FPDF_DOCUMENT handle.
type FPDF_DOCUMENT string

// FPDF_PAGE is an internal reference to a C.FPDF_PAGE handle.
type FPDF_PAGE string

// FPDF_BOOKMARK is an internal reference to a C.FPDF_BOOKMARK handle.
type FPDF_BOOKMARK string

// FPDF_DEST is an internal reference to a C.FPDF_DEST handle.
type FPDF_DEST string

// FPDF_ACTION is an internal reference to a C.FPDF_ACTION handle.
type FPDF_ACTION string

// FPDF_LINK is an internal reference to a C.FPDF_LINK handle.
type FPDF_LINK string

// FPDF_PAGELINK is an internal reference to a C.FPDF_PAGELINK handle.
type FPDF_PAGELINK string

// FPDF_SCHHANDLE is an internal reference to a C.FPDF_SCHHANDLE handle.
type FPDF_SCHHANDLE string

// FPDF_BITMAP is an internal reference to a C.FPDF_BITMAP handle.
type FPDF_BITMAP string

// FPDF_TEXTPAGE is an internal reference to a C.FPDF_TEXTPAGE handle.
type FPDF_TEXTPAGE string

// FPDF_PAGERANGE is an internal reference to a C.FPDF_PAGERANGE handle.
type FPDF_PAGERANGE string

// FPDF_PAGEOBJECT is an internal reference to a C.FPDF_PAGEOBJECT handle.
type FPDF_PAGEOBJECT string

// FPDF_CLIPPATH is an internal reference to a C.FPDF_CLIPPATH handle.
type FPDF_CLIPPATH string

// FPDF_FORMHANDLE is an internal reference to a C.FPDF_FORMHANDLE handle.
type FPDF_FORMHANDLE string
