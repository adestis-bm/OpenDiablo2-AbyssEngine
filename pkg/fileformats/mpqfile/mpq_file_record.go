package mpqfile

// MpqFileRecord represents a file record in an MPQ
type MpqFileRecord struct {
	MpqFile          string
	IsPatch          bool
	UnpatchedMpqFile string
}
