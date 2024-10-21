////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////

// copy & replace
type paramsCopyReplace struct {
	orig  string
	dest  string
	files []string
	reps  []rep
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func creatorCopyReplace() paramsCopyReplace {
	params := paramsCopyReplace{}
	params.files = []string{}
	params.reps = []rep{}
	return params
}

func replacerCopyReplace(target string, reps []rep) paramsCopyReplace {
	params := paramsCopyReplace{}
	params.orig = target
	params.reps = reps
	return params
}

func copierCopyReplace(orig, dest string) paramsCopyReplace {
	params := creatorCopyReplace()
	params.orig = orig
	params.dest = dest
	return params
}

func clonerCopyReplace(params paramsCopyReplace) paramsCopyReplace {
	out := creatorCopyReplace()
	out.orig = params.orig
	out.dest = params.dest
	out.files = params.files
	out.reps = params.reps
	return out
}

////////////////////////////////////////////////////////////////////////////////////////////////////
