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
	"bytes"
	"errors"
	"fmt"

	minio "github.com/minio/minio-go"
)

const endpoint = "localhost:9000"
const accessKey = "U09DX0NB71K0HBYRQGVD"
const secretKey = "PsqdAhZD9tYkZ4+eQL0lOywKMoWnGTDpQNeoF47K"

const bucket = "test"

// Upload a buffer to object storage
func uploadBuffer(name string, buffer []byte) (n int64, err error) {

	// New returns an Amazon S3 compatible client object. API compatibility (v2 or v4) is automatically
	// determined based on the Endpoint value.
	s3Client, err := minio.New(endpoint, accessKey, secretKey, false)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	n, err = s3Client.PutObject(bucket, name, bytes.NewReader(buffer), int64(len(buffer)), minio.PutObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return 0, err
	} else if n < int64(len(buffer)) {
		return 0, errors.New("Error while uploading")
	}

	return n, nil
}
