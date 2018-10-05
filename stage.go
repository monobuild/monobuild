// Copyright © 2017 Sascha Andres <sascha.andres@outlook.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package monobuild

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

// Stage is a build stage containing only independent build configurations
type Stage struct {
	Label          string
	Configurations []*BuildConfiguration
	Log            *logrus.Entry
}

// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
func (stage *Stage) String() string {
	return fmt.Sprintf("%s - %d build configurations", stage.Label, len(stage.Configurations))
}
