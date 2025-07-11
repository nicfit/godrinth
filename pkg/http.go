package godrinth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type HttpError struct {
	Code int
	Msg  string
}

func makeHttpError(response *http.Response) HttpError {
	return HttpError{
		Code: response.StatusCode,
		Msg:  response.Status,
	}
}

func (e HttpError) Error() string {
	return e.Msg
}

func GetMeta(_ context.Context) (Meta, error) {
	var meta = Meta{}

	req, err := http.NewRequest("GET", ModrinthUrl, nil)
	addHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return meta, err
	} else if resp.StatusCode != http.StatusOK {
		return meta, makeHttpError(resp)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return meta, err
	}
	if err := json.Unmarshal(data, &meta); err != nil {
		return meta, err
	}

	return meta, nil
}

func GetProject(_ context.Context, q string) (Project, error) {
	var project = Project{}

	url_q := fmt.Sprintf("%sproject/%s", ModrinthApiUrl, q)
	req, err := http.NewRequest("GET", url_q, nil)
	addHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return project, err
	} else if resp.StatusCode != http.StatusOK {
		return project, makeHttpError(resp)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return project, err
	}
	if err := json.Unmarshal(data, &project); err != nil {
		return project, err
	}

	return project, nil
}

func GetProjects(_ context.Context, ids []string) ([]Project, error) {
	// TODO
	return []Project{}, nil
}

type SearchOptions struct {
	Facets string
	Index  string
	Offset int
	Limit  int
}

func SearchProject(_ context.Context, query string, options *SearchOptions) (SearchResults, error) {
	var results = SearchResults{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%ssearch", ModrinthApiUrl), nil)
	if err != nil {
		return results, err
	}
	addHeaders(req)
	req.URL.RawQuery = url.Values{
		"query":  {query},
		"facets": {options.Facets},
		"limit":  {strconv.Itoa(options.Limit)},
		"offset": {strconv.Itoa(options.Offset)},
		"index":  {options.Index},
	}.Encode()

	client := &http.Client{}
	fmt.Println(req.URL) // FIXME
	resp, err := client.Do(req)
	if err != nil {
		return results, err
	} else if resp.StatusCode != http.StatusOK {
		return results, makeHttpError(resp)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return results, err
	}
	if err := json.Unmarshal(data, &results); err != nil {
		return results, err
	}

	return results, nil
}

func addHeaders(req *http.Request) {
	req.Header.Add("User-Agent", UserAgent)
}
