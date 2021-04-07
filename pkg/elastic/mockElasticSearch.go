package elastic

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ViaQ/log-exploration-api/pkg/logs"
)

type MockedElasticsearchProvider struct {
	App   map[time.Time][]string
	Infra map[time.Time][]string
	Audit map[time.Time][]string
}

func NewMockedElastisearchProvider() *MockedElasticsearchProvider {
	return &MockedElasticsearchProvider{
		App:   map[time.Time][]string{},
		Infra: map[time.Time][]string{},
		Audit: map[time.Time][]string{},
	}
}

func (m *MockedElasticsearchProvider) PutDataIntoIndex(index string, data []string) error {
	switch strings.ToLower(index) {
	case "app":
		m.App[time.Now()] = data
		return nil
	case "infra":
		m.Infra[time.Now()] = data
		return nil
	case "audit":
		m.Audit[time.Now()] = data
		return nil
	default:
		return errors.New("unknown index")
	}
}

func (m *MockedElasticsearchProvider) PutDataAtTime(logTime time.Time, index string, data []string) error {
	switch strings.ToLower(index) {
	case "app":
		m.App[logTime] = data
		return nil
	case "infra":
		m.Infra[logTime] = data
		return nil
	case "audit":
		m.Audit[logTime] = data
		return nil
	default:
		return errors.New("unknown index")
	}
}

func (m *MockedElasticsearchProvider) FilterByIndex(index string) ([]string, error) {
	var lg map[time.Time][]string
	switch strings.ToLower(index) {
	case "app":
		lg = m.App
	case "infra":
		lg = m.Infra
	case "audit":
		lg = m.Audit
	}
	if len(lg) == 0 {
		return nil, logs.NotFoundError()
	}

	result := []string{}
	for _, v := range lg {
		result = append(result, v...)
	}
	return result, nil
}

func (m *MockedElasticsearchProvider) FilterByTime(startTime time.Time, finishTime time.Time) ([]string, error) {
	lg := make(map[time.Time][]string)
	for k, v := range m.App {
		if v != nil {
			lg[k] = v
		}
	}
	for k, v := range m.Infra {
		if v != nil {
			lg[k] = v
		}
	}
	for k, v := range m.Audit {
		if v != nil {
			lg[k] = v
		}
	}

	if len(lg) == 0 {
		return nil, logs.NotFoundError()
	}

	result := []string{}
	for k, v := range lg {
		if k.After(startTime) && k.Before(finishTime) {
			result = append(result, v...)
		}
	}
	return result, nil
}

func (m *MockedElasticsearchProvider) GetAllLogs() ([]string, error) {
	lg := make(map[time.Time][]string)
	for k, v := range m.App {
		if v != nil {
			lg[k] = v
		}
	}
	for k, v := range m.Infra {
		if v != nil {
			lg[k] = v
		}
	}
	for k, v := range m.Audit {
		if v != nil {
			lg[k] = v
		}
	}

	if len(lg) == 0 {
		return nil, logs.NotFoundError()
	}

	result := []string{}
	for _, v := range lg {
		result = append(result, v...)
	}

	return result, nil
}

func (m *MockedElasticsearchProvider) FilterByPodName(podName string) ([]string, error) {
	lg := make(map[time.Time][]string)
	for k, v := range m.App {
		if v != nil {
			lg[k] = v
		}
	}
	for k, v := range m.Infra {
		if v != nil {
			lg[k] = v
		}
	}
	for k, v := range m.Audit {
		if v != nil {
			lg[k] = v
		}
	}

	if len(lg) == 0 {
		return nil, logs.NotFoundError()
	}

	result := []string{}
	for _, v := range lg {
		if strings.Contains(v[0], "pod_name: "+podName) {
			result = append(result, v...)
		}
	}
	if len(result) == 0 {
		return nil, logs.NotFoundError()
	}

	return result, nil
}

func (m *MockedElasticsearchProvider) FilterLogsMultipleParameters(podName string, namespace string, startTime time.Time, finishTime time.Time) ([]string, error) {
	lg := make(map[time.Time][]string)
	for k, v := range m.App {
		if v != nil {
			lg[k] = v
		}
	}
	for k, v := range m.Infra {
		if v != nil {
			lg[k] = v
		}
	}
	for k, v := range m.Audit {
		if v != nil {
			lg[k] = v
		}
	}

	if len(lg) == 0 {
		return nil, logs.NotFoundError()
	}

	result := []string{}
	for k, v := range lg {
		pod := "pod_name: " + podName
		ns := "namespace_name: " + namespace
		fmt.Print(pod, ns)
		if k.After(startTime) && k.Before(finishTime) &&
			strings.Contains(v[0], pod) &&
			strings.Contains(v[0], ns) {
			result = append(result, v...)
		}
	}

	if len(result) == 0 {
		return nil, logs.NotFoundError()
	}

	return result, nil
}
