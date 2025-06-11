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

package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/geekros/server/pkg/router"
	"github.com/gookit/color"
)

type Server struct {
	HttpServer *http.Server
	Router     *router.Router
}

func New() *Server {
	Router := router.New()
	return &Server{
		Router: Router,
	}
}

func (s *Server) Start(port int, mode string, read time.Duration, write time.Duration, callback func(), exit func()) {
	s.HttpServer = &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        s.Router.Init(mode).InitHandler(),
		ReadTimeout:    read * time.Second,
		WriteTimeout:   write * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.HttpServer.ListenAndServe(); err != nil {
			log.Println(color.Yellow.Text(fmt.Sprintf("%v", err)))
		}
	}()

	callback()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	exit()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.HttpServer.Shutdown(ctx); err != nil {
		log.Println(color.Gray.Text(fmt.Sprintf("%v", err)))
	}
}
