package qxLang

const (
	LangZhCN = "zh-CN"
	LangZhTW = "zh-TW"
	LangEnUS = "en-US"
)

var SupportLanMap = map[string]bool{
	LangZhCN: true,
	LangZhTW: true,
	LangEnUS: true,
}

func CheckSupportLang(lang string) bool {
	if _, ok := SupportLanMap[lang]; ok {
		return true
	} else {
		return false
	}
}
