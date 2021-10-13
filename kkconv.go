package kkconv

import "github.com/spiegel-im-spiegel/kkconv/fold"

//Hiragana function converts hiragana character in the string.
func Hiragana(txt string, foldFlag bool) string {
	if foldFlag {
		txt = fold.Convert(txt)
	}
	return replaceHiragana(txt)
}

//Hiragana function converts katakana character in the string.
func Katakana(txt string, foldFlag bool) string {
	if foldFlag {
		txt = fold.Convert(txt)
	}
	return replaceKatakana(txt)
}

//Chokuon function converts youon (拗音) to chokuon (直音).
func Chokuon(txt string, foldFlag bool) string {
	if foldFlag {
		txt = fold.Convert(txt)
	}
	return replaceChokuon(txt)
}

// //Chokuon function converts chokuon (直音) to youon (拗音).
// func Youon(txt string, foldFlag bool) string {
// 	if foldFlag {
// 		txt = fold.Convert(txt)
// 	}
// 	return replaceYouon(txt)
// }

/* Copyright 2020-2021 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
