####################################################################################################
# import
####################################################################################################

# config
import '.just/py.conf'

####################################################################################################
# jobs
####################################################################################################

# create enviroment
[group('python')]
@env:
  echo "\n\033[1;33mBuilding\033[0;37m...\n=================================================="
  python3 -m venv venv

####################################################################################################

# link bpython
[group('python')]
@link ver=pyver:
  echo "\n\033[1;33mLinking\033[0;37m...\n=================================================="
  ln -sf $(which bpython) venv/bin/{{ver}}

####################################################################################################

# build & install app
[group('python')]
@install:
  source .just/pyinstall.sh

####################################################################################################

# watch changes
[group('python')]
watch:
  watchexec --clear --watch {{pyapp}} -- 'source .just/pyinstall.sh'

####################################################################################################
