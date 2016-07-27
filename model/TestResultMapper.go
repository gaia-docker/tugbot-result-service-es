package model

// Static configuration
// {"testcase": {
//    "properties": {
// 	"ContainerId":{"type":"string"},
//      "ExitCode":{"type":"long"},
//      "StartedAt":{"type":""strict_date_optional_time||epoch_millis""},
//      "FinishedAt":{"type":""strict_date_optional_time||epoch_millis""},
//      "ImageName":{"type":"string"},
//      "HostName":{"type":"string"}
//	"TestSet": {
// 	  "properties": {
//		"Name": {"type":"string"},
//		"Time": {"type":"double"},
//		"Tests":
// 		  "properties": {
//			"Failure": { "type": "string" },
//			"Name": { "type": "string" },
//			"Status": { "type": "string" },
//			"Time": { "type": "double" }
// 		}
//	  }
// }
//}
func GetTestResultMapping() map[string]interface{} {

	typeStringMap := map[string]string{"type": "string"}
	typeDoubleMap := map[string]string{"type": "double"}
	typeLongMap := map[string]string{"type": "long"}
	typeDateMap := map[string]string{"type": "strict_date_optional_time||epoch_millis"}
	testsMap := map[string]interface{}{
		"properties": map[string]interface{}{
			"Failure": typeStringMap, "Name": typeStringMap, "Status": typeStringMap, "Time": typeDoubleMap}}
	testSetMap := map[string]interface{}{
		"properties": map[string]interface{}{
			"Name": typeStringMap, "Time": typeDoubleMap, "Tests": testsMap}}
	testCaseMap := map[string]interface{}{
		"properties": map[string]interface{}{
			"ContainerId": typeStringMap, "ExitCode": typeLongMap, "StartedAt": typeDateMap, "FinishedAt": typeDateMap, "ImageName": typeStringMap, "HostName": typeStringMap, "TestSet": testSetMap}}

	return testCaseMap
}
