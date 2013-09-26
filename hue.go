package hue

import (
	"fmt"
	"strings"
	"time"
)

type Api interface {
	GetLights() ([]Light, error)
	GetNewLights() ([]Light, time.Time, error)
	SearchForNewLights() error
	GetLightAttributes(lightId string) (*LightAttributes, error)
	SetLightName(lightId string, name string) error
	SetLightState(lightId string, state LightState) error
	GetConfiguration() (*Configuration, error)

	// TODO: Configuration Api Methods:
	// DeleteUser(user *User) error
	// UpdateConfiguration(*Configuration) error
	// GetDataStore() (DataStore, error)

	// TODO: Groups Api Methods
	// GetGroups() ([]Group, error)
	// GetGroupAttributes(groupId string) (*GroupAttributes, error)
	// SetGroupAttributes(groupId string, attr GroupAttributes) error
	// SetGroupState(groupId string, state GroupState) error
	// CreateGroup - not supported in current hue api
	// DeleteGroup - not supported in current hue api

	// TODO: Schedule Api Methods
	// GetSchedules() ([]Schedule, error)
	// CreateSchedule(schedule Schedule) error
	// GetScheduleAttributes(scheduleId string) (*ScheduleAttributes, error)
	// SetScheduleAttributes(scheduleId string, attr ScheduleAttributes)
	// DeleteSchedule(scheduleId string) error
}

type Light struct {
	Id   string
	Name string
}

type LightState struct {
	On             *bool     `json:"on,omitempty"`
	Brightness     *uint8    `json:"bri,omitempty"`
	Hue            *uint16   `json:"hue,omitempty"`
	Saturation     *uint8    `json:"sat,omitempty"`
	Xy             []float64 `json:"xy,omitempty"`
	ColorTemp      *uint16   `json:"ct,omitempty"`
	Alert          string    `json:"alert,omitempty"`
	Effect         string    `json:"effect,omitempty"`
	TransitionTime *uint16   `json:"transitiontime,omitempty"` /* write only */
	ColorMode      string    `json:"colormode,omitempty"`      /* read only */
	Reachable      bool      `json:"reachable,omitempty"`      /* read only */
}

type LightAttributes struct {
	Name            string      `json:"name"`
	State           *LightState `json:"state"`
	Type            string      `json:"type"`
	ModelId         string      `json:"modelid"`
	SoftwareVersion string      `json:"swversion"`
	PointSymbol     interface{} `json:"pointsymbol"`
}

type Configuration struct {
	// Read and Write
	Name           string              `json:"name,omitempty"`
	ProxyAddress   string              `json:"proxyaddress,omitempty"`
	ProxyPort      *uint16             `json:"proxyport,omitempty"`
	IpAddress      string              `json:"ipaddress,omitempty"`
	Netmask        string              `json:"netmask,omitempty"`
	Gateway        string              `json:"gateway,omitempty"`
	Dhcp           *bool               `json:"dhcp,omitempty"`
	PortalServices *bool               `json:"portalservices,omitempty"`
	LinkButton     *bool               `json:"linkbutton,omitempty"`
	SoftwareUpdate *SoftwareUpdateInfo `json:"swupdate,omitempty"`

	// Read only
	Utc             string                       `json:"utc,omitempty"`
	Whitelist       map[string]map[string]string `json:"whitelist,omitempty"`
	SoftwareVersion string                       `json:"swversion,omitempty"`
	Mac             string                       `json:"mac,omitempty"`
}

type SoftwareUpdateInfo struct {
	UpdateState *uint  `json:"updatestate,omitempty"`
	Url         string `json:"url,omitempty"`
	Text        string `json:"text,omitempty"`
	Notify      *bool  `json:"notify,omitempty"`
}

const (
	UnauthorizedUserErrorType       = 1
	InvalidJsonErrorType            = 2
	ResourceNotAvailableErrorType   = 3
	MethodNotAvailableErrorType     = 4
	MissingParameterErrorType       = 5
	ParameterNotAvailableErrorType  = 6
	InvalidParameterValueErrorType  = 7
	ParameterNotModifiableErrorType = 8
	InternalErrorType               = 901
	LinkButtonNotPressedErrorType   = 101
	DeviceIsOffErrorType            = 201
	GroupTableFullErrorType         = 301
	DeviceGroupTableFullErrorType   = 302
)

type ApiError struct {
	Errors []ApiErrorDetail
}

func (e ApiError) Error() string {
	errors := make([]string, 0, 10)
	for _, error := range e.Errors {
		errors = append(errors, error.Error())
	}

	return strings.Join(errors, " ")
}

type ApiErrorDetail struct {
	Type        int    `json:"type"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

func (e ApiErrorDetail) Error() string {
	return fmt.Sprintf("Hue API Error type '%d' with description '%s'.", e.Type, e.Description)
}