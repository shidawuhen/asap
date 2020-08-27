package global

var (
	//
	globalService map[string](map[string]string)
)

func init() {
	globalService = make(map[string](map[string]string))
}

func SetService(serviceName string, address string) {
	if _, ok := globalService[serviceName];!ok {
		globalService[serviceName] = make(map[string]string)
	}
	globalService[serviceName][address] = address
}

func GetService(serviceName string) (map[string]string) {
	return globalService[serviceName]
}

func GetServiceArr() map[string](map[string]string) {
	return globalService
}

