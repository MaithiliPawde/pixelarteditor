//
// Copyright 2021 Bryan T. Meyers <root@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package util

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// GetPressed returns a list of all the pressed keys
func GetPressed() (keys []ebiten.Key) {
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			keys = append(keys, k)
		}
	}
	return
}

// In checks if an XY coordinate is inside of a bounding box
func In(r image.Rectangle, x, y int) bool {
	return image.Rect(x, y, x+1, y+1).In(r)
}
