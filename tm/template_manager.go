package tm

import (
	"embed"
	"fmt"
	"github.com/labstack/gommon/log"
	"go-pipeline-sample/runmode"
	"net/http"
	"os"
	"path"
	"strings"
)

var uniqueInstance *templateManager

type templateManager struct {
	runMode runmode.RunMode
	statics embed.FS
}

func Init(statics embed.FS) {
	if uniqueInstance == nil {
		uniqueInstance = new(templateManager)
		uniqueInstance.runMode = runmode.CurrentRunMode()
		uniqueInstance.statics = statics
	}
}

func GetInstance() *templateManager {
	return uniqueInstance
}

func (t *templateManager) Exists(relativePath string) bool {
	if t.runMode == runmode.Local {
		_, err := os.Stat(t.makeFilePath(relativePath))
		if err != nil {
			log.Error(err)
			return false
		}
	} else {
		f, err := t.statics.Open(t.makeFilePath(relativePath))
		if err != nil {
			log.Error(err)
			return false
		}
		defer f.Close()
	}

	return true
}

func (t *templateManager) GetFile(relativePath string) ([]byte, string) {
	blob, err := t.readFile(relativePath)
	if err != nil {
		log.Error(err)
		return nil, ""
	}

	contentType := ""
	if strings.HasSuffix(relativePath, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(relativePath, ".js") {
		contentType = "application/javascript"
	} else {
		contentType = http.DetectContentType(blob)
	}

	return blob, contentType
}

func (t *templateManager) readFile(relativePath string) ([]byte, error) {
	if t.runMode == runmode.Local {
		return os.ReadFile(t.makeFilePath(relativePath))
	} else {
		return t.statics.ReadFile(t.makeFilePath(relativePath))
	}
}

func (t *templateManager) makeFilePath(relativePath string) string {
	return path.Clean(fmt.Sprintf("statics/%s", relativePath))
}
