package remotecheck

type Support uint

const (
	SupportUndefined Support = iota // zero value
	SupportNative
	SupportTransition
	SupportInDevelopment
	SupportNotYet
	SupportUnknown
)

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
		return ""

	}
}
