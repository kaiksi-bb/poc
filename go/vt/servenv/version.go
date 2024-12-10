package servenv

// versionName holds the current version of the application.
const versionName = "1.2.3"

// GetVersion returns the current version of the application.
func GetVersion() string {
    return versionName
}
