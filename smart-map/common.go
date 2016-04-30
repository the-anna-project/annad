package smartmap

import (
	"strings"
)

func withoutPrefix(w string) (interface{}, error) {
	if strings.HasPrefix(w, StringPrefix) {
		wo := strings.Replace(w, StringPrefix, "", 1)
		return wo, nil
	}

	if strings.HasPrefix(w, StringSlicePrefix) {
		wo := strings.Replace(w, StringSlicePrefix, "", 1)
		var ss []string
		for _, s := range strings.Split(wo, ",") {
			ss = append(ss, s)
		}
		return ss, nil
	}

	return nil, maskAnyf(prefixNotFoundError, "%#v", w)
}

func withPrefix(wo interface{}) (string, error) {
	switch t := wo.(type) {
	case string:
		w := StringPrefix + t
		return w, nil
	case []string:
		var w = StringSlicePrefix
		for i, s := range t {
			if i > 0 {
				w += ","
			}
			w += s
		}
		return w, nil
	default:
		// TODO test nil
		return "", maskAnyf(unsupportedTypeError, "%T", wo)
	}
}
