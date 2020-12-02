package generator

type value interface {
	getCurrentValue() string
	generateValue()
	init(i string)
}
