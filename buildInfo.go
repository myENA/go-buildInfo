package buildInfo

// BuildInfo contains build information
type BuildInfo struct {
	Name    string `json:"Name,omitempty"`
	Version string `json:"Version,omitempty"`
	Date    string `json:"BuildDate,omitempty"`
	Branch  string `json:"BuildBranch,omitempty"`
	Build   string `json:"BuildNumber,omitempty"`
}
