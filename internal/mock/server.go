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
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const SettingsFile = ".env"

func StartServer() {
	LoadEnv()

	r := NewRouter()
	r.Run()
}

func LoadEnv() {
	// Note, any defined environment variable is used over the ones defined in .env
	if _, err := os.Stat(SettingsFile); err == nil {
		fmt.Printf("Testing setup: Loading (%s)\n", SettingsFile)
		err := godotenv.Load(SettingsFile)
		if err != nil {
			fmt.Printf("Error loading file (%s), error: %v\n", SettingsFile, err)
			return
		}
	}
}
