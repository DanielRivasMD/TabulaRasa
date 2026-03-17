/*
Copyright © 2024 Daniel Rivas <danielrivasmd@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

type moldReplace struct{ old, new string }
type moldForge struct {
	in       string
	out      string
	files    []string
	replaces []moldReplace
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func newMoldConfig(inDir, outFile string, tplFiles []string, replaces ...moldReplace) moldForge {
	return moldForge{in: inDir, out: outFile, files: tplFiles, replaces: replaces}
}

func moldForging(op string, mf moldForge) {
	horus.CheckErr(
		domovoi.ExecSh(mf.Cmd()),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to execute mbombo forge command"),
		horus.WithDetails(map[string]any{"command": mf.Cmd()}),
	)
}

func Replace(old, new string) moldReplace { return moldReplace{old, new} }

func (m moldForge) Cmd() string {
	var files []string
	for _, f := range m.files {
		files = append(files, fmt.Sprintf("--files %s", f))
	}
	fileBlock := strings.Join(files, " \\\n")

	var replaces []string
	for _, r := range m.replaces {
		replaces = append(replaces, fmt.Sprintf(`--replace %s="%s"`, r.old, r.new))
	}
	replaceBlock := strings.Join(replaces, " \\\n")

	return fmt.Sprintf(`mbombo forge \
--in %s \
--out %s \
%s \
%s`, m.in, m.out, fileBlock, replaceBlock)
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source: %w", err)
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return fmt.Errorf("failed to copy data: %w", err)
	}

	info, err := os.Stat(src)
	if err == nil {
		err = os.Chmod(dst, info.Mode())
		if err != nil {
			return fmt.Errorf("failed to set permissions: %w", err)
		}
	}

	return nil
}

func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

func upperFirst(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func templateMapping(outputFile string) string {
	ext := filepath.Ext(outputFile)
	base := strings.TrimSuffix(outputFile, ext)
	return base + "_" + ext[1:]
}

////////////////////////////////////////////////////////////////////////////////////////////////////
