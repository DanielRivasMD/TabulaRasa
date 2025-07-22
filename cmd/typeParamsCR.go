////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

// CopyParams holds information needed to copy a template tree and apply string replacements.
type CopyParams struct {
	Orig  string   // source template directory
	Dest  string   // destination directory
	Files []string // specific files to copy (nil = all)
	Reps  []rep    // replacement rules
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// newCopyParams creates a fresh CopyParams struct.
func newCopyParams(orig, dest string) CopyParams {
	return CopyParams{
		Orig:  orig,
		Dest:  dest,
		Files: []string{},
		Reps:  []rep{},
	}
}

// withReplacements initializes a CopyParams with a set of replacements.
func withReplacements(orig string, reps []rep) CopyParams {
	return CopyParams{
		Orig: orig,
		Reps: reps,
	}
}

// cloneCopyParams makes a shallow copy of an existing CopyParams.
func cloneCopyParams(src CopyParams) CopyParams {
	return CopyParams{
		Orig:  src.Orig,
		Dest:  src.Dest,
		Files: src.Files,
		Reps:  src.Reps,
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
