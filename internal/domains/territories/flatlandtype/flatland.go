package flatlandtype

type (
	Type struct {
		name string
	}
)

func (v Type) Name() string {
	return string(v.name)
}

func (v Type) CanBuildCity() bool {
	return true
}

func New() *Type {
	return &Type{name: "flat_land"}
}
