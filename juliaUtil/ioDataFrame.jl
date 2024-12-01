###################################################################################################

# load packages
begin
  using DataFrames
  using DelimitedFiles
  using XLSX
end;

###################################################################################################

"read dataframe"
function readdf(path; sep = '\t')
  ƒ, п = readdlm(path, sep, header = true)
  DataFrame(ƒ, п |> vec)
end

###################################################################################################

"read dataframe"
function readxlsx(path)
  ʒ = XLSX.readxlsx(path)
  Ω = DataFrame(ʒ[1][2:end, :], :auto)
  for ι ∈ 1:size(ʒ[1][:], 2)
    rename!(Ω, [ι => ʒ[1][1, ι]])
  end
  return Ω
end

###################################################################################################

"write dataframe"
function writedf(path, df::Df; sep = '\t') where Df <: DataFrame
  toWrite = [(df |> names |> permutedims); (df |> Array)]
  writedlm(path, toWrite, sep)
end

###################################################################################################
