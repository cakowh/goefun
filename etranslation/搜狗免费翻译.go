package etranslation

import (
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/ehttp"
	"net/url"
)

// 搜狗免费翻译结构体
type 搜狗免费翻译 struct {
	E语言转转换键值 map[string]string
}

func New搜狗免费翻译() *搜狗免费翻译 {
	return &搜狗免费翻译{
		E语言转转换键值: map[string]string{
			"Auto":                     "",
			"Simplified Chinese":       "zh",
			"Traditional Chinese":      "zh-tw",
			"English":                  "en",
			"Abkhazian":                "ab",
			"Albanian":                 "sq",
			"Akan":                     "ak",
			"Arabic":                   "ar",
			"Aragonese":                "an",
			"Amharic":                  "am",
			"Assamese":                 "as",
			"Azerbaijani":              "az",
			"Asturian":                 "ast",
			"Central Huasteca Nahuatl": "nch",
			"Ewe":                      "ee",
			"Aymara":                   "ay",
			"Irish":                    "ga",
			"Estonian":                 "et",
			"Ojibwa":                   "oj",
			"Occitan":                  "oc",
			"Oriya":                    "or",
			"Oromo":                    "om",
			"Ossetian":                 "os",
			"Tok Pisin":                "tpi",
			"Bashkir":                  "ba",
			"Basque":                   "eu",
			"Belarusian":               "be",
			"Berber languages":         "ber",
			"Bambara":                  "bm",
			"Pangasinan":               "pag",
			"Bulgarian":                "bg",
			"Northern Sami":            "se",
			"Bemba (Zambia)":           "bem",
			"Blin":                     "byn",
			"Bislama":                  "bi",
			"Baluchi":                  "bal",
			"Icelandic":                "is",
			"Polish":                   "pl",
			"Bosnian":                  "bs",
			"Persian":                  "fa",
			"Bhojpuri":                 "bho",
			"Breton":                   "br",
			"Chamorro":                 "ch",
			"Chavacano":                "cbk",
			"Chuvash":                  "cv",
			"Tsonga":                   "ts",
			"Tatar":                    "tt",
			"Danish":                   "da",
			"Shan":                     "shn",
			"Tetum":                    "tet",
			"German":                   "de",
			"Low German":               "nds",
			"Scots":                    "sco",
			"Dhivehi":                  "dv",
			"Kam":                      "kdx",
			"Kadazan Dusun":            "dtp",
			"Russian":                  "ru",
			"Faroese":                  "fo",
			"French":                   "fr",
			"Sanskrit":                 "sa",
			"Filipino":                 "fil",
			"Fijian":                   "fj",
			"Finnish":                  "fi",
			"Friulian":                 "fur",
			"Fur":                      "fvr",
			"Kongo":                    "kg",
			"Khmer":                    "km",
			"Guerrero Nahuatl":         "ngu",
			"Kalaallisut":              "kl",
			"Georgian":                 "ka",
			"Gronings":                 "gos",
			"Gujarati":                 "gu",
			"Guarani":                  "gn",
			"Kazakh":                   "kk",
			"Haitian":                  "ht",
			"Korean":                   "ko",
			"Hausa":                    "ha",
			"Dutch":                    "nl",
			"Montenegrin":              "cnr",
			"Hupa":                     "hup",
			"Gilbertese":               "gil",
			"Rundi":                    "rn",
			"K'iche'":                  "quc",
			"Kirghiz":                  "ky",
			"Galician":                 "gl",
			"Catalan":                  "ca",
			"Czech":                    "cs",
			"Kabyle":                   "kab",
			"Kannada":                  "kn",
			"Kanuri":                   "kr",
			"Kashubian":                "csb",
			"Khasi":                    "kha",
			"Cornish":                  "kw",
			"Xhosa":                    "xh",
			"Corsican":                 "co",
			"Creek":                    "mus",
			"Crimean Tatar":            "crh",
			"Klingon":                  "tlh",
			"Serbo-Croatian":           "hbs",
			"Quechua":                  "qu",
			"Kashmiri":                 "ks",
			"Kurdish":                  "ku",
			"Latin":                    "la",
			"Latgalian":                "ltg",
			"Latvian":                  "lv",
			"Lao":                      "lo",
			"Lithuanian":               "lt",
			"Limburgish":               "li",
			"Lingala":                  "ln",
			"Ganda":                    "lg",
			"Letzeburgesch":            "lb",
			"Rusyn":                    "rue",
			"Kinyarwanda":              "rw",
			"Romanian":                 "ro",
			"Romansh":                  "rm",
			"Romany":                   "rom",
			"Lojban":                   "jbo",
			"Malagasy":                 "mg",
			"Manx":                     "gv",
			"Maltese":                  "mt",
			"Marathi":                  "mr",
			"Malayalam":                "ml",
			"Malay":                    "ms",
			"Mari (Russia)":            "chm",
			"Macedonian":               "mk",
			"Marshallese":              "mh",
			"Kekchí":                   "kek",
			"Maithili":                 "mai",
			"Morisyen":                 "mfe",
			"Maori":                    "mi",
			"Mongolian":                "mn",
			"Bengali":                  "bn",
			"Burmese":                  "my",
			"Hmong":                    "hmn",
			"Umbundu":                  "umb",
			"Navajo":                   "nv",
			"Afrikaans":                "af",
			"Nepali":                   "ne",
			"Niuean":                   "niu",
			"Norwegian":                "no",
			"Pam":                      "pmn",
			"Papiamento":               "pap",
			"Panjabi":                  "pa",
			"Portuguese":               "pt",
			"Pushto":                   "ps",
			"Nyanja":                   "ny",
			"Twi":                      "tw",
			"Cherokee":                 "chr",
			"Japanese":                 "ja",
			"Swedish":                  "sv",
			"Samoan":                   "sm",
			"Sango":                    "sg",
			"Sinhala":                  "si",
			"Upper Sorbian":            "hsb",
			"Esperanto":                "eo",
			"Slovenian":                "sl",
			"Swahili":                  "sw",
			"Somali":                   "so",
			"Slovak":                   "sk",
			"Tagalog":                  "tl",
			"Tajik":                    "tg",
			"Tahitian":                 "ty",
			"Telugu":                   "te",
			"Tamil":                    "ta",
			"Thai":                     "th",
			"Tonga (Tonga Islands)":    "to",
			"Tonga (Zambia)":           "toi",
			"Tigrinya":                 "ti",
			"Tuvalu":                   "tvl",
			"Tuvinian":                 "tyv",
			"Turkish":                  "tr",
			"Turkmen":                  "tk",
			"Walloon":                  "wa",
			"Waray (Philippines)":      "war",
			"Welsh":                    "cy",
			"Venda":                    "ve",
			"Volapük":                  "vo",
			"Wolof":                    "wo",
			"Udmurt":                   "udm",
			"Urdu":                     "ur",
			"Uzbek":                    "uz",
			"Spanish":                  "es",
			"Interlingue":              "ie",
			"Western Frisian":          "fy",
			"Silesian":                 "szl",
			"Hebrew":                   "he",
			"Hiligaynon":               "hil",
			"Hawaiian":                 "haw",
			"Modern Greek":             "el",
			"Lingua Franca Nova":       "lfn",
			"Sindhi":                   "sd",
			"Hungarian":                "hu",
			"Shona":                    "sn",
			"Cebuano":                  "ceb",
			"Syriac":                   "syr",
			"Sundanese":                "su",
			"Armenian":                 "hy",
			"Achinese":                 "ace",
			"Iban":                     "iba",
			"Igbo":                     "ig",
			"Ido":                      "io",
			"Iloko":                    "ilo",
			"Inuktitut":                "iu",
			"Italian":                  "it",
			"Yiddish":                  "yi",
			"Interlingua":              "ia",
			"Hindi":                    "hi",
			"Indonesia":                "id",
			"Ingush":                   "inh",
			"Yoruba":                   "yo",
			"Vietnamese":               "vi",
			"Zaza":                     "zza",
			"Javanese":                 "jv",
			"Chinese":                  "zh",
			"Cantonese":                "yue",
			"Zulu":                     "zu",
		},
	}
}

func (b *搜狗免费翻译) E取初始化参数() []string {
	return make([]string, 0)
}
func (b *搜狗免费翻译) E翻译(text, from, to string) (string, error) {

	语言列表 := New语言列表()
	from = 语言列表.E从区域标识取接口标识(from, b.E语言转转换键值)
	to = 语言列表.E从区域标识取接口标识(to, b.E语言转转换键值)

	eh := ehttp.NewHttp()
	text = url.QueryEscape(text)

	返回文本, err2 := eh.Get("https://fanyi.sogou.com/text?keyword=" + text + "&transfrom=" + from + "&transto=" + to + "&model=general")
	if err2 != nil {
		return "", err2
	}
	//println(返回文本)

	翻译结果 := ecore.StrCut(返回文本, `"dit":"$"`)

	return 翻译结果, nil
}
