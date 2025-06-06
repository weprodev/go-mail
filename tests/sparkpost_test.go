// Copyright 2022 Ainsley Clark. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mail

import (
	"os"
	"testing"

	"github.com/weprodev/go-mail/drivers"
	"github.com/weprodev/go-mail/mail"
)

func Test_SparkPost(t *testing.T) {
	LoadEnv(t)
	cfg := mail.Config{
		URL:         os.Getenv("SPARKPOST_URL"),
		APIKey:      os.Getenv("SPARKPOST_API_KEY"),
		FromAddress: os.Getenv("SPARKPOST_FROM_ADDRESS"),
		FromName:    os.Getenv("SPARKPOST_FROM_NAME"),
	}
	UtilTestSend(t, drivers.NewSparkPost, cfg, "SparkPost")
}
