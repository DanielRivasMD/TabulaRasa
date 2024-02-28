####################################################################################################

_default:
  @just --list

####################################################################################################

# print justfile
@show:
  bat .justfile --language make

####################################################################################################

# edit justfile
@edit:
  micro .justfile

####################################################################################################
# aliases
####################################################################################################

import '.config.just'

####################################################################################################

# build for OSX
osx APP=app:
  @echo "Building..."
  go build -v -o excalibur/{{APP}}

####################################################################################################

# build for linux
linux APP=app:
  @echo "Building..."
  env GOOS=linux GOARCH=amd64 go build -v -o excalibur/{{APP}}

####################################################################################################

# install locally
install APP=app exe=exe als=als:
  @echo "Install..."
  go install
  @cp -v "${HOME}/.go/bin/{{APP}}" "${HOME}/.go/bin/{{als}}"
  @mv -v "${HOME}/.go/bin/{{APP}}" "${HOME}/.go/bin/{{exe}}"
  @if test -e "${HOME}/{{cobra}}"; then rm -rv "${HOME}/{{cobra}}"; fi && cp -v -R "cobra" "${HOME}/{{cobra}}"
  @if test -e "${HOME}/{{just}}"; then rm -rv "${HOME}/{{just}}"; fi && cp -v -R "just" "${HOME}/{{just}}"
  @if test -e "${HOME}/{{todor}}"; then rm -rv "${HOME}/{{todor}}"; fi && cp -v -R "todor" "${HOME}/{{todor}}"

####################################################################################################

# watch changes
dev:
  watchexec --clear --watch cmd -- 'just install'

####################################################################################################
