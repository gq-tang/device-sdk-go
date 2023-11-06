/**
@description provision文件

@copyright    Copyright 2023 seva
@version      1.0.0
@author       tgq <tangguangqiang@rollingstoneiot.com>
@datetime     2023/11/6 13:42
*/

package provision

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gq-tang/device-sdk-go/v2/internal/cache"

	bootstrapContainer "github.com/edgexfoundry/go-mod-bootstrap/v2/bootstrap/container"
	"github.com/edgexfoundry/go-mod-bootstrap/v2/di"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
)

func LoadProvisionWatchers(path string, dic *di.Container) errors.EdgeX {
	if path == "" {
		return nil
	}
	lc := bootstrapContainer.LoggingClientFrom(dic.Get)

	absPath, err := filepath.Abs(path)
	if err != nil {
		return errors.NewCommonEdgeX(errors.KindServerError, "failed to create absolute path", err)
	}

	fileInfo, err := os.ReadDir(absPath)
	if err != nil {
		return errors.NewCommonEdgeX(errors.KindServerError, "failed to read directory", err)
	}

	var addProvsionWatchersReq []requests.AddProvisionWatcherRequest
	pwc := bootstrapContainer.ProvisionWatcherClientFrom(dic.Get)
	lc.Infof("Loading provision watcher from %s", absPath)
	for _, file := range fileInfo {
		var pw dtos.ProvisionWatcher
		fullPath := filepath.Join(absPath, file.Name())
		if strings.HasSuffix(fullPath, yamlExt) || strings.HasSuffix(fullPath, ymlExt) {
			content, err := os.ReadFile(fullPath)
			if err != nil {
				lc.Errorf("Failed to read %s: %v", fullPath, err)
				continue
			}

			err = yaml.Unmarshal(content, &pw)
			if err != nil {
				lc.Errorf("Failed to decode profile %s: %v", file.Name(), err)
				continue
			}
		} else if strings.HasSuffix(fullPath, jsonExt) {
			content, err := os.ReadFile(fullPath)
			if err != nil {
				lc.Errorf("Failed to read %s: %v", fullPath, err)
				continue
			}

			err = json.Unmarshal(content, &pw)
			if err != nil {
				lc.Errorf("Failed to decode profile %s: %v", file.Name(), err)
				continue
			}
		} else {
			continue
		}

		res, err := pwc.ProvisionWatcherByName(context.Background(), pw.Name)
		if err == nil {
			lc.Infof("Provision watcher %s exists, using the existing one", pw.Name)
			_, exist := cache.ProvisionWatchers().ForName(pw.Name)
			if !exist {
				err = cache.ProvisionWatchers().Add(dtos.ToProvisionWatcherModel(res.ProvisionWatcher))
				if err != nil {
					return errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("failed to cache the profile %s", res.ProvisionWatcher.Name), err)
				}
			}
		} else {
			lc.Infof("Provision watcher %s not found in Metadata, adding it ...", pw.Name)
			req := requests.NewAddProvisionWatcherRequest(pw)
			addProvsionWatchersReq = append(addProvsionWatchersReq, req)
		}
	}

	if len(addProvsionWatchersReq) == 0 {
		return nil
	}
	ctx := context.WithValue(context.Background(), common.CorrelationHeader, uuid.NewString()) // nolint:staticcheck
	_, edgexErr := pwc.Add(ctx, addProvsionWatchersReq)
	return edgexErr
}
