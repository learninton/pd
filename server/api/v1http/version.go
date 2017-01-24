// Copyright 2017 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package v1http

import (
	"net/http"

	"github.com/unrolled/render"
)

type version struct {
	Version string `json:"version"`
}

type versionHandler struct {
	rd *render.Render
}

func newVersionHandler(rd *render.Render) *versionHandler {
	return &versionHandler{
		rd: rd,
	}
}

func (h *versionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	version := &version{
		Version: "1.0.0",
	}
	h.rd.JSON(w, http.StatusOK, version)
}