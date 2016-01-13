package backup

type Archiver interface {
	Archive(src, dest string) error
}
