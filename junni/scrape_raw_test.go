package junni

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScrapeJunniRaw0(t *testing.T) {
	url := "https://web.archive.org/web/20240201014844if_/https://www.shogi.or.jp/match/junni/2023/82a/index.html"

	result, err := ScrapeJunniRaw(url)
	if assert.NoError(t, err) {
		assert.Equal(t, url, result.URL)
		assert.Equal(t, "第82期名人戦・順位戦　Ａ級", result.Name)
		assert.Equal(t, "sha256", result.HashAlgorithm)
		expectedHash, _ := hex.DecodeString("88e20b33b09dfbd3656411ad803fa47a094601918112403b324bb51eb29637c5")
		assert.Equal(t, expectedHash, result.Hash)
		expectedTable := [][][]string{
			{{""}, {"1"}, {"<a href=\"/player/pro/235.html\">渡辺\u3000明</a>九段"}, {"4"}, {"4"}, {"●", "佐々木勇"}, {"○", "稲葉"}, {"●", "豊島"}, {"○", "佐藤天"}, {"○", "斎藤慎"}, {"●", "永瀬"}, {"○", "菅井"}, {"●", "中村太"}, {"先", "広瀬"}},
			{{""}, {"2"}, {"<a href=\"/player/pro/255.html\">広瀬章人</a>九段"}, {"3"}, {"5"}, {"○", "中村太"}, {"●", "斎藤慎"}, {"●", "佐々木勇"}, {"●", "菅井"}, {"●", "永瀬"}, {"●", "佐藤天"}, {"○", "豊島"}, {"○", "稲葉"}, {"", "渡辺明"}},
			{{""}, {"3"}, {"<a href=\"/player/pro/264.html\">豊島将之</a>九段"}, {"6"}, {"2"}, {"○", "稲葉"}, {"○", "佐々木勇"}, {"○", "渡辺明"}, {"○", "永瀬"}, {"○", "佐藤天"}, {"○", "中村太"}, {"●", "広瀬"}, {"●", "斎藤慎"}, {"", "菅井"}},
			{{""}, {"4"}, {"<a href=\"/player/pro/276.html\">永瀬拓矢</a>九段"}, {"5"}, {"3"}, {"●", "菅井"}, {"○", "佐藤天"}, {"○", "斎藤慎"}, {"●", "豊島"}, {"○", "広瀬"}, {"○", "渡辺明"}, {"●", "稲葉"}, {"○", "佐々木勇"}, {"先", "中村太"}},
			{{""}, {"5"}, {"<a href=\"/player/pro/286.html\">斎藤慎太郎</a>八段"}, {"3"}, {"5"}, {"●", "佐藤天"}, {"○", "広瀬"}, {"●", "永瀬"}, {"○", "稲葉"}, {"●", "渡辺明"}, {"●", "菅井"}, {"●", "中村太"}, {"○", "豊島"}, {"先", "佐々木勇"}},
			{{""}, {"6"}, {"<a href=\"/player/pro/278.html\">菅井竜也</a>八段"}, {"5"}, {"3"}, {"○", "永瀬"}, {"○", "中村太"}, {"●", "稲葉"}, {"○", "広瀬"}, {"○", "佐々木勇"}, {"○", "斎藤慎"}, {"●", "渡辺明"}, {"●", "佐藤天"}, {"先", "豊島"}},
			{{""}, {"7"}, {"<a href=\"/player/pro/269.html\">稲葉\u3000陽</a>八段"}, {"3"}, {"5"}, {"●", "豊島"}, {"●", "渡辺明"}, {"○", "菅井"}, {"●", "斎藤慎"}, {"○", "中村太"}, {"●", "佐々木勇"}, {"○", "永瀬"}, {"●", "広瀬"}, {"先", "佐藤天"}},
			{{""}, {"8"}, {"<a href=\"/player/pro/263.html\">佐藤天彦</a>九段"}, {"4"}, {"4"}, {"○", "斎藤慎"}, {"●", "永瀬"}, {"●", "中村太"}, {"●", "渡辺明"}, {"●", "豊島"}, {"○", "広瀬"}, {"○", "佐々木勇"}, {"○", "菅井"}, {"", "稲葉"}},
			{{""}, {"9"}, {"<a href=\"/player/pro/280.html\">佐々木勇気</a>八段"}, {"3"}, {"5"}, {"○", "渡辺明"}, {"●", "豊島"}, {"○", "広瀬"}, {"●", "中村太"}, {"●", "菅井"}, {"○", "稲葉"}, {"●", "佐藤天"}, {"●", "永瀬"}, {"", "斎藤慎"}},
			{{""}, {"10"}, {"<a href=\"/player/pro/261.html\">中村太地</a>八段"}, {"4"}, {"4"}, {"●", "広瀬"}, {"●", "菅井"}, {"○", "佐藤天"}, {"○", "佐々木勇"}, {"●", "稲葉"}, {"●", "豊島"}, {"○", "斎藤慎"}, {"○", "渡辺明"}, {"", "永瀬"}},
		}
		assert.Equal(t, expectedTable, result.Table)
	}
}

