package jpconv

import (
	"strings"
	"unicode"
)

var numMap = map[string]string{
	"0": "０", "1": "１", "2": "２", "3": "３", "4": "４",
	"5": "５", "6": "６", "7": "７", "8": "８", "9": "９",
}
var dakuMap = map[string]string{
	"ｶﾞ": "ガ", "ｷﾞ": "ギ", "ｸﾞ": "グ", "ｹﾞ": "ゲ", "ｺﾞ": "ゴ",
	"ｻﾞ": "ザ", "ｼﾞ": "ジ", "ｽﾞ": "ズ", "ｾﾞ": "ゼ", "ｿﾞ": "ゾ",
	"ﾀﾞ": "ダ", "ﾁﾞ": "ヂ", "ﾂﾞ": "ヅ", "ﾃﾞ": "デ", "ﾄﾞ": "ド",
	"ﾊﾞ": "バ", "ﾋﾞ": "ビ", "ﾌﾞ": "ブ", "ﾍﾞ": "ベ", "ﾎﾞ": "ボ",
	"ﾊﾟ": "パ", "ﾋﾟ": "ピ", "ﾌﾟ": "プ", "ﾍﾟ": "ペ", "ﾎﾟ": "ポ",
	"ｳﾞ": "ヴ", "ﾜﾞ": "ヷ", "ｦﾞ": "ヺ",
}

var kataMap = map[string]string{
	"ｱ": "ア", "ｲ": "イ", "ｳ": "ウ", "ｴ": "エ", "ｵ": "オ",
	"ｶ": "カ", "ｷ": "キ", "ｸ": "ク", "ｹ": "ケ", "ｺ": "コ",
	"ｻ": "サ", "ｼ": "シ", "ｽ": "ス", "ｾ": "セ", "ｿ": "ソ",
	"ﾀ": "タ", "ﾁ": "チ", "ﾂ": "ツ", "ﾃ": "テ", "ﾄ": "ト",
	"ﾅ": "ナ", "ﾆ": "ニ", "ﾇ": "ヌ", "ﾈ": "ネ", "ﾉ": "ノ",
	"ﾊ": "ハ", "ﾋ": "ヒ", "ﾌ": "フ", "ﾍ": "ヘ", "ﾎ": "ホ",
	"ﾏ": "マ", "ﾐ": "ミ", "ﾑ": "ム", "ﾒ": "メ", "ﾓ": "モ",
	"ﾔ": "ヤ", "ﾕ": "ユ", "ﾖ": "ヨ",
	"ﾗ": "ラ", "ﾘ": "リ", "ﾙ": "ル", "ﾚ": "レ", "ﾛ": "ロ",
	"ﾜ": "ワ", "ｦ": "ヲ", "ﾝ": "ン",
	"ｧ": "ァ", "ｨ": "ィ", "ｩ": "ゥ", "ｪ": "ェ", "ｫ": "ォ",
	"ｯ": "ッ", "ｬ": "ャ", "ｭ": "ュ", "ｮ": "ョ",
}

var symbolMap = map[string]string{
	"｡": "。", "､": "、", "ｰ": "ー", "｢": "「", "｣": "」", "･": "・",
}

// HiraganaToKatakana はひらがなをカタカナに変換する
func HiraganaToKatakana(s string) string {
	convCase := unicode.SpecialCase{
		unicode.CaseRange{
			0x3041,
			0x3093,
			[unicode.MaxCase]rune{
				0x30a1 - 0x3041,
				0,
				0,
			},
		},
	}
	return strings.ToUpperSpecial(convCase, s)
}

// KatakanaToHiragana はカタカナをひらがなに変換する
func KatakanaToHiragana(s string) string {
	convCase := unicode.SpecialCase{
		unicode.CaseRange{
			0x30a1,
			0x30f3,
			[unicode.MaxCase]rune{
				0,
				0x3041 - 0x30a1,
				0,
			},
		},
	}
	return strings.ToLowerSpecial(convCase, s)
}

// ZenkakuToHankaku は全角カタカナを半角に変換する
func ZenkakuToHankaku(s string) string {
	for k, v := range numMap {
		s = strings.Replace(s, v, k, -1)
	}
	for k, v := range dakuMap {
		s = strings.Replace(s, v, k, -1)
	}
	for k, v := range kataMap {
		s = strings.Replace(s, v, k, -1)
	}
	for k, v := range symbolMap {
		s = strings.Replace(s, v, k, -1)
	}
	return s
}

// HankakuToZenkaku は半角カタカナを全角に変換する
func HankakuToZenkaku(s string) string {
	for k, v := range numMap {
		s = strings.Replace(s, k, v, -1)
	}
	for k, v := range dakuMap {
		s = strings.Replace(s, k, v, -1)
	}
	for k, v := range kataMap {
		s = strings.Replace(s, k, v, -1)
	}
	for k, v := range symbolMap {
		s = strings.Replace(s, k, v, -1)
	}
	return s
}
