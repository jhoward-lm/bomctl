// -----------------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 bomctl a Series of LF Projects, LLC
// SPDX-FileName: internal/pkg/client/client.go
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// -----------------------------------------------------------------------------
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// -----------------------------------------------------------------------------

package client

import (
	"errors"

	"github.com/bomctl/bomctl/internal/pkg/netutil"
	"github.com/bomctl/bomctl/internal/pkg/options"
)

var ErrUnsupportedURL = errors.New("failed to parse URL; see `--help` for valid URL patterns")

type (
	Client interface {
		netutil.Parser
		Name() string
	}

	Fetcher interface {
		Client
		Fetch(fetchURL string, opts *options.FetchOptions) ([]byte, error)
		PrepareFetch(url *netutil.URL, auth *netutil.BasicAuth, opts *options.Options) error
	}

	Pusher interface {
		Client
		AddFile(pushURL, id string, opts *options.PushOptions) error
		PreparePush(pushURL string, opts *options.PushOptions) error
		Push(pushURL string, opts *options.PushOptions) error
	}
)
