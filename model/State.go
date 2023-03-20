package model

type State struct {
	FileID   uint   `json:"file_id"`
	WriterID int64  `json:"writer_id"`
	FileKind string `json:"file_kind"`
}
