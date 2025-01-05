n = gets.to_i
t = gets.split.map(&:to_i)
m = gets.to_i

# 問題の初期合計時間を計算
total_time = t.sum

# 各ドリンクの影響を計算し、出力
m.times do
  p_j, x_j = gets.split.map(&:to_i)
  # ドリンクを飲んだ場合の時間を計算
  adjusted_time = total_time - t[p_j - 1] + x_j
  puts adjusted_time
end
