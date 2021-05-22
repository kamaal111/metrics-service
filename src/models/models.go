package models

import "github.com/go-pg/pg/v10"

type AppsTable struct {
	tableName        struct{} `pg:"apps"`
	ID               int      `pg:"id,pk"`
	BundleIdentifier string   `pg:"bundle_identifier,unique"`
	AccessToken      string   `pg:"access_token"`
}

func (app *AppsTable) Save(pgDB *pg.DB) error {
	_, err := pgDB.Model(app).Insert()
	return err
}

func (app *AppsTable) GetMetrics(pgDB *pg.DB) ([]MetricsTable, error) {
	var metrics []MetricsTable
	err := pgDB.Model(&metrics).Where("app_id = ?", app.ID).Select()
	return metrics, err
}

type MetricsTable struct {
	tableName       struct{}          `pg:"metrics"`
	ID              int               `pg:"id,pk" json:"id"`
	AppVersion      string            `pg:"app_version" json:"app_version"`
	AppBuildVersion string            `pg:"app_build_version" json:"app_build_version"`
	Payload         CollectionMetrics `pg:"payload" json:"payload"`
	AppID           int               `pg:"app_id" json:"app_id"`
}

func (metric *MetricsTable) Save(pgDB *pg.DB) error {
	_, err := pgDB.Model(metric).Insert()
	return err
}

type CollectionMetrics struct {
	LocationActivityMetrics          locationActivityMetrics          `json:"locationActivityMetrics,omitempty"`
	CellularConditionMetrics         cellularConditionMetrics         `json:"cellularConditionMetrics,omitempty"`
	MetaData                         MetaData                         `json:"metaData,omitempty"`
	GPUMetrics                       gpuMetrics                       `json:"gpuMetrics,omitempty"`
	MemoryMetrics                    memoryMetrics                    `json:"memoryMetrics,omitempty"`
	SignpostMetrics                  []signpostMetrics                `json:"signpostMetrics,omitempty"`
	DisplayMetrics                   displayMetrics                   `json:"displayMetrics,omitempty"`
	CPUMetrics                       cpuMetrics                       `json:"cpuMetrics,omitempty"`
	NetworkTransferMetrics           networkTransferMetrics           `json:"networkTransferMetrics,omitempty"`
	DiskIOMetrics                    diskIOMetrics                    `json:"diskIOMetrics,omitempty"`
	ApplicationLaunchMetrics         applicationLaunchMetrics         `json:"applicationLaunchMetrics,omitempty"`
	ApplicationTimeMetrics           applicationTimeMetrics           `json:"applicationTimeMetrics,omitempty"`
	TimeStampEnd                     string                           `json:"timeStampEnd,omitempty"`
	ApplicationResponsivenessMetrics applicationResponsivenessMetrics `json:"applicationResponsivenessMetrics,omitempty"`
	AppVersion                       string                           `json:"appVersion,omitempty"`
	TimeStampBegin                   string                           `json:"timeStampBegin,omitempty"`
}

type locationActivityMetrics struct {
	CumulativeBestAccuracyForNavigationTime string `json:"cumulativeBestAccuracyForNavigationTime,omitempty"`
	CumulativeBestAccuracyTime              string `json:"cumulativeBestAccuracyTime,omitempty"`
	CumulativeHundredMetersAccuracyTime     string `json:"cumulativeHundredMetersAccuracyTime,omitempty"`
	CumulativeNearestTenMetersAccuracyTime  string `json:"cumulativeNearestTenMetersAccuracyTime,omitempty"`
	CumulativeKilometerAccuracyTime         string `json:"cumulativeKilometerAccuracyTime,omitempty"`
	CumulativeThreeKilometersAccuracyTime   string `json:"cumulativeThreeKilometersAccuracyTime,omitempty"`
}

type cellularConditionMetrics struct {
	CellConditionTime cellConditionTime `json:"cellConditionTime,omitempty"`
}

type cellConditionTime struct {
	HistogramNumBuckets int64                     `json:"histogramNumBuckets,omitempty"`
	HistogramValue      map[string]histogramValue `json:"histogramValue,omitempty"`
}

type histogramValue struct {
	BucketCount float64 `json:"bucketCount,omitempty"`
	BucketStart string  `json:"bucketStart,omitempty"`
	BucketEnd   string  `json:"bucketEnd,omitempty"`
}

