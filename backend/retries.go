package backend

import (
	"fmt"
	"net/http"
	"time"
)

type RetryPolicyStep struct {
	// if OnStatusCode.len == 1 it is only one code
	// if == 2 it is an interval. A negative number on the interval is for exclude it from it (example: [-404, 429] is the interval: ]404; 429]
	// The numbers can't be < 300
	OnStatusCode []int

	NbOfRetry int

	TimeInterval time.Duration
}

func (retry RetryPolicyStep) DoesCodeStatusNeedRetry(code int) bool {
	if len(retry.OnStatusCode) == 1 {
		return retry.OnStatusCode[0] == code
	} else if len(retry.OnStatusCode) == 2 {
		left := (retry.OnStatusCode[0] < 0 && -retry.OnStatusCode[0] < code) || (retry.OnStatusCode[0] > 0 && retry.OnStatusCode[0] <= code)
		right := (retry.OnStatusCode[1] < 0 && -retry.OnStatusCode[1] > code) || (retry.OnStatusCode[1] > 0 && retry.OnStatusCode[1] >= code)
		return left && right
	} else {
		panic("Error Usage. retry.OnStatusCode should have a size that is 1 or 2")
	}
}

var defaultPolicy = RetryPolicyStep{
	OnStatusCode: []int{400, 1000},
	NbOfRetry:    1,
	TimeInterval: 1 * time.Second,
}

type RetryPolicy []RetryPolicyStep

func MakeRetryPolicy(retryPolicies ...RetryPolicyStep) RetryPolicy {
	return append(retryPolicies, defaultPolicy)
}
func (retry RetryPolicy) RunPolicy(call func() (*http.Response, error), willResp WillSet[*http.Response], willErr WillSet[error]) {
	go func() {
		resp, err := call()
		nbTry := 0
		for {
			var selectedRetry *RetryPolicyStep = nil
			for _, retryPolicyStep := range retry {
				if retryPolicyStep.DoesCodeStatusNeedRetry(resp.StatusCode) {
					receivedData := make([]byte, 2048)
					n, _ := resp.Body.Read(receivedData)
					content := string(receivedData[:n])
					if n > 150 {
						content = content[:150] + "\n..."
					}
					fmt.Printf("CODE %v\nReceived: %v\n", resp.StatusCode, content)
					selectedRetry = &retryPolicyStep
					time.Sleep(retryPolicyStep.TimeInterval)
					goto Next // we have a selected retry because we have an error
				}
			}
			break // we go out of the for loop without any error, we can go out of that loop
		Next:
			resp, err = call()
			nbTry++
			if nbTry >= selectedRetry.NbOfRetry {
				break
			}
		}
		willResp.Set(resp)
		willErr.Set(err)
	}()
}
