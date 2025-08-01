####################################################################################################

_default:
  @just --list

####################################################################################################

# print justfile
[group('just')]
@show:
  bat .justfile --language make

####################################################################################################

# edit justfile
[group('just')]
@edit:
  micro .justfile

####################################################################################################

# compile roadmap action items
[group('dev')]
@roadmap:
  { echo '\n=================================================='; todor -s; echo '=================================================='; } >> ROADMAP.txt

####################################################################################################
# import
####################################################################################################

# config
import '.just/go.conf'

####################################################################################################
# jobs
####################################################################################################

# build for OSX
[group('go')]
osx app=goapp:
  @echo "\n\033[1;33mBuilding\033[0;37m...\n=================================================="
  go build -v -o excalibur/{{app}}

####################################################################################################

# build for linux
[group('go')]
linux app=goapp:
  @echo "\n\033[1;33mBuilding\033[0;37m...\n=================================================="
  env GOOS=linux GOARCH=amd64 go build -v -o excalibur/{{app}}

####################################################################################################

# install locally
[group('go')]
install app=goapp exe=goexe dir=dir:
  @echo "\n\033[1;33mInstalling\033[0;37m...\n=================================================="
  go install
  @echo "\n\033[1;33mLinking\033[0;37m...\n=================================================="
  @mv -v "${HOME}/go/bin/{{app}}" "${HOME}/go/bin/{{exe}}"
  @echo "\n\033[1;33mCopying\033[0;37m...\n=================================================="
  @if [ ! -d "${HOME}/{{dir}}" ]; then mkdir "${HOME}/{{dir}}"; fi
  @if test -e "${HOME}/{{cobraApp}}"; then rm -r "${HOME}/{{cobraApp}}"; fi && echo "\033[1;33mcobraApp\033[0;37m" && cp -v -R "cobraApp" "${HOME}/{{cobraApp}}"
  @if test -e "${HOME}/{{cobraCmd}}"; then rm -r "${HOME}/{{cobraCmd}}"; fi && echo "\033[1;33mcobraCmd\033[0;37m" && cp -v -R "cobraCmd" "${HOME}/{{cobraCmd}}"
  @if test -e "${HOME}/{{cobraUtil}}"; then rm -r "${HOME}/{{cobraUtil}}"; fi && echo "\033[1;33mcobraUtil\033[0;37m" && cp -v -R "cobraUtil" "${HOME}/{{cobraUtil}}"
  @if test -e "${HOME}/{{just}}"; then rm -r "${HOME}/{{just}}"; fi && echo "\033[1;33mjust\033[0;37m" && cp -v -R "just" "${HOME}/{{just}}"
  @if test -e "${HOME}/{{readme}}"; then rm -r "${HOME}/{{readme}}"; fi && echo "\033[1;33mreadme\033[0;37m" && cp -v -R "readme" "${HOME}/{{readme}}"
  @if test -e "${HOME}/{{todor}}"; then rm -r "${HOME}/{{todor}}"; fi && echo "\033[1;33mtodor\033[0;37m" && cp -v -R "todor" "${HOME}/{{todor}}"

####################################################################################################

# watch changes
[group('go')]
watch:
  watchexec --clear --watch cmd -- 'just install'

####################################################################################################
