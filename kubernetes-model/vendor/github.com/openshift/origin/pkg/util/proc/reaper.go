/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// +build linux

package proc

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/glog"
)

// StartReaper starts a goroutine to reap processes if called from a process
// that has pid 1.
func StartReaper() {
	if os.Getpid() == 1 {
		glog.V(4).Infof("Launching reaper")
		go func() {
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGCHLD)
			for {
				// Wait for a child to terminate
				sig := <-sigs
				glog.V(4).Infof("Signal received: %v", sig)
				for {
					// Reap processes
					cpid, _ := syscall.Wait4(-1, nil, syscall.WNOHANG, nil)
					if cpid < 1 {
						break
					}

					glog.V(4).Infof("Reaped process with pid %d", cpid)
				}
			}
		}()
	}
}
