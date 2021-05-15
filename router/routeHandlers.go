package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-pg/pg/v10"
)

func collectHandler(w http.ResponseWriter, r *http.Request, pgDB *pg.DB) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		errorHandler(w, err.Error(), 500)
		return
	}
	var payload = map[string]string{}
	err = json.Unmarshal([]byte(body), &payload)
	if err != nil {
		errorHandler(w, err.Error(), 500)
		return
	}
	fmt.Println(payload)

	response := struct {
		Success bool `json:"success"`
	}{
		Success: true,
	}
	output, err := json.Marshal(response)
	if err != nil {
		errorHandler(w, err.Error(), 400)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Hello string `json:"Hello"`
	}{
		Hello: "Welcome",
	}
	output, err := json.Marshal(payload)
	if err != nil {
		errorHandler(w, err.Error(), 400)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

// {
// 	"locationActivityMetrics": {
// 	  "cumulativeBestAccuracyForNavigationTime": "20 sec",
// 	  "cumulativeBestAccuracyTime": "30 sec",
// 	  "cumulativeHundredMetersAccuracyTime": "30 sec",
// 	  "cumulativeNearestTenMetersAccuracyTime": "30 sec",
// 	  "cumulativeKilometerAccuracyTime": "20 sec",
// 	  "cumulativeThreeKilometersAccuracyTime": "20 sec"
// 	},
// 	"cellularConditionMetrics": {
// 	  "cellConditionTime": {
// 		"histogramNumBuckets": 3,
// 		"histogramValue": {
// 		  "0": {
// 			"bucketCount": 20,
// 			"bucketStart": "1 bars",
// 			"bucketEnd": "1 bars"
// 		  },
// 		  "1": {
// 			"bucketCount": 30,
// 			"bucketStart": "2 bars",
// 			"bucketEnd": "2 bars"
// 		  },
// 		  "2": {
// 			"bucketCount": 50,
// 			"bucketStart": "3 bars",
// 			"bucketEnd": "3 bars"
// 		  }
// 		}
// 	  }
// 	},
// 	"metaData": {
// 	  "appBuildVersion": "0",
// 	  "osVersion": "iPhone OS 13.1.3 (17A878)",
// 	  "regionFormat": "US",
// 	  "deviceType": "iPhone9,2"
// 	},
// 	"gpuMetrics": {
// 	  "cumulativeGPUTime": "20 sec"
// 	},
// 	"memoryMetrics": {
// 	  "peakMemoryUsage": "200,000 kB",
// 	  "averageSuspendedMemory": {
// 		"averageValue": "100,000 kB",
// 		"standardDeviation": 0,
// 		"sampleCount": 500
// 	  }
// 	},
// 	"signpostMetrics": [
// 	  {
// 		"signpostIntervalData": {
// 		  "histogrammedSignpostDurations": {
// 			"histogramNumBuckets": 3,
// 			"histogramValue": {
// 			  "0": {
// 				"bucketCount": 50,
// 				"bucketStart": "0 ms",
// 				"bucketEnd": "100 ms"
// 			  },
// 			  "1": {
// 				"bucketCount": 60,
// 				"bucketStart": "100 ms",
// 				"bucketEnd": "400 ms"
// 			  },
// 			  "2": {
// 				"bucketCount": 30,
// 				"bucketStart": "400 ms",
// 				"bucketEnd": "700 ms"
// 			  }
// 			}
// 		  },
// 		  "signpostCumulativeCPUTime": "30,000 ms",
// 		  "signpostAverageMemory": "100,000 kB",
// 		  "signpostCumulativeLogicalWrites": "600 kB"
// 		},
// 		"signpostCategory": "TestSignpostCategory1",
// 		"signpostName": "TestSignpostName1",
// 		"totalSignpostCount": 30
// 	  },
// 	  {
// 		"signpostIntervalData": {
// 		  "histogrammedSignpostDurations": {
// 			"histogramNumBuckets": 3,
// 			"histogramValue": {
// 			  "0": {
// 				"bucketCount": 60,
// 				"bucketStart": "0 ms",
// 				"bucketEnd": "200 ms"
// 			  },
// 			  "1": {
// 				"bucketCount": 70,
// 				"bucketStart": "201 ms",
// 				"bucketEnd": "300 ms"
// 			  },
// 			  "2": {
// 				"bucketCount": 80,
// 				"bucketStart": "301 ms",
// 				"bucketEnd": "500 ms"
// 			  }
// 			}
// 		  },
// 		  "signpostCumulativeCPUTime": "50,000 ms",
// 		  "signpostAverageMemory": "60,000 kB",
// 		  "signpostCumulativeLogicalWrites": "700 kB"
// 		},
// 		"signpostCategory": "TestSignpostCategory2",
// 		"signpostName": "TestSignpostName2",
// 		"totalSignpostCount": 40
// 	  }
// 	],
// 	"displayMetrics": {
// 	  "averagePixelLuminance": {
// 		"averageValue": "50 apl",
// 		"standardDeviation": 0,
// 		"sampleCount": 500
// 	  }
// 	},
// 	"cpuMetrics": {
// 	  "cumulativeCPUTime": "100 sec"
// 	},
// 	"networkTransferMetrics": {
// 	  "cumulativeCellularDownload": "80,000 kB",
// 	  "cumulativeWifiDownload": "60,000 kB",
// 	  "cumulativeCellularUpload": "70,000 kB",
// 	  "cumulativeWifiUpload": "50,000 kB"
// 	},
// 	"diskIOMetrics": {
// 	  "cumulativeLogicalWrites": "1,300 kB"
// 	},
// 	"applicationLaunchMetrics": {
// 	  "histogrammedTimeToFirstDrawKey": {
// 		"histogramNumBuckets": 3,
// 		"histogramValue": {
// 		  "0": {
// 			"bucketCount": 50,
// 			"bucketStart": "1,000 ms",
// 			"bucketEnd": "1,010 ms"
// 		  },
// 		  "1": {
// 			"bucketCount": 60,
// 			"bucketStart": "2,000 ms",
// 			"bucketEnd": "2,010 ms"
// 		  },
// 		  "2": {
// 			"bucketCount": 30,
// 			"bucketStart": "3,000 ms",
// 			"bucketEnd": "3,010 ms"
// 		  }
// 		}
// 	  },
// 	  "histogrammedResumeTime": {
// 		"histogramNumBuckets": 3,
// 		"histogramValue": {
// 		  "0": {
// 			"bucketCount": 60,
// 			"bucketStart": "200 ms",
// 			"bucketEnd": "210 ms"
// 		  },
// 		  "1": {
// 			"bucketCount": 70,
// 			"bucketStart": "300 ms",
// 			"bucketEnd": "310 ms"
// 		  },
// 		  "2": {
// 			"bucketCount": 80,
// 			"bucketStart": "500 ms",
// 			"bucketEnd": "510 ms"
// 		  }
// 		}
// 	  }
// 	},
// 	"applicationTimeMetrics": {
// 	  "cumulativeForegroundTime": "700 sec",
// 	  "cumulativeBackgroundTime": "40 sec",
// 	  "cumulativeBackgroundAudioTime": "30 sec",
// 	  "cumulativeBackgroundLocationTime": "30 sec"
// 	},
// 	"timeStampEnd": "2019-10-22 06:59:00 +0000",
// 	"applicationResponsivenessMetrics": {
// 	  "histogrammedAppHangTime": {
// 		"histogramNumBuckets": 3,
// 		"histogramValue": {
// 		  "0": {
// 			"bucketCount": 50,
// 			"bucketStart": "0 ms",
// 			"bucketEnd": "100 ms"
// 		  },
// 		  "1": {
// 			"bucketCount": 60,
// 			"bucketStart": "100 ms",
// 			"bucketEnd": "400 ms"
// 		  },
// 		  "2": {
// 			"bucketCount": 30,
// 			"bucketStart": "400 ms",
// 			"bucketEnd": "700 ms"
// 		  }
// 		}
// 	  }
// 	},
// 	"appVersion": "1.0.0",
// 	"timeStampBegin": "2019-10-21 07:00:00 +0000"
//   }
