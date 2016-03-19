package jpconv

import (
	"testing"
)

func Test_HiraganaToKatakana(t *testing.T) {
	var result string
	result = HiraganaToKatakana("ひらがなをカタカナにへんかんします")
	if result != "ヒラガナヲカタカナニヘンカンシマス" {
		t.Errorf("fail! %s", result)
	}
	result = HiraganaToKatakana("ゔぁいおりん")
	if result != "ヴァイオリン" {
		t.Errorf("fail! %s", result)
	}
}

func Test_KatakanaToHiragana(t *testing.T) {
	var result string
	result = KatakanaToHiragana("カタカナヲひらがなニヘンカンシマス")
	if result != "かたかなをひらがなにへんかんします" {
		t.Errorf("fail! %s", result)
	}
	result = KatakanaToHiragana("ヴァイオリン")
	if result != "ゔぁいおりん" {
		t.Errorf("fail! %s", result)
	}
}

func Test_HankakuToZenkaku(t *testing.T) {
	var result string
	result = HankakuToZenkaku("ﾊﾝｶｸｦｾﾞﾝｶｸﾆﾍﾝｶﾝｼﾏｽ")
	if result != "ハンカクヲゼンカクニヘンカンシマス" {
		t.Error("fail!")
	}
}

func Test_ZenkakuToHankaku(t *testing.T) {
	var result string
	result = ZenkakuToHankaku("ゼンカクヲハンカクニヘンカンシマス")
	if result != "ｾﾞﾝｶｸｦﾊﾝｶｸﾆﾍﾝｶﾝｼﾏｽ" {
		t.Error("fail!")
	}
}
