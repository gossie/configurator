package value

func Sect(v1, v2 Value) Value {
	return v1.sect(v2)
}

type Value interface {
	Subsumes(aValue Value) bool
	subsumedByRange(other intRange) bool
	subsumedBySet(other intValues) bool
	sect(other Value) Value
	sectWithRange(other intRange) Value
	sectWithSet(other intValues) Value
	Final() bool
	String() string
}
