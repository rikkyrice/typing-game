package handler

import (
	"api/internal/common/util"
	"api/internal/domain/model"
	"strings"
	"time"
)

type typingWord struct {
	Name  string     `json:"name"`
	Yomi  string     `json:"yomi"`
	Types [][]string `json:"types"`
}

type typingWordResponse struct {
	Word         *typingWord `json:"word"`
	Meaning      *typingWord `json:"meaning"`
	Explanation  string      `json:"explanation"`
	IsRemembered bool        `json:"isRemembered"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}

// wordとmeaningのどっちも配列にして返す
func toTypingWordResponse(word *model.Word) *typingWordResponse {
	return &typingWordResponse{
		Word: &typingWord{
			Name:  word.Word,
			Yomi:  word.Yomi,
			Types: createTypeWords(word.Word, word.Yomi),
		},
		Meaning: &typingWord{
			Name:  word.Meaning,
			Yomi:  word.MYomi,
			Types: createTypeWords(word.Meaning, word.MYomi),
		},
		Explanation:  word.Explanation,
		IsRemembered: word.IsRemembered,
		CreatedAt:    word.CreatedAt,
		UpdatedAt:    word.UpdatedAt,
	}
}

// Wordsでfor文を作成
// Yomiがあるか、ないか判定 -> アルファベットのみの場合変換の必要がない
// Yomiがない場合、Stringの二次元配列に変換してtypeWordsにpush
// Yomiがある場合、まずparseして細かいひらがなの配列に変換
// その後typing用のアルファベットに変換し、typeWordsにpush
func createTypeWords(word string, yomi string) [][]string {
	if yomi == "" {
		tw := make([][]string, len(word))
		letters := strings.Split(word, "")
		for j, letter := range letters {
			tw[j] = make([]string, 0)
			tw[j] = append(tw[j], letter)
		}
		return tw
	} else {
		parsedSentence := parseKanaSentence(yomi)
		typingSentence := convertToTypingSentences(parsedSentence)
		return typingSentence
	}
}
func parseKanaSentence(str string) []string {
	var res []string
	i := 0
	var uni, bi string
	for i < len([]rune(str)) {
		uni = string([]rune(str)[i : i+1])
		if i+1 < len([]rune(str)) {
			bi = string([]rune(str)[i:i+1]) + string([]rune(str)[i+1:i+2])
		} else {
			bi = ""
		}
		if _, ok := util.Mp[bi]; ok {
			i += 2
			res = append(res, bi)
		} else {
			i++
			res = append(res, uni)
		}
	}
	return res
}
func convertToTypingSentences(str []string) [][]string {
	var res [][]string
	var s, ns string
	for i := 0; i < len(str); i++ {
		s = str[i]
		if i+1 < len(str) {
			ns = str[i+1]
		} else {
			ns = ""
		}
		var tmpList []string
		if s == "ん" {
			var isValidSingleN bool
			var nList = util.Mp[s]
			if len(str)-1 == i {
				isValidSingleN = false
			} else if (i+1 < len(str)) && (ns == "あ" || ns == "い" || ns == "う" || ns == "え" || ns == "お" ||
				ns == "な" || ns == "に" || ns == "ぬ" || ns == "ね" || ns == "の" ||
				ns == "や" || ns == "ゆ" || ns == "よ") {
				isValidSingleN = false
			} else {
				isValidSingleN = true
			}
			for _, t := range nList {
				if !isValidSingleN && t == "n" {
					continue
				}
				tmpList = append(tmpList, t)
			}
		} else if s == "っ" {
			var ltuList = util.Mp[s]
			var nextList = util.Mp[ns]
			var hs []string
			for _, v := range nextList {
				hs = append(hs, string([]rune(v)[0:1]))
			}
			var ltuTypeList = append(hs, ltuList...)
			tmpList = ltuTypeList
		} else if len(s) == 2 && string([]rune(s)[0:1]) != "ん" {
			tmpList = append(tmpList, util.Mp[s]...)
			var fstList = util.Mp[string([]rune(s)[0:1])]
			var sndList = util.Mp[string([]rune(s)[1:2])]
			var resList []string
			for _, fstStr := range fstList {
				for _, sndStr := range sndList {
					var u = fstStr + sndStr
					resList = append(resList, u)
				}
			}
			tmpList = append(tmpList, resList...)
		} else {
			tmpList = util.Mp[s]
		}
		res = append(res, tmpList)
	}
	return res
}
