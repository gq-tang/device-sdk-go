// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
// Copyright (C) 2018-2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package startup

import (
	"context"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"

	"github.com/gorilla/mux"
	"github.com/gq-tang/device-sdk-go/v2/pkg/service"
)

func Bootstrap(serviceName string, serviceVersion string, driver interface{}, lc logger.LoggingClient) {
	ctx, cancel := context.WithCancel(context.Background())
	service.Main(serviceName, serviceVersion, driver, ctx, cancel, mux.NewRouter(), lc)
}
