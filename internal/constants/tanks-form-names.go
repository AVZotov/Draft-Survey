package constants

const (
	WtankType              = "w_tank_type"
	WtankName              = "w_tank_name"
	VolumeCalibrationType1 = "standard_volume_by_trim" //common aproach at most marine vessels
	VolumeCalibrationType2 = "sounding_correction"
	VolumeCalibrationType3 = "volume_correction"
	// Tank identification
	TankName = "tank_name"
	TankID   = "tank_id"
	// Tank measurements — entered manually
	TankSounding = "tank_sounding" // measured by ullage/sounding tape, m
	TankDensity  = "tank_density"  // sea water density, t/m³ (BWT only)
	// Tank volume — calculated or entered manually
	TankVolume = "tank_volume" // m³
	//Tank Water weight - calculated
	TankWeight = "tank_weight" //MT
	// Calibration table type selector
	TankCalibTableType = "tank_calib_table_type"
	// Calibration table boundaries — entered from vessel's calibration tables
	CalibTrimLow   = "calib_ttl" // lower trim boundary (TTL), m
	CalibTrimUpper = "calib_ttu" // upper trim boundary (TTU), m
	// Calibration row fields — entered from vessel's calibration tables
	CalibRowSounding     = "calib_row_sounding"       // sounding value for this row, m
	CalibRowVolTrimLow   = "calib_row_vol_trim_low"   // volume at TTL, m³
	CalibRowVolTrimUpper = "calib_row_vol_trim_upper" // volume at TTU, m³
)
