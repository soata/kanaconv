package kanaconv

import (
	"strings"
	"unicode"
)

var hankaku2zenkaku = strings.NewReplacer(
	"0", "０", "1", "１", "2", "２", "3", "３", "4", "４",
	"5", "５", "6", "６", "7", "７", "8", "８", "9", "９",
	"ｶﾞ", "ガ", "ｷﾞ", "ギ", "ｸﾞ", "グ", "ｹﾞ", "ゲ", "ｺﾞ", "ゴ",
	"ｻﾞ", "ザ", "ｼﾞ", "ジ", "ｽﾞ", "ズ", "ｾﾞ", "ゼ", "ｿﾞ", "ゾ",
	"ﾀﾞ", "ダ", "ﾁﾞ", "ヂ", "ﾂﾞ", "ヅ", "ﾃﾞ", "デ", "ﾄﾞ", "ド",
	"ﾊﾞ", "バ", "ﾋﾞ", "ビ", "ﾌﾞ", "ブ", "ﾍﾞ", "ベ", "ﾎﾞ", "ボ",
	"ﾊﾟ", "パ", "ﾋﾟ", "ピ", "ﾌﾟ", "プ", "ﾍﾟ", "ペ", "ﾎﾟ", "ポ",
	"ｳﾞ", "ヴ", "ﾜﾞ", "ヷ", "ｦﾞ", "ヺ",
	"ｱ", "ア", "ｲ", "イ", "ｳ", "ウ", "ｴ", "エ", "ｵ", "オ",
	"ｶ", "カ", "ｷ", "キ", "ｸ", "ク", "ｹ", "ケ", "ｺ", "コ",
	"ｻ", "サ", "ｼ", "シ", "ｽ", "ス", "ｾ", "セ", "ｿ", "ソ",
	"ﾀ", "タ", "ﾁ", "チ", "ﾂ", "ツ", "ﾃ", "テ", "ﾄ", "ト",
	"ﾅ", "ナ", "ﾆ", "ニ", "ﾇ", "ヌ", "ﾈ", "ネ", "ﾉ", "ノ",
	"ﾊ", "ハ", "ﾋ", "ヒ", "ﾌ", "フ", "ﾍ", "ヘ", "ﾎ", "ホ",
	"ﾏ", "マ", "ﾐ", "ミ", "ﾑ", "ム", "ﾒ", "メ", "ﾓ", "モ",
	"ﾔ", "ヤ", "ﾕ", "ユ", "ﾖ", "ヨ",
	"ﾗ", "ラ", "ﾘ", "リ", "ﾙ", "ル", "ﾚ", "レ", "ﾛ", "ロ",
	"ﾜ", "ワ", "ｦ", "ヲ", "ﾝ", "ン",
	"ｧ", "ァ", "ｨ", "ィ", "ｩ", "ゥ", "ｪ", "ェ", "ｫ", "ォ",
	"ｯ", "ッ", "ｬ", "ャ", "ｭ", "ュ", "ｮ", "ョ",
	"｡", "。", "､", "、", "ｰ", "ー", "｢", "「", "｣", "」", "･", "・",
)