func TestScrapeJunniRaw1(t *testing.T) {
	url := "https://web.archive.org/web/20240119232150if_/https://www.shogi.or.jp/match/junni/2023/82c1/index.html"

	result, err := ScrapeJunniRaw(url)
	if assert.NoError(t, err) {
		assert.Equal(t, url, result.URL)
		assert.Equal(t, "第82期名人戦・順位戦　Ｃ級１組", result.Name)
		assert.Equal(t, "sha256", result.HashAlgorithm)
		expectedTable := [][][]string{
			{{""}, {"1"}, {"伊藤\u3000匠七段"}, {"6"}, {"2"}, {"●", "出口"}, {"○", "西尾"}, {"○", "片上"}, {"●", "古賀"}, {"○", "先崎"}, {"○", "斎藤明"}, {"－"}, {"○", "古森"}, {"○", "村田顕"}, {"先", "阿部健"}, {"", "宮田"}},
			{{""}, {"2"}, {"黒田尭之五段"}, {"4"}, {"4"}, {"●", "斎藤明"}, {"●", "野月"}, {"●", "船江"}, {"○", "畠山成"}, {"－"}, {"○", "宮田"}, {"○", "先崎"}, {"○", "宮本"}, {"●", "古賀"}, {"先", "高橋"}, {"", "西尾"}},
			{{""}, {"3"}, {"出口若武六段"}, {"6"}, {"2"}, {"○", "伊藤匠"}, {"●", "都成"}, {"○", "三枚堂"}, {"○", "飯塚"}, {"－"}, {"○", "北島"}, {"●", "佐藤和"}, {"○", "古賀"}, {"○", "宮本"}, {"", "金井"}, {"先", "服部"}},
			{{""}, {"4"}, {"都成竜馬七段"}, {"6"}, {"2"}, {"○", "北島"}, {"○", "出口"}, {"○", "宮田"}, {"○", "門倉"}, {"－"}, {"○", "阿部健"}, {"●", "西尾"}, {"●", "三枚堂"}, {"○", "畠山成"}, {"先", "窪田"}, {"", "飯塚"}},
			{{""}, {"5"}, {"宮田敦史七段"}, {"4"}, {"4"}, {"○", "日浦"}, {"○", "北島"}, {"●", "都成"}, {"○", "村田顕"}, {"－"}, {"●", "黒田"}, {"●", "門倉"}, {"○", "真田"}, {"●", "三枚堂"}, {"", "野月"}, {"先", "伊藤匠"}},
			{{""}, {"6"}, {"野月浩貴八段"}, {"3"}, {"5"}, {"○", "古森"}, {"○", "黒田"}, {"●", "西田"}, {"●", "窪田"}, {"－"}, {"●", "古賀"}, {"●", "片上"}, {"●", "飯塚"}, {"○", "阪口"}, {"先", "宮田"}, {"", "三枚堂"}},
			{{""}, {"7"}, {"西尾\u3000明七段"}, {"5"}, {"3"}, {"○", "船江"}, {"●", "伊藤匠"}, {"○", "真田"}, {"●", "三枚堂"}, {"－"}, {"●", "千葉"}, {"○", "都成"}, {"○", "日浦"}, {"○", "佐藤和"}, {"", "門倉"}, {"先", "黒田"}},
			{{""}, {"8"}, {"阿部健治郎七段"}, {"2"}, {"6"}, {"●", "三枚堂"}, {"●", "先崎"}, {"○", "門倉"}, {"●", "服部"}, {"－"}, {"●", "都成"}, {"●", "古賀"}, {"●", "斎藤明"}, {"○", "千葉"}, {"", "伊藤匠"}, {"先", "北島"}},
			{{""}, {"9"}, {"三枚堂達也七段"}, {"6"}, {"2"}, {"○", "阿部健"}, {"●", "高橋"}, {"●", "出口"}, {"○", "西尾"}, {"－"}, {"○", "真田"}, {"○", "古森"}, {"○", "都成"}, {"○", "宮田"}, {"", "日浦"}, {"先", "野月"}},
			{{""}, {"10"}, {"片上大輔七段"}, {"4"}, {"5"}, {"●", "服部"}, {"○", "窪田"}, {"●", "伊藤匠"}, {"●", "斎藤明"}, {"○", "飯塚"}, {"●", "畠山成"}, {"○", "野月"}, {"○", "佐藤和"}, {"●", "真田"}, {"－"}, {"", "高橋"}},
			{{""}, {"11"}, {"門倉啓太五段"}, {"4"}, {"4"}, {"○", "村田顕"}, {"○", "阪口"}, {"●", "阿部健"}, {"●", "都成"}, {"－"}, {"●", "先崎"}, {"○", "宮田"}, {"●", "畠山成"}, {"○", "斎藤明"}, {"先", "西尾"}, {"", "古森"}},
			{{""}, {"12"}, {"佐藤和俊七段"}, {"3"}, {"5"}, {"●", "古賀"}, {"●", "村田顕"}, {"－"}, {"○", "西田"}, {"○", "日浦"}, {"●", "船江"}, {"○", "出口"}, {"●", "片上"}, {"●", "西尾"}, {"先", "古森"}, {"", "斎藤明"}},
			{{""}, {"13"}, {"古森悠太五段"}, {"5"}, {"3"}, {"●", "野月"}, {"○", "畠山成"}, {"○", "日浦"}, {"○", "阪口"}, {"－"}, {"○", "高橋"}, {"●", "三枚堂"}, {"●", "伊藤匠"}, {"○", "船江"}, {"", "佐藤和"}, {"先", "門倉"}},
			{{""}, {"rankdown1", "14"}, {"先崎\u3000学九段"}, {"5"}, {"3"}, {"○", "宮本"}, {"○", "阿部健"}, {"○", "高橋"}, {"●", "真田"}, {"●", "伊藤匠"}, {"○", "門倉"}, {"●", "黒田"}, {"○", "西田"}, {"－"}, {"", "北島"}, {"先", "千葉"}},
			{{""}, {"rankdown1", "15"}, {"金井恒太六段"}, {"4"}, {"4"}, {"○", "千葉"}, {"●", "斎藤明"}, {"●", "阪口"}, {"○", "日浦"}, {"○", "船江"}, {"●", "西田"}, {"●", "飯塚"}, {"－"}, {"○", "窪田"}, {"先", "出口"}, {"", "畠山成"}},
			{{""}, {"16"}, {"飯塚祐紀八段"}, {"4"}, {"4"}, {"○", "窪田"}, {"－"}, {"●", "村田顕"}, {"●", "出口"}, {"●", "片上"}, {"●", "日浦"}, {"○", "金井"}, {"○", "野月"}, {"○", "北島"}, {"", "服部"}, {"先", "都成"}},
			{{""}, {"17"}, {"窪田義行七段"}, {"2"}, {"6"}, {"●", "飯塚"}, {"●", "片上"}, {"●", "宮本"}, {"○", "野月"}, {"－"}, {"●", "服部"}, {"○", "阪口"}, {"●", "船江"}, {"●", "金井"}, {"", "都成"}, {"先", "古賀"}},
			{{""}, {"18"}, {"船江恒平六段"}, {"3"}, {"5"}, {"●", "西尾"}, {"●", "西田"}, {"○", "黒田"}, {"－"}, {"●", "金井"}, {"○", "佐藤和"}, {"●", "畠山成"}, {"○", "窪田"}, {"●", "古森"}, {"", "千葉"}, {"先", "村田顕"}},
			{{""}, {"19"}, {"宮本広志五段"}, {"4"}, {"4"}, {"●", "先崎"}, {"●", "服部"}, {"○", "窪田"}, {"○", "千葉"}, {"－"}, {"○", "阪口"}, {"○", "高橋"}, {"●", "黒田"}, {"●", "出口"}, {"", "古賀"}, {"先", "真田"}},
			{{""}, {"20"}, {"真田圭一八段"}, {"2"}, {"6"}, {"－"}, {"●", "古賀"}, {"●", "西尾"}, {"○", "先崎"}, {"●", "村田顕"}, {"●", "三枚堂"}, {"●", "西田"}, {"●", "宮田"}, {"○", "片上"}, {"先", "畠山成"}, {"", "宮本"}},
			{{""}, {"21"}, {"北島忠雄七段"}, {"1"}, {"7"}, {"●", "都成"}, {"●", "宮田"}, {"●", "古賀"}, {"○", "高橋"}, {"－"}, {"●", "出口"}, {"●", "服部"}, {"●", "阪口"}, {"●", "飯塚"}, {"先", "先崎"}, {"", "阿部健"}},
			{{""}, {"22"}, {"西田拓也五段"}, {"6"}, {"2"}, {"○", "畠山成"}, {"○", "船江"}, {"○", "野月"}, {"●", "佐藤和"}, {"－"}, {"○", "金井"}, {"○", "真田"}, {"●", "先崎"}, {"○", "日浦"}, {"先", "斎藤明"}, {"", "阪口"}},
			{{""}, {"23"}, {"阪口\u3000悟六段"}, {"2"}, {"6"}, {"●", "高橋"}, {"●", "門倉"}, {"○", "金井"}, {"●", "古森"}, {"－"}, {"●", "宮本"}, {"●", "窪田"}, {"○", "北島"}, {"●", "野月"}, {"", "村田顕"}, {"先", "西田"}},
			{{""}, {"24"}, {"斎藤明日斗五段"}, {"6"}, {"2"}, {"○", "黒田"}, {"○", "金井"}, {"○", "千葉"}, {"○", "片上"}, {"－"}, {"●", "伊藤匠"}, {"○", "日浦"}, {"○", "阿部健"}, {"●", "門倉"}, {"", "西田"}, {"先", "佐藤和"}},
			{{""}, {"25"}, {"服部慎一郎六段"}, {"8"}, {"0"}, {"○", "片上"}, {"○", "宮本"}, {"○", "畠山成"}, {"○", "阿部健"}, {"－"}, {"○", "窪田"}, {"○", "北島"}, {"○", "村田顕"}, {"○", "高橋"}, {"先", "飯塚"}, {"", "出口"}},
			{{""}, {"26"}, {"古賀悠聖五段"}, {"7"}, {"1"}, {"○", "佐藤和"}, {"○", "真田"}, {"○", "北島"}, {"○", "伊藤匠"}, {"－"}, {"○", "野月"}, {"○", "阿部健"}, {"●", "出口"}, {"○", "黒田"}, {"先", "宮本"}, {"", "窪田"}},
			{{""}, {"rankdown1", "27"}, {"村田顕弘六段"}, {"4"}, {"4"}, {"●", "門倉"}, {"○", "佐藤和"}, {"○", "飯塚"}, {"●", "宮田"}, {"○", "真田"}, {"－"}, {"○", "千葉"}, {"●", "服部"}, {"●", "伊藤匠"}, {"先", "阪口"}, {"", "船江"}},
			{{""}, {"rankdown1", "28"}, {"高橋道雄九段"}, {"2"}, {"6"}, {"○", "阪口"}, {"○", "三枚堂"}, {"●", "先崎"}, {"●", "北島"}, {"－"}, {"●", "古森"}, {"●", "宮本"}, {"●", "千葉"}, {"●", "服部"}, {"", "黒田"}, {"先", "片上"}},
			{{""}, {"rankdown1", "29"}, {"畠山成幸八段"}, {"3"}, {"5"}, {"●", "西田"}, {"●", "古森"}, {"●", "服部"}, {"●", "黒田"}, {"－"}, {"○", "片上"}, {"○", "船江"}, {"○", "門倉"}, {"●", "都成"}, {"", "真田"}, {"先", "金井"}},
			{{""}, {"rankdown1", "30"}, {"千葉幸生七段"}, {"3"}, {"5"}, {"●", "金井"}, {"○", "日浦"}, {"●", "斎藤明"}, {"●", "宮本"}, {"－"}, {"○", "西尾"}, {"●", "村田顕"}, {"○", "高橋"}, {"●", "阿部健"}, {"先", "船江"}, {"", "先崎"}},
			{{""}, {"rankdown1", "31"}, {"日浦市郎八段"}, {"1"}, {"8"}, {"●", "宮田"}, {"●", "千葉"}, {"●", "古森"}, {"●", "金井"}, {"●", "佐藤和"}, {"○", "飯塚"}, {"●", "斎藤明"}, {"●", "西尾"}, {"●", "西田"}, {"先", "三枚堂"}, {"－"}},
		}
		assert.Equal(t, expectedTable, result.Table)
	}
}