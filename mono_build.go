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
	"github.com/sirupsen/logrus"
)

type (
	// MonoBuild contains all data required to run the program
	MonoBuild struct {
		log            *logrus.Entry         // base logging facility
		stages         []*Stage              // build stages are stored here
		baseDirectory  string                // baseDirectory determines the root from where to scan
		configurations []*BuildConfiguration // configurations all found configurations

		DisableParallelism bool   // DisableParallelism can be used to run build independently of markers
		MarkerName         string // MarkerName is the name of the file with the build configuration
	}
)

// NewMonoBuild creates an empty mono build runner
func NewMonoBuild(baseDir string) *MonoBuild {
	return &MonoBuild{
		log:           logrus.WithField("package", "monobuild"),
		baseDirectory: baseDir,
	}
}
