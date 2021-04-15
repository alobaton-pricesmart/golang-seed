package messages

import "github.com/alobaton/i18n"

var i *i18n.I18N

func Init(path, mainLocale string) error {
	var err error

	i = i18n.NewI18N().BindPath(path)
	i, err = i.BindMainLocale(mainLocale)
	if err != nil {
		return err
	}

	i, err = i.Init()
	if err != nil {
		return err
	}

	return nil
}

func Get(key string, args ...interface{}) string {
	result, err := i.Lookup(key, args...)
	if err != nil {
		return key
	}

	return result
}
