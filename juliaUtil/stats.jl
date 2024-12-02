####################################################################################################

# median
function medianAbsDev(data, median)
  dev = 0
  for i in eachindex(data)
    dev += abs(data[i] - median)
  end
  return dev /= length(data)
end

####################################################################################################

