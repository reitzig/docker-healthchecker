def pretty_duration(later, sooner)
    (Time.at(later - sooner)).utc.strftime("%Hh%Mm%Ss")
end
