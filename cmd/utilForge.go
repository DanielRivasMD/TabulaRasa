////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"strings"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

type mbomboReplace struct {
	old string
	new string
}

type mbomboForge struct {
	in       string
	out      string
	files    []string
	replaces []mbomboReplace
}

type filePair struct {
	files []string
	out   string
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func newMbomboConfig(
	inDir, outFile string,
	tplFiles []string,
	replaces ...mbomboReplace,
) mbomboForge {
	return mbomboForge{
		in:       inDir,
		out:      outFile,
		files:    tplFiles,
		replaces: replaces,
	}
}

func mbomboForging(op string, mf mbomboForge) {
	horus.CheckErr(
		domovoi.ExecSh(mf.Cmd()),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to execute mbombo forge command"),
		horus.WithDetails(map[string]any{
			"command": mf.Cmd(),
		}),
	)
}

func Replace(key, val string) mbomboReplace {
	return mbomboReplace{old: key, new: val}
}

func (m mbomboForge) Cmd() string {
	var files []string
	for _, f := range m.files {
		files = append(files, fmt.Sprintf(`--files %s`, f))
	}
	fileBlock := strings.Join(files, " \\\n")

	var replaces []string
	for _, r := range m.replaces {
		replaces = append(replaces, fmt.Sprintf(`--replace %s="%s"`, r.old, r.new))
	}
	replaceBlock := strings.Join(replaces, " \\\n")

	return fmt.Sprintf(
		`mbombo forge \
--in %s \
--out %s \
%s \
%s`,
		m.in,
		m.out,
		fileBlock,
		replaceBlock,
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
