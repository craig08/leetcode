import "strings"

func numUniqueEmails(emails []string) int {
	actualEmails := make(map[string]bool)
	for _, email := range emails {
		at := strings.Index(email, "@")
		localname, domain := email[:at], email[at:]
		if plus := strings.Index(localname, "+"); plus != -1 {
			localname = email[:plus]
		}
		localname = strings.Join(strings.Split(localname, "."), "")
		actualEmails[localname+domain] = true
	}
	return len(actualEmails)
}
