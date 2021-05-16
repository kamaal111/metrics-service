package models

type MetricModel struct {
	tableName struct{}       `pg:"metric"`
	ID        int            `pg:"id,pk"`
	Payload   MetricsPayload `pg:"payload"`
}

// MetricsPayload ...
type MetricsPayload struct {
	LocationActivityMetrics          LocationActivityMetrics          `json:"locationActivityMetrics,omitempty"`
	CellularConditionMetrics         CellularConditionMetrics         `json:"cellularConditionMetrics,omitempty"`
	MetaData                         MetaData                         `json:"metaData,omitempty"`
	GPUMetrics                       GPUMetrics                       `json:"gpuMetrics,omitempty"`
	MemoryMetrics                    MemoryMetrics                    `json:"memoryMetrics,omitempty"`
	SignpostMetrics                  []SignpostMetrics                `json:"signpostMetrics,omitempty"`
	DisplayMetrics                   DisplayMetrics                   `json:"displayMetrics,omitempty"`
	CPUMetrics                       CPUMetrics                       `json:"cpuMetrics,omitempty"`
	NetworkTransferMetrics           NetworkTransferMetrics           `json:"networkTransferMetrics,omitempty"`
	DiskIOMetrics                    DiskIOMetrics                    `json:"diskIOMetrics,omitempty"`
	ApplicationLaunchMetrics         ApplicationLaunchMetrics         `json:"applicationLaunchMetrics,omitempty"`
	ApplicationTimeMetrics           ApplicationTimeMetrics           `json:"applicationTimeMetrics,omitempty"`
	TimeStampEnd                     string                           `json:"timeStampEnd,omitempty"`
	ApplicationResponsivenessMetrics ApplicationResponsivenessMetrics `json:"applicationResponsivenessMetrics,omitempty"`
	AppVersion                       string                           `json:"appVersion,omitempty"`
	TimeStampBegin                   string                           `json:"timeStampBegin,omitempty"`
}

type LocationActivityMetrics struct {
	CumulativeBestAccuracyForNavigationTime string `json:"cumulativeBestAccuracyForNavigationTime,omitempty"`
	CumulativeBestAccuracyTime              string `json:"cumulativeBestAccuracyTime,omitempty"`
	CumulativeHundredMetersAccuracyTime     string `json:"cumulativeHundredMetersAccuracyTime,omitempty"`
	CumulativeNearestTenMetersAccuracyTime  string `json:"cumulativeNearestTenMetersAccuracyTime,omitempty"`
	CumulativeKilometerAccuracyTime         string `json:"cumulativeKilometerAccuracyTime,omitempty"`
	CumulativeThreeKilometersAccuracyTime   string `json:"cumulativeThreeKilometersAccuracyTime,omitempty"`
}

type CellularConditionMetrics struct {
	CellConditionTime CellConditionTime `json:"cellConditionTime,omitempty"`
}

type CellConditionTime struct {
	HistogramNumBuckets int64                     `json:"histogramNumBuckets,omitempty"`
	HistogramValue      map[string]HistogramValue `json:"histogramValue,omitempty"`
}

type HistogramValue struct {
	BucketCount float64 `json:"bucketCount,omitempty"`
	BucketStart string  `json:"bucketStart,omitempty"`
	BucketEnd   string  `json:"bucketEnd,omitempty"`
}

type MetaData struct {
	AppBuildVersion string `json:"appBuildVersion,omitempty"`
	OSVersion       string `json:"osVersion,omitempty"`
	RegionFormat    string `json:"regionFormat,omitempty"`
	DeviceType      string `json:"deviceType,omitempty"`
}

type GPUMetrics struct {
	CumulativeGPUTime string `json:"cumulativeGPUTime,omitempty"`
}

type MemoryMetrics struct {
	PeakMemoryUsage        string                 `json:"peakMemoryUsage,omitempty"`
	AverageSuspendedMemory AverageSuspendedMemory `json:"averageSuspendedMemory,omitempty"`
}

