// Copyright 2025 GEEKROS, Inc.
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

package handler

import (
	"github.com/geekros/structure/pkg/empty"
	"github.com/geekros/utils/pkg/authentication"
	"github.com/geekros/utils/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthToken(c *gin.Context) {

	responseData := responseAgentIndex{}

	roleType := c.DefaultQuery("role_type", "")
	if roleType == "" {
		response.Error(c, empty.EmptyData{})
	}

	data := map[string]interface{}{
		"role": "",
	}

	token, err := authentication.GenerateToken("", data, 24)
	if err != nil {
		response.Warning(c, 10000, "", empty.EmptyData{})
	}

	response.Success(c, empty.EmptyData{})
}
