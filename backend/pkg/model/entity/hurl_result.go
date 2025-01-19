package entity

type HurlTestResult struct {
	Entries  []Entries `json:"entries"`
	Filename string    `json:"filename"`
	Success  bool      `json:"success"`
	Time     int       `json:"time"`
}

type Entries struct {
	CurlCmd string `json:"curl_cmd"`
	Index   int    `json:"index"`
	Line    int    `json:"line"`
	Time    int    `json:"time"`
}
