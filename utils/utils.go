package utils

import (
	"fmt"
	"net/url"
	"regexp"
)


func DecodeUrl(str string) string{
	decodedURL, err := url.QueryUnescape(str)

	if err != nil {
		fmt.Println("Error decoding URL:", err)
		return ""
	 } 

	return decodedURL 
}

func ParseId(str string) string{
    pattern := `signup_job_id=(\d+)`

    re := regexp.MustCompile(pattern)

    match := re.FindStringSubmatch(str)

    if len(match) > 1 {
        jobID := match[1] 
		return jobID
    } 
	return ""
}