/*******************************************************************************
 * Copyright 2021 Samsung Electronics All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *******************************************************************************/

package config

import (
	"github.com/lf-edge/edge-home-orchestration-go/internal/common/logmgr"
	"testing"
)

var (
	log = logmgr.GetInstance()
)

func TestToml(t *testing.T) {
	SetWritable("DEBUG")
	SetService(49986, nil)
	SetRegistry("127.0.0.1", 8500)
	SetDevice(true, "", "", 128, 256, "", "", "./res")
	SetDeviceList("datastorage", "datastorage", "RESTful Device", []string{"rest", "json"})
	SetClients("127.0.0.1", "http", 5000)

	b, err := TomlMarshal()
	if err != nil {
		t.Fatal("Unexpected Error")
	}
	log.Println(string(b))
}
