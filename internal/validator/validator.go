package validator

import (
	"context"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

// Email regex pattern (RFC 5322 simplified)
var EmailRX = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

type Validator interface {
	Valid(context.Context) Evaluator
}

type Evaluator map[string]any

func (e *Evaluator) AddFieldErr(key, msg string) {
	if *e == nil {
		*e = make(map[string]any)
	}

	if _, exist := (*e)[key]; !exist {
		(*e)[key] = msg
	}
}

func (e *Evaluator) CheckField(ok bool, key, msg string) {
	if !ok {
		e.AddFieldErr(key, msg)
	}
}

// HasErrors returns true if there are validation errors
func (e Evaluator) HasErrors() bool {
	return len(e) > 0
}

// String Validators

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// Email Validator

func ValidEmail(value string) bool {
	if len(value) > 254 {
		return false
	}
	return EmailRX.MatchString(value)
}

// Password Validators

// MinPasswordStrength checks if password has minimum 8 chars, at least one uppercase,
// one lowercase, and one number
func MinPasswordStrength(value string) bool {
	if len(value) < 8 {
		return false
	}

	var hasUpper, hasLower, hasNumber bool

	for _, char := range value {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}

	return hasUpper && hasLower && hasNumber
}

// Number Validators

func GreaterThan(value, n float64) bool {
	return value > n
}

func GreaterThanOrEqual(value, n float64) bool {
	return value >= n
}

func LessThan(value, n float64) bool {
	return value < n
}

func LessThanOrEqual(value, n float64) bool {
	return value <= n
}

func PositiveNumber(value float64) bool {
	return value > 0
}

func InRange(value, min, max float64) bool {
	return value >= min && value <= max
}

// ValidBidIncrement checks if the new bid is at least current price + minimum increment
func ValidBidIncrement(bid, current, minIncrement float64) bool {
	return bid >= current+minIncrement
}

// Date Validators

func FutureDate(date time.Time) bool {
	return date.After(time.Now())
}

func PastDate(date time.Time) bool {
	return date.Before(time.Now())
}

func DateAfter(date, reference time.Time) bool {
	return date.After(reference)
}

func DateBefore(date, reference time.Time) bool {
	return date.Before(reference)
}

// ValidAuctionDuration checks if auction has valid duration (end after start, minimum duration)
func ValidAuctionDuration(start, end time.Time, minDurationHours int) bool {
	if !end.After(start) {
		return false
	}
	duration := end.Sub(start)
	return duration >= time.Duration(minDurationHours)*time.Hour
}

// Generic Validators

// PermittedValue checks if value is in the list of permitted values
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	for _, permitted := range permittedValues {
		if value == permitted {
			return true
		}
	}
	return false
}

// Unique checks if all values in a slice are unique
func Unique[T comparable](values []T) bool {
	seen := make(map[T]bool)
	for _, v := range values {
		if seen[v] {
			return false
		}
		seen[v] = true
	}
	return true
}

// URL Validator

func ValidURL(value string) bool {
	u, err := url.Parse(value)
	if err != nil {
		return false
	}
	return u.Scheme != "" && u.Host != ""
}

// ValidHTTPURL checks if URL is valid and uses HTTP/HTTPS scheme
func ValidHTTPURL(value string) bool {
	u, err := url.Parse(value)
	if err != nil {
		return false
	}
	return (u.Scheme == "http" || u.Scheme == "https") && u.Host != ""
}
