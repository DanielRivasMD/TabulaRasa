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
	Ω := make([]rep, 5)
	Ω[0] = rep{old: "YEAR", new: strconv.Itoa(time.Now().Year())}
	Ω[1] = rep{old: "REPOSITORY", new: repo}
	Ω[2] = rep{old: "TOOL", new: strings.ToLower(repo)}
	Ω[3] = rep{old: "AUTHOR", new: author}
	Ω[4] = rep{old: "EMAIL", new: email}
	return Ω
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// bind cobra replace values
func replaceCobraCmd() []rep {
	Ω := make([]rep, 7)
	Ω[0] = rep{old: "YEAR", new: strconv.Itoa(time.Now().Year())}
	Ω[1] = rep{old: "REPOSITORY", new: repo}
	Ω[2] = rep{old: "TOOL", new: strings.ToLower(repo)}
	Ω[3] = rep{old: "AUTHOR", new: author}
	Ω[4] = rep{old: "EMAIL", new: email}
	Ω[5] = rep{old: "CHILD", new: strings.ToLower(child)}
	Ω[6] = rep{old: "PARENT", new: strings.ToLower(parent)}
	return Ω
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// bind deploy just replace values
func replaceDeployJust() []rep {
	Ω := make([]rep, 3)
	Ω[0] = rep{"APP", repo}
	Ω[1] = rep{"EXE", strings.ToLower(repo)}
	Ω[2] = rep{"VER", ver}
	return Ω
}

////////////////////////////////////////////////////////////////////////////////////////////////////
