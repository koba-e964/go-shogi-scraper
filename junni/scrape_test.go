package junni

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRawJunni0(t *testing.T) {
	dummyHash, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f")
	raw := &JunniScrapingRawResult{
		URL:           "https://example.com/test",
		RetrievalTime: "2024-02-01T01:48:44Z",
		Name:          "第xx期名人戦・順位戦　Ａ級",
		HashAlgorithm: "sha256",
		Hash:          dummyHash,
		Table: [][][]string{
			{{""}, {"1"}, {"山田太郎"}, {"0"}, {"1"}, {"●", "山田花"}, {"－"}, {"", "将棋"}},
			{{""}, {"2"}, {"山田花子"}, {"2"}, {"0"}, {"○", "山田太"}, {"○", "将棋"}, {"－"}},
			{{""}, {"3"}, {"将棋太郎"}, {"0"}, {"1"}, {"－"}, {"●", "山田花"}, {"先", "山田太"}},
		},
	}
	result, err := ParseRawJunni(raw)
	if assert.NoError(t, err) {
		assert.Equal(t, raw.URL, result.URL)
		assert.Equal(t, raw.RetrievalTime, result.RetrievalTime)
		assert.Equal(t, raw.Name, result.Name)
		assert.Equal(t, raw.HashAlgorithm, result.HashAlgorithm)
		assert.Equal(t, raw.Hash, result.Hash)
		assert.Equal(t, []string{"山田太郎", "山田花子", "将棋太郎"}, []string{result.Players[0].Name, result.Players[1].Name, result.Players[2].Name})
		expectedResults := []Result{
			{Index1: 0, Index2: 1, Winner: 2, PlayFirst: 0, Round: 1, GameRecordURL: ""},
			{Index1: 1, Index2: 2, Winner: 1, PlayFirst: 0, Round: 2, GameRecordURL: ""},
			{Index1: 0, Index2: 2, Winner: 0, PlayFirst: 2, Round: 3, GameRecordURL: ""},
		}
		assert.Equal(t, expectedResults, result.Results)
	}
}
