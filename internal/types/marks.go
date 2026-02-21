package types

type ReadingMethod string

const (
	ReadingMethodDirect    ReadingMethod = "direct"
	ReadingMethodWaterline ReadingMethod = "waterline"
)

type Mark struct {
	Value  float64
	Method ReadingMethod
}

type Marks struct {
	FwdPort      Mark
	FwdStarboard Mark
	MidPort      Mark
	MidStarboard Mark
	AftPort      Mark
	AftStarboard Mark
}
