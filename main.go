package simplesoapserver

import (
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

func SoapRequestFieldParse(inputs []string, b []byte) (map[string][]string, error) {
	var err error
	var match []string
	var input string
	export := make(map[string][]string)
	for _, input = range inputs {
		match, err = SoapFindField(b, input)
		export[input] = match
	}
	return export, err
}

func SoapFindField(b []byte, mustbefind string) ([]string, error) {
	mached, err := regexp.Match("(?i)<"+mustbefind+">([^<]+)</"+mustbefind+">(?-i)", b)
	if err != nil {
		return nil, err
	}
	if mached == false {
		return nil, errors.New("we have not found: " + mustbefind)
	}
	r := regexp.MustCompile("(?i)<" + mustbefind + ">([^<]+)</" + mustbefind + ">(?-i)")
	matches := r.FindAll(b, -1)
	startEXps := regexp.MustCompile("(?i)<" + mustbefind + ">(?-i)")
	startEXp := startEXps.FindAll(b, 1)
	endEXPs := regexp.MustCompile("(?i)</" + mustbefind + ">(?-i)")
	endEXP := endEXPs.FindAll(b, 1)
	var export []string
	var finded string
	for _, match := range matches {
		finded = strings.ReplaceAll(string(match), string(startEXp[0]), "")
		finded = strings.ReplaceAll(finded, string(endEXP[0]), "")
		export = append(export, finded)
	}
	return export, err
}

func FindSoapMethod(b []byte, mustbefinds []string) (string, error) {
	for _, mustbefind := range mustbefinds {
		mached, err := regexp.Match("(?i)<"+mustbefind+"(?-i)", b)
		if err != nil {
			return "", err
		}
		if mached == true {
			return mustbefind, nil
		}

	}
	return "", errors.New("we have not found")
}

func SoapResponse(method string, KeyValues [][]string) string {
	str := ""
	for _, keyvalue := range KeyValues {
		str += "        <" + keyvalue[0] + ">" + keyvalue[1] + "</" + keyvalue[0] + ">\n"
	}
	s := "<soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\">\n  <soap:Body>\n    <" + method + "Response xmlns=\"http://tempuri.org/\">\n      <" + method + "Result>\n" + str + "      </" + method + "Result>\n    </" + method + "Response>\n  </soap:Body>\n</soap:Envelope>"
	return s
}
