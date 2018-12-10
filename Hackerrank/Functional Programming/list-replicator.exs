#https://www.hackerrank.com/challenges/fp-list-replication/problem
defmodule Solution do
#Enter your code here. Read input from STDIN. Print output to STDOUT
    def main(replicate, data) do
        elements = convert_list(data)
        Enum.each(elements, fn element ->
            for _ <- 1..replicate, do: IO.puts element
        end)
    end
    
    def convert_list(data) do
    data
    |> String.trim 
    |> String.split 
    |> Enum.map(&String.to_integer/1)
    end
end

replicate = IO.gets("") |> String.trim |> String.to_integer
data = IO.read(:stdio, :all)
#:ok = IO.write(data)
Solution.main(replicate, data)

#stdin
#https://hr-testcases-us-east-1.s3.amazonaws.com/755/input00.txt?AWSAccessKeyId=AKIAJ4WZFDFQTZRGO3QA&Expires=1544470588&Signature=4cOtyFOp3sggoohuEp2Z5lEO9bY%3D&response-content-type=text%2Fplain

#expected
#https://hr-testcases-us-east-1.s3.amazonaws.com/755/output00.txt?AWSAccessKeyId=AKIAJ4WZFDFQTZRGO3QA&Expires=1544470613&Signature=Azq7r%2Fy6ZoLGYFqWcxRdjzcpAzQ%3D&response-content-type=text%2Fplain
