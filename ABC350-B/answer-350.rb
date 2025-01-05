# 歯の本数と治療の回数
n, q = gets.split.map(&:to_i)
# 治療される歯の番号
treatments = gets.split.map(&:to_i)
# 歯が生えている
teeth_present = (1..n).to_a
# 歯が抜かれている
teeth_absent = []

# 治療処理
treatments.each do |t|
  if teeth_present.include?(t)
    # 歯が生えている場合
    teeth_present.delete(t)
    teeth_absent.push(t)
  elsif teeth_absent.include?(t)
    # 歯が抜かれている場合
    teeth_absent.delete(t)
    teeth_present.push(t)
  end
end

# 最終的に歯が生えている配列の数を出力
puts teeth_present.length
