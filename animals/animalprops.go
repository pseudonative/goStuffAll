package animals

type Speaker interface {
	Speaks() string
}

func Perform(s Speaker) string {
	return s.Speaks()
}

type Lion struct{}

func (s Lion) Speaks() string {
	return "roars"
}

type Cat struct{}

func (s Cat) Speaks() string {
	return "meows"
}

type Human struct{}

func (s Human) Speaks() string {
	return "talks"
}
