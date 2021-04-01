package remotecheck

// Support represents the Arm support status of an app
type Support uint

const (
	// SupportUndefined is the zero value of Support type
	SupportUndefined Support = iota // zero value
	// SupportNative means an app have native Apple Silicon support
	SupportNative
	// SupportTransition means an app is supported vis Rosetta 2 or Virtual Environment
	SupportTransition
	// SupportInDevelopment means an app does not support Apple Silicon yet but the support is in development
	SupportInDevelopment
	// SupportInDevelopment means an app does not support Apple Silicon
	SupportNotYet
	// SupportUnknown means it's not known if an app supports Apple Silicon
	SupportUnknown
)

// Parse parses support information from string.
func (s *Support) Parse(str string) Support {
	switch str {
	case "✅":
		*s = SupportNative

	case "✳️":
		*s = SupportTransition

	case "⏹":
		*s = SupportInDevelopment

	case "🚫":
		*s = SupportNotYet

	case "🔶":
		*s = SupportUnknown

	}

	return *s
}

func (s Support) String() string {
	switch s {
	case SupportNative:
		return "Supported"

	case SupportTransition:
		return "Supported*"

	case SupportInDevelopment:
		return "Unsupported"

	case SupportNotYet:
		return "Unsupported"

	case SupportUnknown:
		return "Unknown"

	default:
		return "Unknown"

	}
}
