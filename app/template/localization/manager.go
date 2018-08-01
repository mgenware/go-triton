package localization

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/mgenware/go-packagex/filepathx"
	"github.com/mgenware/go-triton/app/defs"
)

type Manager struct {
	defaultDic *Dictionary
	dics       map[string]*Dictionary
}

func NewManagerFromDirectory(dir string, defaultLang string) (*Manager, error) {
	fileNames, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	dics := make(map[string]*Dictionary)
	for _, info := range fileNames {
		if !info.IsDir() {
			d, err := NewDictionaryFromFile(filepath.Join(dir, info.Name()))
			if err != nil {
				return nil, err
			}

			name := filepathx.TrimExt(info.Name())
			dics[name] = d
			log.Printf("Read localization file \"%v\"", name)
		}
	}
	if len(dics) == 0 {
		return nil, fmt.Errorf("No dictionary found in %v", dir)
	}

	defaultDic := dics[defaultLang]
	if defaultDic == nil {
		return nil, fmt.Errorf("Default language \"%v\" not found", defaultLang)
	}

	return &Manager{dics: dics, defaultDic: defaultDic}, nil
}

func (mgr *Manager) DictionaryForLanguage(lang string) *Dictionary {
	dic := mgr.dics[lang]
	if dic == nil {
		return mgr.defaultDic
	}
	return dic
}

func (mgr *Manager) ValueForKeyWithLanguage(lang, key string) string {
	dic := mgr.DictionaryForLanguage(lang)
	if dic == nil {
		return ""
	}
	return dic.Map[key]
}

func (mgr *Manager) ValueForKey(ctx context.Context, key string) string {
	return mgr.ValueForKeyWithLanguage(defs.LanguageFromContext(ctx), key)
}
