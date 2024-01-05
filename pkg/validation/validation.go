// pkg/utils/validation.go

package utils

import (
    "unicode"
    "unicode/utf8"
	"regexp"
)

func ValidateEmail(email string) bool {
    pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    match, _ := regexp.MatchString(pattern, email)
    return match
}
func ValidatePassword(password string) bool {
    if utf8.RuneCountInString(password) < 8 {
        return false // Şifre en az 8 karakter olmalıdır
    }

    var (
        hasUpperCase bool
        hasLowerCase bool
        hasNumber    bool
        hasSymbol    bool
    )

    for _, char := range password {
        switch {
        case unicode.IsUpper(char):
            hasUpperCase = true
        case unicode.IsLower(char):
            hasLowerCase = true
        case unicode.IsNumber(char):
            hasNumber = true
        case unicode.IsPunct(char) || unicode.IsSymbol(char):
            hasSymbol = true
        }
    }

    // Şifre en az bir büyük harf, bir küçük harf, bir rakam ve bir sembol içermelidir
    return hasUpperCase && hasLowerCase && hasNumber && hasSymbol
}
