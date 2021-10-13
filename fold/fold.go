package fold

import "golang.org/x/text/width"

//Convert function converts character width to Fold form.
func Convert(txt string) string {
	return convertStringWidth(width.Fold, txt)
}

//ConvertWiden function converts character width to Wide form.
func ConvertWiden(txt string) string {
	return convertStringWidth(width.Widen, txt)
}

//ConvertNarrow function converts character width to Narrow form.
func ConvertNarrow(txt string) string {
	return convertStringWidth(width.Narrow, txt)
}

func convertStringWidth(f width.Transformer, txt string) string {
	if f == width.Narrow {
		return NewReplaceerHalfWidthkana().Replace(f.String(NewReplaceerkanaNFD().Replace(txt)))
	}
	return NewReplaceerkanaNFC().Replace(f.String(txt))
}

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
