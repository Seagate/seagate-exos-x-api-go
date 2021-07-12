/*
 * Copyright (c) 2021 Enix, SAS
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 *
 * Authors:
 * Paul Laffitte <paul.laffitte@enix.fr>
 * Alexandre Buisine <alexandre.buisine@enix.fr>
 */

package mock

import (
	"net/http"

	"github.com/Seagate/seagate-exos-x-api-go/internal/mock/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	auth := controllers.NewAuthController()

	api := router.Group("api")
	{
		api.GET("/login/:hash", auth.Login)
	}

	router.Use(func(c *gin.Context) {
		if sessionKey, ok := c.Request.Header["Sessionkey"]; ok && len(sessionKey) > 0 {
			if _, ok := auth.Tokens[sessionKey[0]]; !ok {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}
	})

	router.NoRoute(func(c *gin.Context) {
		c.String(400, "")
	})

	return router
}
