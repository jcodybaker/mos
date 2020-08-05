//
// Copyright (c) 2014-2019 Cesanta Software Limited
// All rights reserved
//
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
//
package cc32xx

import (
	"os/exec"
	"regexp"

	"github.com/juju/errors"
	glog "k8s.io/klog/v2"
)

func GetUSBSerialNumberForPort(port string) (string, error) {
	out, err := exec.Command("udevadm", "info", "--name", port).Output()
	if err != nil {
		return "", errors.Trace(err)
	}
	glog.V(1).Infof("udevadm output:\n%s", out)
	m := regexp.MustCompile(` ID_SERIAL_SHORT=(\S+)`).FindSubmatch(out)
	if m == nil {
		return "", errors.Errorf("No serial number in udevadm output")
	}
	return string(m[1]), nil
}
