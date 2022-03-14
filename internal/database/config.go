package database

var (
	kind Kind = KindDummy
)

func GetKind() Kind {
	return kind
}
