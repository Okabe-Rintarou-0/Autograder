package canvas

import (
	"autograder/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"

	"autograder/pkg/config"
	"autograder/pkg/model/entity/canvas"
)

const (
	BASEURL = "https://oc.sjtu.edu.cn"
)

type DAOImpl struct {
	innerCli *http.Client
	token    string
}

func NewDAO() *DAOImpl {
	return &DAOImpl{
		innerCli: &http.Client{},
		token:    config.Instance.CanvasToken,
	}
}

func (c *DAOImpl) ListCourses(ctx context.Context) ([]*canvas.Course, error) {
	URL := fmt.Sprintf(
		"%s/api/v1/courses?include[]=teachers&include[]=term",
		BASEURL,
	)
	course, err := listItems[canvas.Course](c.innerCli, URL, c.token)
	if err != nil {
		return nil, err
	}
	return utils.Filter(course, func(v *canvas.Course) bool {
		if v.AccessRestrictedByDate != nil && *v.AccessRestrictedByDate {
			return false
		}
		for _, enrollment := range v.Enrollments {
			if enrollment != nil &&
				(enrollment.Role == canvas.TaEnrollment || enrollment.Role == canvas.TeacherEnrollment) {
				return true
			}
		}
		return false
	}), nil
}

func (c *DAOImpl) ListCourseUsers(ctx context.Context, courseID int64) ([]*canvas.User, error) {
	URL := fmt.Sprintf(
		"%s/api/v1/courses/%d/users",
		BASEURL, courseID,
	)
	return listItems[canvas.User](c.innerCli, URL, c.token)
}

func listItems[T any](cli *http.Client, URL, token string) ([]*T, error) {
	var allItems []*T
	page := 1
	for {
		items, err := listItemsWithPage[T](cli, URL, token, page)
		if err != nil {
			return nil, err
		}
		if len(items) == 0 {
			break
		}
		page++
		allItems = append(allItems, items...)
	}

	return allItems, nil
}

func listItemsWithPage[T any](cli *http.Client, URL, token string, page int) ([]*T, error) {
	query := url.Values{}
	query.Add("page", fmt.Sprint(page))
	query.Add("per_page", "100")

	return get[[]*T](cli, URL, token, &query)
}

func get[T any](cli *http.Client, URL, token string, query *url.Values) (T, error) {
	var data T
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		logrus.Errorf("[Canvas Client] get request error: %v", err)
		return data, err
	}
	parsedURL, err := url.Parse(URL)
	if err != nil {
		return data, err
	}
	passedQuery := parsedURL.Query()
	if query != nil {
		if passedQuery != nil {
			for key, vals := range passedQuery {
				for _, val := range vals {
					query.Add(key, val)
				}
			}
		}
		req.URL.RawQuery = query.Encode()
	}
	logrus.Infof("[Canvas Client] get request: %v", req.URL)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := cli.Do(req)
	if err != nil {
		logrus.Errorf("[Canvas Client] get request error: %v", err)
		return data, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		logrus.Errorf("[Canvas Client] decode response body error: %v", err)
		return data, err
	}
	return data, nil
}
