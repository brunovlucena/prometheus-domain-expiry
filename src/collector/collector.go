package collector

// Prometheur Exporter - Collector.
// Author: Bruno Lucena <bvlg900f@gmail.com>

import (
	"fmt"
	"math"
	"os/exec"
	"regexp"
	"time"

	"github.com/brunovlucena/prometheus-domain-expiry/src/utils"
)

// VerifyExpire ...
func VerifyExpire(host string) int {
	// Get expiration Date
	command := fmt.Sprintf("whois %s | %s", host, "grep 'Registry Expiry Date'")
	out, err := exec.Command("bash", "-c", command).Output()
	// Error
	utils.FailOnError(err, "An error occured")
	// Convert
	return convertDate(string(out))
}

// converts the Registry Expire Date to days
func convertDate(out string) int {
	// E.g. Registry Expiry Date: 2020-09-14T04:00:00Z
	r := regexp.MustCompile(`\d.*`)
	// Match
	out = r.FindString(out)
	format := "2006-01-02T15:04:05Z"
	then, _ := time.Parse(format, out)
	date := time.Now()
	diff := date.Sub(then)
	return int(math.Abs(diff.Hours() / 24))
}