var zenkaku2hankaku = strings.NewReplacer(
	"０", "0", "１", "1", "２", "2", "３", "3", "４", "4",
	"５", "5", "６", "6", "７", "7", "８", "8", "９", "9",
	"ガ", "ｶﾞ", "ギ", "ｷﾞ", "グ", "ｸﾞ", "ゲ", "ｹﾞ", "ゴ", "ｺﾞ",
	"ザ", "ｻﾞ", "ジ", "ｼﾞ", "ズ", "ｽﾞ", "ゼ", "ｾﾞ", "ゾ", "ｿﾞ",
	"ダ", "ﾀﾞ", "ヂ", "ﾁﾞ", "ヅ", "ﾂﾞ", "デ", "ﾃﾞ", "ド", "ﾄﾞ",
	"バ", "ﾊﾞ", "ビ", "ﾋﾞ", "ブ", "ﾌﾞ", "ベ", "ﾍﾞ", "ボ", "ﾎﾞ",
	"パ", "ﾊﾟ", "ピ", "ﾋﾟ", "プ", "ﾌﾟ", "ペ", "ﾍﾟ", "ポ", "ﾎﾟ",
	"ヴ", "ｳﾞ", "ヷ", "ﾜﾞ", "ヺ", "ｦﾞ",
	"ア", "ｱ", "イ", "ｲ", "ウ", "ｳ", "エ", "ｴ", "オ", "ｵ",
	"カ", "ｶ", "キ", "ｷ", "ク", "ｸ", "ケ", "ｹ", "コ", "ｺ",
	"サ", "ｻ", "シ", "ｼ", "ス", "ｽ", "セ", "ｾ", "ソ", "ｿ",
	"タ", "ﾀ", "チ", "ﾁ", "ツ", "ﾂ", "テ", "ﾃ", "ト", "ﾄ",
	"ナ", "ﾅ", "ニ", "ﾆ", "ヌ", "ﾇ", "ネ", "ﾈ", "ノ", "ﾉ",
	"ハ", "ﾊ", "ヒ", "ﾋ", "フ", "ﾌ", "ヘ", "ﾍ", "ホ", "ﾎ",
	"マ", "ﾏ", "ミ", "ﾐ", "ム", "ﾑ", "メ", "ﾒ", "モ", "ﾓ",
	"ヤ", "ﾔ", "ユ", "ﾕ", "ヨ", "ﾖ",
	"ラ", "ﾗ", "リ", "ﾘ", "ル", "ﾙ", "レ", "ﾚ", "ロ", "ﾛ",
	"ワ", "ﾜ", "ヲ", "ｦ", "ン", "ﾝ",
	"ァ", "ｧ", "ィ", "ｨ", "ゥ", "ｩ", "ェ", "ｪ", "ォ", "ｫ",
	"ッ", "ｯ", "ャ", "ｬ", "ュ", "ｭ", "ョ", "ｮ",
	"。", "｡", "、", "､", "ー", "ｰ", "「", "｢", "」", "｣", "・", "･",
)

var smart = strings.NewReplacer(
	"ａ", "a", "ｂ", "b", "ｃ", "c", "ｄ", "d", "ｅ", "e", "ｆ", "f",
	"ｇ", "g", "ｈ", "h", "ｉ", "i", "ｊ", "j", "ｋ", "k", "ｌ", "l",
	"ｍ", "m", "ｎ", "n", "ｏ", "o", "ｐ", "p", "ｑ", "q", "ｒ", "r",
	"ｓ", "s", "ｔ", "t", "ｕ", "u", "ｖ", "v", "ｗ", "w", "ｘ", "x",
	"ｙ", "y", "ｚ", "z",
	"Ａ", "A", "Ｂ", "B", "Ｃ", "C", "Ｄ", "D", "Ｅ", "E", "Ｆ", "F",
	"Ｇ", "G", "Ｈ", "H", "Ｉ", "I", "Ｊ", "J", "Ｋ", "K", "Ｌ", "L",
	"Ｍ", "M", "Ｎ", "N", "Ｏ", "O", "Ｐ", "P", "Ｑ", "Q", "Ｒ", "R",
	"Ｓ", "S", "Ｔ", "T", "Ｕ", "U", "Ｖ", "V", "Ｗ", "W", "Ｘ", "X",
	"Ｙ", "Y", "Ｚ", "Z",
	"０", "0", "１", "1", "２", "2", "３", "3", "４", "4",
	"５", "5", "６", "6", "７", "7", "８", "8", "９", "9",
	"ｶﾞ", "ガ", "ｷﾞ", "ギ", "ｸﾞ", "グ", "ｹﾞ", "ゲ", "ｺﾞ", "ゴ",
	"ｻﾞ", "ザ", "ｼﾞ", "ジ", "ｽﾞ", "ズ", "ｾﾞ", "ゼ", "ｿﾞ", "ゾ",
	"ﾀﾞ", "ダ", "ﾁﾞ", "ヂ", "ﾂﾞ", "ヅ", "ﾃﾞ", "デ", "ﾄﾞ", "ド",
	"ﾊﾞ", "バ", "ﾋﾞ", "ビ", "ﾌﾞ", "ブ", "ﾍﾞ", "ベ", "ﾎﾞ", "ボ",
	"ﾊﾟ", "パ", "ﾋﾟ", "ピ", "ﾌﾟ", "プ", "ﾍﾟ", "ペ", "ﾎﾟ", "ポ",
	"ｳﾞ", "ヴ", "ﾜﾞ", "ヷ", "ｦﾞ", "ヺ",
	"ｱ", "ア", "ｲ", "イ", "ｳ", "ウ", "ｴ", "エ", "ｵ", "オ",
	"ｶ", "カ", "ｷ", "キ", "ｸ", "ク", "ｹ", "ケ", "ｺ", "コ",
	"ｻ", "サ", "ｼ", "シ", "ｽ", "ス", "ｾ", "セ", "ｿ", "ソ",
	"ﾀ", "タ", "ﾁ", "チ", "ﾂ", "ツ", "ﾃ", "テ", "ﾄ", "ト",
	"ﾅ", "ナ", "ﾆ", "ニ", "ﾇ", "ヌ", "ﾈ", "ネ", "ﾉ", "ノ",
	"ﾊ", "ハ", "ﾋ", "ヒ", "ﾌ", "フ", "ﾍ", "ヘ", "ﾎ", "ホ",
	"ﾏ", "マ", "ﾐ", "ミ", "ﾑ", "ム", "ﾒ", "メ", "ﾓ", "モ",
	"ﾔ", "ヤ", "ﾕ", "ユ", "ﾖ", "ヨ",
	"ﾗ", "ラ", "ﾘ", "リ", "ﾙ", "ル", "ﾚ", "レ", "ﾛ", "ロ",
	"ﾜ", "ワ", "ｦ", "ヲ", "ﾝ", "ン",
	"ｧ", "ァ", "ｨ", "ィ", "ｩ", "ゥ", "ｪ", "ェ", "ｫ", "ォ",
	"ｯ", "ッ", "ｬ", "ャ", "ｭ", "ュ", "ｮ", "ョ",
	"｡", "。", "､", "、", "ｰ", "ー", "｢", "「", "｣", "」", "･", "・",
	"－", "-", "−", "-", "‐", "-", "＋", "+", "＝", "=", "，", ",", "％", "%", "＊", "*",
	"（", "(", "）", ")", "!", "！", "”", `"`, "＃", "#", "＄", "$", "％", "%",
	"＆", "&", "’", "'",
)

