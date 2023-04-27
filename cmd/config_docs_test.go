// Copyright © 2023 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
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

//go:build !docs
// +build !docs

package cmd

import (
	"context"
	"crypto/sha1"
	"os"
	"path/filepath"
	"testing"

	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestConfigDocsUpToDate(t *testing.T) {
	// Initialize config of all plugins
	InitConfig()
	generatedConfig, err := config.GenerateConfigMarkdown(context.Background(), "", config.GetKnownKeys())
	assert.NoError(t, err)
	configOnDisk, err := os.ReadFile(filepath.Join("..", "config.md"))
	assert.NoError(t, err)

	generatedConfigHash := sha1.New()
	generatedConfigHash.Write(generatedConfig)
	configOnDiskHash := sha1.New()
	configOnDiskHash.Write(configOnDisk)
	assert.Equal(t, configOnDiskHash.Sum(nil), generatedConfigHash.Sum(nil), "The config reference docs generated by the code did not match the config.md file in git. Did you forget to run `make docs`?")
}
