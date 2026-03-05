package types

type ReadingMethod string

const (
	ReadingMethodDirect    ReadingMethod = "direct"
	ReadingMethodWaterline ReadingMethod = "waterline"
)

type Mark struct {
	Value  *float64      `json:"value"`
	Method ReadingMethod `json:"method"`
}

type Marks struct {
	FwdPort      Mark `json:"fwd_port"`
	FwdStarboard Mark `json:"fwd_starboard"`
	MidPort      Mark `json:"mid_port"`
	MidStarboard Mark `json:"mid_starboard"`
	AftPort      Mark `json:"aft_port"`
	AftStarboard Mark `json:"aft_starboard"`
}