var kanaCase = unicode.SpecialCase{
	// ァ-ヴ
	unicode.CaseRange{
		0x3040,
		0x3094,
		[unicode.MaxCase]rune{
			0x30a1 - 0x3041,
			0,
			0,
		},
	},
	// ぁ-ゔ
	unicode.CaseRange{
		0x30a0,
		0x30f4,
		[unicode.MaxCase]rune{
			0,
			0x3041 - 0x30a1,
			0,
		},
	},
}

var zenhanCase = unicode.SpecialCase{
	// 0-9
	unicode.CaseRange{
		0x0030,
		0x0039,
		[unicode.MaxCase]rune{
			0xff10 - 0x0030,
			0,
			0,
		},
	},
	// ０-９
	unicode.CaseRange{
		0xff10,
		0xff19,
		[unicode.MaxCase]rune{
			0,
			0x0030 - 0xff10,
			0,
		},
	},
	// ｧ-ﾝ
	unicode.CaseRange{
		0xff67,
		0xff9D,
		[unicode.MaxCase]rune{
			0x30a1 - 0xff67,
			0,
			0,
		},
	},
	// ァ-ン
	unicode.CaseRange{
		0x30a1,
		0x30f3,
		[unicode.MaxCase]rune{
			0,
			0xff67 - 0x30a1,
			0,
		},
	},
}

// HiraganaToKatakana はひらがなをカタカナに変換する
func HiraganaToKatakana(s string) string {
	return strings.ToUpperSpecial(kanaCase, s)
}

// KatakanaToHiragana はカタカナをひらがなに変換する
func KatakanaToHiragana(s string) string {
	return strings.ToLowerSpecial(kanaCase, s)
}

// ZenkakuToHankaku は全角文字を半角に変換する
func ZenkakuToHankaku(s string) string {
	return zenkaku2hankaku.Replace(s)
}

// HankakuToZenkaku は半角文字を全角に変換する
func HankakuToZenkaku(s string) string {
	return hankaku2zenkaku.Replace(s)
}

// SmartConv はカタカナは全角に、英数字記号は半角に統一する
func SmartConv(s string) string {
	return smart.Replace(s)
}
