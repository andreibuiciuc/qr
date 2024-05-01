package qr

import (
	"fmt"
	"regexp"
	"strconv"
)

type versioner struct{}

const (
	mode_NUMERIC      string = "numeric"
	mode_ALPHANUMERIC string = "alphanumeric"
	mode_BYTE         string = "byte"
)

const (
	ec_LOW      rune = 'L'
	ec_MEDIUM   rune = 'M'
	ec_QUARTILE rune = 'Q'
	ec_HIGH     rune = 'H'
)

const (
	indicator_NUMERIC      string = "0001"
	indicator_ALPHANUMERIC string = "0010"
	indicator_BYTE         string = "0100"
)

var qrModeRegexes = map[string]string{
	mode_NUMERIC:      "^\\d+$",
	mode_ALPHANUMERIC: "^[\\dA-Z $%*+\\-./:]+$",
	mode_BYTE:         "^[\\x00-\\xff]+$",
}

var qrModeIndices = map[string]int{
	mode_NUMERIC:      0,
	mode_ALPHANUMERIC: 1,
	mode_BYTE:         2,
}

var qrModeIndicators = map[string]string{
	mode_NUMERIC:      indicator_NUMERIC,
	mode_ALPHANUMERIC: indicator_ALPHANUMERIC,
	mode_BYTE:         indicator_BYTE,
}

var qrCountIndLengths = map[string]int{
	mode_NUMERIC:      10,
	mode_ALPHANUMERIC: 9,
	mode_BYTE:         8,
}

var qrCapacities = map[int]map[rune][]int{
	1: {
		ec_LOW:      {41, 25, 17},
		ec_MEDIUM:   {34, 20, 14},
		ec_QUARTILE: {27, 16, 11},
		ec_HIGH:     {17, 10, 7},
	},
	2: {
		ec_LOW:      {77, 47, 32},
		ec_MEDIUM:   {63, 38, 26},
		ec_QUARTILE: {48, 29, 20},
		ec_HIGH:     {34, 20, 14},
	},
	3: {
		ec_LOW:      {127, 77, 53},
		ec_MEDIUM:   {101, 61, 42},
		ec_QUARTILE: {77, 47, 32},
		ec_HIGH:     {58, 35, 24},
	},
	4: {
		ec_LOW:      {187, 114, 78},
		ec_MEDIUM:   {149, 90, 62},
		ec_QUARTILE: {111, 67, 46},
		ec_HIGH:     {82, 50, 34},
	},
	5: {
		ec_LOW:      {255, 154, 106},
		ec_MEDIUM:   {202, 122, 84},
		ec_QUARTILE: {144, 87, 60},
		ec_HIGH:     {106, 64, 44},
	},
}

func newVersioner() *versioner {
	return &versioner{}
}

func (v *versioner) getMode(s string) (string, error) {
	matched, err := regexp.MatchString(qrModeRegexes[mode_NUMERIC], s)
	if err != nil {
		return "", err
	} else if matched {
		return mode_NUMERIC, nil
	}

	matched, err = regexp.MatchString(qrModeRegexes[mode_ALPHANUMERIC], s)
	if err != nil {
		return "", err
	} else if matched {
		return mode_ALPHANUMERIC, nil
	}

	matched, err = regexp.MatchString(qrModeRegexes[mode_BYTE], s)
	if err != nil {
		return "", err
	} else if matched {
		return mode_BYTE, nil
	}
	return string(""), fmt.Errorf("invalid input pattern")
}

func (v *versioner) getVersion(s string, mode string, lvl rune) (int, error) {
	version := 1
	for version <= len(qrCapacities) {
		if len(s) <= qrCapacities[version][lvl][qrModeIndices[mode]] {
			return int(version), nil
		}
		version += 1
	}
	return 0, fmt.Errorf("cannot compute qr int")
}

func (v *versioner) getModeIndicator(mode string) string {
	return qrModeIndicators[mode]
}

func (v *versioner) getCountIndicator(s string, vrs int, string string) string {
	sLenBin := strconv.FormatInt(int64(len(s)), 2)
	return padLeft(sLenBin, "0", qrCountIndLengths[string])
}
