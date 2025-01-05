# 入力を取得
n, a, b = gets.split.map(&:to_i)

# 全体のマス目を構築
(0...(a * n)).each do |i|
  row = ""
  (0...(b * n)).each do |j|
    # タイルの位置に基づいて黒白を判定
    if ((i / a) + (j / b)) % 2 == 0
      row += "."
    else
      row += "#"
    end
  end
  puts row
end
