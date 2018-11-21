/*
 * Minio Cloud Storage, (C) 2018 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"math/rand"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	data, _ := Asset("data/LogglySample.log")
	logLines := strings.Split(string(data), "\n")

	start, t, total := time.Now(), time.Now(), int64(0)

	// Simulate a few days of logs
	const days = 2
	const minutes = 3

	for minute := 0; minute < 10; /*60*24*days*/ minute += minutes {
		buffer := aggregate(logLines, 1000000)

		buffer[0] = byte(minute)
		object := deriveName(t, buffer, "log")

		n, _ := uploadBuffer(object, buffer)

		fmt.Println(humanize.IBytes(uint64(n)), "uploaded:", object)

		t = t.Add(time.Minute*minutes + (time.Microsecond)*time.Duration(rand.Intn(10000000))) // (rand.Int31())))
		total += n
	}

	elasped := time.Since(start)
	fmt.Printf("%s uploaded in %.1f seconds (%.1f Gbit/sec)\n", humanize.IBytes(uint64(total)), elasped.Seconds(), float64(total*8/1024/1024/1024)/elasped.Seconds())
}
