####################################################################################################

function similarity(ɒ; prs::SimilarParams)
  # preallocate
  ln = length(ɒ) - prs.wsize + 1
  Ω = zeros(ln, ln)

  # iterate
  ι = 0
  while ι <= length(ɒ) - prs.wsize
    ξ = ι + prs.wsize
    global ι += 1
    global ο = 0
    while ο <= length(ɒ) - prs.wsize
      θ = ο + prs.wsize
      global ο += 1
      if ι == ο
        Ω[ι, ο] = 0
      elseif ι < ο
        Ω[ο, ι] = prs.σ(Float64.(ɒ[ι:ξ]), Float64.(ɒ[ο:θ]))
      elseif ι > ο
        Ω[ο, ι] = Ω[ι, ο]
      end
    end
  end

  return Ω
end

####################################################################################################

