package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ghophp/buildbot-dashboard/container"
	"github.com/ghophp/render"
)

const (
	buildingState   string = "building"
	failedState     string = "failed"
	successfulState string = "successful"
	warningState    string = "warnings"
	exceptionState  string = "exception"
)

var validStates []string = []string{
	buildingState,
	failedState,
	successfulState,
	warningState,
	exceptionState,
}

type (
	BuildersHandler struct {
		c *container.ContainerBag
	}

	Builder struct {
		Id         string   `json:"id"`
		State      string   `json:"state"`
		Reason     string   `json:"reason"`
		Blame      []string `json:"blame"`
		Number     int      `json:"number"`
		Slave      string   `json:"slave"`
		LastUpdate string   `json:"last_update"`
	}

	DetailedBuilder struct {
		Blame  []string  `json:"blame"`
		Number int       `json:"number"`
		Reason string    `json:"reason"`
		Slave  string    `json:"slave"`
		Times  []float64 `json:"times"`
		Text   []string  `json:"text"`
		Error  string    `json:"error"`
	}
)

func NewBuildersHandler(c *container.ContainerBag) *BuildersHandler {
	return &BuildersHandler{
		c: c,
	}
}

func isValidState(v string) bool {
	for _, s := range validStates {
		if v == s {
			return true
		}
	}
	return false
}

func GetBuilder(c *container.ContainerBag, id string, builder Builder) (Builder, error) {
	var b map[string]DetailedBuilder

	data, err := c.Buildbot.FetchBuilder(id)
	if err != nil {
		return builder, err
	}

	if err := json.Unmarshal(data, &b); err != nil {
		return builder, err
	}

	if current, ok := b["-1"]; ok && current.Error == "" {

		builder.Id = id
		builder.Blame = current.Blame
		builder.Number = current.Number
		builder.Slave = current.Slave
		builder.Reason = current.Reason
		builder.State = buildingState
		builder.LastUpdate = strconv.Itoa(int(time.Now().Unix()))

		if len(current.Times) > 0 {
			builder.LastUpdate = strconv.FormatFloat(current.Times[0], 'f', 6, 64)
		}

		if len(current.Text) > 0 {
			for _, v := range current.Text {
				if isValidState(v) {
					builder.State = v
					break
				}
			}
		}

		return builder, nil
	}

	return builder, fmt.Errorf("[GetBuilder] %s", "no last build defined")
}

// GetBuilders will fetch the builders from buildbot, if the fresh parameter is equal true
// it will not respect cache
func GetBuilders(c *container.ContainerBag, fresh bool) (map[string]Builder, error) {
	var data map[string]Builder

	dataBytes, err := c.Cache.GetCache(c.HashedUrl)
	if fresh || err != nil {
		dataBytes, err = c.Buildbot.FetchBuilders()
		if err != nil {
			return nil, err
		}

		if err := c.Cache.SetCache(c.HashedUrl, dataBytes); err != nil {
			return nil, err
		}
	}

	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return nil, err
	}

	if c.FilterRegex != nil {
		for key, _ := range data {
			if !c.FilterRegex.MatchString(key) {
				delete(data, key)
			}
		}
	}

	return data, nil
}

func (h BuildersHandler) ServeHTTP(req *http.Request, r render.Render) {
	fresh, err := strconv.ParseBool(req.URL.Query().Get("fresh"))
	if err != nil {
		fresh = false
	}

	if builders, err := GetBuilders(h.c, fresh); err == nil {
		r.JSON(200, builders)
	} else {
		r.Error(500)
	}
}
