package handy

import "net/http"

// HTTPRequestAsString gets a parameter coming from a http request as string, truncated to maxLenght
// Only maxLenght >= 1 is considered. Otherwise, it's ignored
func HTTPRequestAsString(r *http.Request, key string, maxLenght int) string {
	s := r.FormValue(key)

	if s=="" {
		return ""
	}

	if ( maxLenght > 0 ) && ( len([]rune(s)) >= maxLenght ) {
		return s[0:maxLenght]
	}

	return s
}

// HTTPRequestAsinteger gets a parameter coming from a http request as an integer
// It tries to guess if it's a signed/negative integer
func HTTPRequestAsInteger(r *http.Request, key string) int {
	s := r.FormValue(key)

	if s=="" {
		return 0
	}

	neg := s[0:1] == "-"

	i := StringAsInteger(s)

	if neg && ( i > 0 ) {
		return i*-1
	}

	return i
}

// HTTPRequestAsFloat gets a parameter coming from a http request as float64 number
// You have to inform the decimal separator symbol.
// If decimalSeparator is period, engine considers thousandSeparator is comma, and vice-versa.
func HTTPRequestAsFloat64(r *http.Request, key string, decimalSeparator rune) float64 {
	s := r.FormValue(key)

	if s=="" {
		return 0
	}

	thousandSeparator := Tif(decimalSeparator == ',', '.', ',').(rune)

	return StringAsFloat(s, decimalSeparator, thousandSeparator)
}

