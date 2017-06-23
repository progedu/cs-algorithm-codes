#! /usr/bin/env ruby
#
# 演習 1-3 の入力生成スクリプトです。
# これ自体は教材の一部ではなく、このプログラムが生成するファイルを教材として
# 使うことを想定しています（なので受講者に解説していない Ruby で記述している）。

def printCase(filename, a, x)
  open(filename, "w") do |fp|
    fp.puts a.size
    first = true
    a.each do |e|
      fp.print ' ' unless first
      fp.print e
      first = false
    end
    fp.puts
    fp.puts x
  end
end

def generateArray(n)
  a = [1]
  (n - 1).times do
    last = a[-1]
    a.push(last + rand(3) + 1)
  end
  a.shuffle
end

def main()
  printCase("ex1-3-n10.in", [1, 2, 3, 5, 8, 10, 15, 24, 36, 51].shuffle, 24)

  n = 1000
  a = generateArray(n)
  printCase("ex1-3-n1000.in", a, a[rand(n)])

  n = 1000000
  a = generateArray(n)
  printCase("ex1-3-n1000000.in", a, a[rand(n)])

end

main()
