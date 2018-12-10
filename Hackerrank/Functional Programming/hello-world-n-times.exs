#https://www.hackerrank.com/challenges/fp-hello-world-n-times/problem
defmodule Solution do
#Enter your code here. Read input from STDIN. Print output to STDOUT
    def main(input) do
        for _ <- 1..input, do: IO.puts "Hello World"
    end
end

input = IO.gets("") |> String.to_integer
Solution.main(input)

#stdin
#https://hr-testcases-us-east-1.s3.amazonaws.com/748/input00.txt?AWSAccessKeyId=AKIAJ4WZFDFQTZRGO3QA&Expires=1544467849&Signature=lSI1bLxUZoKwyHGenSQCUWNUI44%3D&response-content-type=text%2Fplain

#expected 
#https://hr-testcases-us-east-1.s3.amazonaws.com/748/output00.txt?AWSAccessKeyId=AKIAJ4WZFDFQTZRGO3QA&Expires=1544467914&Signature=jNTmo5OVk41oeJCd8k8E7vYM8jc%3D&response-content-type=text%2Fplain
