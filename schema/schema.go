/*
Copyright © 2020 Ramon Moraes <ramonmoraes8080@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package schema

import (
	"gitlab.com/velvetkeyboard/go-quickbackup/config"
	"gitlab.com/velvetkeyboard/go-quickbackup/utils"
)

type Schema struct {
	Name  string
	Files []string
}

func (s *Schema) Init(cfg *config.Configuration, schemaName string) {
	s.Name = schemaName
	s.Files = utils.WalkDirs(cfg.Schemas[schemaName])
}
