package utils

func StringArrayToInterface(stringArray []string) []interface{} {
    interfaces := make([]interface{}, len(stringArray))
    for i, s := range stringArray {
        interfaces[i] = s
    }
    return interfaces
}
