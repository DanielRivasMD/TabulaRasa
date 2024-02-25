#!/bin/bash
####################################################################################################

# declarations
repository=$1
tool=$2
author_email=$3

####################################################################################################

# execute
sd 'YEAR' $(awk '{print $NF}' <(date)) $(rg --files-with-matches "YEAR")
sd 'REPOSITORY' "${repository}" $(rg --files-with-matches "REPOSITORY")
sd 'TOOL' "${tool}" $(rg --files-with-matches "TOOL")
sd 'AUTHOR_EMAIL' "${author_email}" $(rg --files-with-matches "AUTHOR_EMAIL")

####################################################################################################
