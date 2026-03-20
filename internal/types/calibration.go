package types

import "github.com/AVZotov/draft-survey/internal/constants"

type CalibrationTableType string

const (
	CalibrationTypeVolumeByTrim       CalibrationTableType = constants.VolumeCalibrationType1
	CalibrationTypeSoundingCorrection CalibrationTableType = constants.VolumeCalibrationType2
	CalibrationTypeVolumeCorrection   CalibrationTableType = constants.VolumeCalibrationType3
)

type CalibrationRow struct {
	Sounding        *float64 `json:"calib_row_sounding"`
	VolumeTrimLow   *float64 `json:"calib_row_vol_trim_low"`
	VolumeTrimUpper *float64 `json:"calib_row_vol_trim_upper"`
}

type VolumeCalibrationData struct {
	TableType      CalibrationTableType `json:"tank_calib_table_type"`
	TableTrimLow   *float64             `json:"calib_ttl"`
	TableTrimUpper *float64             `json:"calib_ttu"`
	TrimRows       []CalibrationRow     `json:"trim_rows"`
	ListRows       []CalibrationRow     `json:"list_rows"`
	VolumeRows     []CalibrationRow     `json:"volume_rows"` // Table 2 for Type2 and Type3
}
