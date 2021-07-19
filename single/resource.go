package single

import (
	"crypto/md5"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"woodpecker/checker"
	"woodpecker/core"
	"woodpecker/logger"
	"woodpecker/notifier"
)

func LoadAllResourceFile(engine *Engine) error {
	dir, err := ioutil.ReadDir(engine.Config.ResourceDirName)
	if err != nil {
		return err
	}
	for _, f := range dir {
		if f.IsDir() {
			continue
		}
		if !strings.HasSuffix(f.Name(), ".yaml") {
			continue
		}
		path := filepath.Join(engine.Config.ResourceDirName, f.Name())
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		fingerprint := fmt.Sprintf("%x", md5.Sum(content))
		logger.Info(fmt.Sprintf("load resource file: %s\n", filepath.Base(path)))
		engine.ResourceFileMap[path] = &core.ResourceFile{
			FilePath:    path,
			Fingerprint: fingerprint,
		}
	}
	return err
}

// load new resource file
// ignore deleted file
func ReloadNewResourceFile(engine *Engine) error {
	dir, err := ioutil.ReadDir(engine.Config.ResourceDirName)
	if err != nil {
		return err
	}
	for _, f := range dir {
		if !f.IsDir() {
			continue
		}
		path := filepath.Join(engine.Config.ResourceDirName, f.Name())
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		fingerprint := fmt.Sprintf("%x", md5.Sum(content))
		old := engine.ResourceFileMap[path]
		if old != nil && old.Fingerprint == fingerprint {
			return nil
		}
		logger.Info(fmt.Sprintf("load resource file: %s", filepath.Base(path)))
		engine.ResourceFileMap[path] = &core.ResourceFile{
			FilePath:    path,
			Fingerprint: fingerprint,
		}
		return nil
	}
	return err

}

func LoadAllResource(engine *Engine) error {
	for _, rf := range engine.ResourceFileMap {
		path := rf.FilePath
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		decoder := yaml.NewDecoder(file)
		resource := core.Resource{}
		err = decoder.Decode(&resource)
		for err == nil {
			_ = loadResource(engine, &resource)
			resource := core.Resource{}
			err = decoder.Decode(&resource)
		}
		if err == io.EOF {
			_ = loadResource(engine, &resource)
			continue
		}
		logger.Error("load resource from file error", path, err)
	}
	return nil
}

func loadResource(engine *Engine, resource *core.Resource) error {
	switch resource.Kind {
	case "notifier":
		loadNotifier(engine, resource)
	case "checker":
		loadChecker(engine, resource)
	default:
		panic("unknown source: " + resource.Kind)
	}
	return nil
}

func loadNotifier(engine *Engine, r *core.Resource) {
	template := notifier.NotifierTemplate{}
	err := core.Interface2Interface(r, &template)
	if err != nil {
		panic(err)
	}
	nf, err := notifier.NewNotifier(&template)
	engine.NotifierMap[template.Metadata.Name] = nf
}

func loadChecker(engine *Engine, r *core.Resource) {
	template := checker.CheckerTemplate{}
	err := core.Interface2Interface(r, &template)
	if err != nil {
		panic(err)
	}
	c, err := checker.NewChecker(&template)
	if err != nil {
		panic(err)
	}
	engine.CheckerMap[template.Metadata.Name] = c
}
