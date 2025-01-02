n = gets.to_i

# 特定の値に基づいて処理を分岐
case n
when 0...100
  puts "00"
when 100..5000
  puts format("%02d", n / 100)
when 6000..30000
  puts n / 1000 + 50
when 35000..70000
  puts (n / 1000 - 30) / 5 + 80
else
  puts "89"
end
