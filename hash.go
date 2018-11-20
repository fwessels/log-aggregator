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
	"encoding/hex"
	"github.com/minio/highwayhash"
	"log"
)

// Return hash of buffer
func hashBuffer(buffer []byte) string {
	key, err := hex.DecodeString("000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f")
	if err != nil {
		log.Fatalf("Failed to decode key: %v", err)
	}

	//
	// Alternatively use highwayhash.New128 for key of 32 characters
	//
	h, err := highwayhash.New64(key[:])
	if err != nil {
		log.Fatalf("Failed to create highwayhash instance:  %v", err)
	}

	// Write contents to hash algorithm
	h.Write(buffer)

	return hex.EncodeToString(h.Sum(nil))
}
