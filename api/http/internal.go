// Copyright 2022 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package http

import (
	"encoding/json"
	"net/http"

	"github.com/mendersoftware/iot-manager/app"
	"github.com/mendersoftware/iot-manager/client"
	"github.com/mendersoftware/iot-manager/model"

	"github.com/gin-gonic/gin"
	"github.com/mendersoftware/go-lib-micro/identity"
	"github.com/mendersoftware/go-lib-micro/rest.utils"
	"github.com/pkg/errors"
)

const (
	ParamTenantID = "tenant_id"
	ParamDeviceID = "device_id"
)

type InternalHandler APIHandler

type internalDevice model.DeviceEvent

func (dev *internalDevice) UnmarshalJSON(b []byte) error {
	type deviceAlias struct {
		// device_id kept for backward compatibility
		ID string `json:"device_id"`
		model.DeviceEvent
	}
	var aDev deviceAlias
	err := json.Unmarshal(b, &aDev)
	if err != nil {
		return err
	}
	if aDev.ID != "" {
		aDev.DeviceEvent.ID = aDev.ID
	}
	*dev = internalDevice(aDev.DeviceEvent)
	return nil
}

// POST /tenants/:tenant_id/devices
// code: 204 - device provisioned to iothub
//
//	500 - internal server error
func (h *InternalHandler) ProvisionDevice(c *gin.Context) {
	tenantID := c.Param(ParamTenantID)
	var device internalDevice
	if err := c.ShouldBindJSON(&device); err != nil {
		rest.RenderError(c,
			http.StatusBadRequest,
			errors.Wrap(err, "malformed request body"))
		return
	}
	if device.ID == "" {
		rest.RenderError(c, http.StatusBadRequest, errors.New("missing device ID"))
		return
	}

	ctx := identity.WithContext(c.Request.Context(), &identity.Identity{
		Subject: device.ID,
		Tenant:  tenantID,
	})
	err := h.app.ProvisionDevice(ctx, model.DeviceEvent(device))
	switch cause := errors.Cause(err); cause {
	case nil, app.ErrNoCredentials:
		c.Status(http.StatusNoContent)
	case app.ErrDeviceAlreadyExists:
		rest.RenderError(c, http.StatusConflict, cause)
	default:
		rest.RenderError(c, http.StatusInternalServerError, err)
	}
}

func (h *InternalHandler) DecommissionDevice(c *gin.Context) {
	deviceID := c.Param(ParamDeviceID)
	tenantID := c.Param(ParamTenantID)

	ctx := identity.WithContext(c.Request.Context(), &identity.Identity{
		Subject: deviceID,
		Tenant:  tenantID,
	})
	err := h.app.DecommissionDevice(ctx, deviceID)
	switch errors.Cause(err) {
	case nil, app.ErrNoCredentials:
		c.Status(http.StatusNoContent)
	case app.ErrDeviceNotFound:
		rest.RenderError(c, http.StatusNotFound, err)
	default:
		rest.RenderError(c, http.StatusInternalServerError, err)
	}
}

type BulkResult struct {
	Error bool       `json:"error"`
	Items []BulkItem `json:"items"`
}

type BulkItem struct {
	// Status code for the operation (translates to HTTP status)
	Status int `json:"status"`
	// Description in case of error
	Description string `json:"description,omitempty"`
	// Parameters used for producing BulkItem
	Parameters map[string]interface{} `json:"parameters"`
}

const (
	maxBulkItems = 100
)

// PUT /tenants/:tenant_id/devices/status/{status}
func (h *InternalHandler) BulkSetDeviceStatus(c *gin.Context) {
	var schema []struct {
		DeviceID string `json:"id"`
	}
	status := model.Status(c.Param("status"))
	if err := status.Validate(); err != nil {
		rest.RenderError(c, http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBindJSON(&schema); err != nil {
		rest.RenderError(c,
			http.StatusBadRequest,
			errors.Wrap(err, "invalid request body"),
		)
		return
	} else if len(schema) > maxBulkItems {
		rest.RenderError(c,
			http.StatusBadRequest,
			errors.New("too many bulk items: max 100 items per request"),
		)
		return
	}
	ctx := identity.WithContext(
		c.Request.Context(),
		&identity.Identity{
			Tenant: c.Param("tenant_id"),
		},
	)
	res := BulkResult{
		Error: false,
		Items: make([]BulkItem, len(schema)),
	}
	for i, item := range schema {
		res.Items[i].Parameters = map[string]interface{}{
			"id": item.DeviceID,
		}
		err := h.app.SetDeviceStatus(ctx, item.DeviceID, status)
		if err != nil {
			res.Error = true
			if e, ok := errors.Cause(err).(client.HTTPError); ok {
				res.Items[i].Status = e.Code()
				res.Items[i].Description = e.Error()
			} else {
				res.Items[i].Status = http.StatusInternalServerError
				res.Items[i].Description = err.Error()
			}
		} else {
			res.Items[i].Status = http.StatusOK
		}
	}
	c.JSON(http.StatusOK, res)
}
