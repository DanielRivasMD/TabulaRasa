####################################################################################################

# load dependency
using DataFrames

####################################################################################################

"""automatic mapper"""
function walker(inPath)
  # preallocate
  Ω = begin
    levels = 4
    DataFrame([String[] for _ in 1:levels], [string("lv$(n)") for n in 0:(levels - 1)])
  end

  # walk from root
  for lv1 ∈ readdir(inPath)
    if isdir(string(inPath, "/", lv1))
      for lv2 ∈ readdir(string(inPath, "/", lv1))
        if isdir(string(inPath, "/", lv1, "/", lv2))
          for lv3 ∈ readdir(string(inPath, "/", lv1, "/", lv2))
            push!(Ω, [inPath, lv1, lv2, lv3])
          end
        end
      end
    end
  end

  return Ω
end

####################################################################################################

"""recursive constructor"""
function builder(df::DataFrame, dir)
  # construct paths dataframe
  Ω = df[:, Not(end)] |> unique
  Ω[:, :lv0] .= dir

  # build directories
  for ρ ∈ eachrow(Ω)
    mkpath(join(collect(ρ), "/"))
  end

  return Ω
end

####################################################################################################

