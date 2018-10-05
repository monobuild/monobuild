package methods

import "github.com/sirupsen/logrus"

// SetLogLevel sets the log level bases on debug, info or *
// for any value other than debug or info warn is used
func SetLogLevel(desiredLogLevel string) {
	switch desiredLogLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
		break
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
		break
	default:
		logrus.SetLevel(logrus.WarnLevel)
		break
	}
}