type MetaData struct {
	AppBuildVersion      string `json:"appBuildVersion,omitempty"`
	OSVersion            string `json:"osVersion,omitempty"`
	RegionFormat         string `json:"regionFormat,omitempty"`
	DeviceType           string `json:"deviceType,omitempty"`
	PlatformArchitecture string `json:"platformArchitecture,omitempty"`
}

type gpuMetrics struct {
	CumulativeGPUTime string `json:"cumulativeGPUTime,omitempty"`
}

type memoryMetrics struct {
	PeakMemoryUsage        string                 `json:"peakMemoryUsage,omitempty"`
	AverageSuspendedMemory averageSuspendedMemory `json:"averageSuspendedMemory,omitempty"`
}

type averageSuspendedMemory struct {
	AverageValue      string  `json:"averageValue,omitempty"`
	StandardDeviation float64 `json:"standardDeviation,omitempty"`
	SampleCount       float64 `json:"sampleCount,omitempty"`
}

type signpostMetrics struct {
	SignpostIntervalData signpostIntervalData `json:"signpostIntervalData,omitempty"`
	SignpostCategory     string               `json:"signpostCategory,omitempty"`
	SignpostName         string               `json:"signpostName,omitempty"`
	TotalSignpostCount   float64              `json:"totalSignpostCount,omitempty"`
}

type signpostIntervalData struct {
	HistogrammedSignpostDurations   histogrammedSignpostDurations `json:"histogrammedSignpostDurations,omitempty"`
	SignpostCumulativeCPUTime       string                        `json:"signpostCumulativeCPUTime,omitempty"`
	SignpostAverageMemory           string                        `json:"signpostAverageMemory,omitempty"`
	SignpostCumulativeLogicalWrites string                        `json:"signpostCumulativeLogicalWrites,omitempty"`
}

type histogrammedSignpostDurations struct {
	HistogramNumBuckets float64                   `json:"histogramNumBuckets,omitempty"`
	HistogramValue      map[string]histogramValue `json:"histogramValue,omitempty"`
}

type displayMetrics struct {
	AveragePixelLuminance averagePixelLuminance `json:"averagePixelLuminance,omitempty"`
}

type averagePixelLuminance struct {
	AverageValue      string  `json:"averageValue,omitempty"`
	StandardDeviation float64 `json:"standardDeviation,omitempty"`
	SampleCount       float64 `json:"sampleCount,omitempty"`
}

type cpuMetrics struct {
	CumulativeCPUTime string `json:"cumulativeCPUTime,omitempty"`
}

type networkTransferMetrics struct {
	CumulativeCellularDownload string `json:"cumulativeCellularDownload,omitempty"`
	CumulativeWifiDownload     string `json:"cumulativeWifiDownload,omitempty"`
	CumulativeCellularUpload   string `json:"cumulativeCellularUpload,omitempty"`
	CumulativeWifiUpload       string `json:"cumulativeWifiUpload,omitempty"`
}

type diskIOMetrics struct {
	CumulativeLogicalWrites string `json:"cumulativeLogicalWrites,omitempty"`
}

type applicationLaunchMetrics struct {
	HistogrammedTimeToFirstDrawKey histogrammedTimeToFirstDrawKey `json:"histogrammedTimeToFirstDrawKey,omitempty"`
	HistogrammedResumeTime         histogrammedResumeTime         `json:"histogrammedResumeTime,omitempty"`
}

type histogrammedTimeToFirstDrawKey struct {
	HistogramNumBuckets int64                     `json:"histogramNumBuckets,omitempty"`
	HistogramValue      map[string]histogramValue `json:"histogramValue,omitempty"`
}

type histogrammedResumeTime struct {
	HistogramNumBuckets int64                     `json:"histogramNumBuckets,omitempty"`
	HistogramValue      map[string]histogramValue `json:"histogramValue,omitempty"`
}

type applicationTimeMetrics struct {
	CumulativeForegroundTime         string `json:"cumulativeForegroundTime,omitempty"`
	CumulativeBackgroundTime         string `json:"cumulativeBackgroundTime,omitempty"`
	CumulativeBackgroundAudioTime    string `json:"cumulativeBackgroundAudioTime,omitempty"`
	CumulativeBackgroundLocationTime string `json:"cumulativeBackgroundLocationTime,omitempty"`
}

type applicationResponsivenessMetrics struct {
	HistogrammedAppHangTime histogrammedAppHangTime `json:"histogrammedAppHangTime,omitempty"`
}

type histogrammedAppHangTime struct {
	HistogramNumBuckets int64                     `json:"histogramNumBuckets,omitempty"`
	HistogramValue      map[string]histogramValue `json:"histogramValue,omitempty"`
}
