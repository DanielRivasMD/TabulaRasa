#!/bin/bash
####################################################################################################

# declarations
repository=$1
author_email=$2

####################################################################################################

# execute
sd 'YEAR' $(awk '{print $NF}' <(date)) $(rg --files-with-matches "YEAR")
sd 'REPOSITORY' "${repository}" $(rg --files-with-matches "REPOSITORY")
sd 'AUTHOR_EMAIL' "${author_email}" $(rg --files-with-matches "AUTHOR_EMAIL")

####################################################################################################
