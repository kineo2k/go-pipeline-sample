package runmode

import "os"

type RunMode string

const (
	Local      = RunMode("local")
	Production = RunMode("production")
)

func isUnacceptable(rawValue string) bool {
	switch RunMode(rawValue) {
	case Local:
		fallthrough
	case Production:
		return false
	default:
		return true
	}
}

func CurrentRunMode() RunMode {
	argsWithoutProgramArguments := os.Args[1:]
	if len(argsWithoutProgramArguments) != 1 {
		panic("실행 모드 파라미터 누락 (local, production)\n")
	}

	rawValue := argsWithoutProgramArguments[0]
	if isUnacceptable(rawValue) {
		panic("실행 모드 파라미터 오류 (local, production)\n")
	}

	return RunMode(rawValue)
}
