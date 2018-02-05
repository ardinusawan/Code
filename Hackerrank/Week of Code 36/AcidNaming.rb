#!/bin/ruby

def acidNaming(acid_name)
    isIC=false
    isHydroIC=false
    if acid_name[-2,2] == 'ic'
        isIC=true
    end
    if acid_name[0,5] == 'hydro'
        isHydroIC=true
    end
    if isIC and isHydroIC
        return 'non-metal acid'
    elsif isIC and not isHydroIC
        return 'polyatomic acid'
    else return 'not an acid'
    end
end

n = gets.strip.to_i
for a0 in (0..n-1)
    acid_name = gets.strip
    result = acidNaming(acid_name)
    puts result
end

