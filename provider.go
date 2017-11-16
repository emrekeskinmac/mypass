package mypass

type Provider interface {
	Save(username, name, hex string) error
	Find(username, name string) (hex []string, err error)
	Delete(username, name string) (hex []string, err error)
}
