package messages

import "github.com/alobaton/i18n"

var translate *i18n.Translate

func Init(path, mainLocale string) error {
	var err error
	translate, err = i18n.NewTranslate().BindPath(path).BindMainLocale(mainLocale).Init()
	if err != nil {
		return err
	}

	return nil
}

func Get(key string, args ...interface{}) string {
	result, err := translate.Lookup(key, args...)
	if err != nil {
		return key
	}

	return result
}
