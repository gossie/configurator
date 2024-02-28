package value

type Value interface {
	Subsumes(aValue Value) bool
	subsumedByRange(other intRange) bool
	subsumedBySet(other intValues) bool
	Sect(other Value) Value
	sectWithRange(other intRange) Value
	sectWithSet(other intValues) Value
	Diff(other Value) Value
	diffFromRange(other intRange) Value
	diffFromSet(other intValues) Value
	Final() bool
	String() string
}
