#https://www.hackerrank.com/challenges/fp-filter-array/problem
defmodule Solution do
    def main(limit, data) do
        list = convert_list(data)
        for d <- list, d < limit, do: IO.puts d
    end
    
    def convert_list(data) do
        data
        |> String.trim
        |> String.split
        |> Enum.map(&String.to_integer/1)
    end
end

limit = IO.gets("") |> String.trim |> String.to_integer
data = IO.read(:stdio, :all)
Solution.main(limit, data)

