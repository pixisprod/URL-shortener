package hash

type Generator interface {
	Generate() (string, error)
}