type AverageSuspendedMemory struct {
	AverageValue      string  `json:"averageValue,omitempty"`
	StandardDeviation float64 `json:"standardDeviation,omitempty"`
	SampleCount       float64 `json:"sampleCount,omitempty"`
}

type SignpostMetrics struct {
	SignpostIntervalData SignpostIntervalData `json:"signpostIntervalData,omitempty"`
	SignpostCategory     string               `json:"signpostCategory,omitempty"`
	SignpostName         string               `json:"signpostName,omitempty"`
	TotalSignpostCount   float64              `json:"totalSignpostCount,omitempty"`
}

type SignpostIntervalData struct {
	HistogrammedSignpostDurations   HistogrammedSignpostDurations `json:"histogrammedSignpostDurations,omitempty"`
	SignpostCumulativeCPUTime       string                        `json:"signpostCumulativeCPUTime,omitempty"`
	SignpostAverageMemory           string                        `json:"signpostAverageMemory,omitempty"`
	SignpostCumulativeLogicalWrites string                        `json:"signpostCumulativeLogicalWrites,omitempty"`
}

type HistogrammedSignpostDurations struct {
	HistogramNumBuckets float64                   `json:"histogramNumBuckets,omitempty"`
	HistogramValue      map[string]HistogramValue `json:"histogramValue,omitempty"`
}

type DisplayMetrics struct {
	AveragePixelLuminance AveragePixelLuminance `json:"averagePixelLuminance,omitempty"`
}

type AveragePixelLuminance struct {
	AverageValue      string  `json:"averageValue,omitempty"`
	StandardDeviation float64 `json:"standardDeviation,omitempty"`
	SampleCount       float64 `json:"sampleCount,omitempty"`
}

type CPUMetrics struct {
	CumulativeCPUTime string `json:"cumulativeCPUTime,omitempty"`
}

type NetworkTransferMetrics struct {
	CumulativeCellularDownload string `json:"cumulativeCellularDownload,omitempty"`
	CumulativeWifiDownload     string `json:"cumulativeWifiDownload,omitempty"`
	CumulativeCellularUpload   string `json:"cumulativeCellularUpload,omitempty"`
	CumulativeWifiUpload       string `json:"cumulativeWifiUpload,omitempty"`
}

type DiskIOMetrics struct {
	CumulativeLogicalWrites string `json:"cumulativeLogicalWrites,omitempty"`
}

type ApplicationLaunchMetrics struct {
	HistogrammedTimeToFirstDrawKey HistogrammedTimeToFirstDrawKey `json:"histogrammedTimeToFirstDrawKey,omitempty"`
	HistogrammedResumeTime         HistogrammedResumeTime         `json:"histogrammedResumeTime,omitempty"`
}

type HistogrammedTimeToFirstDrawKey struct {
	HistogramNumBuckets int64                     `json:"histogramNumBuckets,omitempty"`
	HistogramValue      map[string]HistogramValue `json:"histogramValue,omitempty"`
}

type HistogrammedResumeTime struct {
	HistogramNumBuckets int64                     `json:"histogramNumBuckets,omitempty"`
	HistogramValue      map[string]HistogramValue `json:"histogramValue,omitempty"`
}

type ApplicationTimeMetrics struct {
	CumulativeForegroundTime         string `json:"cumulativeForegroundTime,omitempty"`
	CumulativeBackgroundTime         string `json:"cumulativeBackgroundTime,omitempty"`
	CumulativeBackgroundAudioTime    string `json:"cumulativeBackgroundAudioTime,omitempty"`
	CumulativeBackgroundLocationTime string `json:"cumulativeBackgroundLocationTime,omitempty"`
}

type ApplicationResponsivenessMetrics struct {
	HistogrammedAppHangTime HistogrammedAppHangTime `json:"histogrammedAppHangTime,omitempty"`
}

type HistogrammedAppHangTime struct {
	HistogramNumBuckets int64                     `json:"histogramNumBuckets,omitempty"`
	HistogramValue      map[string]HistogramValue `json:"histogramValue,omitempty"`
}
