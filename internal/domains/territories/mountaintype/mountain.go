package mountaintype

type (
	Type struct {
		name string
	}
)

func (v Type) Name() string {
	return string(v.name)
}


func (v Type) CanBuildCity() bool {
	return false
}

func New() *Type {
	return &Type{name: "mountain"}
}
