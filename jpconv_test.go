package jpconv

import (
	"testing"
)

func Test_HiraganaToKatakana(t *testing.T) {
	result := HiraganaToKatakana("ひらがなをかたかなにへんかんします")
	if result != "ヒラガナヲカタカナニヘンカンシマス" {
		t.Errorf("fail! %s", result)
	}
}

func Test_KatakanaToHiragana(t *testing.T) {
	result := KatakanaToHiragana("カタカナヲヒラガナニヘンカンシマス")
	if result != "かたかなをひらがなにへんかんします" {
		t.Errorf("fail! %s", result)
	}
}

func Test_HankakuToZenkaku(t *testing.T) {
	if HankakuToZenkaku("ﾊﾝｶｸｦｾﾞﾝｶｸﾆﾍﾝｶﾝｼﾏｽ") != "ハンカクヲゼンカクニヘンカンシマス" {
		t.Error("fail!")
	}
}

func Test_ZenkakuToHankaku(t *testing.T) {
	if ZenkakuToHankaku("ゼンカクヲハンカクニヘンカンシマス") != "ｾﾞﾝｶｸｦﾊﾝｶｸﾆﾍﾝｶﾝｼﾏｽ" {
		t.Error("fail!")
	}
}
