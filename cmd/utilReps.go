////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"strconv"
	"strings"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// bind cobra replace values
func replaceCobraApp() []rep {
	out := make([]rep, 5)
	out[0] = rep{old: "YEAR", new: strconv.Itoa(time.Now().Year())}
	out[1] = rep{old: "REPOSITORY", new: repo}
	out[2] = rep{old: "TOOL", new: strings.ToLower(repo)}
	out[3] = rep{old: "AUTHOR", new: author}
	out[4] = rep{old: "EMAIL", new: email}
	return out
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// bind cobra replace values
func replaceCobraCmd() []rep {
	out := make([]rep, 7)
	out[0] = rep{old: "YEAR", new: strconv.Itoa(time.Now().Year())}
	out[1] = rep{old: "REPOSITORY", new: repo}
	out[2] = rep{old: "TOOL", new: strings.ToLower(repo)}
	out[3] = rep{old: "AUTHOR", new: author}
	out[4] = rep{old: "EMAIL", new: email}
	out[5] = rep{old: "CHILD", new: strings.ToLower(child)}
	out[6] = rep{old: "PARENT", new: strings.ToLower(parent)}
	return out
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// bind deploy just replace values
func replaceDeployJust() []rep {
	out := make([]rep, 3)
	out[0] = rep{"APP", repo}
	out[1] = rep{"EXE", strings.ToLower(repo)}
	out[2] = rep{"VER", ver}
	return out
}

////////////////////////////////////////////////////////////////////////////////////////////////////
