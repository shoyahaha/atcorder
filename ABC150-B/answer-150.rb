n = gets.to_i
s = gets.chomp

count = 0

# 文字列を1文字ずつ確認
i = 0
while i < n - 2
  if s[i] == 'A' && s[i + 1] == 'B' && s[i + 2] == 'C'
    count += 1
    # "ABC" を見つけたら次の部分に進む
    i += 3
  else
    # 次の文字に進む
    i += 1
  end
end

# 結果を出力
puts count
