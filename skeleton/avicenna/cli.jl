####################################################################################################

module XXX_ROOT2_XXXCLI

####################################################################################################

using ArgParse
using Avicenna.Flow: Cache, launch
using ..XXX_ROOT2_XXXFlow: flow

####################################################################################################

export run

####################################################################################################

function main(args = ARGS)
  s = ArgParseSettings()
  @add_arg_table! s begin
    "--no-cache"
    help = "Disable caching"
    action = :store_true
    "--verbose"
    help = "Enable verbose diagnostics"
    action = :store_false
  end

  config = Dict()

  cache = Cache("cache/_", !parsed["no-cache"])
  result = launch(flow, config, cache = cache)

  if parsed["verbose"]
  end
  return result
end

####################################################################################################

end

####################################################################################################
